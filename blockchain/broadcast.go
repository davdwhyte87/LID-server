package blockchain

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"kura_coin/models"
	// "kura_coin/models"
)

// This set of code handles broadcasting requests from this server to other servers
// Broadcast handlers take in paramerets and make new requests

// BCreateWallet ... Broadcast a create wallet command throughout the network
func BCreateWallet(data models.CreateWallet) {
	// data models.CreateWallet
	// select 3 random servers

	print("Broadcasting!!!!!")
	servers := GetServers()
	// n := RandomInt(0, len(servers))
	requestBody, err := json.Marshal(map[string]string{
		"Address":    data.Address,
		"PrivateKey": data.PrivateKey,
	})

	if err != nil {
		print(err.Error())
		return
	}

	// shuffedServers := shuffle(servers)
	x := len(servers)
	for x > 0 {
		x = x - 1
		n := RandomInt(0, len(servers)-1)
		http.Post(servers[n]+"wallet/create", "application/json", bytes.NewBuffer(requestBody))
		print("currently on: " + servers[x] + "wallet/create")
	}
}

// BroadCastTransfer ...
func BroadCastTransfer(data models.TransferReq) {

	print("Broadcasting!!!!!")
	servers := GetServers()
	// n := RandomInt(0, len(servers))
	requestBody, err := json.Marshal(map[string]string{
		
		"Amount":                  data.Amount,
		
	})

	if err != nil {
		print(err.Error())
		return
	}

	// shuffedServers := shuffle(servers)
	x := 3
	for x > 0 {
		x = x - 1
		n := RandomInt(0, len(servers)-1)
		_, err := http.Post(servers[n]+"wallet/transfer", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			print(err.Error())
		}
		// print("currently on: " + servers[x] + "wallet/create")
	}
}

// GetServers ...
func GetServers() []string {
	fileName := "server_list.json"
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)

	type Data []string
	var data Data
	decoder.Decode(&data)

	// fmt.Printf("%+v\n", data)
	// for x := range data {
	// 	print("link")
	// 	print(data[x])
	// }
	return data
}

func RandomInt(min int, max int) int {
	rand.Seed(time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := r.Perm(max - min + 1)
	return p[min]
}

func shuffle(vals []string) []string {
	// vals := []int{10, 12, 14, 16, 18, 20}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		//   vals = vals[:n-1]
	}
	return vals
}
