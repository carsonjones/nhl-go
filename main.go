package main

import (
	"fmt"
	"go-nhl/nhl"
	"sort"
	"strings"
	"time"
)

func formatGameTime(utcTime string) (string, error) {
	// Parse the UTC time
	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %v", err)
	}

	// Load EST location
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return "", fmt.Errorf("error loading EST location: %v", err)
	}

	// Load CST location
	cst, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return "", fmt.Errorf("error loading CST location: %v", err)
	}

	// Convert to EST and CST
	estTime := t.In(est)
	cstTime := t.In(cst)

	// Format the times
	return fmt.Sprintf("%s EST (%s CST)", 
		estTime.Format("3:04 PM"),
		cstTime.Format("3:04 PM")), nil
}

func displayGames(games *nhl.FilteredScoreboardResponse) {
	fmt.Printf("\nGames for %s:\n", games.Date)
	for _, game := range games.Games {
		gameTime, err := formatGameTime(game.StartTimeUTC)
		if err != nil {
			fmt.Printf("Error formatting game time: %v\n", err)
			continue
		}

		fmt.Printf("%s at %s - %s\n", 
			game.AwayTeam.Name.Default,
			game.HomeTeam.Name.Default,
			gameTime)
		
		if game.GameState == "LIVE" || game.GameState == "OFF" {
			fmt.Printf("Score: %s %d, %s %d\n",
				game.AwayTeam.Name.Default, game.AwayTeam.Score,
				game.HomeTeam.Name.Default, game.HomeTeam.Score)
			
			if game.GameState == "LIVE" {
				fmt.Printf("Period: %d (%s)\n", 
					game.Period,
					game.PeriodDescriptor.PeriodType)
			}
		} else {
			fmt.Printf("Game Status: %s\n", game.GameState)
		}
		fmt.Println()
	}
}

func displayRoster(roster *nhl.RosterResponse, teamAbbr string) {
	fmt.Printf("\nRoster for %s:\n\n", teamAbbr)

	fmt.Println("Forwards:")
	fmt.Println("---------")
	for _, player := range roster.Forwards {
		fmt.Printf("#%d %s %s - %s\n",
			player.JerseyNumber,
			player.FirstName.Default,
			player.LastName.Default,
			player.Position)
	}

	fmt.Println("\nDefensemen:")
	fmt.Println("-----------")
	for _, player := range roster.Defensemen {
		fmt.Printf("#%d %s %s - %s\n",
			player.JerseyNumber,
			player.FirstName.Default,
			player.LastName.Default,
			player.Position)
	}

	fmt.Println("\nGoalies:")
	fmt.Println("--------")
	for _, player := range roster.Goalies {
		fmt.Printf("#%d %s %s - %s\n",
			player.JerseyNumber,
			player.FirstName.Default,
			player.LastName.Default,
			player.Position)
	}
}

func formatSeasonID(seasonID int) string {
	start := seasonID / 10000
	return fmt.Sprintf("%d-%d", start, start+1)
}

func formatTimeOnIce(seconds int) string {
	minutes := seconds / 60
	return fmt.Sprintf("%d:%02d", minutes/60, minutes%60)
}

func displayPlayerStats(stats interface{}) {
	switch s := stats.(type) {
	case *nhl.SkaterStatsResponse:
		displaySkaterStats(s)
	case *nhl.GoalieStatsResponse:
		displayGoalieStats(s)
	default:
		fmt.Printf("Unknown stats type: %T\n", stats)
	}
}

func displaySkaterStats(stats *nhl.SkaterStatsResponse) {
	if len(stats.Data) == 0 {
		fmt.Println("No stats available")
		return
	}

	// Sort seasons in descending order (most recent first)
	sort.Slice(stats.Data, func(i, j int) bool {
		return stats.Data[i].SeasonID > stats.Data[j].SeasonID
	})

	// Display the most recent season's stats
	current := stats.Data[0]
	fmt.Printf("\nStats for %s (%s):\n", current.FullName, formatSeasonID(current.SeasonID))
	fmt.Printf("Team: %s\n", current.TeamAbbrev)
	fmt.Printf("Position: %s\n", current.PositionCode)
	fmt.Printf("Shoots/Catches: %s\n", current.ShootsCatches)

	fmt.Printf("\nCurrent Season Stats:\n")
	fmt.Printf("Games Played: %d\n", current.GamesPlayed)
	fmt.Printf("Goals: %d\n", current.Goals)
	fmt.Printf("Assists: %d\n", current.Assists)
	fmt.Printf("Points: %d (%.2f per game)\n", current.Points, current.PointsPerGame)
	fmt.Printf("Plus/Minus: %d\n", current.PlusMinus)
	fmt.Printf("PIM: %d\n", current.PenaltyMinutes)

	fmt.Printf("\nScoring Breakdown:\n")
	fmt.Printf("Even Strength: %d goals, %d points\n", current.EvenStrengthGoals, current.EvenStrengthPoints)
	fmt.Printf("Power Play: %d goals, %d points\n", current.PowerPlayGoals, current.PowerPlayPoints)
	fmt.Printf("Short Handed: %d goals, %d points\n", current.ShortHandedGoals, current.ShortHandedPoints)
	fmt.Printf("Game Winners: %d\n", current.GameWinningGoals)
	fmt.Printf("Overtime Goals: %d\n", current.OvertimeGoals)

	fmt.Printf("\nShooting:\n")
	fmt.Printf("Shots: %d\n", current.Shots)
	fmt.Printf("Shooting %%: %.1f\n", current.ShootingPct)
	fmt.Printf("TOI/Game: %.1f\n", current.TimeOnIcePerGame)
	
	if current.FaceoffWinPct > 0 {
		fmt.Printf("Faceoff Win %%: %.1f\n", current.FaceoffWinPct)
	}

	// Show career stats summary
	if len(stats.Data) > 1 {
		var totalGames, totalGoals, totalAssists, totalPoints int
		for _, season := range stats.Data {
			totalGames += season.GamesPlayed
			totalGoals += season.Goals
			totalAssists += season.Assists
			totalPoints += season.Points
		}

		fmt.Printf("\nCareer Totals (%d seasons):\n", len(stats.Data))
		fmt.Printf("Games: %d\n", totalGames)
		fmt.Printf("Goals: %d\n", totalGoals)
		fmt.Printf("Assists: %d\n", totalAssists)
		fmt.Printf("Points: %d (%.2f per game)\n", 
			totalPoints, float64(totalPoints)/float64(totalGames))
	}
}

func displayGoalieStats(stats *nhl.GoalieStatsResponse) {
	if len(stats.Data) == 0 {
		fmt.Println("No stats available")
		return
	}

	// Sort seasons in descending order (most recent first)
	sort.Slice(stats.Data, func(i, j int) bool {
		return stats.Data[i].SeasonID > stats.Data[j].SeasonID
	})

	// Display the most recent season's stats
	current := stats.Data[0]
	fmt.Printf("\nStats for %s (%s):\n", current.FullName, formatSeasonID(current.SeasonID))
	fmt.Printf("Team: %s\n", current.TeamAbbrev)
	fmt.Printf("Catches: %s\n", current.ShootsCatches)

	fmt.Printf("\nCurrent Season Stats:\n")
	fmt.Printf("Games Played: %d\n", current.GamesPlayed)
	fmt.Printf("Games Started: %d\n", current.GamesStarted)
	fmt.Printf("Record: %d-%d-%d\n", current.Wins, current.Losses, current.OvertimeLosses)
	fmt.Printf("Goals Against Average: %.2f\n", current.GoalsAgainstAvg)
	fmt.Printf("Save Percentage: %.3f\n", current.SavePctg)
	fmt.Printf("Shutouts: %d\n", current.Shutouts)

	fmt.Printf("\nDetailed Stats:\n")
	fmt.Printf("Shots Against: %d\n", current.ShotsAgainst)
	fmt.Printf("Saves: %d\n", current.Saves)
	fmt.Printf("Goals Against: %d\n", current.GoalsAgainst)
	fmt.Printf("Time on Ice: %s\n", formatTimeOnIce(current.TimeOnIce))
	
	if current.Points > 0 {
		fmt.Printf("\nScoring:\n")
		fmt.Printf("Goals: %d\n", current.Goals)
		fmt.Printf("Assists: %d\n", current.Assists)
		fmt.Printf("Points: %d\n", current.Points)
	}

	// Show career stats summary
	if len(stats.Data) > 1 {
		var totalGames, totalWins, totalLosses, totalOTL int
		var totalShots, totalSaves, totalGoalsAgainst, totalShutouts int
		totalTimeOnIce := 0

		for _, season := range stats.Data {
			totalGames += season.GamesPlayed
			totalWins += season.Wins
			totalLosses += season.Losses
			totalOTL += season.OvertimeLosses
			totalShots += season.ShotsAgainst
			totalSaves += season.Saves
			totalGoalsAgainst += season.GoalsAgainst
			totalShutouts += season.Shutouts
			totalTimeOnIce += season.TimeOnIce
		}

		careerSavePct := float64(totalSaves) / float64(totalShots)
		careerGAA := float64(totalGoalsAgainst*3600) / float64(totalTimeOnIce)

		fmt.Printf("\nCareer Totals (%d seasons):\n", len(stats.Data))
		fmt.Printf("Games Played: %d\n", totalGames)
		fmt.Printf("Record: %d-%d-%d\n", totalWins, totalLosses, totalOTL)
		fmt.Printf("Save Percentage: %.3f\n", careerSavePct)
		fmt.Printf("Goals Against Average: %.2f\n", careerGAA)
		fmt.Printf("Shutouts: %d\n", totalShutouts)
	}
}

func displaySeasonStats(stats []nhl.SeasonTotal, gameType nhl.GameType) {
	if len(stats) == 0 {
		fmt.Println("No stats available")
		return
	}

	// Display the most recent season's stats
	current := stats[0]
	fmt.Printf("\nStats for %s (%d-%d):\n", current.TeamName.Default, current.Season/10000, (current.Season/10000)+1)
	fmt.Printf("Game Type: %s\n", getGameTypeName(gameType))

	fmt.Printf("\nCurrent Season Stats:\n")
	fmt.Printf("Games Played: %d\n", current.GamesPlayed)
	fmt.Printf("Goals: %d\n", current.Goals)
	fmt.Printf("Assists: %d\n", current.Assists)
	fmt.Printf("Points: %d (%.2f per game)\n", current.Points, float64(current.Points)/float64(current.GamesPlayed))
	fmt.Printf("Plus/Minus: %d\n", current.PlusMinus)
	fmt.Printf("PIM: %d\n", current.PenaltyMinutes)

	fmt.Printf("\nScoring Breakdown:\n")
	fmt.Printf("Power Play Goals: %d\n", current.PowerPlayGoals)
	fmt.Printf("Power Play Points: %d\n", current.PowerPlayPoints)
	fmt.Printf("Short Handed Goals: %d\n", current.ShorthandedGoals)
	fmt.Printf("Short Handed Points: %d\n", current.ShorthandedPoints)
	fmt.Printf("Game Winners: %d\n", current.GameWinningGoals)
	fmt.Printf("Overtime Goals: %d\n", current.OTGoals)

	fmt.Printf("\nShooting:\n")
	fmt.Printf("Shots: %d\n", current.Shots)
	fmt.Printf("Shooting %%: %.1f\n", current.ShootingPctg)
	fmt.Printf("TOI/Game: %s\n", current.AvgTOI)
	
	if current.FaceoffWinningPctg > 0 {
		fmt.Printf("Faceoff Win %%: %.1f\n", current.FaceoffWinningPctg)
	}

	// Show career stats summary
	if len(stats) > 1 {
		var totalGames, totalGoals, totalAssists, totalPoints int
		for _, season := range stats {
			totalGames += season.GamesPlayed
			totalGoals += season.Goals
			totalAssists += season.Assists
			totalPoints += season.Points
		}

		fmt.Printf("\nCareer %s Totals (%d seasons):\n", getGameTypeName(gameType), len(stats))
		fmt.Printf("Games: %d\n", totalGames)
		fmt.Printf("Goals: %d\n", totalGoals)
		fmt.Printf("Assists: %d\n", totalAssists)
		fmt.Printf("Points: %d (%.2f per game)\n", 
			totalPoints, float64(totalPoints)/float64(totalGames))
	}
}

func getGameTypeName(gameType nhl.GameType) string {
	switch gameType {
	case nhl.GameTypeRegularSeason:
		return "Regular Season"
	case nhl.GameTypePlayoffs:
		return "Playoff"
	case nhl.GameTypeAllStar:
		return "All-Star"
	default:
		return "Unknown"
	}
}

func getCurrentSeasonID() int {
	now := time.Now()
	year := now.Year()
	// NHL season typically starts in October
	if now.Month() < time.July {
		year-- // If we're in the first half of the year, we're in the previous year's season
	}
	return year * 10000
}

func main() {
	client := nhl.NewClient()

	exampleGetCurrentSchedule := false
	exampleGetScheduleByDate := false
	exampleGetRoster := false
	examplePlayerSearch := false
	exampleSkaterSearch := false
	exampleGoalieSearch := false
	exampleSeasonStats := true

	if exampleGetCurrentSchedule {
		// Get today's schedule with default sort (ascending - earliest games first)
		scores, err := client.GetCurrentSchedule()
		if err != nil {
			fmt.Printf("Error getting current schedule: %v\n", err)
			return
		}
		fmt.Println("Games sorted by start time (earliest first - default):")
		displayGames(scores)
	}

	if exampleGetScheduleByDate {
		// Example: Get schedule for a specific date with explicit descending sort
		date := "2025-02-01"
		scores, err := client.GetScheduleByDate(date, nhl.SortByDateDesc)
		if err != nil {
			fmt.Printf("Error getting schedule for date %s: %v\n", date, err)
			return
		}
		fmt.Println("\nGames sorted by start time (latest first):")
		displayGames(scores)
	}

	if exampleGetRoster {
		// Example: Get roster for teams using different identifier types
		identifiers := []string{
			"TOR",                  // by abbreviation
			"Montreal Canadiens",   // by full name
			"6",                    // by ID (Boston Bruins)
		}
		for _, identifier := range identifiers {
			roster, err := client.GetTeamRoster(identifier)
			if err != nil {
				fmt.Printf("Error getting roster for %s: %v\n", identifier, err)
				continue
			}
			displayRoster(roster, identifier)
			fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
		}
	}

	if examplePlayerSearch {
		// Example: Generic player search (will be replaced by specific skater/goalie searches)
		searchName := "Matthews"
		players, err := client.SearchPlayer(searchName)
		if err != nil {
			fmt.Printf("Error searching for player %s: %v\n", searchName, err)
			return
		}

		if len(players) == 0 {
			fmt.Printf("No players found matching '%s'\n", searchName)
			return
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
			fmt.Printf("Error getting stats for player %d: %v\n", player.PlayerID, err)
			return
		}

		displaySeasonStats(stats, nhl.GameTypeRegularSeason)
	}

	if exampleSkaterSearch {
		// Example: Search for a skater and display their stats
		searchName := "Matthews"  // Search for Auston Matthews
		players, err := client.SearchPlayer(searchName)
		if err != nil {
			fmt.Printf("Error searching for skater %s: %v\n", searchName, err)
			return
		}

		if len(players) == 0 {
			fmt.Printf("No players found matching '%s'\n", searchName)
			return
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
			return
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
			fmt.Printf("Error getting regular season stats for skater %d: %v\n", player.PlayerID, err)
			return
		}
		displaySeasonStats(stats, nhl.GameTypeRegularSeason)

		// Get playoff stats
		fmt.Println("\nPlayoff Stats:")
		stats, err = client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
			GameType: nhl.GameTypePlayoffs,
		})
		if err != nil {
			fmt.Printf("Error getting playoff stats for skater %d: %v\n", player.PlayerID, err)
			return
		}
		displaySeasonStats(stats, nhl.GameTypePlayoffs)
	}

	if exampleGoalieSearch {
		// Example: Search for a goalie and display their stats
		searchName := "Woll"  // Search for Joseph Woll
		players, err := client.SearchPlayer(searchName)
		if err != nil {
			fmt.Printf("Error searching for goalie %s: %v\n", searchName, err)
			return
		}

		if len(players) == 0 {
			fmt.Printf("No players found matching '%s'\n", searchName)
			return
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
			return
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
			fmt.Printf("Error getting regular season stats for goalie %d: %v\n", player.PlayerID, err)
			return
		}
		displaySeasonStats(stats, nhl.GameTypeRegularSeason)

		// Get playoff stats
		fmt.Println("\nPlayoff Stats:")
		stats, err = client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
			GameType: nhl.GameTypePlayoffs,
		})
		if err != nil {
			fmt.Printf("Error getting playoff stats for goalie %d: %v\n", player.PlayerID, err)
			return
		}
		displaySeasonStats(stats, nhl.GameTypePlayoffs)
	}

	if exampleSeasonStats {
		// Example: Get stats for a specific season
		searchName := "Matthews"  // Search for Auston Matthews
		players, err := client.SearchPlayer(searchName)
		if err != nil {
			fmt.Printf("Error searching for player %s: %v\n", searchName, err)
			return
		}

		if len(players) == 0 {
			fmt.Printf("No players found matching '%s'\n", searchName)
			return
		}

		// Get the first matching player
		player := players[0]
		fmt.Printf("\nShowing stats for %s %s:\n",
			player.FirstName.Default,
			player.LastName.Default)

		// Get all seasons to show what's available
		allStats, err := client.GetFilteredPlayerStats(player.PlayerID, nil)
		if err != nil {
			fmt.Printf("Error getting player stats: %v\n", err)
			return
		}

		fmt.Println("\nAvailable NHL Seasons:")
		for _, season := range allStats {
			fmt.Printf("- %d-%d (%s): %d games played, %d goals, %d points\n", 
				season.Season/10000, 
				(season.Season/10000)+1,
				getGameTypeName(nhl.GameType(season.GameTypeID)),
				season.GamesPlayed,
				season.Goals,
				season.Points)
		}

		// Calculate current and previous season IDs
		currentSeasonID := getCurrentSeasonID()
		previousSeasonID := currentSeasonID - 10000

		// Example seasons to show
		seasons := []struct {
			seasonID int
			gameType nhl.GameType
		}{
			{currentSeasonID, nhl.GameTypeRegularSeason},    // Current season
			{previousSeasonID, nhl.GameTypeRegularSeason},   // Previous season
			{previousSeasonID, nhl.GameTypePlayoffs},        // Previous season playoffs
		}

		// Show stats for each season
		for _, s := range seasons {
			fmt.Printf("\nStats for %d-%d %s:\n", 
				s.seasonID/10000, 
				(s.seasonID/10000)+1,
				getGameTypeName(s.gameType))
			
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

			displaySeasonStats(stats, s.gameType)
		}

		// Example: Compare specific stat across seasons
		fmt.Printf("\nGoals per season (Regular Season):\n")
		regularSeasonStats, err := client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
			GameType: nhl.GameTypeRegularSeason,
		})
		if err != nil {
			fmt.Printf("Error getting regular season stats: %v\n", err)
			return
		}

		for _, season := range regularSeasonStats {
			fmt.Printf("%d-%d: %d goals in %d games (%.2f goals per game)\n",
				season.Season/10000,
				(season.Season/10000)+1,
				season.Goals,
				season.GamesPlayed,
				float64(season.Goals)/float64(season.GamesPlayed))
		}
	}
}
