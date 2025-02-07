package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/g-e-e-z/gogetr/requests"
)

func rmain() {

	pwd, _ := os.Getwd()
	requestsDir := filepath.Join(pwd, "requests_dir")
	fmt.Printf("Requests Dir: %s\n", requestsDir)

	// Load all requests from the given directory
	groupRequests, err := requests.LoadAllRequests(requestsDir)
	if err != nil {
		log.Fatalf("Error loading group requests: %v", err)
	}

	// Iterate over groups and their requests
	for group, groupRequests := range groupRequests {
		fmt.Printf("Group: %s\n", group)
		for i, req := range groupRequests.Requests {
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

			// Execute the request
			response, err := req.Execute()
			if err != nil {
				log.Panic(err)
			}
			req.ParseResponse(response)
		}
	}
}

