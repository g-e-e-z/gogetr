package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvFiles loads environment variables from a slice of files
func LoadEnvFiles(envFiles []string) error {
	for _, envFile := range envFiles {
		err := godotenv.Overload(envFile) // Overload merges the env variables from multiple files
		if err != nil {
			log.Printf("Error loading env file %s: %v", envFile, err)
		}
	}
	return nil
}

