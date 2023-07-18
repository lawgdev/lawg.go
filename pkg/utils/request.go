package lawg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Options struct {
	UA     *string         `json:"ua,omitempty"`
	Method string          `json:"method"`
	Token  string          `json:"token"`
	Data   json.RawMessage `json:"data,omitempty"`
}

const API_URL = "https://api.lawg.dev/v1"

func getUA(ua *string) string {
	if ua != nil {
		return *ua
	}
	
	return "lawg.js; (+https://github.com/lawgdev/lawg.js)"
}

func Request(url string, options Options) (*http.Response, error) {
	constructedURL := fmt.Sprintf("%s/%s", API_URL, url)
	reqBody, err := json.Marshal(options.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %v", err)
	}

	req, err := http.NewRequest(options.Method, constructedURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", getUA(options.UA))
	req.Header.Set("Authorization", options.Token)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	return resp, nil
}
