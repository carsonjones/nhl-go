package nhl

import (
	"encoding/json"
	"fmt"
	"go-nhl/internal/formatters"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStatsLeaders(t *testing.T) {
	testCases := []struct {
		name           string
		seasonID       int
		expectedSeason int
	}{
		{
			name:           "Current season",
			seasonID:       0,
			expectedSeason: formatters.GetCurrentSeasonID(),
		},
		{
			name:           "Previous season",
			seasonID:       20222023,
			expectedSeason: 20222023,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request method
				if r.Method != http.MethodGet {
					t.Errorf("Expected 'GET' request, got '%s'", r.Method)
				}

				// Verify request path
				expectedPath := fmt.Sprintf("/stats/rest/en/leaders/skaters/points?season=%d", tc.expectedSeason)
				if r.URL.Path != "/stats/rest/en/leaders/skaters/points" {
					t.Errorf("Expected request to '%s', got '%s'", expectedPath, r.URL.Path)
				}

				// Verify season parameter
				season := r.URL.Query().Get("season")
				expectedSeason := fmt.Sprintf("%d", tc.expectedSeason)
				if season != expectedSeason {
					t.Errorf("Expected season parameter '%s', got '%s'", expectedSeason, season)
				}

				// Return mock response
				response := StatsLeadersResponse{
					Points: []StatsLeaderPlayer{
						{
							FirstName:  LanguageNames{Default: "Nikita"},
							LastName:   LanguageNames{Default: "Kucherov"},
							TeamAbbrev: "TBL",
							Value:      144,
						},
					},
				}
				json.NewEncoder(w).Encode(response)
			}))
			defer server.Close()

			// Create client with test server URL
			client := &Client{
				baseURL:    server.URL,
				httpClient: server.Client(),
			}

			// Call GetStatsLeaders
			leaders, err := client.GetStatsLeaders(tc.seasonID)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Verify response
			if len(leaders.Points) != 1 {
				t.Errorf("Expected 1 points leader, got %d", len(leaders.Points))
			}
			if leaders.Points[0].FirstName.Default != "Nikita" {
				t.Errorf("Expected first name 'Nikita', got '%s'", leaders.Points[0].FirstName.Default)
			}
			if leaders.Points[0].LastName.Default != "Kucherov" {
				t.Errorf("Expected last name 'Kucherov', got '%s'", leaders.Points[0].LastName.Default)
			}
			if leaders.Points[0].TeamAbbrev != "TBL" {
				t.Errorf("Expected team abbrev 'TBL', got '%s'", leaders.Points[0].TeamAbbrev)
			}
			if leaders.Points[0].Value != 144 {
				t.Errorf("Expected value 144, got %d", leaders.Points[0].Value)
			}
		})
	}
}

func TestGetStatsLeadersError(t *testing.T) {
	// Create test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create client with test server URL
	client := &Client{
		baseURL:    server.URL,
		httpClient: server.Client(),
	}

	// Call GetStatsLeaders and verify error
	_, err := client.GetStatsLeaders(0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
