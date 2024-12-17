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

// VideoEndpoints handles all video-related API calls for the Pexels API.
type VideoEndpoints struct {
	FetchWrapper *fetchwrapper.FetchWrapper // FetchWrapper is used to send HTTP requests to the Pexels API.
}

// NewVideoEndpoints initializes a new instance of VideoEndpoints with the provided FetchWrapper.
// This function allows interaction with the video-related endpoints in the Pexels API.
func NewVideoEndpoints(fetchWrapper *fetchwrapper.FetchWrapper) VideoEndpoints {
	return VideoEndpoints{FetchWrapper: fetchWrapper}
}

// handleError logs the error with a custom message indicating the action that failed.
// It returns a formatted error with additional context.
func (ve *VideoEndpoints) handleError(action string, err error) error {
	log.Printf("Error %s: %v", action, err)
	return fmt.Errorf("error %s: %w", action, err)
}

// prepareCleanedParams removes empty or zero-value parameters from the given map.
// This helps to avoid sending unnecessary parameters in the API request.
func (ve *VideoEndpoints) prepareCleanedParams(params map[string]interface{}) map[string]interface{} {
	return utils.CleanParams(params)
}

// unmarshalResponse unmarshals the given JSON response body into the target structure.
// This method is used to convert the raw JSON data into Go structs.
func (ve *VideoEndpoints) unmarshalResponse(body []byte, target interface{}) error {
	return json.Unmarshal(body, target)
}

// Search searches for videos based on the provided query and optional filters.
// It returns a list of videos matching the search criteria with pagination support.
func (ve *VideoEndpoints) Search(params *types.VideoSearchParams) (*types.VideosResponse, error) {
	if params == nil {
		params = &types.VideoSearchParams{}
	}

	// Prepare query parameters
	paramsMap := map[string]interface{}{
		"query":       params.Query,       // The search query string (e.g., 'nature')
		"orientation": params.Orientation, // Optional filter for video orientation
		"size":        params.Size,        // Optional filter for video size (small, medium, large)
		"locale":      params.Locale,      // Optional filter for video locale
		"page":        params.Page,        // Page number for pagination
		"per_page":    params.PerPage,     // Number of items per page
	}

	cleanedParams := ve.prepareCleanedParams(paramsMap)

	// Fetch video search results
	body, err := ve.FetchWrapper.Fetch(SearchVideoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching video search results: %v", err)
		return nil, fmt.Errorf("error fetching video search results: %w", err)
	}

	// Unmarshal response into VideosResponse struct
	var response types.VideosResponse
	if err := ve.unmarshalResponse(body, &response); err != nil {
		return nil, ve.handleError("unmarshaling video search response", err)
	}

	log.Printf("Successfully fetched video search results for query: %s", params.Query)
	return &response, nil
}

// Popular fetches a list of popular videos based on optional filter parameters.
// This function allows you to retrieve high-quality, trending videos from the Pexels library.
func (ve *VideoEndpoints) Popular(params *types.VideoFilterParams) (*types.VideosResponse, error) {
	if params == nil {
		params = &types.VideoFilterParams{}
	}

	// Prepare query parameters based on filter criteria
	paramsMap := map[string]interface{}{
		"min_width":    params.MinWidth,    // Minimum video width
		"min_height":   params.MinHeight,   // Minimum video height
		"min_duration": params.MinDuration, // Minimum video duration in seconds
		"max_duration": params.MaxDuration, // Maximum video duration in seconds
		"page":         params.Page,        // Page number for pagination
		"per_page":     params.PerPage,     // Number of items per page
	}

	cleanedParams := ve.prepareCleanedParams(paramsMap)

	log.Println("Fetching popular videos with filters...")
	// Fetch popular videos based on filters
	body, err := ve.FetchWrapper.Fetch(PopularVideoEndpoint, cleanedParams)
	if err != nil {
		log.Printf("Error fetching popular videos: %v", err)
		return nil, fmt.Errorf("error fetching popular videos: %w", err)
	}

	// Unmarshal response into VideosResponse struct
	var response types.VideosResponse
	if err := ve.unmarshalResponse(body, &response); err != nil {
		return nil, ve.handleError("unmarshaling popular videos response", err)
	}

	log.Printf("Successfully retrieved popular videos.")
	return &response, nil
}

// GetVideo fetches detailed information about a specific video by its ID.
// It returns metadata such as the video's dimensions, duration, and download links.
func (ve *VideoEndpoints) GetVideo(videoID int) (*types.Video, error) {
	log.Printf("Fetching details for video ID: %d", videoID)

	// Construct endpoint URL for a specific video by its ID
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
