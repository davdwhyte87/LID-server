package main

import (
	"log"
	"net/http"
	"os"

	// "github.com/davdwhyte87/LID-server/blockchain"
	"github.com/davdwhyte87/LID-server/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {

	// blockchain.BCreateWallet()
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	walletRouter := v1.PathPrefix("/wallet").Subrouter()

	walletRouter.HandleFunc("/create", controllers.CreateWallet).Methods("POST")
	walletRouter.HandleFunc("/print", controllers.PrintBlocks).Methods("POST")
	
	port := os.Getenv("PORT")
	print("Server Running on Port: " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
