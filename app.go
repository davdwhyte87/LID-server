package main

import (
	"log"
	"net"
	"strings"

	// "net/http"
	"os"

	// "kura_coin/blockchain"
	// "kura_coin/controllers"
	// "kura_coin/utils"
	// "github.com/gorilla/mux"
	"kura_coin/blockchain"
	"kura_coin/handlers"

	//"kura_coin/models"
	"kura_coin/utils"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func handleClient(conn net.Conn) {
	utils.Logger.Debug().Msg("New request")
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	// Read data from the client
	n, err := conn.Read(buffer)
	if err != nil {
		utils.Logger.Error().Str("err", err.Error()).Msg("Error reading net data")
		println("Error:", err.Error())
		return
	}
	message := string(buffer[:n])
	// Process and use the data (here, we'll just print it)

	// break message down into signature and message

	// req := models.Request{}
	// err = json.Unmarshal([]byte(message), &req)
	// if err != nil {
	// 	println(err.Error())
	// 	//json: Unmarshal(non-pointer main.Request)
	// }

	datapit := strings.Split(message, "\n")

	// utils.Logger.Debug().Str("length",string(rune(len(datapit))))
	// if len(datapit) != 3 {
	// 	response := models.ErrorResponse{}
	// 	response.Code = 0
	// 	response.Message = "Invalid request structure for protocol"

	// 	utils.RespondTCP(response, conn)
	// 	return
	// }
	routeTCPActions(datapit[0], datapit[len(datapit)-1], conn)
	// messageSignature := datapit[1]
	// senderPublicKey := datapit[2]
	// if !blockchain.VerifyMessage(senderPublicKey, messageSignature, datapit[len(datapit)-1]) {
	// 	println("Error with verifyin message ")
	// }
}

func routeTCPActions(action string, message string, conn net.Conn) {

	switch action {
	case "CreateWallet":
		utils.Logger.Debug().Msg("creating wallet req")
		handlers.CreateWallet(message, conn)
	case "GetBalance":
		utils.Logger.Debug().Msg("getting balance req")
		handlers.GetBalance(message, conn)
	}
}

func main() {
	//utils.Pay()
	// initialize logger
	utils.InitLogger()
	prk1, pk1, _ := blockchain.CreateWallet("oasis34", "onome do that tin")
	prk2, pk2, _ := blockchain.CreateWallet("oasis34", "onome do that tin")

	println(prk1, pk1)
	println(prk2, pk2)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		println("Error:", err)
		return
	}
	defer listener.Close()
	println("Server is listening on port " + port)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
	// r := mux.NewRouter()

	// v1 := r.PathPrefix("/api/v1").Subrouter()
	// walletRouter := v1.PathPrefix("/wallet").Subrouter()

	// walletRouter.HandleFunc("/create", controllers.CreateWallet).Methods("POST")
	// walletRouter.HandleFunc("/print", controllers.PrintBlocks).Methods("POST")
	// walletRouter.HandleFunc("/transfer", controllers.TransferLID).Methods("POST")

	// port := os.Getenv("PORT")
	// print("Server Running on the Port: " + port)
	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }
}
