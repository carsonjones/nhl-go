package nhl

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
	if client.baseURL != BaseURLWeb {
		t.Errorf("NewClient() baseURL = %v, want %v", client.baseURL, BaseURLWeb)
	}
	if client.httpClient == nil {
		t.Fatal("NewClient() httpClient is nil")
	}
	if client.httpClient.Timeout != time.Second*30 {
		t.Errorf("NewClient() timeout = %v, want %v", client.httpClient.Timeout, time.Second*30)
	}
}

func TestGetTeamByIdentifier(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name       string
		identifier string
		wantTeam   bool
		wantErr    bool
	}{
		{
			name:       "Valid team abbreviation",
			identifier: "TOR",
			wantTeam:   true,
			wantErr:    false,
		},
		{
			name:       "Valid team name",
			identifier: "Toronto Maple Leafs",
			wantTeam:   true,
			wantErr:    false,
		},
		{
			name:       "Valid team ID",
			identifier: "10",
			wantTeam:   true,
			wantErr:    false,
		},
		{
			name:       "Invalid team",
			identifier: "INVALID",
			wantTeam:   false,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			team, err := client.GetTeamByIdentifier(tt.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTeamByIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (team != nil) != tt.wantTeam {
				t.Errorf("GetTeamByIdentifier() team = %v, want %v", team, tt.wantTeam)
			}
		})
	}
}

func TestGetTeams(t *testing.T) {
	client := NewClient()
	teams, err := client.GetTeams()
	if err != nil {
		t.Fatalf("GetTeams() error = %v", err)
	}
	if teams == nil {
		t.Fatal("GetTeams() returned nil")
	}
	if len(teams.Teams) == 0 {
		t.Error("GetTeams() returned empty teams list")
	}

	// Test caching
	teams2, err := client.GetTeams()
	if err != nil {
		t.Fatalf("GetTeams() second call error = %v", err)
	}
	if teams2 != teams {
		t.Error("GetTeams() cache not working, got different instances")
	}
}

// mockHTTPClient helps test API endpoints
type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetTeamSchedule(t *testing.T) {
	client := NewClient()
	team, err := client.GetTeamByIdentifier("TOR")
	if err != nil {
		t.Fatalf("Failed to get team: %v", err)
	}

	// Test with nil team
	_, err = client.GetTeamSchedule(nil, 20232024)
	if err == nil {
		t.Error("GetTeamSchedule() with nil team should return error")
	}

	// Test with valid team but invalid season
	_, err = client.GetTeamSchedule(team, -1)
	if err == nil {
		t.Error("GetTeamSchedule() with invalid season should return error")
	}

	// Test with valid team and season
	schedule, err := client.GetTeamSchedule(team, 20232024)
	if err != nil {
		t.Errorf("GetTeamSchedule() error = %v", err)
		return
	}
	if schedule == nil {
		t.Error("GetTeamSchedule() returned nil schedule")
	}
	if len(schedule.Games) == 0 {
		t.Error("GetTeamSchedule() returned empty games list")
	}
}

func TestSearchPlayer(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name      string
		query     string
		wantCount int
		wantErr   bool
	}{
		{
			name:      "Valid player name - Matthews",
			query:     "Matthews",
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "Valid player name - common surname",
			query:     "Smith",
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "Empty query",
			query:     "",
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:      "Invalid player name",
			query:     "xxxxxxxxxxx",
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := client.SearchPlayer(tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(results) < tt.wantCount {
				t.Errorf("SearchPlayer() got %v results, want at least %v", len(results), tt.wantCount)
			}
		})
	}
}

func TestGetPlayerStats(t *testing.T) {
	client := NewClient()

	// First get a player ID through search
	results, err := client.SearchPlayer("Matthews")
	if err != nil || len(results) == 0 {
		t.Fatal("Failed to get test player ID")
	}
	playerID := results[0].PlayerID

	tests := []struct {
		name       string
		playerID   int
		isGoalie   bool
		reportType string
		filter     *StatsFilter
		wantErr    bool
	}{
		{
			name:       "Valid player regular stats",
			playerID:   playerID,
			isGoalie:   false,
			reportType: "regularSeason",
			filter:     nil,
			wantErr:    false,
		},
		{
			name:       "Valid player playoff stats",
			playerID:   playerID,
			isGoalie:   false,
			reportType: "playoffs",
			filter:     nil,
			wantErr:    false,
		},
		{
			name:       "Invalid player ID",
			playerID:   -1,
			isGoalie:   false,
			reportType: "regularSeason",
			filter:     nil,
			wantErr:    true,
		},
		{
			name:       "Invalid report type",
			playerID:   playerID,
			isGoalie:   false,
			reportType: "invalid",
			filter:     nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := client.GetPlayerStats(tt.playerID, tt.isGoalie, tt.reportType, tt.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && stats == nil {
				t.Error("GetPlayerStats() returned nil stats")
			}
		})
	}
}

func TestGetPlayerSeasonStats(t *testing.T) {
	client := NewClient()

	// First get a player ID through search
	results, err := client.SearchPlayer("Matthews")
	if err != nil || len(results) == 0 {
		t.Fatal("Failed to get test player ID")
	}
	playerID := results[0].PlayerID

	tests := []struct {
		name     string
		playerID int
		wantErr  bool
	}{
		{
			name:     "Valid player ID",
			playerID: playerID,
			wantErr:  false,
		},
		{
			name:     "Invalid player ID",
			playerID: -1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := client.GetPlayerSeasonStats(tt.playerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerSeasonStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && stats == nil {
				t.Error("GetPlayerSeasonStats() returned nil stats")
			}
		})
	}
}

func TestGetFilteredPlayerStats(t *testing.T) {
	client := NewClient()

	// First get a player ID through search
	results, err := client.SearchPlayer("Matthews")
	if err != nil || len(results) == 0 {
		t.Fatal("Failed to get test player ID")
	}
	playerID := results[0].PlayerID

	currentSeason := time.Now().Year() * 10000
	if time.Now().Month() < time.October {
		currentSeason = (time.Now().Year() - 1) * 10000
	}
	currentSeason = currentSeason + (currentSeason/10000 + 1)

	tests := []struct {
		name     string
		playerID int
		filter   *StatsFilter
		wantErr  bool
	}{
		{
			name:     "Valid player with nil filter",
			playerID: playerID,
			filter:   nil,
			wantErr:  false,
		},
		{
			name:     "Valid player with season filter",
			playerID: playerID,
			filter: &StatsFilter{
				SeasonID: currentSeason,
			},
			wantErr: false,
		},
		{
			name:     "Invalid player ID",
			playerID: -1,
			filter:   nil,
			wantErr:  true,
		},
		{
			name:     "Valid player with invalid season",
			playerID: playerID,
			filter: &StatsFilter{
				SeasonID: -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := client.GetFilteredPlayerStats(tt.playerID, tt.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilteredPlayerStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && stats == nil {
				t.Error("GetFilteredPlayerStats() returned nil stats")
			}
		})
	}
}
