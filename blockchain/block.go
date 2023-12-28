package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	// "kura_coin/utils"
	"kura_coin/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type Block struct {
	Amount         int
	ID             string
	Hash           string
	Sender         string
	Reciever       string
	PrevHash       string
	PrivateKeyHash string
}

var walletDataPath = "./data/"

func saveBlock(address string, block Block) error {
	var err error
	walletPath := walletDataPath + address
	_, errstat := os.Stat(walletPath)
	err = errstat
	utils.Logger.Debug().Msg(walletPath)
	if os.IsNotExist(errstat) {
		errDir := os.MkdirAll(walletPath, 0755)
		err = errDir
		if errDir != nil {
			utils.Logger.Error().Msg("Error creating wallet dir")
			log.Print("Error creating temp dir")
			return err
		}

		// add wallet data if its a new wallet
		err = addWalletData(address)
		if err  != nil {
			utils.Logger.Error().Str("err",err.Error()).Msg("error adding wallet data")
			return err
		}
		
		var chain []Block
		var dataBuffer bytes.Buffer
		enc := gob.NewEncoder(&dataBuffer)
		chain = append(chain, block)
		errenc := enc.Encode(chain)
		err = errenc
		if errenc != nil {
			utils.Logger.Error().Str("err",errenc.Error()).Msg("error encoding chain data")
			return err
		}
		errWr := ioutil.WriteFile(walletPath+"/chain.bin", dataBuffer.Bytes(), 0700)
		err = errWr
		if errWr != nil {
			utils.Logger.Error().Str("err",errWr.Error()).Msg("error writing chain data")
			return err
		}
	} else {
		var chain []Block

		// get current chain
		chain = retrieveChain(address)
		var dataBuffer bytes.Buffer
		enc := gob.NewEncoder(&dataBuffer)
		chain = append(chain, block)
		errenc := enc.Encode(chain)
		err = errenc
		if errenc != nil {
			return err
		}
		errWr := ioutil.WriteFile(walletPath+"/chain.bin", dataBuffer.Bytes(), 0700)
		err = errWr
		if errWr != nil {
			utils.Logger.Error().Str("err",errWr.Error()).Msg("error writing chain data")
			return err
		}
	}

	// err = errors.New("HAHAHA")
	return err
}

func addWalletData(address string) error {
	var err error
	walletPath := walletDataPath + address
	_, errstat := os.Stat(walletPath)
	err = errstat
	print(walletPath)
	if os.IsNotExist(errstat) {
		errDir := os.MkdirAll(walletPath, 0755)
		err = errDir
		if errDir != nil {
			utils.Logger.Error().Msg("error creating wallet.bin")
			log.Print("Error creating temp dir")
			return err
		}

		
		var dataBuffer bytes.Buffer
		enc := gob.NewEncoder(&dataBuffer)
		wallet := Wallet{}
		wallet.Address = address
		wallet.CreatedAt = time.Now().String()
		errenc := enc.Encode(wallet)
		err = errenc
		if errenc != nil {
			utils.Logger.Error().Str("err",errenc.Error()).Msg("error encoding wallet data")
			return err
		}
		errWr := ioutil.WriteFile(walletPath+"/wallet.bin", dataBuffer.Bytes(), 0700)
		err = errWr
		if errWr != nil {
			utils.Logger.Error().Str("err",errWr.Error()).Msg("error writing to file for wallet")
			return err
		}
	} else {
	
	}
	return err
}


func GetChain(address string) []Block{
	return retrieveChain(address)
}

func retrieveChain(address string) []Block {

	db := NewKvStore[[]Block](address, "chain.bin")
	chain, err :=db.GetData()
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("error getting data")
		return nil
	}
	return chain
}

// PrintChain ...
func PrintChain(address string) []Block {
	chain := retrieveChain(address)
	for x := range chain {
		block := chain[x]
		log.Print("Amount")
		log.Print(int(block.Amount))
		log.Print("Reciever")
		log.Print(block.Reciever)
		log.Print("Sender")
		log.Print(block.Sender)
		log.Print("HAsh")
		log.Print(block.Hash)
		log.Print("PrevHash")
		log.Print(block.PrevHash)
	}

	return chain
}

// func main() {
// 	// saveBlock("damolaregunn")
// 	// retrieveChain("damolaregunn")
// 	// CreateWallet("1234567")
// 	print(verifyChain("NEsEargUrYdn", "1234567"))
// 	// printChain("dgonrLeddidn")
// 	// GetBalance("gVOsLGOAHEOA")
// 	// PrintBlocks("dgonrLeddidn")
// 	// getLastBlock("AnYedOAyFidA")
// 	// Transfer("NEsEargUrYdn", "dgonrLeddidn", 2, "1234567", "1234567" )
// }

// PrintBlocks ...
func PrintBlocks(address string) {
	// get wallet by adress
	walletPath := datapath + address

	// walletPathReciever := dataPath + reciever
	db, err := leveldb.OpenFile(walletPath, nil)
	if err != nil {
		log.Print("db" + err.Error())
		return
	}

	iter := db.NewIterator(nil, nil)

	log.Print("Showing blocks for address  " + address)
	for iter.Next() {
		var block Block
		print("inhere")
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		// _ := iter.Key()
		value := iter.Value()
		dec := gob.NewDecoder(bytes.NewBuffer(value))

		errdec := dec.Decode(&block)
		check(errdec)
		log.Print("Amount")
		log.Print(int(block.Amount))
		log.Print("Reciever")
		log.Print(block.Reciever)
		log.Print("Sender")
		log.Print(block.Sender)
		log.Print("HAsh")
		log.Print(block.Hash)
		log.Print("PrevHash")
		log.Print(block.PrevHash)

	}

	iter.Release()
	err = iter.Error()
	db.Close()
}

func check(e error) {
	if e != nil {
		log.Print(e.Error())
		return
	}
}

func getLastBlock(address string) Block {
	chain := retrieveChain(address)
	length := len(chain)
	lastBlock := chain[length-1]
	return lastBlock
}

// func VerifyChain() {

// }

// VerifyChain ... Verify a chain with hash recreation
func VerifyChain(address string) bool {
	// this function verifies the last block in the chain by checking the previous one

	chain := retrieveChain(address)

	// if the chain has one block return true
	if len(chain) == 1 {
		block := chain[0]
		stringForHash := "00000000000" + strconv.Itoa(10) + address
		shaEngine := sha256.New()
		shaEngine.Write([]byte(stringForHash))

		hashCalc := hex.EncodeToString(shaEngine.Sum(nil))
		if hashCalc == block.Hash {
			return true
		} else {
			return false
		}
	}

	// get last block
	length := len(chain)
	lastBlock := chain[length-1]

	// get previous block (2nd to last)
	prevBlock := chain[length-2]

	// calculate a hypothetic hash of the last prock from prev block data
	stringForHashRec := prevBlock.Hash + strconv.Itoa(lastBlock.Amount) + address
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHashRec))

	hashCalc := hex.EncodeToString(shaEngine.Sum(nil))

	if hashCalc == lastBlock.Hash {
		return true
	}

	return false
}

// VerifyPrivateKey ... Checks if the hash of the supplied key is equal to the hash of the
// previous blocks key
func VerifyPrivateKey(address string, privateKey string) bool {
	chain := retrieveChain(address)

	// if len(chain) == 1 {
	// 	return true
	// }

	// get last block
	length := len(chain)
	lastBlock := chain[length-1]

	// hash supplied private key
	stringForHashRec := privateKey
	shaEngine := sha256.New()
	shaEngine.Write([]byte(stringForHashRec))

	hashCalc := hex.EncodeToString(shaEngine.Sum(nil))

	if hashCalc == lastBlock.PrivateKeyHash {
		return true
	}

	return false
}

// func verifyChain2(address ) bool{
// 	// this function verifies the last block in the chain by checking the previous one

// 	chain := retrieveChain(address)

// 	// if the chain has one block return true
// 	if len(chain) == 1 {
// 		block := chain[0]
// 		stringForHash := "00000000000" + strconv.Itoa(10) + address + privateKey
// 		shaEngine := sha256.New()
// 		shaEngine.Write([]byte(stringForHash))

// 		hashCalc := hex.EncodeToString(shaEngine.Sum(nil))
// 		if hashCalc == block.Hash {
// 			return true
// 		} else {
// 			return false
// 		}
// 	}

// 	// get last block
// 	length := len(chain)
// 	lastBlock := chain[length-1]

// 	// get previous block (2nd to last)
// 	prevBlock := chain[length-2]

// 	// calculate a hypothetic hash of the last prock from prev block data
// 	stringForHashRec := prevBlock.Hash + strconv.Itoa(lastBlock.Amount)
// 	shaEngine := sha256.New()
// 	shaEngine.Write([]byte(stringForHashRec))

// 	hashCalc := hex.EncodeToString(shaEngine.Sum(nil))

// 	if hashCalc == lastBlock.Hash {
// 		return true
// 	}

// 	return false
// }

// BlockExists ...
func BlockExists(address string, ID string) bool {
	chain := retrieveChain(address)

	for x := range chain {
		if ID == chain[x].ID {
			return true
		}
	}

	return false
}
