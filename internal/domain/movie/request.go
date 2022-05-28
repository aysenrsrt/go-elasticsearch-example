package movie

import "time"

type CreateRequest struct {
	Name        string    `json:"name"`
	IMDB        float32   `json:"imdb"`
	Actors      []string  `json:"actors"`
	Author      string    `json:"author"`
	ReleaseDate time.Time `json:"releaseDate"`
}
