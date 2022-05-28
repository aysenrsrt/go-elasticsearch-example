package index

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go-elasticsearch-ex/internal/document"
)

type Index struct {
	es *elasticsearch.Client
}

func New(es *elasticsearch.Client) *Index {
	return &Index{es: es}
}

func (i *Index) makeIndexRequest(d document.Document) esapi.IndexRequest {
	data, _ := json.Marshal(d)
	return esapi.IndexRequest{
		Index:      d.Index(),
		DocumentID: d.ID(),
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
}

func (i *Index) makeSearchRequest(match map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"match": match,
		},
	}
}
