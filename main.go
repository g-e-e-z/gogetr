package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/g-e-e-z/gogetr/requests"
	"github.com/joho/godotenv"
)

func loadRequestsFromFile(requestsFile string) (*requests.Requests, error) {
	data, err := os.ReadFile(requestsFile)
	if err != nil {
		return nil, err
	}

	var requests requests.Requests
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return nil, err
	}

	return &requests, nil
}

// Load requests from a group directory and replace environment variables in them
func loadRequestsFromGroup(groupDir string, envFiles []string) (*requests.Requests, error) {
	// Load the environment variables for the group
	err := loadEnvFiles(envFiles)
	if err != nil {
		return nil, err
	}

	// Load the requests.json file
	requestsFile := filepath.Join(groupDir, "requests.json")
	data, err := os.ReadFile(requestsFile)
	if err != nil {
		return nil, err
	}

	var requests requests.Requests
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return nil, err
	}

	// Replace environment variables in the request fields
	for i := range requests.Requests {
		requests.Requests[i].URL = replaceEnvVariables(requests.Requests[i].URL)
		for headerKey, headerValue := range requests.Requests[i].Headers {
			requests.Requests[i].Headers[headerKey] = replaceEnvVariables(headerValue)
		}
		for paramKey, paramValue := range requests.Requests[i].QueryParams {
			requests.Requests[i].QueryParams[paramKey] = replaceEnvVariables(paramValue)
		}
		if requests.Requests[i].Body != nil {
			*requests.Requests[i].Body = replaceEnvVariables(*requests.Requests[i].Body)
		}
	}

	return &requests, nil
}
func listGroups(configDir string) ([]string, error) {
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

// Load all requests from the config directory
func loadAllRequests(configDir string) (map[string]*requests.Requests, error) {
	groupRequests := make(map[string]*requests.Requests)
	groups, err := listGroups(configDir)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		groupDir := filepath.Join(configDir, group)
		files, err := os.ReadDir(groupDir)
		if err != nil {
			panic(err)
		}
		// TODO: Environment selection/ inheritance/ overlaoding will need to be
		// adjusted in the future
		//
		// Define the environment files to load
		envFiles := []string{
			filepath.Join(configDir, "default.env"),
		}
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".env" {
				envFiles = append(envFiles, filepath.Join(groupDir, file.Name()))
			}
		}
		requests, err := loadRequestsFromGroup(groupDir, envFiles)
		if err != nil {
			log.Printf("Error loading requests from group %s: %v", group, err)
			continue
		}
		groupRequests[group] = requests
	}

	return groupRequests, nil
}

// Load environment variables from a series of files
func loadEnvFiles(envFiles []string) error {
	for _, envFile := range envFiles {
		err := godotenv.Overload(envFile) // Overload allows merging of multiple .env files
		if err != nil {
			log.Printf("Error loading env file %s: %v", envFile, err)
		}
	}
	return nil
}

// Replace placeholders in a string with corresponding environment variables
func replaceEnvVariables(input string) string {
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

func main() {
	pwd, _ := os.Getwd()

	requestsDir := filepath.Join(pwd, "requests_dir")
	fmt.Printf("Requests Dir: %s\n", requestsDir)

	groupRequests, err := loadAllRequests(requestsDir)
	if err != nil {
		log.Fatalf("Error loading group requests: %v", err)
	}

	for group, requests := range groupRequests {
		fmt.Printf("Group: %s\n", group)
		for i, req := range requests.Requests {
			fmt.Printf("  Request %d:\n", i+1)
			fmt.Printf("    Method: %s\n", req.Method)
			fmt.Printf("    URL: %s\n", req.URL)
			fmt.Printf("    Headers: %v\n", req.Headers)
			fmt.Printf("    Query Params: %v\n", req.QueryParams)
			if req.Body != nil {
				fmt.Printf("    Body: %s\n", *req.Body)
			} else {
				fmt.Printf("    Body: null\n")
			}
			response, err := req.Execute()
			if err != nil {
				log.Panic(err)
			}
			req.ParseResponse(response)
		}
	}
}
