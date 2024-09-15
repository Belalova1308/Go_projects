package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Restaurant struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("Error opening CSV file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var bulkRequest bytes.Buffer
	bulkSize := 16358
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 6)
		if len(parts) < 6 {
			log.Fatalf("Error parsing line, expected 6 fields but got %d: %s", len(parts), line)
		}

		restaurant := Restaurant{
			ID:        parts[0],
			Name:      parts[1],
			Address:   parts[2],
			Phone:     parts[3],
			Longitude: parts[4],
			Latitude:  parts[5],
		}

		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "places", "_id" : "%s" } }%s`, restaurant.ID, "\n"))
		data, err := json.Marshal(restaurant)
		if err != nil {
			log.Fatalf("Error marshaling restaurant to JSON: %s", err)
		}
		bulkRequest.Write(meta)
		bulkRequest.Write(data)
		bulkRequest.WriteString("\n")

		count++
		if count%bulkSize == 0 {
			sendBulkRequest(es, bulkRequest)
			bulkRequest.Reset()
		}
	}
	if bulkRequest.Len() > 0 {
		sendBulkRequest(es, bulkRequest)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}

func sendBulkRequest(es *elasticsearch.Client, bulkRequest bytes.Buffer) {
	res, err := es.Bulk(bytes.NewReader(bulkRequest.Bytes()))
	if err != nil {
		log.Fatalf("Error performing bulk indexing: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error in bulk response: %s", res.String())
	}
	fmt.Println("Bulk indexing request completed")
}
