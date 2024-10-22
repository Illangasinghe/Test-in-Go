package protocol_helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"test-in-go/config"
)

func PostRequest(endpoint string, payload interface{}) (*http.Response, error) {
	// Convert payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Create the API request
	apiURL := config.GetEnv("API_URL") + endpoint
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create API request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "utf-8")
	username := config.GetEnv("API_USER")
	password := config.GetEnv("API_PASS")
	req.SetBasicAuth(username, password) // Set Basic Authentication

	// Execute the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send API request: %v", err)
	}

	return resp, nil
}
