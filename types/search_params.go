package types

// PhotoSearchParams represents the search parameters for photo search.
type PhotoSearchParams struct {
	Query       string `json:"query"`                 // Required: Search query term
	Orientation string `json:"orientation,omitempty"` // Optional: Photo orientation ("landscape", "portrait", "square")
	Size        string `json:"size,omitempty"`        // Optional: Photo size ("large", "medium", "small")
	Color       string `json:"color,omitempty"`       // Optional: Filter by color
	Locale      string `json:"locale,omitempty"`      // Optional: Locale for the search ("en-US", "es-ES")
	Page        int    `json:"page,omitempty"`        // Optional: Page number for pagination
	PerPage     int    `json:"per_page,omitempty"`    // Optional: Results per page (default 15, max 80)
}

// VideoSearchParams represents the search parameters for video search.
type VideoSearchParams struct {
	Query       string `json:"query"`       // Required: Search query term
	Orientation string `json:"orientation"` // Optional: Video orientation ("landscape", "portrait")
	Size        string `json:"size"`        // Optional: Video size
	Locale      string `json:"locale"`      // Optional: Locale for search
	Page        int    `json:"page"`        // Optional: Page number for pagination
	PerPage     int    `json:"per_page"`    // Optional: Results per page
}

// VideoFilterParams represents additional filters for video search.
type VideoFilterParams struct {
	MinWidth    int `json:"min_width,omitempty"`    // Minimum video width in pixels
	MinHeight   int `json:"min_height,omitempty"`   // Minimum video height in pixels
	MinDuration int `json:"min_duration,omitempty"` // Minimum video duration in seconds
	MaxDuration int `json:"max_duration,omitempty"` // Maximum video duration in seconds
	Page        int `json:"page,omitempty"`         // Page number for pagination
	PerPage     int `json:"per_page,omitempty"`     // Results per page
}

// PaginationParams represents the pagination parameters for fetching paginated results.
type PaginationParams struct {
	PerPage int `json:"per_page,omitempty"` // Number of results per page
	Page    int `json:"page,omitempty"`     // Page number for pagination
}

// MediaParams encapsulates parameters for fetching collection media
type MediaParams struct {
	CollectionID string           // ID of the collection
	MediaType    string           // Specify media type (photos or videos)
	Sort         string           // Specify sort order (asc or desc)
	Pagination   PaginationParams // Pagination parameters (page and per_page)
}
