package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/e30chris/bitcoin-balance-app/pkg/handlers"
)

func main() {
	http.HandleFunc("/balance", handlers.GetBalanceHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}