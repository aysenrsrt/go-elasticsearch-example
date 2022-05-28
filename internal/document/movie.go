package document

import (
	"github.com/google/uuid"
	"go-elasticsearch-ex/internal/domain/movie"
	"time"
)

type Movie struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	IMDB        float32   `json:"imdb"`
	Actors      []string  `json:"actors"`
	Author      string    `json:"author"`
	ReleaseDate time.Time `json:"releaseDate"`
}

func New(r movie.CreateRequest) Movie {
	return Movie{
		Id:          uuid.New().String(),
		Name:        r.Name,
		IMDB:        r.IMDB,
		Actors:      r.Actors,
		Author:      r.Author,
		ReleaseDate: r.ReleaseDate,
	}
}

func (m Movie) ID() string {
	return m.Id
}

func (m Movie) Index() string {
	return "movie"
}
