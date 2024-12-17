package client

import (
	"log"

	"github.com/kumarsgoyal/pexels-go/client/endpoints"
	"github.com/kumarsgoyal/pexels-go/client/fetchwrapper"
)

// Base URLs for different endpoints
const (
	PhotoBaseURL      = "https://api.pexels.com/v1/"
	VideoBaseURL      = "https://api.pexels.com/videos/"
	CollectionBaseURL = "https://api.pexels.com/v1/collections/"
)

// PexelsClient holds references to the endpoints for photos, videos, and collections
type PexelsClient struct {
	Photos      endpoints.PhotoEndpoints
	Videos      endpoints.VideoEndpoints
	Collections endpoints.CollectionEndpoints
}

// NewClient creates a new instance of PexelsClient and initializes the endpoints
func NewClient(apiKey string) *PexelsClient {
	log.Println("Initializing Pexels Client...")

	// Initialize fetch wrappers for each service with API key and base URLs
	photoFetchWrapper := createFetchWrapper(PhotoBaseURL, apiKey)
	videoFetchWrapper := createFetchWrapper(VideoBaseURL, apiKey)
	collectionFetchWrapper := createFetchWrapper(CollectionBaseURL, apiKey)

	// Return a new PexelsClient with initialized endpoints
	return &PexelsClient{
		Photos:      endpoints.NewPhotoEndpoints(photoFetchWrapper),
		Videos:      endpoints.NewVideoEndpoints(videoFetchWrapper),
		Collections: endpoints.NewCollectionEndpoints(collectionFetchWrapper),
	}
}

// Helper function to create a new fetch wrapper with provided base URL and API key
func createFetchWrapper(baseURL, apiKey string) *fetchwrapper.FetchWrapper {
	return fetchwrapper.NewFetchWrapper(baseURL, apiKey)
}
