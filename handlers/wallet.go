package handlers

import (
	"encoding/json"
	"net"

	"github.com/davdwhyte87/LID-server/blockchain"
	"github.com/davdwhyte87/LID-server/models"
	"github.com/davdwhyte87/LID-server/utils"
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
	if blockchain.WalletExists(request.WalletName) {
		response := models.ErrorResponse{}
		response.Code = 3
		response.Message ="Wallet Exists"
		utils.RespondTCP(response, conn)
		return
	}

	// create new block 
	

	// create wallet
	_, publicKey, err := blockchain.CreateWallet(request.WalletName, request.PassPhrase)
	if err != nil {
		println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	response := models.CreateWalletResponse{}
	response.PublicKey = publicKey

	
	// turn reponse data to bytes
	responseByte, err := json.Marshal(response)
	if err != nil {
		println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	conn.Write(responseByte)
}
