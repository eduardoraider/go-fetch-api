package external

import (
	"encoding/json"
	"github.com/eduardoraider/go-fetch-api/pkg/entity"
	"net/http"
)

type MovieResponse struct {
	Page       int            `json:"page"`
	PerPage    int            `json:"per_page"`
	Total      int            `json:"total"`
	TotalPages int            `json:"total_pages"`
	Data       []entity.Movie `json:"data"`
}

func FetchMovies(title string) (MovieResponse, error) {
	var movieResponse MovieResponse

	res, err := http.Get("https://jsonmock.hackerrank.com/api/movies/search/?Title=" + title)
	if err != nil {
		return MovieResponse{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&movieResponse)
	if err != nil {
		return MovieResponse{}, err
	}

	return movieResponse, nil
}
