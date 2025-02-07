package requests

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/g-e-e-z/gogetr/config"
	"github.com/g-e-e-z/gogetr/utils"
)

// LoadRequestsFromFile loads requests from a specific file
func LoadRequestsFromFile(requestsFile string) (*Requests, error) {
	data, err := os.ReadFile(requestsFile)
	if err != nil {
		return nil, err
	}

	var requests Requests
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return nil, err
	}

	return &requests, nil
}

// LoadRequestsFromGroup loads requests from a group directory and replaces placeholders with environment variables
func LoadRequestsFromGroup(groupDir string, envFiles []string) (*Requests, error) {
	// Load environment variables for the group
	err := config.LoadEnvFiles(envFiles)
	if err != nil {
		return nil, err
	}

	// Load the requests from requests.json
	requestsFile := filepath.Join(groupDir, "requests.json")
	requests, err := LoadRequestsFromFile(requestsFile)
	if err != nil {
		return nil, err
	}

	// Replace environment variables in URLs, headers, and body
	for i := range requests.Requests {
		requests.Requests[i].URL = utils.ReplaceEnvVariables(requests.Requests[i].URL)
		for headerKey, headerValue := range requests.Requests[i].Headers {
			requests.Requests[i].Headers[headerKey] = utils.ReplaceEnvVariables(headerValue)
		}
		for paramKey, paramValue := range requests.Requests[i].QueryParams {
			requests.Requests[i].QueryParams[paramKey] = utils.ReplaceEnvVariables(paramValue)
		}
		if requests.Requests[i].Body != nil {
			*requests.Requests[i].Body = utils.ReplaceEnvVariables(*requests.Requests[i].Body)
		}
	}

	return requests, nil
}

// LoadAllRequests loads requests for all groups
func LoadAllRequests(configDir string) (map[string]*Requests, error) {
	groupRequests := make(map[string]*Requests)
	groups, err := utils.ListGroups(configDir)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		groupDir := filepath.Join(configDir, group)
		files, err := os.ReadDir(groupDir)
		if err != nil {
			log.Println("Error reading group directory:", err)
			continue
		}

		// Define the environment files to load
		envFiles := []string{
			filepath.Join(configDir, "default.env"),
		}
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".env" {
				envFiles = append(envFiles, filepath.Join(groupDir, file.Name()))
			}
		}

		// Load requests for the group
		requests, err := LoadRequestsFromGroup(groupDir, envFiles)
		if err != nil {
			log.Printf("Error loading requests from group %s: %v", group, err)
			continue
		}
		groupRequests[group] = requests
	}

	return groupRequests, nil
}

