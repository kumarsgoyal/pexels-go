package client

import (
	"log"

	"github.com/kumarsgoyal/pexels-go/client/endpoints"
	"github.com/kumarsgoyal/pexels-go/client/fetchwrapper"
)

// Base URLs for different Pexels API endpoints
const (
	PhotoBaseURL      = "https://api.pexels.com/v1/"             // Base URL for photo-related endpoints
	VideoBaseURL      = "https://api.pexels.com/videos/"         // Base URL for video-related endpoints
	CollectionBaseURL = "https://api.pexels.com/v1/collections/" // Base URL for collections
)

// PexelsClient serves as the central client for interacting with the Pexels API.
// It holds references to individual service endpoints for Photos, Videos, and Collections.
type PexelsClient struct {
	Photos      endpoints.PhotoEndpoints      // Photo-related API operations
	Videos      endpoints.VideoEndpoints      // Video-related API operations
	Collections endpoints.CollectionEndpoints // Collection-related API operations
}

// NewClient initializes a new PexelsClient with the given API key.
// It sets up fetch wrappers for each type of service (photos, videos, collections).
func NewClient(apiKey string) *PexelsClient {
	log.Println("Initializing Pexels Client...")

	// Create fetch wrappers with the appropriate base URL and API key
	photoFetchWrapper := createFetchWrapper(PhotoBaseURL, apiKey)
	videoFetchWrapper := createFetchWrapper(VideoBaseURL, apiKey)
	collectionFetchWrapper := createFetchWrapper(CollectionBaseURL, apiKey)

	// Initialize and return the PexelsClient with specific endpoints
	return &PexelsClient{
		Photos:      endpoints.NewPhotoEndpoints(photoFetchWrapper),
		Videos:      endpoints.NewVideoEndpoints(videoFetchWrapper),
		Collections: endpoints.NewCollectionEndpoints(collectionFetchWrapper),
	}
}

// createFetchWrapper is a helper function that constructs a new FetchWrapper
// with the provided base URL and API key for a specific service.
func createFetchWrapper(baseURL, apiKey string) *fetchwrapper.FetchWrapper {
	return fetchwrapper.NewFetchWrapper(baseURL, apiKey)
}
