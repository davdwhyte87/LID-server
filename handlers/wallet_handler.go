package handlers

import (
	"encoding/json"
	"net"

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

	store := blockchain.NewKvStore[[]blockchain.Block](request.WalletName, "")
	if store.DbExist() {
		response := models.ErrorResponse{}
		response.Code = 3
		response.Message = "Wallet Exists"
		utils.RespondTCP(response, conn)
		return
	}

	// create new block

	// create wallet
	_, publicKey, err := blockchain.CreateWallet(request.WalletName, request.PassPhrase)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error creating wallet")
		println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}

	response := models.CreateWalletResponse{}
	response.PublicKey = publicKey

	utils.RespondTCP(response, conn)

	// turn reponse data to bytes
	// responseByte, err := json.Marshal(response)
	// if err != nil {
	// 	utils.Logger.Error().Str("err",err.Error()).Msg("error encoding response")
	// 	println(err.Error())
	// 	conn.Write([]byte(err.Error()))
	// 	return
	// }
	// conn.Write(responseByte)
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
	var chain []blockchain.Block
	// _, err = blockchain.GetData(request.WalletName, "/chain.bin", chain)
	// if err != nil {
	// 	utils.Logger.Error().Str("err", err.Error()).Msg("error getting chain")
	// }

	store := blockchain.NewKvStore[[]blockchain.Block](request.WalletName, "/chain.bin")
	data, _ := store.GetData(chain)
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
	var balance int64
	for _, block := range chain {
		balance = balance + int64(block.Amount)
	}

	response := models.GetBalanceResp{}
	response.Balance = balance
	response.Code = 1

	utils.RespondTCP(response, conn)
}
