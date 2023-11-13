package main

import (
	"log"
	"net"
	// "net/http"
	"os"

	// "github.com/davdwhyte87/LID-server/blockchain"
	// "github.com/davdwhyte87/LID-server/controllers"
	// "github.com/davdwhyte87/LID-server/utils"
	// "github.com/gorilla/mux"
	"github.com/davdwhyte87/LID-server/blockchain"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
   
	// Create a buffer to read data into
	buffer := make([]byte, 1024)
   
	for {
	 // Read data from the client
	 n, err := conn.Read(buffer)
	 if err != nil {
	  println("Error:", err)
	  return
	 }
   
	 // Process and use the data (here, we'll just print it)
	 println("Received: %s\n", string(buffer[:n]))
	}
   }
func main() {
	//utils.Pay()
	prk1, pk1, _ := blockchain.CreateWallet("oasis34", "onome do that tin")
	prk2, pk2, _ := blockchain.CreateWallet("oasis34", "onome do that tin")

	println(prk1, pk1)
	println(prk2, pk2)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
    if err != nil {
        println("Error:", err)
        return
    }
    defer listener.Close()
    println("Server is listening on port "+port)

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
