package utils

import (
	"os"
	// "path/filepath"
	"regexp"
)

// ListGroups returns a list of group directories
func ListGroups(configDir string) ([]string, error) {
	var groups []string

	// Read all directories in the specified path
	files, err := os.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	// Filter out directories (we are only interested in directories)
	for _, file := range files {
		if file.IsDir() {
			groups = append(groups, file.Name())
		}
	}

	return groups, nil
}

// ReplaceEnvVariables replaces placeholders in a string with corresponding environment variables
func ReplaceEnvVariables(input string) string {
	// Regex to find placeholders like {{api_key}}
	re := regexp.MustCompile(`\{\{([a-zA-Z0-9_]+)\}\}`)
	return re.ReplaceAllStringFunc(input, func(placeholder string) string {
		// Extract the variable name (without the surrounding {{ and }})
		varName := placeholder[2 : len(placeholder)-2]
		// Get the environment variable, defaulting to an empty string if not found
		envValue := os.Getenv(varName)
		return envValue
	})
}

