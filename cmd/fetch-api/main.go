package main

import (
	"github.com/eduardoraider/go-fetch-api/internal/api"
	"log"
	"net/http"
)

func main() {
	// Start the HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", api.NewHandler()))
}
