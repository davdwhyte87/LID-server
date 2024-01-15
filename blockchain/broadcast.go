package blockchain

import (
	
	"encoding/json"
	"fmt"
	"math/rand"
	
	"os"
	
	"time"

	"kura_coin/models"
	"kura_coin/utils"
	// "kura_coin/models"
)

// This set of code handles broadcasting requests from this server to other servers
// Broadcast handlers take in paramerets and make new requests

// BCreateWallet ... Broadcast a create wallet command throughout the network
func BCreateWallet(data models.CreateWalletRequest) {
	// data models.CreateWallet
	// select 3 random servers

	print("Broadcasting!!!!!")
	servers := GetServers()
	// n := RandomInt(0, len(servers))

	data.IsBroadcasted = true
	// shuffedServers := shuffle(servers)
	x := len(servers)
	for x > 0 {
		x = x - 1
		//n := RandomInt(0, len(servers)-1)
		// mm := `CreateWallet\n{"Address":"moola","PassPhrase":"beach need love"}`
		nt := "\n"
	
		message :=  fmt.Sprintf(`CreateWallet%v{"Address":"%s","PassPhrase":"beach need love","IsBroadcasted":%v}`,nt,  data.Address, true) 
		utils.Logger.Debug().Msg("Sending data to "+servers[x].Address)
		err :=utils.SendTCP(servers[x].Address, message )
		if err != nil {
			utils.Logger.Error().Str("err", err.Error()).Msg("Error sending message "+ servers[x].Address)
			return 
		}
		///http.Post(servers[n].address+"wallet/create", "application/json", bytes.NewBuffer(requestBody))
		//print("currently on: " + servers[x].address + "wallet/create")
	}
}

// BroadCastTransfer ...
func BroadCastTransfer(request models.TransferReq) {

	request.IsBroadcasted = true
	
	servers := GetServers()
	x := len(servers)
	for x > 0 {
		x = x - 1
		//n := RandomInt(0, len(servers)-1)
		// mm := `CreateWallet\n{"Address":"moola","PassPhrase":"beach need love"}`
		nt := "\n"
	
		message :=  fmt.Sprintf(`Transfer%v{"Sender":"%s","Receiver":"%s", "Amount":"%s","IsBroadcasted":%v}`,nt, request.Sender, request.Receiver, request.Amount, request.IsBroadcasted) 
		utils.Logger.Debug().Msg("Sending data to "+servers[x].Address)
		err :=utils.SendTCP(servers[x].Address, message )
		if err != nil {
			utils.Logger.Error().Str("err", err.Error()).Msg("Error sending message "+ servers[x].Address)
			return 
		}
	}
}

type NodeNetData struct{
	Address string  `json:"address"`
	PublicKey string  `json:"publicKey"`
}
// GetServers ...
func GetServers() []NodeNetData {
	fileName := "server_list.json"
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)

	type Data []string

	var data []NodeNetData
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
