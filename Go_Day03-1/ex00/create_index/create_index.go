package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	index := "places"
	_, err = es.Indices.Delete([]string{index})
	if err != nil {
		log.Fatalf("Error deleting the index: %s", err)
	}
	res, _ := es.Indices.Create(
		index,
		es.Indices.Create.WithBody(strings.NewReader(`{
		"mappings": {
			"properties": {
				"name": {
					"type":  "text"
				},
				"address": {
					"type":  "text"
				},
				"phone": {
					"type":  "text"
				},
				"location": {
					"type": "geo_point"
				}
			}
		}
	}`)),
	)
	defer res.Body.Close()
	log.Println("Index creation response:", res)

	getRes, err := es.Indices.Get([]string{index})
	if err != nil {
		log.Fatalf("Error retrieving index details: %s", err)
	}
	defer getRes.Body.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(getRes.Body); err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	var indexDetails map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &indexDetails); err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	formattedOutput, err := json.MarshalIndent(indexDetails, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %s", err)
	}
	log.Println("Index details:", string(formattedOutput))
}
