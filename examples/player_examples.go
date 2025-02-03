package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/client"
)

// SearchPlayer demonstrates searching for players and displaying their basic info
func SearchPlayer(client *nhl.Client, searchName string) error {
	players, err := client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for player %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	fmt.Printf("\nFound %d players matching '%s':\n", len(players), searchName)
	for i, player := range players {
		fmt.Printf("%d. %s %s (#%d) - %s %s\n",
			i+1,
			player.FirstName.Default,
			player.LastName.Default,
			player.JerseyNumber,
			player.TeamAbbrev,
			player.Position)
	}

	// Get stats for the first player found
	player := players[0]
	stats, err := client.GetFilteredPlayerStats(player.PlayerID, nil)
	if err != nil {
		return fmt.Errorf("error getting stats for player %d: %v", player.PlayerID, err)
	}

	display.SeasonStats(stats, nhl.GameTypeRegularSeason)
	return nil
}

// SearchSkater demonstrates searching for skaters and displaying their stats
func SearchSkater(client *nhl.Client, searchName string) error {
	players, err := client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for skater %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	// Filter for skaters only
	var skaters []nhl.PlayerSearchResult
	for _, player := range players {
		if player.Position != "G" {
			skaters = append(skaters, player)
		}
	}

	if len(skaters) == 0 {
		fmt.Printf("No skaters found matching '%s'\n", searchName)
		return nil
	}

	fmt.Printf("\nFound %d skaters matching '%s':\n", len(skaters), searchName)
	for i, player := range skaters {
		fmt.Printf("%d. %s %s (#%d) - %s %s\n",
			i+1,
			player.FirstName.Default,
			player.LastName.Default,
			player.JerseyNumber,
			player.TeamAbbrev,
			player.Position)
	}

	// Get stats for the first skater found
	player := skaters[0]

	// Get regular season stats
	fmt.Println("\nRegular Season Stats:")
	stats, err := client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypeRegularSeason,
	})
	if err != nil {
		return fmt.Errorf("error getting regular season stats for skater %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypeRegularSeason)

	// Get playoff stats
	fmt.Println("\nPlayoff Stats:")
	stats, err = client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypePlayoffs,
	})
	if err != nil {
		return fmt.Errorf("error getting playoff stats for skater %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypePlayoffs)
	return nil
}

// TODO: showing player stats for goalie. Needs to be goalie stats
// SearchGoalie demonstrates searching for goalies and displaying their stats
func SearchGoalie(client *nhl.Client, searchName string) error {
	players, err := client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for goalie %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	// Filter for goalies only
	var goalies []nhl.PlayerSearchResult
	for _, player := range players {
		if player.Position == "G" {
			goalies = append(goalies, player)
		}
	}

	if len(goalies) == 0 {
		fmt.Printf("No goalies found matching '%s'\n", searchName)
		return nil
	}

	fmt.Printf("\nFound %d goalies matching '%s':\n", len(goalies), searchName)
	for i, player := range goalies {
		fmt.Printf("%d. %s %s (#%d) - %s %s\n",
			i+1,
			player.FirstName.Default,
			player.LastName.Default,
			player.JerseyNumber,
			player.TeamAbbrev,
			player.Position)
	}

	// Get stats for the first goalie found
	player := goalies[0]

	// Get regular season stats
	fmt.Println("\nRegular Season Stats:")
	stats, err := client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypeRegularSeason,
	})
	if err != nil {
		return fmt.Errorf("error getting regular season stats for goalie %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypeRegularSeason)

	// Get playoff stats
	fmt.Println("\nPlayoff Stats:")
	stats, err = client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypePlayoffs,
	})
	if err != nil {
		return fmt.Errorf("error getting playoff stats for goalie %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypePlayoffs)
	return nil
}
