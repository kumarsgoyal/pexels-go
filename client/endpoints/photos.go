package endpoints

import (
	"fmt"
	"log"

	"github.com/kumarsgoyal/pexels-go/client/fetchwrapper"
	"github.com/kumarsgoyal/pexels-go/types"
	"github.com/kumarsgoyal/pexels-go/utils"
)

const (
	SearchPhotoEndpoint  = "search"
	CuratedPhotoEndpoint = "curated"
	PhotoEndpoint        = "photos"
)

// PhotoEndpoints handles all photo-related API calls.
type PhotoEndpoints struct {
	FetchWrapper *fetchwrapper.FetchWrapper
}

// NewPhotoEndpoints initializes a new instance of PhotoEndpoints.
func NewPhotoEndpoints(fetchWrapper *fetchwrapper.FetchWrapper) PhotoEndpoints {
	return PhotoEndpoints{FetchWrapper: fetchWrapper}
}

// Helper function to clean up empty or zero-value parameters
func (pe *PhotoEndpoints) prepareCleanedParams(params map[string]interface{}) map[string]interface{} {
	return utils.CleanParams(params)
}

// Helper function to unmarshal responses
func (pe *PhotoEndpoints) unmarshalResponse(body []byte, target interface{}) error {
	return utils.UnmarshalResponse(body, target)
}

// Helper function for error logging
func (pe *PhotoEndpoints) handleError(action string, err error) error {
	log.Printf("Error %s: %v", action, err)
	return fmt.Errorf("error %s: %w", action, err)
}

// Search retrieves photos based on a search query and optional filters.
func (pe *PhotoEndpoints) Search(params *types.PhotoSearchParams) (*types.PhotosResponse, error) {
	if params == nil {
		params = &types.PhotoSearchParams{}
	}

	// Prepare query parameters
	paramsMap := map[string]interface{}{
		"query":       params.Query,
		"orientation": params.Orientation,
		"size":        params.Size,
		"color":       params.Color,
		"locale":      params.Locale,
		"page":        params.Page,
		"per_page":    params.PerPage,
	}

	cleanedParams := pe.prepareCleanedParams(paramsMap)

	// Fetch data
	body, err := pe.FetchWrapper.Fetch(SearchPhotoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching search results: %v", err)
		return nil, fmt.Errorf("error fetching search results: %w", err)
	}

	// Unmarshal response
	var response types.PhotosResponse
	if err := pe.unmarshalResponse(body, &response); err != nil {
		return nil, pe.handleError("unmarshaling search response", err)
	}
	log.Printf("Successfully fetched search photos for query: %s", params.Query)
	return &response, nil
}

// Curated fetches a curated list of photos based on pagination parameters.
func (pe *PhotoEndpoints) Curated(params *types.PaginationParams) (*types.PhotosResponse, error) {
	log.Println("Fetching curated photos...")

	// If params are nil, initialize with defaults
	if params == nil {
		params = &types.PaginationParams{}
	}

	// Set default pagination values if not provided
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PerPage == 0 {
		params.PerPage = 15
	}

	// Map query parameters
	paramsMap := map[string]interface{}{
		"page":     params.Page,
		"per_page": params.PerPage,
	}

	cleanedParams := pe.prepareCleanedParams(paramsMap)

	// Fetch data
	body, err := pe.FetchWrapper.Fetch(CuratedPhotoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching curated photos: %v", err)
		return nil, fmt.Errorf("error fetching curated photos: %w", err)
	}

	// Unmarshal response
	var response types.PhotosResponse
	if err := pe.unmarshalResponse(body, &response); err != nil {
		return nil, pe.handleError("unmarshaling curated photos response", err)
	}

	log.Printf("Successfully fetched curated photos.")
	return &response, nil
}

// GetPhoto fetches a specific photo by its ID.
func (pe *PhotoEndpoints) GetPhoto(photoID int) (*types.Photo, error) {
	log.Printf("Fetching photo with ID: %d", photoID)

	// Construct endpoint
	endpoint := fmt.Sprintf("%s/%d", PhotoEndpoint, photoID)

	// Fetch data
	body, err := pe.FetchWrapper.Fetch(endpoint, nil)
	if err != nil {
		log.Printf("Error fetching photo with ID %d: %v", photoID, err)
		return nil, fmt.Errorf("error fetching photo with ID %d: %w", photoID, err)
	}

	// Unmarshal response
	var photo types.Photo
	if err := pe.unmarshalResponse(body, &photo); err != nil {
		return nil, pe.handleError(fmt.Sprintf("unmarshaling photo response for ID %d", photoID), err)
	}

	log.Printf("Successfully fetched photo: %+v", photo)
	return &photo, nil
}
