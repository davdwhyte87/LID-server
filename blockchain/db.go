package blockchain

import (
	"bytes"
	"encoding/gob"

	"os"

	//"reflect"

	"kura_coin/utils"
)

// create a database
func (ks *KvStore[k]) CreateDb() error {
	_, errstat := os.Stat(ks.WalletPath)
	if os.IsNotExist(errstat) {
		utils.Logger.Debug().Str("wallet path", ks.WalletPath)
		errDir := os.MkdirAll(ks.WalletPath, 0755)
		if errDir != nil {
			utils.Logger.Error().Str("err", errDir.Error()).Msg("Error creating wallet dir")
			return errDir
		}
	}

	return nil
}

// save data to database
func (ks *KvStore[k]) SaveData(data k) error {
	var err error
	walletPath := ks.WalletPath
	_, errstat := os.Stat(walletPath)
	err = errstat
	utils.Logger.Debug().Str("save wallet apth ", walletPath)
	if os.IsNotExist(errstat) {
		utils.Logger.Error().Msg("Error saving data ")
	}

	var dataBuffer bytes.Buffer
	enc := gob.NewEncoder(&dataBuffer)
	errenc := enc.Encode(data)
	if errenc != nil {
		utils.Logger.Error().Str("err", errenc.Error()).Msg("error encoding data")
		return err
	}
	utils.Logger.Debug().Str("save db path ", ks.DbPath).Msg("")
	errWr := os.WriteFile(ks.DbPath, dataBuffer.Bytes(), 0700)
	err = errWr
	if errWr != nil {
		utils.Logger.Error().Str("err", errWr.Error()).Msg("error writing chain data")
		return err
	}

	// err = errors.New("HAHAHA")
	return err
}

type KvStoreer[k any] interface {
	GetData(string, string, k)
}

type KvStore[k any] struct {
	data       k
	Path       string
	Address    string
	DbPath     string
	WalletPath string
}

func NewKvStore[k any](address string, path string) *KvStore[k] {
	return &KvStore[k]{
		Path:       path,
		Address:    address,
		DbPath:     walletDataPath + address + "/" + path,
		WalletPath: walletDataPath + address,
	}
}

// get data from db
func (ks *KvStore[k]) GetData() (k, error) {
	dbPath := ks.DbPath
	_, err := os.Stat(dbPath)
	var data k
	// print(walletPath)
	if os.IsNotExist(err) {
		//errDir := os.MkdirAll(walletPath, 0755)
		utils.Logger.Error().Str("err", err.Error()).Msg("database does not exist")
		return data, err
	}

	//var chain []Block
	dataByte, errRead := os.ReadFile(dbPath)
	if errRead != nil {
		utils.Logger.Error().Str("err", errRead.Error()).Msg("Error reading database")
		return data, errRead
	}
	utils.Logger.Debug().Msg("wa")
	dec := gob.NewDecoder(bytes.NewBuffer(dataByte))
	//var obj = reflect.ValueOf(&data)

	errdec := dec.Decode(&data)

	if errdec != nil {
		utils.Logger.Error().Str("err", errdec.Error()).Msg("Error decoding database data")
		return data, errdec
	}

	return data, nil
}

// check is the database exists
func (ks *KvStore[k]) DbExist() bool {
	walletPath := ks.WalletPath
	_, errstat := os.Stat(walletPath)
	utils.Logger.Debug().Msg(walletPath)
	if os.IsNotExist(errstat) {
		return false
	}
	return true
}
