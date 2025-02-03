package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/nhl"
)

// GetLeagueLeaders demonstrates retrieving and displaying NHL stats leaders
func GetLeagueLeaders(client *nhl.Client) error {
	// Get current season leaders
	leaders, err := client.GetStatsLeaders(0)
	if err != nil {
		return fmt.Errorf("error getting stats leaders: %v", err)
	}

	display.StatsLeaders(leaders, 0)

	// Get previous season leaders
	prevSeasonLeaders, err := client.GetStatsLeaders(20222023)
	if err != nil {
		return fmt.Errorf("error getting previous season stats leaders: %v", err)
	}

	display.StatsLeaders(prevSeasonLeaders, 20222023)
	return nil
}
