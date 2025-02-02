package display

import (
	"fmt"
	"go-nhl/internal/formatters"
	"go-nhl/nhl"
	"sort"
)

// PlayerStats displays player statistics based on their type
func PlayerStats(stats interface{}) {
	switch s := stats.(type) {
	case *nhl.SkaterStatsResponse:
		SkaterStats(s)
	case *nhl.GoalieStatsResponse:
		GoalieStats(s)
	default:
		fmt.Printf("Unknown stats type: %T\n", stats)
	}
}

// SkaterStats displays statistics for a skater
func SkaterStats(stats *nhl.SkaterStatsResponse) {
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
	fmt.Printf("\nStats for %s (%s):\n", current.FullName, formatters.FormatSeasonID(current.SeasonID))
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

// GoalieStats displays statistics for a goalie
func GoalieStats(stats *nhl.GoalieStatsResponse) {
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
	fmt.Printf("\nStats for %s (%s):\n", current.FullName, formatters.FormatSeasonID(current.SeasonID))
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
	fmt.Printf("Time on Ice: %s\n", formatters.FormatTimeOnIce(current.TimeOnIce))

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

// SeasonStats displays season statistics for a player
func SeasonStats(stats []nhl.SeasonTotal, gameType nhl.GameType) {
	if len(stats) == 0 {
		fmt.Println("No stats available")
		return
	}

	// Display the most recent season's stats
	current := stats[0]
	fmt.Printf("\nStats for %s (%d-%d):\n", current.TeamName.Default, current.Season/10000, (current.Season/10000)+1)
	fmt.Printf("Game Type: %s\n", GetGameTypeName(gameType))

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

		fmt.Printf("\nCareer %s Totals (%d seasons):\n", GetGameTypeName(gameType), len(stats))
		fmt.Printf("Games: %d\n", totalGames)
		fmt.Printf("Goals: %d\n", totalGoals)
		fmt.Printf("Assists: %d\n", totalAssists)
		fmt.Printf("Points: %d (%.2f per game)\n",
			totalPoints, float64(totalPoints)/float64(totalGames))
	}
}

// GetGameTypeName returns a human-readable name for a game type
func GetGameTypeName(gameType nhl.GameType) string {
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
