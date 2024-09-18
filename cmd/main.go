package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic to handle the HTTP request and response
	// You can use the mempool.space API to fetch the balance of a Bitcoin address
	// Example code:
	// bitcoinAddress := r.FormValue("address")
	// balance := GetBitcoinBalance(bitcoinAddress)
	// fmt.Fprintf(w, "Balance of %s: %f BTC", bitcoinAddress, balance)
	fmt.Fprintf(w, "Hello, World!")
}