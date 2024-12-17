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

// PhotoEndpoints handles all photo-related API calls for the Pexels API.
type PhotoEndpoints struct {
	FetchWrapper *fetchwrapper.FetchWrapper // FetchWrapper is used to send HTTP requests to the Pexels API.
}

// NewPhotoEndpoints initializes a new instance of PhotoEndpoints with the given FetchWrapper.
// This is used to interact with the photos endpoints in the Pexels API.
func NewPhotoEndpoints(fetchWrapper *fetchwrapper.FetchWrapper) PhotoEndpoints {
	return PhotoEndpoints{FetchWrapper: fetchWrapper}
}

// prepareCleanedParams removes empty or zero-value parameters from the given map.
// It ensures that only valid parameters are sent to the API.
func (pe *PhotoEndpoints) prepareCleanedParams(params map[string]interface{}) map[string]interface{} {
	return utils.CleanParams(params)
}

// unmarshalResponse unmarshals the given JSON response body into the target structure.
// This method is used to convert raw JSON data into Go structs.
func (pe *PhotoEndpoints) unmarshalResponse(body []byte, target interface{}) error {
	return utils.UnmarshalResponse(body, target)
}

// handleError logs and returns a formatted error for a given action.
// This function helps in handling errors consistently within the PhotoEndpoints methods.
func (pe *PhotoEndpoints) handleError(action string, err error) error {
	log.Printf("Error %s: %v", action, err)
	return fmt.Errorf("error %s: %w", action, err)
}

// Search retrieves photos based on a search query and optional filters such as orientation, size, color, etc.
// It returns a list of photos matching the search criteria, with pagination support.
func (pe *PhotoEndpoints) Search(params *types.PhotoSearchParams) (*types.PhotosResponse, error) {
	if params == nil {
		params = &types.PhotoSearchParams{}
	}

	// Prepare query parameters
	paramsMap := map[string]interface{}{
		"query":       params.Query,       // Search query string
		"orientation": params.Orientation, // Optional filter for photo orientation
		"size":        params.Size,        // Optional filter for photo size (small, medium, large)
		"color":       params.Color,       // Optional filter for photo color
		"locale":      params.Locale,      // Optional filter for photo locale
		"page":        params.Page,        // Page number for pagination
		"per_page":    params.PerPage,     // Number of items per page
	}

	cleanedParams := pe.prepareCleanedParams(paramsMap)

	// Fetch search results
	body, err := pe.FetchWrapper.Fetch(SearchPhotoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching search results: %v", err)
		return nil, fmt.Errorf("error fetching search results: %w", err)
	}

	// Unmarshal response into PhotosResponse struct
	var response types.PhotosResponse
	if err := pe.unmarshalResponse(body, &response); err != nil {
		return nil, pe.handleError("unmarshaling search response", err)
	}

	log.Printf("Successfully fetched search photos for query: %s", params.Query)
	return &response, nil
}

// Curated fetches a curated list of photos based on pagination parameters.
// Curated photos are hand-picked by the Pexels team and include high-quality photos for various themes.
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

	// Map query parameters for pagination
	paramsMap := map[string]interface{}{
		"page":     params.Page,
		"per_page": params.PerPage,
	}

	cleanedParams := pe.prepareCleanedParams(paramsMap)

	// Fetch curated photos with pagination
	body, err := pe.FetchWrapper.Fetch(CuratedPhotoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching curated photos: %v", err)
		return nil, fmt.Errorf("error fetching curated photos: %w", err)
	}

	// Unmarshal the response into PhotosResponse struct
	var response types.PhotosResponse
	if err := pe.unmarshalResponse(body, &response); err != nil {
		return nil, pe.handleError("unmarshaling curated photos response", err)
	}

	log.Printf("Successfully fetched curated photos.")
	return &response, nil
}

// GetPhoto fetches a specific photo by its ID. This function returns detailed information about a photo.
// It includes metadata such as the photographer's name, photo dimensions, and download links.
func (pe *PhotoEndpoints) GetPhoto(photoID int) (*types.Photo, error) {
	log.Printf("Fetching photo with ID: %d", photoID)

	// Construct the endpoint URL to fetch the specific photo by ID
	endpoint := fmt.Sprintf("%s/%d", PhotoEndpoint, photoID)

	// Fetch the photo data
	body, err := pe.FetchWrapper.Fetch(endpoint, nil)
	if err != nil {
		log.Printf("Error fetching photo with ID %d: %v", photoID, err)
		return nil, fmt.Errorf("error fetching photo with ID %d: %w", photoID, err)
	}

	// Unmarshal response into Photo struct
	var photo types.Photo
	if err := pe.unmarshalResponse(body, &photo); err != nil {
		return nil, pe.handleError(fmt.Sprintf("unmarshaling photo response for ID %d", photoID), err)
	}

	log.Printf("Successfully fetched photo: %+v", photo)
	return &photo, nil
}
