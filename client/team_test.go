package nhl_test

import (
	"go-nhl/client"
	"testing"
)

func TestGetTeams(t *testing.T) {
	client := nhl.NewClient()
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

func TestGetTeamByIdentifier(t *testing.T) {
	client := nhl.NewClient()

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
