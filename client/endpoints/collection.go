package endpoints

import (
	"fmt"
	"log"

	"github.com/kumarsgoyal/pexels-go/client/fetchwrapper"
	"github.com/kumarsgoyal/pexels-go/types"
	"github.com/kumarsgoyal/pexels-go/utils"
)

const (
	FeaturedCollectionEndpoint = "/featured"
	ParamType                  = "type"
	ParamSort                  = "sort"
	ParamPage                  = "page"
	ParamPerPage               = "per_page"
)

// CollectionEndpoints handles all collection-related API calls for the Pexels API.
type CollectionEndpoints struct {
	FetchWrapper *fetchwrapper.FetchWrapper // FetchWrapper is used to send HTTP requests to the Pexels API.
}

// NewCollectionEndpoints initializes a new instance of CollectionEndpoints with the given FetchWrapper.
// This is used to interact with the collections endpoints in the Pexels API.
func NewCollectionEndpoints(fetchWrapper *fetchwrapper.FetchWrapper) CollectionEndpoints {
	return CollectionEndpoints{FetchWrapper: fetchWrapper}
}

// handleError logs and returns a formatted error for a given action.
// This function helps in handling errors consistently within the CollectionEndpoints methods.
func (ce *CollectionEndpoints) handleError(action string, err error) error {
	log.Printf("Error %s: %v", action, err)
	return fmt.Errorf("error %s: %w", action, err)
}

// prepareCleanedParams removes empty or zero-value parameters from the given map.
// It helps ensure that only valid parameters are sent to the API.
func (ce *CollectionEndpoints) prepareCleanedParams(params map[string]interface{}) map[string]interface{} {
	return utils.CleanParams(params)
}

// unmarshalResponse unmarshals the given JSON response body into the target structure.
// This method is used to convert raw JSON data into Go structs.
func (ce *CollectionEndpoints) unmarshalResponse(body []byte, target interface{}) error {
	return utils.UnmarshalResponse(body, target)
}

// All fetches all collections with pagination support. It uses the provided pagination parameters
// to retrieve the collections in a paginated manner.
func (ce *CollectionEndpoints) All(params types.PaginationParams) (*types.CollectionsResponse, error) {
	log.Printf("Fetching collections with page %d and %d per page...", params.Page, params.PerPage)

	// Prepare query parameters for pagination
	paramsMap := map[string]interface{}{
		"page":     params.Page,
		"per_page": params.PerPage,
	}

	cleanedParams := ce.prepareCleanedParams(paramsMap)

	// Fetch collections using the FetchWrapper
	body, err := ce.FetchWrapper.Fetch("", cleanedParams)
	if err != nil {
		log.Printf("Error fetching collections: %v", err)
		return nil, fmt.Errorf("error fetching collections: %w", err)
	}

	// Unmarshal the response into the CollectionsResponse struct
	var response types.CollectionsResponse
	if err := ce.unmarshalResponse(body, &response); err != nil {
		return nil, ce.handleError("unmarshaling collections response", err)
	}
	log.Printf("Successfully retrieved collections. Total collections: %d", len(response.Collections))
	return &response, nil
}

// Featured fetches featured collections with pagination support.
// It returns a list of featured collections based on the given pagination parameters.
func (ce *CollectionEndpoints) Featured(params types.PaginationParams) (*types.CollectionsResponse, error) {
	log.Println("Fetching featured collections...")

	// Prepare query parameters based on PaginationParams
	paramsMap := map[string]interface{}{
		"page":     params.Page,
		"per_page": params.PerPage,
	}

	cleanedParams := ce.prepareCleanedParams(paramsMap)

	// Fetch featured collections with pagination
	body, err := ce.FetchWrapper.Fetch(FeaturedCollectionEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching featured collections: %v", err)
		return nil, err
	}

	// Unmarshal the response into the CollectionsResponse struct
	var response types.CollectionsResponse
	if err := ce.unmarshalResponse(body, &response); err != nil {
		return nil, ce.handleError("unmarshaling featured collections response", err)
	}
	log.Printf("Successfully retrieved featured collections.")
	return &response, nil
}

// Media fetches media (photos or videos) from a specific collection.
// It allows filtering by media type (photos or videos), sort order, and pagination.
func (ce *CollectionEndpoints) Media(params types.MediaParams) (*types.MediaResponse, error) {
	log.Printf("Fetching media for collection ID: %s", params.CollectionID)

	// Prepare query parameters
	paramsMap := map[string]interface{}{
		"type":     params.MediaType,          // Specify media type (photos or videos)
		"sort":     params.Sort,               // Specify sort order (asc or desc)
		"page":     params.Pagination.Page,    // Pagination page
		"per_page": params.Pagination.PerPage, // Pagination per_page
	}

	cleanedParams := ce.prepareCleanedParams(paramsMap)

	// Fetch media for the collection with pagination and filters
	body, err := ce.FetchWrapper.Fetch(params.CollectionID, cleanedParams)
	if err != nil {
		log.Printf("Error fetching media for collection ID %s: %v", params.CollectionID, err)
		return nil, fmt.Errorf("error fetching media: %w", err)
	}

	// Define response structure to handle both photos and videos
	var response types.MediaResponse
	if err := ce.unmarshalResponse(body, &response); err != nil {
		return nil, ce.handleError("unmarshaling media response", err)
	}

	log.Printf("Successfully retrieved media for collection ID: %s", params.CollectionID)
	return &response, nil
}
