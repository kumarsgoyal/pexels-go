package types

// VideosResponse represents the response for the video search endpoint.
type VideosResponse struct {
	Page         int     `json:"page"`                // Page number of the response
	PerPage      int     `json:"per_page"`            // Number of results per page
	TotalResults int     `json:"total_results"`       // Total number of video results
	URL          string  `json:"url"`                 // URL to the search results
	Videos       []Video `json:"videos"`              // Array of video results
	PrevPage     string  `json:"prev_page,omitempty"` // Optional: URL for the previous page
	NextPage     string  `json:"next_page,omitempty"` // Optional: URL for the next page
}

// Video represents a single video resource from the Pexels API.
type Video struct {
	ID            int            `json:"id"`             // Video ID
	Width         int            `json:"width"`          // Video width
	Height        int            `json:"height"`         // Video height
	URL           string         `json:"url"`            // URL to the video page
	Image         string         `json:"image"`          // URL to a still image representing the video
	FullRes       interface{}    `json:"full_res"`       // Full resolution video (can be null or another type)
	Tags          []string       `json:"tags"`           // Tags associated with the video
	Duration      int            `json:"duration"`       // Duration of the video in seconds
	User          User           `json:"user"`           // Information about the videographer
	VideoFiles    []VideoFile    `json:"video_files"`    // List of video file resolutions
	VideoPictures []VideoPicture `json:"video_pictures"` // List of preview images for the video
}

// User represents the videographer who shot the video.
type User struct {
	ID   int    `json:"id"`   // User ID
	Name string `json:"name"` // User's name
	URL  string `json:"url"`  // User's profile URL
}

// VideoFile represents different sized versions of the video.
type VideoFile struct {
	ID       int     `json:"id"`        // File ID
	Quality  string  `json:"quality"`   // Quality (e.g., 'hd' or 'sd')
	FileType string  `json:"file_type"` // File type (e.g., "video/mp4")
	Width    int     `json:"width"`     // Video width
	Height   int     `json:"height"`    // Video height
	FPS      float64 `json:"fps"`       // Frames per second
	Link     string  `json:"link"`      // URL to the video file
}

// VideoPicture represents preview pictures of the video.
type VideoPicture struct {
	ID      int    `json:"id"`      // Image ID
	Picture string `json:"picture"` // URL to the preview image
	NR      int    `json:"nr"`      // Sequence number of the picture
}
