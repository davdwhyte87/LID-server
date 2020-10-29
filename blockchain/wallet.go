package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"log"
	"os"

	// "os"
	"strconv"

	"github.com/davdwhyte87/LID-server/models"
	"github.com/davdwhyte87/LID-server/utils"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var datapath = "./temp/"

// CreateWallet ...
func CreateWallet(name string, privateKey string) (string, error) {
	var err error
	var retWalletName string
	var walletName string
	if name != "" {
		if WalletExists(name) {
			print("wallet existss/......")
			return name, err
		}
	}

	log.Print("Createing Wallet ....... ")
	walletName = name
	if name == "" {
		walletName = utils.StringWithCharset(12)
	}
	retWalletName = walletName
	// initail block

	var dataBuffer bytes.Buffer
	enc := gob.NewEncoder(&dataBuffer)
	var block Block
	block.Amount = 10
	block.Sender = "00000000000"
	block.Reciever = walletName
	block.PrevHash = "00000000000"

	// calculate hash
	stringForHash := block.PrevHash + strconv.Itoa(block.Amount) + walletName
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHash))

	block.Hash = hex.EncodeToString(shaEngine.Sum(nil))

	// hash supplied private key
	stringForHashRec := privateKey
	shaEngine2 := sha256.New()
	shaEngine2.Write([]byte(stringForHashRec))

	block.PrivateKeyHash = hex.EncodeToString(shaEngine2.Sum(nil))

	err = enc.Encode(block)
	if err != nil {
		return retWalletName, err
	}
	// save hash in database
	// err = db.Put([]byte(generateTransID()), dataBuffer.Bytes(), nil)
	err = saveBlock(walletName, block)

	return retWalletName, err

}

// func AddTransfer(sender string, reciever string, amount int, sPrivateKey string, rPrivateKey string) (string, string, error) {
// 	// get senders previous block
// 	prevBlock := getLastBlock(sender)
// 	print(prevBlock.Amount)
// 	// os.Exit(2)

// 	var err error
// 	var senderBlockID string
// 	var recieverBlockID string
// 	//get reciever previous bloick
// 	prevBlockReciever := getLastBlock(reciever)

// 	// Perform verifications for chains and private keys to confirm identity

// 	// validate sender chain
// 	vSenChain := VerifyChain(sender)
// 	if !vSenChain {
// 		err = errors.New("Error verifying sender")
// 		return senderBlockID, recieverBlockID, err
// 	}
// 	//verify reciever chain()
// 	vRecChain := VerifyChain(reciever)
// 	if !vRecChain {
// 		err = errors.New("Error verifying reciever")
// 		return senderBlockID, recieverBlockID, err
// 	}

// 	// verify senders identity with supplied private key
// 	vPrivateKey := VerifyPrivateKey(sender, sPrivateKey)
// 	if !vPrivateKey {
// 		err = errors.New("Error verifying sender private key")
// 		return senderBlockID, recieverBlockID, err
// 	}

// 	if prevBlock.Amount < amount {
// 		log.Fatal("Not enough funds")
// 		err = errors.New("Not enought funds for transfer")
// 		return senderBlockID, recieverBlockID, err
// 	}

// }

// Transfer ...
func Transfer(data models.TrxData) (models.TrxData, error) {
	// get senders previous block
	prevBlock := getLastBlock(data.Sender)
	print(prevBlock.Amount)
	// os.Exit(2)

	var err error
	// var senderBlockID string
	// var recieverBlockID string
	//get reciever previous bloick
	prevBlockReciever := getLastBlock(data.Reciever)

	// Perform verifications for chains and private keys to confirm identity

	// validate sender chain
	vSenChain := VerifyChain(data.Sender)
	if !vSenChain {
		err = errors.New("Error verifying sender")
		return data, err
	}
	//verify reciever chain()
	vRecChain := VerifyChain(data.Reciever)
	if !vRecChain {
		err = errors.New("Error verifying reciever")
		return data, err
	}

	// verify senders identity with supplied private key
	vPrivateKey := VerifyPrivateKey(data.Sender, data.SPrivateKey)
	if !vPrivateKey {
		err = errors.New("Error verifying sender private key")
		return data, err
	}

	// open wallets

	ops := new(opt.Options)
	ops.NoSync = false

	// take money from the sender

	// check if there is enough money
	if !data.IsBroadcasted {
		if prevBlock.Amount < data.Amount {
			err = errors.New("Not enought funds for transfer")
			return data, err
		}
	}

	var dataBuffer bytes.Buffer
	enc := gob.NewEncoder(&dataBuffer)
	// var tempBlockHolder []Block
	var senderBlock Block
	if !data.IsBroadcasted {
		senderBlock.Amount = prevBlock.Amount - data.Amount
	}

	if data.IsBroadcasted {
		senderBlock.Amount = data.SenderLastBlockAmount
	}

	senderBlock.PrevHash = prevBlock.Hash
	senderBlock.Sender = data.Sender
	senderBlock.Reciever = data.Reciever
	senderBlock.PrivateKeyHash = prevBlock.PrivateKeyHash
	if !data.IsBroadcasted {
		senderBlock.ID = utils.StringWithCharset(12)
	}

	if data.IsBroadcasted {
		senderBlock.ID = data.SenderBlockID
	}

	// calculate hash
	stringForHash := senderBlock.PrevHash + strconv.Itoa(senderBlock.Amount) + data.Sender
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHash))

	senderBlock.Hash = hex.EncodeToString(shaEngine.Sum(nil))

	print(senderBlock.Amount)
	// add block to temp block holder set
	// tempBlockHolder = append(tempBlockHolder, senderBlock)
	// print(block.Hash)
	errenc := enc.Encode(senderBlock)
	if errenc != nil {
		err = errenc
		log.Fatal(errenc)
		return data, err
	}
	// check(errenc)

	// get transaction id
	// transID := utils.StringWithCharset(24)

	err = saveBlock(data.Sender, senderBlock)
	check(err)

	// credit reciever

	var recieverBlock Block
	if !data.IsBroadcasted {
		recieverBlock.Amount = prevBlockReciever.Amount + data.Amount
	}
	if data.IsBroadcasted {
		recieverBlock.Amount = data.RecieverLastBlockAmount
	}
	recieverBlock.PrevHash = prevBlockReciever.Hash
	recieverBlock.Sender = data.Sender
	recieverBlock.Reciever = data.Reciever
	recieverBlock.PrivateKeyHash = prevBlockReciever.PrivateKeyHash

	if !data.IsBroadcasted {
		recieverBlock.ID = utils.StringWithCharset(12)
	}

	if data.IsBroadcasted {
		recieverBlock.ID = data.RecieverBlockID
	}

	// calculate hash
	stringForHashRec := recieverBlock.PrevHash + strconv.Itoa(recieverBlock.Amount) + data.Reciever
	shaEngineRe := sha256.New()
	shaEngineRe.Write([]byte(stringForHashRec))

	recieverBlock.Hash = hex.EncodeToString(shaEngineRe.Sum(nil))

	print(recieverBlock.Amount)
	var recBlockDataBuffer bytes.Buffer
	encDataRec := gob.NewEncoder(&recBlockDataBuffer)
	errencRec := encDataRec.Encode(recieverBlock)
	if errencRec != nil {
		err = errencRec
		return data, err
	}
	check(errencRec)

	// get transID
	// transIDRec := utils.StringWithCharset(12)

	// save on disk

	err = saveBlock(data.Reciever, recieverBlock)
	check(err)
	data.RecieverBlockID = recieverBlock.ID
	data.SenderBlockID = senderBlock.ID
	data.RecieverLastBlockAmount = recieverBlock.Amount
	data.SenderLastBlockAmount = senderBlock.Amount
	return data, err
}

// GetBalance ...
func GetBalance(address string) {
	block := getLastBlock(address)
	log.Print(block.Amount)
	log.Print("...........")
	log.Print(address)
}

// WalletExists ... Check if a wallet exists on a single server
func WalletExists(address string) bool {
	walletPath := walletDataPath + address + "/"
	_, errstat := os.Stat(walletPath)
	if os.IsNotExist(errstat) {
		print(errstat.Error())
		return false
	}
	// print(errstat.Error())
	return true
}
