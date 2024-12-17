package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// CleanParams removes empty or zero-value parameters from the provided map.
// It ensures that only non-nil, non-empty, and non-zero parameters are included.
func CleanParams(paramsMap map[string]interface{}) map[string]interface{} {
	cleanedParamsMap := make(map[string]interface{})
	for key, value := range paramsMap {
		// Check if value is not nil, empty string, zero number, empty array, or empty map
		if !isZeroValue(value) {
			cleanedParamsMap[key] = value
		}
	}
	return cleanedParamsMap
}

// Helper function to check for zero-value of various types
func isZeroValue(value interface{}) bool {
	if value == nil {
		return true
	}

	// Use reflection to check for zero values (for non-numeric types, empty strings, slices, etc.)
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		// Empty string, slice, map, or channel
		return v.Len() == 0
	case reflect.Ptr:
		// Nil pointer
		return v.IsNil()
	default:
		// Zero value for numeric types
		return v.IsZero()
	}
}

// Helper function to unmarshal responses
func UnmarshalResponse(body []byte, target interface{}) error {
	// Log the raw response body for debugging purposes
	fmt.Printf("Raw response body: %s\n", string(body))

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}
	return nil
}
