package storage

import (
	"context"
	"fmt"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

// StorageES is a struct that implements the Storage interface by using ElasticSearch as the storage
type StorageES struct {
	esClient *es8.Client
}

func NewStorageES(hosts []string) (Storage, error) {
	esClient, err := es8.NewClient(
		es8.Config{
			Addresses: hosts,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %s", err)
	}

	return &StorageES{
		esClient: esClient,
	}, nil
}

func (s *StorageES) Store(model string, data map[string]interface{}) error {
	req := esapi.IndexRequest{
		Index:      "your-index-name",
		DocumentID: "your-document-id", // You can generate or use a specific ID
		Body:       esutil.NewJSONReader(data),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), s.esClient)
	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.String())
	}

	return nil
}

func (s *StorageES) StoreDB(sql string, data map[string]interface{}) error {
	return nil
}
