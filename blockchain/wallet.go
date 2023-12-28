package blockchain

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"strings"
	"time"

	// "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"log"
	"os"

	// "os"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	mrand "math/rand"
	"strconv"

	"kura_coin/models"
	"kura_coin/utils"

	"github.com/syndtr/goleveldb/leveldb/opt"
	"gonum.org/v1/gonum/mathext/prng"
)

var datapath = "./temp/"

type Wallet struct {
	Address   string
	CreatedAt string
}

// CreateWallet ...
func CreateWallet(name string, passPhrase string) (string, string, error) {
	var err error
	// generate seed
	slice := []byte(passPhrase)
	tt := binary.LittleEndian.Uint64(slice)
	prng.NewMT19937().Seed(tt)
	seed := int64(tt)
	var myRand = mrand.New(mrand.NewSource(seed))

	// generate key with seeded rand reader
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), myRand)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error generating ecdsa keys")
		panic(err)
	}
	encPriv, encPub := encode(privateKey, &privateKey.PublicKey)

	// create database
	db := NewKvStore[any](name, "")
	createDbError := db.CreateDb()
	if createDbError != nil {
		utils.Logger.Error().Str("err", createDbError.Error()).Msg("error creating database")
		return "", "", createDbError
	}
	// create new block
	var block Block
	block.Amount = 10
	block.Sender = "00000000000"
	block.Reciever = name
	block.PrevHash = "00000000000"

	// calculate hash
	stringForHash := block.PrevHash + strconv.Itoa(block.Amount) + encPub
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHash))
	block.Hash = hex.EncodeToString(shaEngine.Sum(nil))

	// create wallet data
	store := NewKvStore[Wallet](name, "wallet.bin")
	wallet := Wallet{}
	wallet.Address = name
	wallet.CreatedAt = time.Now().GoString()
	svErr := store.SaveData(wallet)
	if svErr != nil {
		utils.Logger.Error().Str("err", svErr.Error()).Msg("error saving data")
		return "", "", svErr
	}

	// create block
	var chain []Block
	var dataBuffer bytes.Buffer
	enc := gob.NewEncoder(&dataBuffer)
	chain = append(chain, block)
	errenc := enc.Encode(chain)
	err = errenc
	if errenc != nil {
		utils.Logger.Error().Str("err", errenc.Error()).Msg("error encoding chain data")
		return "", "", err
	}
	// err = saveBlock(name, block)
	storeb := NewKvStore[[]Block](name, "chain.bin")
	err = storeb.SaveData(chain)
	return encPriv, encPub, err
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}
func decodePrivateKey(key string) *ecdsa.PrivateKey {
	block, _ := pem.Decode([]byte(key))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	return privateKey
}

func decodePublicKey(key string) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode([]byte(key))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return publicKey
}

func SignMessage(pkey string, message string) string {
	// convert string to key
	privateKey := decodePrivateKey(pkey)

	// hist data

	signhash := sha256.New().Sum([]byte(message))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, signhash)
	//signature := r.Bytes()
	//signature = append(signature, s.Bytes()...)
	if err != nil {
		print(err.Error())
		return ""
	}
	newSign := r.String() + "," + s.String()
	newSign = base64.StdEncoding.EncodeToString([]byte(newSign))
	// hex := base64.StdEncoding.EncodeToString(signature)
	//print(r.String())
	return newSign
}

func VerifyMessage(pKey string, signHashString string, message string) bool {
	//get puvlic key
	publicKey := decodePublicKey(pKey)
	data, err := base64.StdEncoding.DecodeString(signHashString)
	if err != nil {
		print(err.Error())
	}
	signString := string(data[:])
	signStringSlice := strings.Split(signString, ",")

	print()
	hash := sha256.New().Sum([]byte(message))
	r := new(big.Int)
	r.SetString(signStringSlice[0], 10)
	s := new(big.Int)
	s.SetString(signStringSlice[1], 10)
	verified := ecdsa.Verify(publicKey, hash, r, s)

	if verified {
		print("YEsss ! it is ok")
		return true
	} else {
		print("No not verified")
		return false
	}
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
	// vSenChain := VerifyChain(data.Sender)
	// if !vSenChain {
	// 	err = errors.New("Error verifying sender")
	// 	return data, err
	// }
	//verify reciever chain()
	// vRecChain := VerifyChain(data.Reciever)
	// if !vRecChain {
	// 	err = errors.New("Error verifying reciever")
	// 	return data, err
	// }

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
