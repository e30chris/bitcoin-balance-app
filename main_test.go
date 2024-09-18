package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBalanceHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("POST", "/balance", strings.NewReader("address=1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(balanceHandler)

	// Call the handler with our request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `<h1>Balance for Address: 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa</h1>`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
