package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPRequestOptions struct {
	Method      string
	URL         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        interface{}
}

func HttpRequest(options HTTPRequestOptions) ([]byte, error) {
	// Prepare request body if present
	var requestBody []byte
	if options.Body != nil {
		var err error
		requestBody, err = json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request object
	req, err := http.NewRequest(options.Method, options.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Add headers to the request
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// Add query parameters if present
	query := req.URL.Query()
	for key, value := range options.QueryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// Create HTTP client
	client := &http.Client{}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check if the response status code is not OK
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return respBody, nil
}
