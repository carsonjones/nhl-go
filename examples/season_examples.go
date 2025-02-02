package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/internal/formatters"
	"go-nhl/nhl"
)

// GetSeasonStats demonstrates retrieving and displaying season stats for a player
func GetSeasonStats(client *nhl.Client, searchName string) error {
	players, err := client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for player %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	// Get the first matching player
	player := players[0]
	fmt.Printf("\nShowing stats for %s %s:\n",
		player.FirstName.Default,
		player.LastName.Default)

	// Get all seasons to show what's available
	allStats, err := client.GetFilteredPlayerStats(player.PlayerID, nil)
	if err != nil {
		return fmt.Errorf("error getting player stats: %v", err)
	}

	fmt.Println("\nAvailable NHL Seasons:")
	for _, season := range allStats {
		fmt.Printf("- %d-%d (%s): %d games played, %d goals, %d points\n",
			season.Season/10000,
			(season.Season/10000)+1,
			display.GetGameTypeName(nhl.GameType(season.GameTypeID)),
			season.GamesPlayed,
			season.Goals,
			season.Points)
	}

	// Calculate current and previous season IDs
	currentSeasonID := formatters.GetCurrentSeasonID()
	previousSeasonID := currentSeasonID - 10000

	// Example seasons to show
	seasons := []struct {
		seasonID int
		gameType nhl.GameType
	}{
		{currentSeasonID, nhl.GameTypeRegularSeason},  // Current season
		{previousSeasonID, nhl.GameTypeRegularSeason}, // Previous season
		{previousSeasonID, nhl.GameTypePlayoffs},      // Previous season playoffs
	}

	// Show stats for each season
	for _, s := range seasons {
		fmt.Printf("\nStats for %d-%d %s:\n",
			s.seasonID/10000,
			(s.seasonID/10000)+1,
			display.GetGameTypeName(s.gameType))

		stats, err := client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
			GameType: s.gameType,
			SeasonID: s.seasonID,
		})
		if err != nil {
			fmt.Printf("Error getting stats: %v\n", err)
			continue
		}

		if len(stats) == 0 {
			fmt.Println("No stats available")
			continue
		}

		display.SeasonStats(stats, s.gameType)
	}

	// Example: Compare specific stat across seasons
	fmt.Printf("\nGoals per season (Regular Season):\n")
	regularSeasonStats, err := client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypeRegularSeason,
	})
	if err != nil {
		return fmt.Errorf("error getting regular season stats: %v", err)
	}

	for _, season := range regularSeasonStats {
		fmt.Printf("%d-%d: %d goals in %d games (%.2f goals per game)\n",
			season.Season/10000,
			(season.Season/10000)+1,
			season.Goals,
			season.GamesPlayed,
			float64(season.Goals)/float64(season.GamesPlayed))
	}

	return nil
}
