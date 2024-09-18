package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "html/template"
    "log"
)

const blockchainAPI = "https://blockchain.info/q/addressbalance/"

type BalanceResponse struct {
    Balance int64 `json:"balance"`
}

func main() {
    http.HandleFunc("/", formHandler)
    http.HandleFunc("/balance", balanceHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Bitcoin Balance Checker</title>
        </head>
        <body>
            <h1>Enter Bitcoin Address</h1>
            <form action="/balance" method="post">
                <label for="address">Bitcoin Address:</label>
                <input type="text" id="address" name="address" required>
                <input type="submit" value="Check Balance">
            </form>
        </body>
        </html>
    `
    fmt.Fprint(w, tmpl)
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    address := r.FormValue("address")
    if address == "" {
        http.Error(w, "Address is required", http.StatusBadRequest)
        return
    }

    resp, err := http.Get(blockchainAPI + address)
    if err != nil {
        http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var balance int64
    if err := json.NewDecoder(resp.Body).Decode(&balance); err != nil {
        http.Error(w, "Failed to parse response", http.StatusInternalServerError)
        return
    }

    tmpl := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Bitcoin Balance Checker</title>
        </head>
        <body>
            <h1>Balance for Address: {{.Address}}</h1>
            <p>Balance: {{.Balance}} Satoshis</p>
            <a href="/">Check another address</a>
        </body>
        </html>
    `
    data := struct {
        Address string
        Balance int64
    }{
        Address: address,
        Balance: balance,
    }

    t, err := template.New("balance").Parse(tmpl)
    if err != nil {
        http.Error(w, "Failed to render template", http.StatusInternalServerError)
        return
    }
    t.Execute(w, data)
}
