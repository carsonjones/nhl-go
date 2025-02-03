package nhl_test

import (
	"go-nhl/nhl"
	"testing"
	"time"
)

func TestSearchPlayer(t *testing.T) {
	client := nhl.NewClient()

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
	client := nhl.NewClient()

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
		filter     *nhl.StatsFilter
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
	client := nhl.NewClient()

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
	client := nhl.NewClient()

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
		filter   *nhl.StatsFilter
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
			filter: &nhl.StatsFilter{
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
			filter: &nhl.StatsFilter{
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
