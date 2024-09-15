package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var candyTypes = map[string]int{
	"CE": 5,
	"AA": 15,
	"NT": 10,
	"DE": 7,
	"YR": 23,
}

type orderRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}
type orderResponse struct {
	Change int    `json:"change"`
	Thank  string `json:"candyCount"`
}
type notEnough struct {
	Amount int `json:"amount"`
}

func main() {
	http.HandleFunc("/buy_candy", handleRequest)
	log.Println("Server is running on https://localhost:3333...")
	if err := http.ListenAndServeTLS(":3333", "localhost/cert.pem", "localhost/key.pem", nil); err != nil {
		log.Fatal("Error starting HTTPS server:", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order orderRequest

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	price, ok := candyTypes[order.CandyType]
	if !ok || order.CandyCount < 0 || order.Money < 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	var responseOrder orderResponse
	var needMoney notEnough
	if order.Money >= price*order.CandyCount {
		w.WriteHeader(http.StatusCreated)
		responseOrder.Change = order.Money - price*order.CandyCount
		responseOrder.Thank = "Thank you!"
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		if err := enc.Encode(responseOrder); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
	} else {
		needMoney.Amount = -(order.Money - price*order.CandyCount)
		http.Error(w, fmt.Sprintf("You need %d more money!", needMoney.Amount), http.StatusPaymentRequired)
	}
}
