package nhl_test

import (
	"go-nhl/internal/display"
	"go-nhl/client"
	"testing"
)

func TestGetStandings(t *testing.T) {
	client := nhl.NewClient()

	standings, err := client.GetStandings()
	if err != nil {
		t.Fatalf("GetStandings() error = %v", err)
	}

	if standings == nil {
		t.Fatal("GetStandings() returned nil")
	}

	if len(standings.Standings) == 0 {
		t.Error("GetStandings() returned empty standings")
	}

	// Test team data
	for _, team := range standings.Standings {
		// Basic info
		if team.TeamName.Default == "" {
			t.Error("Team name is empty")
		}
		if team.Conference == "" {
			t.Error("Conference name is empty")
		}
		if team.Division == "" {
			t.Error("Division name is empty")
		}

		// Record calculations
		if team.Points != (team.Wins*2 + team.OtLosses) {
			t.Errorf("Points calculation incorrect: got %d, want %d", team.Points, team.Wins*2+team.OtLosses)
		}

		// Home/Away records should add up to total record
		totalHomeAway := team.HomeWins + team.HomeLosses + team.HomeOtLosses +
			(team.Wins - team.HomeWins) + (team.Losses - team.HomeLosses) + (team.OtLosses - team.HomeOtLosses)
		if totalHomeAway != team.GamesPlayed {
			t.Errorf("Home/Away games don't match total: got %d, want %d", totalHomeAway, team.GamesPlayed)
		}

		// L10 record should not exceed 10 games
		l10Total := team.L10Wins + team.L10Losses + team.L10OtLosses
		if l10Total > 10 {
			t.Errorf("L10 record exceeds 10 games: %d", l10Total)
		}

		// Goal differential should match goals for/against
		if team.GoalDifferential != team.GoalsFor-team.GoalsAgainst {
			t.Errorf("Goal differential incorrect: got %d, want %d",
				team.GoalDifferential, team.GoalsFor-team.GoalsAgainst)
		}

		// Regulation wins should not exceed total wins
		if team.RegulationWins > team.Wins {
			t.Errorf("Regulation wins (%d) exceeds total wins (%d)",
				team.RegulationWins, team.Wins)
		}

		// Points percentage should be between 0 and 1
		ptsPercentage := float64(team.Points) / float64(team.GamesPlayed*2)
		if ptsPercentage < 0 || ptsPercentage > 1 {
			t.Errorf("Points percentage out of range: %f", ptsPercentage)
		}
	}
}

func TestGetStandingsByDate(t *testing.T) {
	client := nhl.NewClient()

	tests := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{
			name:    "Valid date",
			date:    "2024-02-01",
			wantErr: false,
		},
		{
			name:    "Invalid date format",
			date:    "invalid",
			wantErr: true,
		},
		{
			name:    "Future date",
			date:    "2025-01-01",
			wantErr: false, // Should return projected standings
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			standings, err := client.GetStandingsByDate(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStandingsByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && standings == nil {
				t.Error("GetStandingsByDate() returned nil standings")
			}
		})
	}
}

func TestStandingsSorting(t *testing.T) {
	// Create test data
	standings := &nhl.StandingsResponse{
		Standings: []nhl.StandingsTeam{
			{
				TeamName:         nhl.LanguageNames{Default: "Team A"},
				Points:           80,
				RegulationWins:   30,
				GoalDifferential: 20,
			},
			{
				TeamName:         nhl.LanguageNames{Default: "Team B"},
				Points:           80,
				RegulationWins:   30,
				GoalDifferential: 25,
			},
			{
				TeamName:         nhl.LanguageNames{Default: "Team C"},
				Points:           80,
				RegulationWins:   35,
				GoalDifferential: 15,
			},
			{
				TeamName:         nhl.LanguageNames{Default: "Team D"},
				Points:           85,
				RegulationWins:   25,
				GoalDifferential: 10,
			},
		},
	}

	// Expected order after sorting:
	// 1. Team D (85 points)
	// 2. Team C (80 points, 35 regulation wins)
	// 3. Team B (80 points, 30 regulation wins, +25 goal diff)
	// 4. Team A (80 points, 30 regulation wins, +20 goal diff)

	// Sort the teams using the shared sorting function
	display.SortTeams(standings.Standings)

	// Verify the order
	expectedOrder := []string{"Team D", "Team C", "Team B", "Team A"}
	for i, expected := range expectedOrder {
		if standings.Standings[i].TeamName.Default != expected {
			t.Errorf("Wrong team at position %d: got %s, want %s",
				i+1, standings.Standings[i].TeamName.Default, expected)
		}
	}
}
