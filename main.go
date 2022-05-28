package main

import (
	"go-elasticsearch-ex/internal"
	"go-elasticsearch-ex/internal/document"
	"go-elasticsearch-ex/internal/domain/movie"
	"go-elasticsearch-ex/internal/index"
	"log"
	"time"
)

func main() {
	es, err := internal.NewElasticSearchClient()
	if err != nil {
		panic(err)
	}

	doc := document.New(movie.CreateRequest{
		Name:        "Harry Potter ve Ölüm Yadigarları",
		IMDB:        10.0,
		Actors:      []string{"Emma Watson", "Daniel Rascliffle", "Rupert Grint", "Tom Felton"},
		Author:      "J.K. Rowling",
		ReleaseDate: time.Date(2010, 11, 17, 0, 0, 0, 0, time.UTC),
	})

	m := index.NewMovie(es)
	_, err = m.Index(doc)
	if err != nil {
		panic(err)
	}

	result, err := m.Search(doc)
	if err != nil {
		panic(err)
	}

	log.Printf(
		"%d hits; took: %dms",
		int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(result["took"].(float64)),
	)

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
}
