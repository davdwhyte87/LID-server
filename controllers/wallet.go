package controllers

import (
	"net/http"

	"github.com/davdwhyte87/LID-server/models"
	"github.com/davdwhyte87/LID-server/utils"
	"github.com/davdwhyte87/LID-server/blockchain"
)

// CreateWallet ...
func CreateWallet(w http.ResponseWriter, r *http.Request){
	var reqData models.CreateWallet
	err := utils.DecodeReq(r, &reqData)
	if err != nil{
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}


	// incase a propagation occurs there will be noned toprocess and propagate again
	if blockchain.WalletExists(reqData.Address) {
		utils.RespondWithJSON(w, http.StatusCreated, reqData.Address)
		return
	}

	// create wallet
	walletName, err :=blockchain.CreateWallet(reqData.Address, reqData.PrivateKey)
	if err != nil{
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
func PrintBlocks(w http.ResponseWriter, r *http.Request){
	var reqData models.PrintBlockReq
	err := utils.DecodeReq(r, &reqData)
	if err != nil{
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	chain := blockchain.PrintChain(reqData.Address)
	utils.RespondWithJSON(w, http.StatusOK, chain)
}



// TransferLID ...
func TransferLID(w http.ResponseWriter, r *http.Request){
	var reqData models.TransferReq
	err := utils.DecodeReq(r, &reqData)
	if err != nil{
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	if !blockchain.WalletExists(reqData.SenderAddress) {
		utils.RespondWithError(w, http.StatusCreated, "Sender Wallet does not exist")
		return
	}

	if !blockchain.WalletExists(reqData.RecieverAddress){
		utils.RespondWithError(w, http.StatusCreated, "Reciever Wallet does not exist")
		return
	}


	// strat transfer transaction

	blockchain.Transfer(reqData.SenderAddress, reqData.RecieverAddress, reqData.Amount, reqData.SenderPrivateKey, reqData.RecieverPrivateKey)

}