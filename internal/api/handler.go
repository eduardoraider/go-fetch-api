package api

import (
	"encoding/json"
	"github.com/eduardoraider/go-fetch-api/internal/external"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the value of TITLE from the environment
	title := os.Getenv("TITLE")

	// Fetch data from the external API
	movies, err := external.FetchMovies(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Marshal the movies into JSON format and send it as the response
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/movies", moviesHandler)
	return mux
}
