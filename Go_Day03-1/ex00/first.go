package main

import (
	"fmt"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

type Restaurant struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func GetESCClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),

		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}

func main() {
	client
}
