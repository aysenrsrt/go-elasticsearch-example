package internal

import (
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

func NewElasticSearchClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Username: "elastic",
		Password: "",
	})
	if err != nil {
		return nil, err
	}

	info, err := es.Info()
	if err != nil {
		return nil, err
	}

	defer info.Body.Close()

	if info.IsError() {
		return nil, errors.New(info.String())
	}

	r := make(map[string]interface{}, 0)
	if err := json.NewDecoder(info.Body).Decode(&r); err != nil {
		return nil, err
	}

	log.Println(strings.Repeat("~", 37))
	log.Printf("Client Version: %s", elasticsearch.Version)
	log.Printf("Server Version: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	return es, nil
}
