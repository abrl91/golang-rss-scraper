package auth

import (
	"errors"
	"net/http"
)

// extracts the API key from the headers
func GetApiKey(headers http.Header) (string, error) {
	apiKey := headers.Get("X-API-Key")

	if apiKey == "" {
		return "", errors.New("API Key is required")
	}

	return apiKey, nil
}
