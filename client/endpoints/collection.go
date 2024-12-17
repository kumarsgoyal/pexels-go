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

// CollectionEndpoints handles all collection-related API calls.
type CollectionEndpoints struct {
	FetchWrapper *fetchwrapper.FetchWrapper
}

// NewCollectionEndpoints initializes a new instance of CollectionEndpoints.
func NewCollectionEndpoints(fetchWrapper *fetchwrapper.FetchWrapper) CollectionEndpoints {
	return CollectionEndpoints{FetchWrapper: fetchWrapper}
}

// Helper function for error logging
func (ce *CollectionEndpoints) handleError(action string, err error) error {
	log.Printf("Error %s: %v", action, err)
	return fmt.Errorf("error %s: %w", action, err)
}

// Helper function to clean up empty or zero-value parameters
func (ce *CollectionEndpoints) prepareCleanedParams(params map[string]interface{}) map[string]interface{} {
	return utils.CleanParams(params)
}

// Helper function to unmarshal responses
func (ce *CollectionEndpoints) unmarshalResponse(body []byte, target interface{}) error {
	return utils.UnmarshalResponse(body, target)
}

// All fetches all collections with pagination.
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

// Featured fetches featured collections with pagination.
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
