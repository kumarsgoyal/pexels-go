package types

// PhotosResponse represents the response for the photo search endpoint.
type PhotosResponse struct {
	TotalResults int     `json:"total_results"` // Total number of results
	Page         int     `json:"page"`          // Current page number
	PerPage      int     `json:"per_page"`      // Number of results per page
	Photos       []Photo `json:"photos"`        // Array of photos
	NextPage     string  `json:"next_page"`     // URL to the next page of results
}

// Photo represents an individual photo in the response.
type Photo struct {
	ID              int      `json:"id"`               // Photo ID
	Width           int      `json:"width"`            // Photo width
	Height          int      `json:"height"`           // Photo height
	URL             string   `json:"url"`              // URL to the photo page
	Photographer    string   `json:"photographer"`     // Photographer's name
	PhotographerID  int      `json:"photographer_id"`  // Photographer's ID
	PhotographerURL string   `json:"photographer_url"` // Photographer's profile URL
	AvgColor        string   `json:"avg_color"`        // Average color of the photo
	Src             PhotoSrc `json:"src"`              // Source URLs for the photo in different sizes
	Liked           bool     `json:"liked"`            // Whether the photo is liked
	Alt             string   `json:"alt"`              // Alternative text for the photo
}

// PhotoSrc represents the source URLs for a photo in different resolutions and orientations.
type PhotoSrc struct {
	Original  string `json:"original"`  // Original resolution URL
	Large2x   string `json:"large2x"`   // Larger resolution URL
	Large     string `json:"large"`     // Large resolution URL
	Medium    string `json:"medium"`    // Medium resolution URL
	Small     string `json:"small"`     // Small resolution URL
	Portrait  string `json:"portrait"`  // Portrait orientation URL
	Landscape string `json:"landscape"` // Landscape orientation URL
	Tiny      string `json:"tiny"`      // Tiny size URL
}
