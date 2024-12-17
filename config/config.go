package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	PexelAPIKey string `json:"pexelApiKey"`
}

// LoadConfig loads the config from the specified file path
func LoadConfig(filePath string) (*Config, error) {
	log.Printf("Attempting to open the config file: %s", filePath)

	// Open the config file
	file, err := openConfigFile(filePath)
	if err != nil {
		return nil, err
	}

	// Close the file when the function returns
	defer closeConfigFile(file)

	// Decode the JSON from the config file
	config, err := decodeConfig(file)
	if err != nil {
		return nil, err
	}

	// Validate the config for missing API key
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	log.Println("Config file decoded and validated successfully.")
	return config, nil
}

// Helper function to open the config file
func openConfigFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening config file: %s", err)
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	log.Println("Config file opened successfully.")
	return file, nil
}

// Helper function to close the config file
func closeConfigFile(file *os.File) {
	log.Println("Closing the config file.")
	file.Close()
}

// Helper function to decode the JSON into the Config struct
func decodeConfig(file *os.File) (*Config, error) {
	log.Println("Decoding the JSON...")

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Printf("Error decoding config JSON: %s", err)
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	return &config, nil
}

// Helper function to validate the config data
func validateConfig(config *Config) error {
	if config.PexelAPIKey == "" {
		log.Println("API key is missing in the config file.")
		return fmt.Errorf("api key is missing in the config file")
	}
	return nil
}
