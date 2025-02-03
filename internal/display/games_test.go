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
		Summary: nhl.GameSummary{
			Scoring: []nhl.PeriodSummary{
				{
					PeriodDescriptor: nhl.PeriodDescriptor{Number: 1},
					Goals: []nhl.GoalEvent{
						{
							TimeInPeriod: "10:00",
							Name:         nhl.LanguageNames{Default: "Player 1"},
							TeamAbbrev:   nhl.LanguageNames{Default: "HOME"},
							Assists:      []nhl.AssistEvent{},
						},
					},
				},
			},
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
		ID:       2023020204,
		GameDate: "2024-02-01",
		HomeTeam: nhl.DetailedTeam{
			Abbrev:      "HOME",
			Score:       3,
			ShotsOnGoal: 30,
		},
		AwayTeam: nhl.DetailedTeam{
			Abbrev:      "AWAY",
			Score:       2,
			ShotsOnGoal: 25,
		},
		PlayerByGameStats: nhl.PlayerGameStats{
			HomeTeam: nhl.TeamPlayerStats{
				Forwards: []nhl.PlayerStats{
					{
						Name:              nhl.LanguageNames{Default: "Player 1"},
						Goals:             1,
						Assists:           1,
						Points:            2,
						PlusMinus:         1,
						TOI:               "20:00",
						SOG:               3,
						Hits:              2,
						BlockedShots:      1,
						PIM:               2,
						FaceoffWinningPct: 60.0,
					},
				},
				Defense: []nhl.PlayerStats{
					{
						Name:         nhl.LanguageNames{Default: "Player 2"},
						Hits:         3,
						PIM:          0,
						BlockedShots: 2,
					},
				},
				Goalies: []nhl.GoalieGameStats{
					{
						Name:         nhl.LanguageNames{Default: "Goalie 1"},
						TOI:          "60:00",
						Saves:        28,
						GoalsAgainst: 2,
						SavePctg:     0.933,
						Decision:     "W",
					},
				},
			},
			AwayTeam: nhl.TeamPlayerStats{
				Forwards: []nhl.PlayerStats{
					{
						Name:              nhl.LanguageNames{Default: "Player 3"},
						Goals:             1,
						Assists:           0,
						Points:            1,
						PlusMinus:         -1,
						TOI:               "18:00",
						SOG:               2,
						Hits:              1,
						BlockedShots:      0,
						PIM:               4,
						FaceoffWinningPct: 40.0,
					},
				},
				Defense: []nhl.PlayerStats{
					{
						Name:         nhl.LanguageNames{Default: "Player 4"},
						Hits:         2,
						PIM:          2,
						BlockedShots: 1,
					},
				},
				Goalies: []nhl.GoalieGameStats{
					{
						Name:         nhl.LanguageNames{Default: "Goalie 2"},
						TOI:          "58:00",
						Saves:        27,
						GoalsAgainst: 3,
						SavePctg:     0.900,
						Decision:     "L",
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
		Plays: []nhl.PlayEvent{
			{
				TypeDescKey:   "shot-on-goal",
				TimeInPeriod:  "10:00",
				TimeRemaining: "10:00",
				PeriodDescriptor: nhl.PeriodDescriptor{
					Number: 1,
				},
				Details: nhl.EventDetails{
					ShootingPlayerID: 1,
					GoalieInNetID:    2,
				},
			},
		},
		RosterSpots: []nhl.RosterSpot{
			{
				PlayerID:      1,
				FirstName:     nhl.LanguageNames{Default: "Player"},
				LastName:      nhl.LanguageNames{Default: "One"},
				SweaterNumber: 91,
			},
			{
				PlayerID:      2,
				FirstName:     nhl.LanguageNames{Default: "Goalie"},
				LastName:      nhl.LanguageNames{Default: "One"},
				SweaterNumber: 31,
			},
		},
	}

	// Test display function (no error should occur)
	display.GamePlayByPlay(pbp)
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
