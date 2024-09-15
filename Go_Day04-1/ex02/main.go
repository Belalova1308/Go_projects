package main

/*
#include "cow.c"
*/
import "C"
import (
	"encoding/json"
	"log"
	"net/http"
	"unsafe"
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
	err := http.ListenAndServeTLS(":3333", "localhost/cert.pem", "localhost/key.pem", nil)
	if err != nil {
		log.Fatal("Error starting HTTPS server:", err)
	}
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var order orderRequest
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	price, ok := candyTypes[order.CandyType]
	if !ok || order.CandyCount <= 0 || order.Money < 0 {
		http.Error(w, "Invalid candy type or count", http.StatusBadRequest)
		return
	}

	var responseOrder orderResponse
	if order.Money >= price*order.CandyCount {
		w.WriteHeader(http.StatusCreated)
		responseOrder.Change = order.Money - price*order.CandyCount

		thankYouMessage := "Thank you!"
		cMessage := C.CString(thankYouMessage)
		defer C.free(unsafe.Pointer(cMessage))

		cowMessage := C.ask_cow(cMessage)
		defer C.free(unsafe.Pointer(cowMessage))

		responseOrder.Thank = C.GoString(cowMessage)

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "   ")
		err := enc.Encode(responseOrder)
		if err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Not enough money!", http.StatusPaymentRequired)
	}
}
