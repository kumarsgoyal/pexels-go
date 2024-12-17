package endpoints

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kumarsgoyal/pexels-go/client/fetchwrapper"
	"github.com/kumarsgoyal/pexels-go/types"
	"github.com/kumarsgoyal/pexels-go/utils"
)

const (
	SearchVideoEndpoint  = "search"
	PopularVideoEndpoint = "popular"
	VideoEndpoint        = "videos"
)

// VideoEndpoints handles all video-related API calls.
type VideoEndpoints struct {
	FetchWrapper *fetchwrapper.FetchWrapper
}

// NewVideoEndpoints initializes a new instance of VideoEndpoints.
func NewVideoEndpoints(fetchWrapper *fetchwrapper.FetchWrapper) VideoEndpoints {
	return VideoEndpoints{FetchWrapper: fetchWrapper}
}

// Helper function for error logging
func (ve *VideoEndpoints) handleError(action string, err error) error {
	log.Printf("Error %s: %v", action, err)
	return fmt.Errorf("error %s: %w", action, err)
}

// Helper function to clean up empty or zero-value parameters
func (ve *VideoEndpoints) prepareCleanedParams(params map[string]interface{}) map[string]interface{} {
	return utils.CleanParams(params)
}

// Helper function to unmarshal responses
func (ve *VideoEndpoints) unmarshalResponse(body []byte, target interface{}) error {
	return json.Unmarshal(body, target)
}

// Search searches for videos based on a query and optional filters.
func (ve *VideoEndpoints) Search(params *types.VideoSearchParams) (*types.VideosResponse, error) {
	if params == nil {
		params = &types.VideoSearchParams{}
	}

	// Prepare query parameters
	paramsMap := map[string]interface{}{
		"query":       params.Query,
		"orientation": params.Orientation,
		"size":        params.Size,
		"locale":      params.Locale,
		"page":        params.Page,
		"per_page":    params.PerPage,
	}

	cleanedParams := ve.prepareCleanedParams(paramsMap)

	// Fetch data
	body, err := ve.FetchWrapper.Fetch(SearchVideoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching video search results: %v", err)
		return nil, fmt.Errorf("error fetching video search results: %w", err)
	}

	// Unmarshal response
	var response types.VideosResponse
	if err := ve.unmarshalResponse(body, &response); err != nil {
		return nil, ve.handleError("unmarshaling video search response", err)
	}

	log.Printf("Successfully fetched video search results for query: %s", params.Query)
	return &response, nil
}

// Popular fetches a list of popular videos based on filters.
func (ve *VideoEndpoints) Popular(params *types.VideoFilterParams) (*types.VideosResponse, error) {
	if params == nil {
		params = &types.VideoFilterParams{}
	}

	// Prepare query parameters
	paramsMap := map[string]interface{}{
		"min_width":    params.MinWidth,
		"min_height":   params.MinHeight,
		"min_duration": params.MinDuration,
		"max_duration": params.MaxDuration,
		"page":         params.Page,
		"per_page":     params.PerPage,
	}

	cleanedParams := ve.prepareCleanedParams(paramsMap)

	log.Println("Fetching popular videos with filters...")
	body, err := ve.FetchWrapper.Fetch(PopularVideoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching popular videos: %v", err)
		return nil, fmt.Errorf("error fetching popular videos: %w", err)
	}

	// Unmarshal response
	var response types.VideosResponse
	if err := ve.unmarshalResponse(body, &response); err != nil {
		return nil, ve.handleError("unmarshaling popular videos response", err)
	}

	log.Printf("Successfully retrieved popular videos.")
	return &response, nil
}

// GetVideo fetches details about a specific video by its ID.
func (ve *VideoEndpoints) GetVideo(videoID int) (*types.Video, error) {
	log.Printf("Fetching details for video ID: %d", videoID)

	// Convert videoID to a string and build the endpoint
	endpoint := fmt.Sprintf("%s/%d", VideoEndpoint, videoID)

	// Fetch video details
	body, err := ve.FetchWrapper.Fetch(endpoint, nil)
	if err != nil {
		log.Printf("Error fetching video details for ID %d: %v", videoID, err)
		return nil, fmt.Errorf("error fetching video details for ID %d: %w", videoID, err)
	}

	// Parse the response into a Video struct
	var video types.Video
	if err := ve.unmarshalResponse(body, &video); err != nil {
		log.Printf("Error unmarshaling video details for ID %d: %v", videoID, err)
		return nil, fmt.Errorf("error unmarshaling video details for ID %d: %w", videoID, err)
	}

	log.Printf("Successfully retrieved details for video ID: %d", videoID)
	return &video, nil
}
