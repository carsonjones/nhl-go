package nhl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// HTTPClient interface for making HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents an NHL API client
type Client struct {
	baseURL    string
	httpClient HTTPClient
	teams      *TeamsResponse // Cache of teams
	cacheMutex sync.RWMutex
}

// NewClient creates a new NHL API client
func NewClient() *Client {
	return &Client{
		baseURL: BaseURLWeb,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

// get performs a GET request and unmarshals the response into v
func (c *Client) get(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return nil
}
