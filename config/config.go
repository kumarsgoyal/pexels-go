package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Config holds the structure of the configuration file
type Config struct {
	PexelAPIKey string `json:"pexelApiKey"` // API key field expected in the JSON config
}

// LoadConfig loads the config from the specified file path
func LoadConfig(filePath string) (*Config, error) {
	log.Printf("Attempting to open the config file: %s", filePath)

	// Open the config file
	file, err := openConfigFile(filePath)
	if err != nil {
		return nil, err
	}

	// Ensure the file is closed when the function returns
	defer closeConfigFile(file)

	// Decode the JSON configuration into the Config struct
	config, err := decodeConfig(file)
	if err != nil {
		return nil, err
	}

	// Validate the loaded configuration to ensure required fields are present
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	log.Println("Config file decoded and validated successfully.")
	return config, nil
}

// Helper function to open the config file
func openConfigFile(filePath string) (*os.File, error) {
	// Attempt to open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening config file: %s", err)
		return nil, fmt.Errorf("failed to open config file: %w", err) // Wrap the error for better debugging
	}
	log.Println("Config file opened successfully.")
	return file, nil
}

// Helper function to close the config file
func closeConfigFile(file *os.File) {
	log.Println("Closing the config file.")
	file.Close() // Close the file to release system resources
}

// Helper function to decode the JSON into the Config struct
func decodeConfig(file *os.File) (*Config, error) {
	log.Println("Decoding the JSON...")

	// Initialize an empty Config struct
	var config Config
	decoder := json.NewDecoder(file) // Create a new JSON decoder for the file

	// Decode JSON content into the struct
	if err := decoder.Decode(&config); err != nil {
		log.Printf("Error decoding config JSON: %s", err)
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	return &config, nil
}

// Helper function to validate the config data
func validateConfig(config *Config) error {
	// Check if the API key is missing or empty
	if config.PexelAPIKey == "" {
		log.Println("API key is missing in the config file.")
		return fmt.Errorf("api key is missing in the config file")
	}
	return nil
}
