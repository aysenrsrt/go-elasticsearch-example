package index

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go-elasticsearch-ex/internal/document"
)

type Movie struct {
	i *Index
}

func NewMovie(es *elasticsearch.Client) *Movie {
	return &Movie{i: New(es)}
}

func (m *Movie) Index(d document.Movie) (map[string]interface{}, error) {
	req := m.i.makeIndexRequest(d)

	res, err := req.Do(context.Background(), m.i.es)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("[%s] Error indexing document ID=%s", res.Status(), d.ID())
	}

	var r map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}

func (m *Movie) Search(d document.Movie) (map[string]interface{}, error) {
	var buf bytes.Buffer
	query := m.i.makeSearchRequest(map[string]interface{}{"name": d.Name})
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := m.i.es.Search(
		m.i.es.Search.WithContext(context.Background()),
		m.i.es.Search.WithIndex(d.Index()),
		m.i.es.Search.WithBody(&buf),
		m.i.es.Search.WithTrackTotalHits(true),
		m.i.es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	r := make(map[string]interface{}, 0)
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}
