package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type orderRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type orderResponse struct {
	Change int    `json:"change"`
	Thank  string `json:"thank"`
}

func main() {
	candyType := flag.String("k", "", "type of candy")
	candyCount := flag.Int("c", 0, "amount of candy")
	money := flag.Int("m", 0, "money for candy")
	flag.Parse()

	if *candyType == "" || *candyCount <= 0 || *money <= 0 {
		fmt.Println("Invalid input. Use flags -k (candy type), -c (candy count), and -m (money amount).")
		os.Exit(1)
	}

	order := orderRequest{
		Money:      *money,
		CandyType:  *candyType,
		CandyCount: *candyCount,
	}
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatal("Error marshalling order data:", err)
	}

	caCert, err := os.ReadFile("minica.pem")
	if err != nil {
		log.Fatal("Error loading CA certificate:", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Post("https://localhost:3333/buy_candy", "application/json", bytes.NewBuffer(orderJSON))
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	if resp.StatusCode == http.StatusCreated {
		var response orderResponse
		if err := json.Unmarshal(body, &response); err != nil {
			log.Fatal("Error decoding response:", err)
		}
		fmt.Printf("Thank you! Your change is %d\n", response.Change)
	} else {
		fmt.Printf("Error: %s\n", string(body))
	}
}
