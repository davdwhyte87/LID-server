package handlers

import (
	"encoding/json"
	"net"
	"strconv"

	"kura_coin/blockchain"
	"kura_coin/models"
	"kura_coin/utils"
)

func CreateWallet(message string, conn net.Conn) {
	// perse message
	request := models.CreateWalletRequest{}
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		println(err.Error())
		utils.Logger.Error().Str("err", err.Error()).Msg("error decoding request")
		conn.Write([]byte(err.Error()))
		return
		//json: Unmarshal(non-pointer main.Request)
	}
	// check if wallet exists

	// if blockchain.WalletExists(request.WalletName) {
	// 	response := models.ErrorResponse{}
	// 	response.Code = 3
	// 	response.Message = "Wallet Exists"
	// 	utils.RespondTCP(response, conn)
	// 	return
	// }

	store := blockchain.NewKvStore[[]blockchain.Block](request.Address, "")
	if store.DbExist() {
		response := models.ErrorResponse{}
		response.Code = 3
		response.Message = "Wallet Exists"
		utils.RespondTCP(response, conn)
		
	
	}

	// create new block

	// create wallet
	_, publicKey, err := blockchain.CreateWallet(request.Address, request.PassPhrase)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error creating wallet")
		println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}

	response := models.CreateWalletResponse{}
	response.PublicKey = publicKey

	utils.RespondTCP(response, conn)

	// broadcast 
	if !request.IsBroadcasted {
		blockchain.BCreateWallet(request)
	}
	
}

func GetBalance(message string, conn net.Conn) {
	// perse message
	request := models.GetBalanceReq{}
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error decoding req")
		conn.Write([]byte(err.Error()))
		return
		//json: Unmarshal(non-pointer main.Request)
	}

	// errResponse := models.ErrorResponse{}

	// get wallet chain
	//chain :=blockchain.GetChain(request.WalletName)

	// _, err = blockchain.GetData(request.WalletName, "/chain.bin", chain)
	// if err != nil {
	// 	utils.Logger.Error().Str("err", err.Error()).Msg("error getting chain")
	// }

	store := blockchain.NewKvStore[[]blockchain.Block](request.WalletName, "/chain.bin")
	data, _ := store.GetData()
	utils.Logger.Debug().Any("data", data).Msg("data")
	//print(data)
	// if chain == nil || len(chain) <1 {
	// 	utils.Logger.Error().Msg("error getting chain")
	// 	errResponse.Code = 0
	// 	errResponse.Message = "Cannot get chain"
	// 	utils.RespondTCP(errResponse, conn)
	// 	return
	// }

	// // voting happens here
	var balance float64
	for _, block := range data {
		balance = balance + block.Amount
	}

	response := models.GetBalanceResp{}
	response.Balance = balance
	response.Code = 1

	utils.RespondTCP(response, conn)
}

func TransferHandler(message string, conn net.Conn) {
	request := models.TransferReq{}
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error decoding req")
		conn.Write([]byte(err.Error()))
		return
		//json: Unmarshal(non-pointer main.Request)
	}

	trxData := models.TrxData{}
	amounti, err := strconv.ParseFloat(request.Amount[:len(request.Amount)-1], 64)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error converting amount to int64")
		return	
	}
	trxData.Amount = amounti
	trxData.Sender = request.Sender
	trxData.Reciever = request.Receiver

	_, err = blockchain.Transfer(trxData)

	errResponse := models.ErrorResponse{}
	errResponse.Code = 0
	errResponse.Message = "Error making transfer"
	if err != nil {
		utils.RespondTCP(errResponse, conn)
		return
	}

	response := models.GeneralResponse{}
	response.Code = 1
	response.Message = "Transfer Successfull"
	utils.RespondTCP(response, conn)
	// broadcast trassaction
	if !request.IsBroadcasted {
		blockchain.BroadCastTransfer(request)
	}
}
