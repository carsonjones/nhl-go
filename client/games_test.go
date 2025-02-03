package nhl_test

import (
	"go-nhl/client"
	"testing"
)

func TestGetGameDetails(t *testing.T) {
	client := nhl.NewClient()

	// Test with a valid game ID
	details, err := client.GetGameDetails(2023020204)
	if err != nil {
		t.Fatalf("GetGameDetails() error = %v", err)
	}

	// Validate game details
	if details == nil {
		t.Fatal("GetGameDetails() returned nil")
	}

	// Basic info validation
	if details.ID == 0 {
		t.Error("Game ID is 0")
	}
	if details.GameDate == "" {
		t.Error("Game date is empty")
	}
	if details.Venue.Default == "" {
		t.Error("Venue is empty")
	}

	// Team info validation
	if details.HomeTeam.CommonName.Default == "" {
		t.Error("Home team name is empty")
	}
	if details.AwayTeam.CommonName.Default == "" {
		t.Error("Away team name is empty")
	}

	// Game state validation
	validStates := map[string]bool{
		"LIVE":      true,
		"OFF":       true,
		"FINAL":     true,
		"SCHEDULED": true,
	}
	if !validStates[details.GameState] {
		t.Errorf("Invalid game state: %s", details.GameState)
	}
}

func TestGetGameBoxscore(t *testing.T) {
	client := nhl.NewClient()

	boxscore, err := client.GetGameBoxscore(2023020204)
	if err != nil {
		t.Fatalf("GetGameBoxscore() error = %v", err)
	}

	// Validate boxscore
	if boxscore == nil {
		t.Fatal("GetGameBoxscore() returned nil")
	}

	// Basic info validation
	if boxscore.ID == 0 {
		t.Error("Game ID is 0")
	}
	if boxscore.GameDate == "" {
		t.Error("Game date is empty")
	}

	// Team stats validation
	if boxscore.HomeTeam.Score < 0 {
		t.Error("Invalid home team score")
	}
	if boxscore.AwayTeam.Score < 0 {
		t.Error("Invalid away team score")
	}

	// Player stats validation
	validatePlayerStats(t, boxscore.PlayerByGameStats.HomeTeam, "Home")
	validatePlayerStats(t, boxscore.PlayerByGameStats.AwayTeam, "Away")
}

func validatePlayerStats(t *testing.T, stats nhl.TeamPlayerStats, teamType string) {
	// Check forwards
	for i, player := range stats.Forwards {
		if player.TOI == "" {
			t.Errorf("%s team forward %d has no time on ice", teamType, i)
		}
	}

	// Check defense
	for i, player := range stats.Defense {
		if player.TOI == "" {
			t.Errorf("%s team defense %d has no time on ice", teamType, i)
		}
	}

	// Check goalies
	for i, goalie := range stats.Goalies {
		if goalie.TOI == "" {
			t.Errorf("%s team goalie %d has no time on ice", teamType, i)
		}
	}
}

func TestGetGamePlayByPlay(t *testing.T) {
	client := nhl.NewClient()

	pbp, err := client.GetGamePlayByPlay(2023020204)
	if err != nil {
		t.Fatalf("GetGamePlayByPlay() error = %v", err)
	}

	// Validate play-by-play
	if pbp == nil {
		t.Fatal("GetGamePlayByPlay() returned nil")
	}

	// Validate plays
	for i, play := range pbp.Plays {
		if play.TimeInPeriod == "" {
			t.Errorf("Play %d has no time in period", i)
		}
		if play.PeriodDescriptor.Number == 0 {
			t.Errorf("Play %d has invalid period number", i)
		}
	}

	// Validate roster spots
	for i, player := range pbp.RosterSpots {
		if player.PlayerID == 0 {
			t.Errorf("Roster spot %d has invalid player ID", i)
		}
		if player.FirstName.Default == "" || player.LastName.Default == "" {
			t.Errorf("Roster spot %d has invalid player name", i)
		}
	}
}
