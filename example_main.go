package main

import (
	"log"

	"github.com/kumarsgoyal/pexels-go/client"
	"github.com/kumarsgoyal/pexels-go/config"
	"github.com/kumarsgoyal/pexels-go/types"
	"github.com/kumarsgoyal/pexels-go/utils"
)

const (
	APICONFIG = ".apiConfig"
)

func main() {
	// Configure logging
	utils.ConfigureLogging()

	log.Println("Starting the Pexels client application...")

	// Load the API key from the .apiConfig file
	log.Println("Loading configuration from the .apiconfig file...")

	cfg, err := config.LoadConfig(APICONFIG)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	log.Println("Configuration loaded successfully.")

	// Initialize the Pexels client with the loaded API key
	pexelsClient := client.NewClient(cfg.PexelAPIKey)

	searchParams := &types.PhotoSearchParams{
		Query:       "elephant",
		Orientation: "landscape",
		Size:        "large",
		Page:        1,
		PerPage:     5,
	}

	// Example of using the client to search for photos
	photosResponse, err := pexelsClient.Photos.Search(searchParams)
	if err != nil {
		log.Fatalf("Error searching photos: %v", err)
	}

	if len(photosResponse.Photos) == 0 {
		log.Fatal("No photos found")
	}

	for _, photo := range photosResponse.Photos {
		log.Printf("Photo ID: %d, Photographer: %s, URL: %s", photo.ID, photo.Photographer, photo.URL)
	}
	log.Println("Pexels client application initialized successfully.")
}
