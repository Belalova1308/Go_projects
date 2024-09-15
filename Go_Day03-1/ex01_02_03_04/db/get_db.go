package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"tidy/types"

	"github.com/elastic/go-elasticsearch/v8"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]types.Place, int, error)
}

type ElasticsearchStore struct {
	Client *elasticsearch.Client
}

func NewElasticsearchStore(client *elasticsearch.Client) *ElasticsearchStore {
	return &ElasticsearchStore{Client: client}
}

func (es *ElasticsearchStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	var buf bytes.Buffer
	query := fmt.Sprintf(`
	{
		"from": %d,
		"size": %d,
		"query": {
			"match_all": {}
		}
	}`, offset, limit)
	buf.WriteString(query)

	res, err := es.Client.Search(
		es.Client.Search.WithContext(context.Background()),
		es.Client.Search.WithIndex("places"),
		es.Client.Search.WithBody(&buf),
	)

	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	var r struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source types.Place `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	places := make([]types.Place, len(r.Hits.Hits))
	for i, hit := range r.Hits.Hits {
		places[i] = hit.Source
	}

	return places, r.Hits.Total.Value, nil
}
