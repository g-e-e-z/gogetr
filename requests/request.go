package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request represents an HTTP request
type Request struct {
	Name        string            `json:"name"`
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"query_params"`
	Body        *string           `json:"body"`
}

// ResponseWithTime adds a ResponseTime field to http.Response
type ResponseWithTime struct {
	*http.Response
	ResponseTime time.Duration
}

// Requests holds an array of Requests
type Requests struct {
	Requests []Request `json:"requests"`
}

// NewRequest creates and returns a new Request instance
func NewRequest(name, method, url string, headers map[string]string, queryParams map[string]string, body *string) *Request {
	return &Request{
		Name:        name,
		Method:      method,
		URL:         url,
		Headers:     headers,
		QueryParams: queryParams,
		Body:        body,
	}
}

// Execute sends the HTTP request and returns the response or error
func (r *Request) Execute() (*ResponseWithTime, error) {
	// Build the URL with query parameters (if any)
	urlWithParams := r.buildURLWithParams()

	// Create a new HTTP request
	var reqBody *bytes.Buffer
	if r.Body != nil {
		reqBody = bytes.NewBuffer([]byte(*r.Body)) // If there's a body, convert it to a buffer
	} else {
		reqBody = bytes.NewBuffer([]byte("")) // If there's a body, convert it to a buffer
	}

	req, err := http.NewRequest(r.Method, urlWithParams, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	// Send the HTTP request
	client := &http.Client{
		Timeout: 30 * time.Second, // Timeout after 30 seconds
	}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	res := &ResponseWithTime{
		Response:     resp,
		ResponseTime: time.Since(start),
	}
	return res, nil
}

// buildURLWithParams constructs the full URL with query parameters appended
func (r *Request) buildURLWithParams() string {
	// If there are no query parameters, return the original URL
	if len(r.QueryParams) == 0 {
		return r.URL
	}

	// Parse the base URL to append query parameters
	parsedURL, err := url.Parse(r.URL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return r.URL
	}

	// Add query parameters to the URL
	query := parsedURL.Query()
	for key, value := range r.QueryParams {
		query.Set(key, value)
	}

	parsedURL.RawQuery = query.Encode()
	return parsedURL.String()
}

func mapprint(m map[string]string) string {
	var builder strings.Builder
	for key, val := range m {
		fmt.Fprintf(&builder, "\t%s: %s\n", key, val)
	}
	return builder.String()
}

// ViewerFormat formats the request to be displayed in the tui
func (r *Request) ViewerFormat() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "METHOD: %s\nURL: %s\n", r.Method, r.URL)
	fmt.Fprintf(&builder, "HEADERS:\n%s", mapprint(r.Headers))
	fmt.Fprintf(&builder, "QUERY_PARAMS:\n%s", mapprint(r.QueryParams))
	// TODO: Fix this formatting of json strings
	if r.Body != nil {
		jsonData, err := json.MarshalIndent(r.Body, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(&builder, "BODY:\n%s", jsonData)
	}
	return builder.String()
}

// ParseResponse reads and prints the response body
func (r *Request) ParseResponse(resp *ResponseWithTime) string {
    var builder strings.Builder
	// Ensure the response body is closed after reading
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
        fmt.Fprintf(&builder, "ERROR READING RESPONSE BODY: %s\n", err)
		return builder.String()
	}

	// Print response details

	fmt.Fprintf(&builder, "RESPONSE STATUS: %s\n", resp.Status)
	fmt.Fprintf(&builder, "RESPONSE TIME: %s\n", resp.ResponseTime)
	fmt.Fprintf(&builder, "RESPONSE BODY: %s\n", string(body))
    return builder.String()
}
