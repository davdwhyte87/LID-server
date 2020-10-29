package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"os"

	// "os"
	"strconv"

	"github.com/davdwhyte87/LID-server/blockchain"
	"github.com/davdwhyte87/LID-server/utils"
	"github.com/syndtr/goleveldb/leveldb"
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
	stringForHash := block.PrevHash + strconv.Itoa(block.Amount) + walletName + privateKey
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHash))

	block.Hash = hex.EncodeToString(shaEngine.Sum(nil))
	print(block.Hash)
	err = enc.Encode(block)
	if err != nil {
		return retWalletName, err
	}
	// save hash in database
	// err = db.Put([]byte(generateTransID()), dataBuffer.Bytes(), nil)
	err = saveBlock(walletName, block)

	return retWalletName, err

}

// Transfer ...
func Transfer(sender string, reciever string, amount int, sPrivateKey string, rPrivateKey string) {
	// get senders previous block
	prevBlock := getLastBlock(sender)
	print(prevBlock.Amount)
	// os.Exit(2)
	//get reciever previous bloick
	prevBlockReciever := getLastBlock(reciever)

	walletPathSender := datapath + sender
	walletPathReciever := datapath + reciever

	// Perform verifications for chains and private keys to confirm identity

	// validate sender chain
	vSenChain := blockchain.VerifyChain(sender)
	if !vSenChain {
		return
	}
	//verify reciever chain()
	vRecChain := blockchain.VerifyChain(reciever)
	if !vRecChain {
		return
	}

	// verify senders identity with supplied private key
	vPrivateKey := blockchain.VerifyPrivateKey(sender, sPrivateKey)
	if !vPrivateKey {
		return
	}

	// open wallets

	ops := new(opt.Options)
	ops.NoSync = false

	dbSender, err := leveldb.OpenFile(walletPathSender, ops)
	if err != nil {
		log.Print("db" + err.Error())
		return
	}

	dbReciever, err := leveldb.OpenFile(walletPathReciever, ops)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// take money from the sender

	// check if there is enough money
	if prevBlock.Amount < amount {
		log.Fatal("Not enough funds")
		return
	}

	var dataBuffer bytes.Buffer
	enc := gob.NewEncoder(&dataBuffer)
	// var tempBlockHolder []Block
	var senderBlock Block
	senderBlock.Amount = prevBlock.Amount - amount
	senderBlock.PrevHash = prevBlock.Hash
	senderBlock.Sender = sender
	senderBlock.Reciever = reciever

	// calculate hash
	stringForHash := senderBlock.PrevHash + strconv.Itoa(senderBlock.Amount) + sender + sPrivateKey
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHash))

	senderBlock.Hash = hex.EncodeToString(shaEngine.Sum(nil))

	print(senderBlock.Amount)
	// add block to temp block holder set
	// tempBlockHolder = append(tempBlockHolder, senderBlock)
	// print(block.Hash)
	errenc := enc.Encode(senderBlock)
	if errenc != nil {
		log.Fatal(errenc)
	}
	// check(errenc)

	// get transaction id
	transID := utils.StringWithCharset(24)

	err = dbSender.Put([]byte(transID), dataBuffer.Bytes(), nil)
	saveBlock(sender, senderBlock)
	check(err)
	dbSender.Close()

	// credit reciever

	var recieverBlock Block
	recieverBlock.Amount = prevBlockReciever.Amount + amount
	recieverBlock.PrevHash = prevBlockReciever.Hash
	recieverBlock.Sender = sender
	recieverBlock.Reciever = reciever

	// calculate hash
	stringForHashRec := recieverBlock.PrevHash + strconv.Itoa(recieverBlock.Amount) + reciever + rPrivateKey
	shaEngineRe := sha256.New()
	shaEngineRe.Write([]byte(stringForHashRec))

	recieverBlock.Hash = hex.EncodeToString(shaEngineRe.Sum(nil))

	print("reciever amt")
	print(recieverBlock.Amount)
	var recBlockDataBuffer bytes.Buffer
	encDataRec := gob.NewEncoder(&recBlockDataBuffer)
	errencRec := encDataRec.Encode(recieverBlock)
	check(errencRec)

	// get transID
	transIDRec := utils.StringWithCharset(12)

	// save on disk

	err = dbReciever.Put([]byte(transIDRec), recBlockDataBuffer.Bytes(), nil)
	saveBlock(reciever, recieverBlock)
	check(err)
	dbReciever.Close()

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
