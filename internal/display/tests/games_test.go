package display_test

import (
	"go-nhl/internal/display"
	"go-nhl/nhl"
	"testing"
)

func TestGameDetails(t *testing.T) {
	// Create test data
	game := &nhl.GameDetails{
		ID:           2023020204,
		GameDate:     "2024-02-01",
		StartTimeUTC: "2024-02-01T00:00:00Z",
		Venue: nhl.Venue{
			Default: "Test Arena",
		},
		GameState: "FINAL",
		HomeTeam: nhl.DetailedTeam{
			CommonName:  nhl.LanguageNames{Default: "Home Team"},
			Abbrev:      "HOME",
			Score:       3,
			ShotsOnGoal: 30,
			PlaceName:   nhl.LanguageNames{Default: "Home City"},
		},
		AwayTeam: nhl.DetailedTeam{
			CommonName:  nhl.LanguageNames{Default: "Away Team"},
			Abbrev:      "AWAY",
			Score:       2,
			ShotsOnGoal: 25,
			PlaceName:   nhl.LanguageNames{Default: "Away City"},
		},
		Clock: nhl.GameClock{
			TimeRemaining: "00:00",
		},
	}

	boxscore := &nhl.BoxscoreResponse{
		ID:       2023020204,
		GameDate: "2024-02-01",
		PlayerByGameStats: nhl.PlayerGameStats{
			HomeTeam: nhl.TeamPlayerStats{
				Forwards: []nhl.PlayerStats{
					{
						Hits:              2,
						PIM:               2,
						FaceoffWinningPct: 60.0,
					},
				},
				Defense: []nhl.PlayerStats{
					{
						Hits: 3,
						PIM:  0,
					},
				},
			},
			AwayTeam: nhl.TeamPlayerStats{
				Forwards: []nhl.PlayerStats{
					{
						Hits:              1,
						PIM:               4,
						FaceoffWinningPct: 40.0,
					},
				},
				Defense: []nhl.PlayerStats{
					{
						Hits: 2,
						PIM:  2,
					},
				},
			},
		},
	}

	// Test display function (no error should occur)
	display.GameDetails(game, boxscore)
}

func TestGameBoxscore(t *testing.T) {
	// Create test data
	boxscore := &nhl.BoxscoreResponse{
		GameID:   2023020204,
		GameDate: "2024-02-01",
		HomeTeam: nhl.BoxscoreTeamStats{
			TeamStats: nhl.TeamStats{
				Goals:       3,
				ShotsOnGoal: 30,
				Hits:        20,
				FaceoffPct:  55.5,
				PIM:         6,
			},
		},
		AwayTeam: nhl.BoxscoreTeamStats{
			TeamStats: nhl.TeamStats{
				Goals:       2,
				ShotsOnGoal: 25,
				Hits:        22,
				FaceoffPct:  44.5,
				PIM:         8,
			},
		},
		Periods: []nhl.PeriodStats{
			{
				PeriodNumber: 1,
				HomeScore:    1,
				AwayScore:    1,
				Goals: []nhl.GoalSummary{
					{
						Period:       1,
						TimeInPeriod: "10:00",
						GoalType:     "EVEN",
						ScoredBy:     nhl.PlayerBrief{Name: nhl.LanguageNames{Default: "Player 1"}},
						AssistedBy:   []nhl.PlayerBrief{{Name: nhl.LanguageNames{Default: "Player 2"}}},
					},
				},
				Penalties: []nhl.PenaltySummary{
					{
						Period:       1,
						TimeInPeriod: "15:00",
						Type:         "Tripping",
						Minutes:      2,
						Player:       nhl.PlayerBrief{Name: nhl.LanguageNames{Default: "Player 3"}},
					},
				},
			},
		},
	}

	// Test display function (no error should occur)
	display.GameBoxscore(boxscore)
}

func TestGamePlayByPlay(t *testing.T) {
	// Create test data
	pbp := &nhl.PlayByPlayResponse{
		GameID: 2023020204,
		Plays: []nhl.Play{
			{
				Period:       1,
				TimeInPeriod: "10:00",
				Type:         "SHOT",
				Description:  "Shot by Player 1",
				Details: nhl.PlayDetails{
					EventOwner: "HOME",
					Shooter:    nhl.PlayerBrief{Name: nhl.LanguageNames{Default: "Player 1"}},
					Goalie:     nhl.PlayerBrief{Name: nhl.LanguageNames{Default: "Player 2"}},
				},
			},
		},
	}

	// Test display function (no error should occur)
	display.GamePlayByPlay(pbp)
}

func TestGameStory(t *testing.T) {
	// Create test data
	story := &nhl.GameStoryResponse{
		GameID:      2023020204,
		Headline:    "Test Headline",
		SubHeadline: "Test Subheadline",
		Story:       "Test story content.",
	}

	// Test display function (no error should occur)
	display.GameStory(story)
}

func TestFormatAssists(t *testing.T) {
	tests := []struct {
		name     string
		assists  []nhl.PlayerBrief
		expected string
	}{
		{
			name:     "No assists",
			assists:  []nhl.PlayerBrief{},
			expected: "Unassisted",
		},
		{
			name: "Single assist",
			assists: []nhl.PlayerBrief{
				{Name: nhl.LanguageNames{Default: "Player 1"}},
			},
			expected: "Player 1",
		},
		{
			name: "Multiple assists",
			assists: []nhl.PlayerBrief{
				{Name: nhl.LanguageNames{Default: "Player 1"}},
				{Name: nhl.LanguageNames{Default: "Player 2"}},
			},
			expected: "Player 1, Player 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := display.FormatAssists(tt.assists)
			if result != tt.expected {
				t.Errorf("FormatAssists() = %v, want %v", result, tt.expected)
			}
		})
	}
}
