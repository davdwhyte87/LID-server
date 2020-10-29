package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


func generateTransID() string {
	
	maxrange := 777777
	minrange := 10000 - 666
	randnum := rand.Intn(maxrange-minrange) + maxrange
	randstring := StringWithCharset(24)
	idname := strconv.Itoa(randnum) + randstring
	return idname
}

// DecodeReq ... This helps decode a json request body into an interface
func DecodeReq(r *http.Request, model interface{}) interface{} {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, model)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if  err != nil {
		fmt.Printf("%+v\n", err.Error())
		return err
	}
	return err
}


// StringWithCharset ...
func StringWithCharset(length int) string {
	charset := "JESUSiskingskingdomMoneyALPHAaandOMEGAsLordOFLordssLOVEECONOMY"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}