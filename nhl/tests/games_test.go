package nhl_test

import (
	"go-nhl/nhl"
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
	if details.HomeTeam.Name.Default == "" {
		t.Error("Home team name is empty")
	}
	if details.AwayTeam.Name.Default == "" {
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
	if boxscore.GameID == 0 {
		t.Error("Game ID is 0")
	}
	if boxscore.GameDate == "" {
		t.Error("Game date is empty")
	}

	// Team stats validation
	validateTeamStats(t, boxscore.HomeTeam.TeamStats, "Home")
	validateTeamStats(t, boxscore.AwayTeam.TeamStats, "Away")

	// Period stats validation
	if len(boxscore.Periods) == 0 {
		t.Error("No period stats found")
	}
	for i, period := range boxscore.Periods {
		if period.PeriodNumber == 0 {
			t.Errorf("Period %d has invalid number", i+1)
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

	// Basic info validation
	if pbp.GameID == 0 {
		t.Error("Game ID is 0")
	}

	// Plays validation
	if len(pbp.Plays) == 0 {
		t.Error("No plays found")
	}
	for i, play := range pbp.Plays {
		if play.Period == 0 {
			t.Errorf("Play %d has invalid period", i+1)
		}
		if play.TimeInPeriod == "" {
			t.Errorf("Play %d has no time", i+1)
		}
		if play.Description == "" {
			t.Errorf("Play %d has no description", i+1)
		}
	}
}

func TestGetGameStory(t *testing.T) {
	client := nhl.NewClient()

	story, err := client.GetGameStory(2023020204)
	if err != nil {
		t.Fatalf("GetGameStory() error = %v", err)
	}

	// Validate game story
	if story == nil {
		t.Fatal("GetGameStory() returned nil")
	}

	// Content validation
	if story.GameID == 0 {
		t.Error("Game ID is 0")
	}
	if story.Headline == "" {
		t.Error("Headline is empty")
	}
	if story.Story == "" {
		t.Error("Story content is empty")
	}
}

// Helper function to validate team stats
func validateTeamStats(t *testing.T, stats nhl.TeamStats, teamType string) {
	t.Helper()

	if stats.Goals < 0 {
		t.Errorf("%s team has invalid goals: %d", teamType, stats.Goals)
	}
	if stats.ShotsOnGoal < 0 {
		t.Errorf("%s team has invalid shots on goal: %d", teamType, stats.ShotsOnGoal)
	}
	if stats.FaceoffPct < 0 || stats.FaceoffPct > 100 {
		t.Errorf("%s team has invalid faceoff percentage: %f", teamType, stats.FaceoffPct)
	}
	if stats.Hits < 0 {
		t.Errorf("%s team has invalid hits: %d", teamType, stats.Hits)
	}
	if stats.PIM < 0 {
		t.Errorf("%s team has invalid penalty minutes: %d", teamType, stats.PIM)
	}
}
