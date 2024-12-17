package utils

import (
	"log"
	"os"
)

const (
	LogFilePath = "pexels_client.log"
)

// ConfigureLogging sets up global logging settings
// logFilePath allows specifying a custom log file location.
func ConfigureLogging() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Open or create the log file
	logFile, err := os.OpenFile(LogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to set up log file at %s: %v", LogFilePath, err)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Printf("Failed to close log file: %v", err)
		}
	}()

	log.SetOutput(logFile)
	log.Printf("Logging initialized. Writing logs to %s", LogFilePath)
}
