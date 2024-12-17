package types

// CollectionsResponse represents the response from the API when fetching collections.
type CollectionsResponse struct {
	Collections  []Collection `json:"collections"`         // Array of collections
	Page         int          `json:"page"`                // Current page number
	PerPage      int          `json:"per_page"`            // Number of collections per page
	TotalResults int          `json:"total_results"`       // Total number of collections
	PrevPage     string       `json:"prev_page,omitempty"` // Optional: URL for the previous page
	NextPage     string       `json:"next_page,omitempty"` // Optional: URL for the next page
}

// Collection represents a collection of photos or videos in Pexels.
type Collection struct {
	ID          string `json:"id"`                    // Collection ID
	Title       string `json:"title"`                 // Collection title
	Description string `json:"description,omitempty"` // Optional: Collection description
	Private     bool   `json:"private"`               // Indicates if the collection is private
	MediaCount  int    `json:"media_count"`           // Total number of media items
	PhotosCount int    `json:"photos_count"`          // Number of photos in the collection
	VideosCount int    `json:"videos_count"`          // Number of videos in the collection
}

// MediaResponse defines the structure of the response when fetching media from a collection.
type MediaResponse struct {
	ID           string      `json:"id"`                  // Collection ID
	Media        []MediaItem `json:"media"`               // Array of media items (photos or videos)
	Page         int         `json:"page"`                // Current page number
	PerPage      int         `json:"per_page"`            // Results per page
	TotalResults int         `json:"total_results"`       // Total number of media items in the collection
	PrevPage     string      `json:"prev_page,omitempty"` // Optional: URL for the previous page
	NextPage     string      `json:"next_page,omitempty"` // Optional: URL for the next page
}

// MediaItem can be a photo or a video, representing media in the response.
type MediaItem struct {
	Type            string         `json:"type"`                       // "Photo" or "Video"
	ID              int            `json:"id"`                         // Media ID
	Width           int            `json:"width"`                      // Width of the media
	Height          int            `json:"height"`                     // Height of the media
	URL             string         `json:"url"`                        // URL to the media
	Photographer    string         `json:"photographer,omitempty"`     // Photographer's name (for photos)
	PhotographerURL string         `json:"photographer_url,omitempty"` // Photographer's URL
	PhotographerID  int            `json:"photographer_id,omitempty"`  // Photographer's ID
	AvgColor        string         `json:"avg_color,omitempty"`        // Average color of the photo
	Src             *PhotoSrc      `json:"src,omitempty"`              // Photo sources (only for "Photo" type)
	VideoFiles      []VideoFile    `json:"video_files,omitempty"`      // List of video files (only for "Video" type)
	VideoPictures   []VideoPicture `json:"video_pictures,omitempty"`   // Video preview pictures (for "Video")
	User            *User          `json:"user,omitempty"`             // User information (for videos)
	Liked           bool           `json:"liked,omitempty"`            // If the photo is liked
	Duration        int            `json:"duration,omitempty"`         // Duration of the video
}
