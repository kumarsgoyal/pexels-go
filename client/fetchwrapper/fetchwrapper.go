package fetchwrapper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// FetchWrapper struct to hold the base URL and API key for making requests
type FetchWrapper struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

// NewFetchWrapper creates and returns a new FetchWrapper instance
func NewFetchWrapper(baseURL, apiKey string) *FetchWrapper {
	return &FetchWrapper{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client:  &http.Client{Timeout: 30 * time.Second},
	}
}

// Helper function to construct a query string from parameters
func (fw *FetchWrapper) constructQueryString(params map[string]interface{}) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, fmt.Sprintf("%v", value))
	}
	return values.Encode()
}

// Helper function to create an HTTP request with the appropriate headers
func (fw *FetchWrapper) createRequest(endpoint, queryString string) (*http.Request, error) {
	fullURL := fmt.Sprintf("%s%s?%s", fw.BaseURL, endpoint, queryString)
	log.Printf("Making request to: %s", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	// Set headers
	req.Header.Add("Authorization", fw.APIKey)
	req.Header.Add("User-Agent", "Pexels-Go/1.0")

	return req, nil
}

// Fetch function performs a GET request and returns the decoded JSON response
func (fw *FetchWrapper) Fetch(endpoint string, params map[string]interface{}) ([]byte, error) {
	// Construct query string from parameters
	queryString := fw.constructQueryString(params)

	// Create request
	req, err := fw.createRequest(endpoint, queryString)
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := fw.Client.Do(req)
	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Log response status
	log.Printf("Received response with status: %s", resp.Status)

	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Received non-OK response: %d", resp.StatusCode)
		return nil, fmt.Errorf("received non-OK response: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err
	}

	return body, nil
}
