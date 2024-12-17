package types

// ErrorResponse represents the error response returned by the Pexels API.
type ErrorResponse struct {
	Error string `json:"error"` // Error message from Pexels API
}
