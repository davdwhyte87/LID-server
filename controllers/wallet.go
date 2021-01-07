package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davdwhyte87/LID-server/blockchain"
	"github.com/davdwhyte87/LID-server/models"
	"github.com/davdwhyte87/LID-server/utils"
)

// CreateWallet ...
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	var reqData models.CreateWallet
	err := utils.DecodeReq(r, &reqData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	// incase a propagation occurs there will be noned toprocess and propagate again
	if reqData.Address != "" {
		if blockchain.WalletExists(reqData.Address) {
			utils.RespondWithJSON(w, http.StatusCreated, reqData.Address)
			return
		}
	}

	// create wallet
	walletName, err := blockchain.CreateWallet(reqData.Address, reqData.PrivateKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error Creating Wallet")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, walletName)

	// broadcast new wallet creation
	reqData.Address = walletName
	blockchain.BCreateWallet(reqData)
	// print("still working ..")
	return
}

// PrintBlocks ...
func PrintBlocks(w http.ResponseWriter, r *http.Request) {
	var reqData models.PrintBlockReq
	err := utils.DecodeReq(r, &reqData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	chain := blockchain.PrintChain(reqData.Address)
	utils.RespondWithJSON(w, http.StatusOK, chain)
}

// TransferLID ...
func TransferLID(w http.ResponseWriter, r *http.Request) {
	var reqData models.TransferReq
	err := utils.DecodeReq(r, &reqData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	if !blockchain.WalletExists(reqData.SenderAddress) {
		utils.RespondWithError(w, http.StatusCreated, "Sender Wallet does not exist")
		print("Sender Wallet does not exist")
		return
	}

	if !blockchain.WalletExists(reqData.RecieverAddress) {
		utils.RespondWithError(w, http.StatusCreated, "Reciever Wallet does not exist")
		print("Reciever Wallet does not exist")
		return
	}

	// check if req is coming from user or server
	if reqData.SenderBlockID != "" {
		if blockchain.BlockExists(reqData.SenderAddress, reqData.SenderBlockID) {
			utils.RespondWithError(w, http.StatusCreated, "Sender Block does not exist")
			return
		}
	}

	// check if req is coming from user or server
	if reqData.RecieverBlockID != "" {
		// check if block exists
		if blockchain.BlockExists(reqData.RecieverAddress, reqData.RecieverBlockID) {
			utils.RespondWithError(w, http.StatusCreated, "RecieverBlock does not exist")
			return
		}
	}

	// strat transfer transaction
	print("Transfering......")
	amount, _ := strconv.Atoi(reqData.Amount)
	var data models.TrxData
	data.Amount = amount
	if reqData.RecieverBlockID == "" {
		data.IsBroadcasted = false
	} else {
		data.IsBroadcasted = true
	}

	data.Sender = reqData.SenderAddress
	data.Reciever = reqData.RecieverAddress
	data.RecieverBlockID = reqData.RecieverBlockID
	data.SenderBlockID = reqData.SenderBlockID
	data.SPrivateKey = reqData.SenderPrivateKey
	data.RPrivateKey = reqData.RecieverPrivateKey

	if data.IsBroadcasted {
		data.SenderLastBlockAmount, _ = strconv.Atoi(reqData.SenderLastBlockAmount)
		print("LastBlock/..........")
		print(reqData.RecieverLastBlockAmount)
		print("   .     ")
		data.RecieverLastBlockAmount, _ = strconv.Atoi(reqData.RecieverLastBlockAmount)
	}

	// make trnasfer
	newData, errTransfer := blockchain.Transfer(data)

	if errTransfer != nil {
		utils.RespondWithError(w, http.StatusCreated, errTransfer.Error())
		return
	}

	reqData.SenderBlockID = newData.SenderBlockID
	reqData.RecieverBlockID = newData.RecieverBlockID
	reqData.RecieverLastBlockAmount = strconv.Itoa(newData.RecieverLastBlockAmount)
	reqData.SenderLastBlockAmount = strconv.Itoa(newData.SenderLastBlockAmount)
	print("New data           ")
	print(reqData.RecieverLastBlockAmount)
	print("          ")
	utils.RespondWithOk(w, "Transfer Complete")

	// BreadCast to other servers
	blockchain.BroadCastTransfer(reqData)
	return

}

// UserGetBalance ...
// this function gets the balance from the server and requests balance data from
// other servers. It then compares the balance, making a concensus to determin a legitimate figure
func UserGetBalance(w http.ResponseWriter, r *http.Request) {
	var reqData models.GetBalanceReq
	err := utils.DecodeReq(r, &reqData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	// primaryBalance := blockchain.GetBalance(reqData.Address)

	// make requests to other servers

	servers := blockchain.GetServers()
	requestBody, errDe := json.Marshal(map[string]string{
		"Address":        reqData.Address,
	
	})

	if errDe != nil {
		print(errDe.Error())
		return
	}

	// shuffedServers := shuffle(servers)
	x := 3
	for x > 0 {
		x = x - 1
		n := utils.RandomInt(0, len(servers)-1)
		_, err := http.Post(servers[n]+"wallet/transfer", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			print(err.Error())
		}
		// print("currently on: " + servers[x] + "wallet/create")
	}
}
