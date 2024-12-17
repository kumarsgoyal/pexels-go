package fetchwrapper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// FetchWrapper struct holds the base URL, API key, and HTTP client for making requests.
type FetchWrapper struct {
	BaseURL string       // The base URL of the API endpoint
	APIKey  string       // The API key for authenticating requests
	Client  *http.Client // The HTTP client used to make requests
}

// NewFetchWrapper initializes a new FetchWrapper instance with the provided base URL and API key.
// It also sets the default timeout for HTTP requests.
func NewFetchWrapper(baseURL, apiKey string) *FetchWrapper {
	return &FetchWrapper{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client:  &http.Client{Timeout: 30 * time.Second}, // Set a 30-second timeout for all requests
	}
}

// constructQueryString converts a map of parameters into a URL-encoded query string.
// This is used to append parameters to the URL for the GET request.
func (fw *FetchWrapper) constructQueryString(params map[string]interface{}) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, fmt.Sprintf("%v", value)) // Convert value to string and add to URL values
	}
	return values.Encode() // Return the URL-encoded query string
}

// createRequest builds an HTTP GET request for the provided endpoint and query string.
// It adds necessary headers, including the API key and user-agent, for authenticating the request.
func (fw *FetchWrapper) createRequest(endpoint, queryString string) (*http.Request, error) {
	// Construct full URL with endpoint and query string
	fullURL := fmt.Sprintf("%s%s?%s", fw.BaseURL, endpoint, queryString)
	log.Printf("Making request to: %s", fullURL) // Log the URL being requested

	// Create the HTTP GET request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	// Set authorization and user-agent headers
	req.Header.Add("Authorization", fw.APIKey)
	req.Header.Add("User-Agent", "Pexels-Go/1.0")

	return req, nil
}

// Fetch performs a GET request to the given endpoint with the specified parameters.
// It returns the decoded JSON response or an error if the request fails.
func (fw *FetchWrapper) Fetch(endpoint string, params map[string]interface{}) ([]byte, error) {
	// Construct query string from parameters
	queryString := fw.constructQueryString(params)

	// Create the HTTP request
	req, err := fw.createRequest(endpoint, queryString)
	if err != nil {
		return nil, err // Return error if request creation failed
	}

	// Execute the request using the HTTP client
	resp, err := fw.Client.Do(req)
	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err // Return error if the request execution failed
	}
	defer resp.Body.Close() // Ensure that the response body is closed when done

	// Log the status code of the response
	log.Printf("Received response with status: %s", resp.Status)

	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Received non-OK response: %d", resp.StatusCode)
		return nil, fmt.Errorf("received non-OK response: %d", resp.StatusCode)
	}

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err // Return error if reading the body failed
	}

	return body, nil // Return the response body
}
