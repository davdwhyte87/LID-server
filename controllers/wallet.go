package controllers

import (
	"net/http"

	"kura_coin/blockchain"
	"kura_coin/models"
	"kura_coin/utils"
)


type WalletController struct {

}

// CreateWallet ...
func  CreateWallet(w http.ResponseWriter, r *http.Request) {
	var reqData models.CreateWallet
	err := utils.DecodeReq(r, &reqData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot decode on server")
		return
	}

	// incase a propagation occurs there will be noned toprocess and propagate again
	if reqData.Address != "" {
		// if blockchain.WalletExists(reqData.Address) {
		// 	utils.RespondWithJSON(w, http.StatusCreated, reqData.Address)
		// 	return
		// }
	}

	// create wallet
	walletName,_, err := blockchain.CreateWallet(reqData.Address, reqData.PrivateKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error Creating Wallet")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, walletName)

	// broadcast new wallet creation
	reqData.Address = walletName
	// blockchain.BCreateWallet(reqData)
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


