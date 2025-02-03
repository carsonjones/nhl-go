package display

import (
	"bytes"
	"fmt"
	"go-nhl/internal/formatters"
	"go-nhl/nhl"
	"io"
	"os"
	"strings"
	"testing"
)

func TestStatsLeaders(t *testing.T) {
	testCases := []struct {
		name     string
		seasonID int
	}{
		{
			name:     "Current season",
			seasonID: 0,
		},
		{
			name:     "Previous season",
			seasonID: 20222023,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test data
			leaders := &nhl.StatsLeadersResponse{
				Points: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Nikita"},
						LastName:   nhl.LanguageNames{Default: "Kucherov"},
						TeamAbbrev: "TBL",
						Value:      144,
					},
				},
				Goals: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Auston"},
						LastName:   nhl.LanguageNames{Default: "Matthews"},
						TeamAbbrev: "TOR",
						Value:      69,
					},
				},
				Assists: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Connor"},
						LastName:   nhl.LanguageNames{Default: "McDavid"},
						TeamAbbrev: "EDM",
						Value:      100,
					},
				},
				GoalsPp: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Sam"},
						LastName:   nhl.LanguageNames{Default: "Reinhart"},
						TeamAbbrev: "FLA",
						Value:      27,
					},
				},
				GoalsSh: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Travis"},
						LastName:   nhl.LanguageNames{Default: "Konecny"},
						TeamAbbrev: "PHI",
						Value:      6,
					},
				},
				PlusMinus: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Gustav"},
						LastName:   nhl.LanguageNames{Default: "Forsling"},
						TeamAbbrev: "FLA",
						Value:      56,
					},
				},
				FaceoffLeaders: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Nico"},
						LastName:   nhl.LanguageNames{Default: "Sturm"},
						TeamAbbrev: "SJS",
						Value:      0.601,
					},
				},
				TOI: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "John"},
						LastName:   nhl.LanguageNames{Default: "Carlson"},
						TeamAbbrev: "WSH",
						Value:      1553.5, // 25:53
					},
				},
				PenaltyMins: []nhl.StatsLeaderPlayer{
					{
						FirstName:  nhl.LanguageNames{Default: "Liam"},
						LastName:   nhl.LanguageNames{Default: "O'Brien"},
						TeamAbbrev: "ARI",
						Value:      153,
					},
				},
			}

			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call display function
			StatsLeaders(leaders, tc.seasonID)

			// Restore stdout
			w.Close()
			os.Stdout = old

			// Read captured output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Verify season header
			expectedSeason := tc.seasonID
			if expectedSeason == 0 {
				expectedSeason = formatters.GetCurrentSeasonID()
			}
			expectedHeader := fmt.Sprintf("NHL Stats Leaders (%s)", formatters.FormatSeasonID(expectedSeason))
			if !strings.Contains(output, expectedHeader) {
				t.Errorf("Expected output to contain header '%s', but it didn't", expectedHeader)
			}

			// Verify output contains expected data
			expectedStrings := []string{
				"Nikita Kucherov           TBL             144",
				"Auston Matthews           TOR              69",
				"Connor McDavid            EDM             100",
				"Sam Reinhart              FLA              27",
				"Travis Konecny            PHI               6",
				"Gustav Forsling           FLA              56",
				"Nico Sturm                SJS              60.1",
				"John Carlson              WSH             25:53",
				"Liam O'Brien              ARI             153",
			}

			for _, expected := range expectedStrings {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't", expected)
				}
			}
		})
	}
}

func TestStatsLeadersEmpty(t *testing.T) {
	// Test with empty response
	leaders := &nhl.StatsLeadersResponse{}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call display function
	StatsLeaders(leaders, 0)

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify season header
	expectedHeader := fmt.Sprintf("NHL Stats Leaders (%s)", formatters.FormatSeasonID(formatters.GetCurrentSeasonID()))
	if !strings.Contains(output, expectedHeader) {
		t.Errorf("Expected output to contain header '%s', but it didn't", expectedHeader)
	}

	// Verify headers are still present
	expectedHeaders := []string{
		"Points Leaders",
		"Goals Leaders",
		"Assists Leaders",
		"Power Play Goals Leaders",
		"Short-handed Goals Leaders",
		"Plus/Minus Leaders",
		"Faceoff Leaders",
		"Time on Ice Leaders",
		"Penalty Minutes Leaders",
	}

	for _, header := range expectedHeaders {
		if !strings.Contains(output, header) {
			t.Errorf("Expected output to contain header '%s', but it didn't", header)
		}
	}
}
