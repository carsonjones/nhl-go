package nhl_test

import (
	"go-nhl/nhl"
	"testing"
)

func TestGetTeamSchedule(t *testing.T) {
	client := nhl.NewClient()
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
