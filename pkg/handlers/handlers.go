package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetBalanceHandler handles the HTTP request for retrieving the balance of a Bitcoin address.
func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	// Read the Bitcoin address from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	address := string(body)

	// Make a request to the mempool.space API to fetch the balance information
	resp, err := http.Get(fmt.Sprintf("https://mempool.space/api/address/%s", address))
	if err != nil {
		http.Error(w, "Failed to fetch balance information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Write the response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}