package nhl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// mockHTTPClient is a mock implementation of HTTPClient for testing
type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// mockResponse creates a mock HTTP response with the given status code and body
func mockResponse(statusCode int, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewReader(jsonBody)),
	}, nil
}

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
	if client.baseURL != BaseURLWeb {
		t.Errorf("NewClient().baseURL = %v, want %v", client.baseURL, BaseURLWeb)
	}
	if client.httpClient == nil {
		t.Error("NewClient().httpClient is nil")
	}
}

func TestClientGet(t *testing.T) {
	type testStruct struct {
		Field string `json:"field"`
	}

	tests := []struct {
		name       string
		url        string
		response   interface{}
		statusCode int
		wantErr    bool
	}{
		{
			name: "Successful request",
			url:  "https://api.example.com/test",
			response: testStruct{
				Field: "test value",
			},
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "Non-200 status code",
			url:        "https://api.example.com/error",
			response:   nil,
			statusCode: http.StatusNotFound,
			wantErr:    true,
		},
		{
			name:       "Invalid JSON response",
			url:        "https://api.example.com/invalid",
			response:   "invalid json",
			statusCode: http.StatusOK,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient()
			client.httpClient = &mockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					if req.URL.String() != tt.url {
						return nil, fmt.Errorf("unexpected URL: got %v, want %v", req.URL.String(), tt.url)
					}
					return mockResponse(tt.statusCode, tt.response)
				},
			}

			var result testStruct
			err := client.get(tt.url, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				expected := tt.response.(testStruct)
				if result.Field != expected.Field {
					t.Errorf("get() result = %v, want %v", result, expected)
				}
			}
		})
	}
}
