 package main

// import (
// 	// "encoding/binary"
// 	"bytes"
// 	"crypto/sha256"

// 	// "io/ioutil"
// 	"math/rand"
// 	"strconv"
// 	"time"

// 	// "encoding"
// 	"encoding/gob"
// 	"encoding/hex"
// 	"encoding/json"
// 	"io"

// 	// "io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/syndtr/goleveldb/leveldb"
// )

// var path = "sup.json"
// var walletsData = "./data/"
// func init() {

// }

// func stringWithCharset(length int, charset string) string {
// 	var seededRand *rand.Rand = rand.New(
// 		rand.NewSource(time.Now().UnixNano()))
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[seededRand.Intn(len(charset))]
// 	}
// 	return string(b)
// }

// func generateTransID() string {
// 	maxrange := 777777
// 	minrange := 10000 - 666
// 	randnum := rand.Intn(maxrange-minrange) + maxrange
// 	randstring := stringWithCharset(24, "JESUSiskingskingdomMoneyALPHAaandOMEGAsLordOFLordssLOVEECONOMY")
// 	idname := strconv.Itoa(randnum) + randstring
// 	return idname
// }

// type Block struct {
// 	Sender   string
// 	Reciever string
// 	Amount   int
// 	PrevHash string
// 	Hash     string
// }

// func createNewWallet() {
// 	// maxrange := 777777
// 	// minrange := 10000 - 666
// 	// randnum := rand.Intn(maxrange-minrange) + maxrange
// 	randstring := stringWithCharset(24, "JESUSiskingskingdomMoneyALPHAaandOMEGAsLordOFLordssLOVEECONOMY")
// 	// walletName := strconv.Itoa(randnum) + randstring
// 	walletName := randstring
// 	print("Wallet name    " + walletName)

// 	//create wallet on disk
// 	walletPath := "./temp/" + walletName
// 	db, err := leveldb.OpenFile(walletPath, nil)
// 	check(err)

// 	// initail block

// 	var dataBuffer bytes.Buffer
// 	enc := gob.NewEncoder(&dataBuffer)
// 	var block Block
// 	block.Amount = 0
// 	block.Sender = "00000000000"
// 	block.Reciever = walletName
// 	block.PrevHash = "00000000000"

// 	// calculate hash
// 	stringForHash := block.PrevHash + strconv.Itoa(block.Amount) + walletName
// 	shaEngine := sha256.New()
// 	shaEngine.Write([]byte(stringForHash))

// 	block.Hash = hex.EncodeToString(shaEngine.Sum(nil))
// 	print(block.Hash)
// 	errenc := enc.Encode(block)
// 	check(errenc)
// 	// save hash in database
// 	err = db.Put([]byte(generateTransID()), dataBuffer.Bytes(), nil)
// 	check(err)
// 	db.Close()
// }

// func getLastBlock() Block {

// 	db, err := leveldb.OpenFile("./temp/HVAViHUOnEdnsLoMPEdSsogP", nil)
// 	if err != nil {
// 		print("  hello  ")
// 		log.Fatal(err.Error())
// 	}
// 	var block Block
// 	iter := db.NewIterator(nil, nil)
// 	print(len(iter.Key()))
// 	for iter.Next() {
// 		print("inhere")
// 		// Remember that the contents of the returned slice should not be modified, and
// 		// only valid until the next call to Next.
// 		// _ := iter.Key()
// 		value := iter.Value()
// 		dec := gob.NewDecoder(bytes.NewBuffer(value))

// 		errdec := dec.Decode(&block)
// 		check(errdec)
// 		log.Print(int(block.Amount))
// 		break
// 	}
// 	iter.Release()
// 	err = iter.Error()
// 	db.Close()
// 	return block
// }

// var dataPath = "./temp/"





// func saveBlock(address string, data byte){
// 	// walletPath := walletsData + address + ".bin"
// 	// _, err := os.Stat(walletPath)

// 	// print(walletPath)
// 	// if os.IsNotExist(err) {
// 	// 	errDir := os.MkdirAll(walletPath, 0755)
// 	// 	if errDir != nil {
// 	// 		log.Print("Error creating temp dir")
// 	// 		return
// 	// 	}
// 	// }
// 	// var chain []Block
// 	// var block Block

// 	// block.Amount = 100
// 	// block.Hash = "kajdkf"
// 	// block
// 	// errWr := ioutil.WriteFile(walletPath, tempDataBuffer.Bytes(), 0700)	
// }





// func addBlock(sender string, reciever string, amount int) {
// 	prevBlock := getLastBlock()

// 	walletPathSender := dataPath + sender
// 	// walletPathReciever := dataPath + reciever
// 	dbSender, err := leveldb.OpenFile(walletPathSender, nil)
// 	if err != nil {
// 		log.Print("db"+err.Error())
// 		return
// 	}
	
// 	// dbReciever, err := leveldb.OpenFile(walletPathReciever, nil)
// 	// if err != nil{
// 	// 	log.Print(err.Error())
// 	// 	return
// 	// }


// 	// take money from the sender

// 	print("  hello  ")
// 	// check if there is enough money
// 	if prevBlock.Amount < amount {
// 		log.Fatal("Not enough funds")
// 		return
// 	}
// 	print("hello")
// 	var dataBuffer bytes.Buffer
// 	enc := gob.NewEncoder(&dataBuffer)
// 	// var tempBlockHolder []Block
// 	var senderBlock Block
// 	senderBlock.Amount = -(amount)
// 	senderBlock.PrevHash = prevBlock.Hash

// 	// calculate hash
// 	stringForHash := senderBlock.PrevHash + strconv.Itoa(senderBlock.Amount) + sender
// 	shaEngine := sha256.New()
// 	shaEngine.Write([]byte(stringForHash))

// 	senderBlock.Hash = hex.EncodeToString(shaEngine.Sum(nil))

// 	// add block to temp block holder set
// 	// tempBlockHolder = append(tempBlockHolder, senderBlock)
// 	// print(block.Hash)
// 	errenc := enc.Encode(senderBlock)
// 	check(errenc)

// 	// save block to temporary block
// 	transID := generateTransID()

// 	err = dbSender.Put([]byte(transID), dataBuffer.Bytes(), nil)
// 	check(err)
// 	dbSender.Close()
// }

// func validateChain(address string) {
// 	walletPathSender := dataPath + address
// 	db, err := leveldb.OpenFile(walletPathSender, nil)
// 	if err != nil{
// 		log.Print(err.Error())
// 		return
// 	}

// 	var block Block
// 	iter := db.NewIterator(nil, nil)
// 	print(len(iter.Key()))

// 	count := 0
// 	var prevBlock Block
// 	for iter.Next() {
// 		print("inhere")
// 		// Remember that the contents of the returned slice should not be modified, and
// 		// only valid until the next call to Next.
// 		// _ := iter.Key()
// 		value := iter.Value()

// 		if count == 0{
// 			// first block
// 			// only get the block 

// 			// get prev hash
// 			// dec := gob.NewDecoder(bytes.NewBuffer(value))
// 			// errdec := dec.Decode(&prevBlock)
// 			// check(errdec)

// 		} else {
// 			// calculate hash, compare with previous
			
// 			// get current block

// 			if iter.Last() {




// 				dec := gob.NewDecoder(bytes.NewBuffer(value))
// 				errdec := dec.Decode(&block)
// 				check(errdec)

// 				iter.Prev()


// 			}
			


// 			// calculate the hash of the previous block
// 			stringForHash := prevBlock.PrevHash + strconv.Itoa(prevBlock.Amount) + prevBlock.Sender
// 			shaEngine := sha256.New()
// 			shaEngine.Write([]byte(stringForHash))
		
// 			hashOfPrevBlcok := hex.EncodeToString(shaEngine.Sum(nil))

// 			// get prev hash of current block and compare

// 			if block.PrevHash != hashOfPrevBlcok{
// 				print("  Invalid  ")
// 				break

// 			}

// 		}




	
// 		log.Print(int(block.Amount))
// 		break
// 	}
// 	iter.Release()
// 	err = iter.Error()
// 	db.Close()
// }

// // file, err := os.Create(walletPathSender+"/tempBlock/"+ transID)
// // if err != nil {
// // 	log.Print(err.Error())
// // }
// // defer file.Close()

// // _, err := os.Stat(walletPathSender + "/tempBlock")
// // print(walletPathSender + "/tempBlock")
// // if os.IsNotExist(err) {
// // 	errDir := os.MkdirAll(walletPathSender+"/tempBlock", 0755)
// // 	if errDir != nil {
// // 		log.Print("Error creating temp dir")
// // 		return
// // 	}
// // }
// // errWr := ioutil.WriteFile(walletPathSender+"/tempBlock/"+transID, tempDataBuffer.Bytes(), 0700)
// // io.Copy(file, &dataBuffer)
// func main() {
// 	// createNewWallet()
// 	// getLastBlock()
// 	addBlock("HVAViHUOnEdnsLoMPEdSsogP", "dhckabenk", 1)
// 	// db, err := leveldb.OpenFile("./tmp/lenee", nil)
// 	// check(err)
// 	// // err = db.Put([]byte("name"), []byte("david ojodo"), nil)

// 	// type Data struct {
// 	// 	Name string
// 	// 	Age  int
// 	// }
// 	// var dataBuffer bytes.Buffer
// 	// enc := gob.NewEncoder(&dataBuffer)

// 	// errenc := enc.Encode(Data{Name: "Prince", Age: 6387484})
// 	// check(errenc)

// 	// err = db.Put([]byte("Ppre"), dataBuffer.Bytes(), nil)
// 	// check(err)

// 	// // data, err := db.Get([]byte("nenjii"), nil)
// 	// // dec := gob.NewDecoder(bytes.NewBuffer(data))
// 	// // var resdata Data
// 	// // errdec := dec.Decode(&resdata)
// 	// // check(errdec)
// 	// // log.Print(string(resdata.Name))

// 	// iter := db.NewIterator(nil, nil)
// 	// for iter.Next() {
// 	// 	// Remember that the contents of the returned slice should not be modified, and
// 	// 	// only valid until the next call to Next.
// 	// 	// _ := iter.Key()
// 	// 	value := iter.Value()
// 	// 	dec := gob.NewDecoder(bytes.NewBuffer(value))
// 	// 	var resdata Data
// 	// 	errdec := dec.Decode(&resdata)
// 	// 	check(errdec)
// 	// 	log.Print(string(resdata.Name))

// 	// }
// 	// iter.Release()
// 	// err = iter.Error()

// 	// message := Message{Title: "jnjncjn", Body:"djnskjdnjks"}
// 	// buff := new(bytes.Buffer)
// 	// encoder:= json.NewEncoder(buff)
// 	// encoder.Encode(message)
// 	// print(buff.String())
// 	// // b, _:= json.Marshal(message)
// 	// // log.Print(string(b))
// 	// file, err := os.Create(path)
// 	// if err != nil {
// 	// 	log.Print("Error ooo")
// 	// }
// 	// defer file.Close()

// 	// io.Copy(file, buff)
// 	// print(buff.String())
// }

// func createWallet() {
// 	file, err := os.Create(path)
// 	if err != nil {
// 		log.Print("could not create file")
// 		return
// 	}
// 	log.Print("Wallet created")
// 	// addInitialBlock()
// 	defer file.Close()
// }

// // func addInitialBlock() {

// // 	file, err := os.Open(path)
// // 	if err != nil {
// // 		log.Print("Error ooo")
// // 	}
// // 	defer file.Close()

// // 	block := Block{from: "00000000", amount: "jjjj"}
// // 	err = binary.Write(file, binary.LittleEndian, block)
// // 	if err != nil {
// // 		log.Print(err.Error())
// // 		log.Print("Failed to write block")
// // 		return
// // 	}
// // }

// func check(e error) {
// 	if e != nil {
// 		log.Print(e.Error())
// 		return
// 	}
// }

// type Message struct {
// 	Title  string
// 	Body   string
// 	blocks []Block
// }

// // func addInitialBlock3() {
// // 	db, err := badger.Open(badger.DefaultOptions("./tmp/badger"))
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	defer db.Close()

// // 	// Start a writable transaction.
// // 	txn := db.NewTransaction(true)

// // 	defer txn.Discard()

// // 	data, errdata := txn.Get([]byte("answer"))

// // 	check(errdata)
// // 	// dt := data.ValueCopy()
// // 	// print(data.String())
// // 	// geterr := db.View(func(txn *badger.Txn) error {
// // 	// 	item, err := txn.Get([]byte("answer"))
// // 	// 	check(err)
// // 	// 	data, errdata :=
// // 	// 	check(errdata)
// // 	// })
// // 	// Use the transaction...

// // 	// seterr := txn.Set([]byte("answer"), []byte("42"))
// // 	// if err != nil {
// // 	// 	check(seterr)
// // 	// }

// // 	// // Commit the transaction and check for error.
// // 	// if commerr := txn.Commit(); err != nil {
// // 	// 	check(commerr)
// // 	// }
// // }
// func addInitialBlock2() {

// 	message := Message{Title: "jnjncjn", Body: "djnskjdnjks"}
// 	// message.blocks = []Block{ {prevHash: "djkdnkjd"} }

// 	buff := new(bytes.Buffer)
// 	encoder := json.NewEncoder(buff)
// 	encoder.Encode(message)

// 	// b, _:= json.Marshal(message)
// 	// log.Print(string(b))
// 	file, err := os.Create(path)
// 	if err != nil {
// 		log.Print("Error ooo")
// 	}
// 	defer file.Close()

// 	io.Copy(file, buff)
// 	// print(buff.String())
// }

// func readJSONToken(fileName string) []map[string]interface{} {
// 	file, _ := os.Open(fileName)
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)

// 	filteredData := []map[string]interface{}{}

// 	// Read the array open bracket
// 	decoder.Token()

// 	data := map[string]interface{}{}
// 	for decoder.More() {
// 		decoder.Decode(&data)

// 		filteredData = append(filteredData, data)

// 	}

// 	return filteredData
// }
