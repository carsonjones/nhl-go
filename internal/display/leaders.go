package display

import (
	"fmt"
	"go-nhl/internal/formatters"
	"go-nhl/client"
	"strings"
)

// StatsLeaders displays the NHL stats leaders
func StatsLeaders(leaders *nhl.StatsLeadersResponse, seasonID int) {
	// If seasonID is 0, use current season
	if seasonID == 0 {
		seasonID = formatters.GetCurrentSeasonID()
	}

	fmt.Printf("\nNHL Stats Leaders (%s)\n", formatters.FormatSeasonID(seasonID))
	fmt.Println("================")

	// Points Leaders
	fmt.Println("\nPoints Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "PTS")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.Points {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}

	// Goals Leaders
	fmt.Println("\nGoals Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "G")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.Goals {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}

	// Assists Leaders
	fmt.Println("\nAssists Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "A")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.Assists {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}

	// Power Play Goals Leaders
	fmt.Println("\nPower Play Goals Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "PPG")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.GoalsPp {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}

	// Short-handed Goals Leaders
	fmt.Println("\nShort-handed Goals Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "SHG")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.GoalsSh {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}

	// Plus/Minus Leaders
	fmt.Println("\nPlus/Minus Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "+/-")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.PlusMinus {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}

	// Faceoff Leaders
	fmt.Println("\nFaceoff Leaders")
	fmt.Printf("%-25s %-15s %5s\n", "Player", "Team", "FO%")
	fmt.Println(strings.Repeat("-", 47))
	for _, player := range leaders.FaceoffLeaders {
		fmt.Printf("%-25s %-15s %5.1f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value*100)
	}

	// Time on Ice Leaders
	fmt.Println("\nTime on Ice Leaders")
	fmt.Printf("%-25s %-15s %8s\n", "Player", "Team", "TOI")
	fmt.Println(strings.Repeat("-", 50))
	for _, player := range leaders.TOI {
		minutes := int(player.Value / 60)
		seconds := int(player.Value) % 60
		fmt.Printf("%-25s %-15s %02d:%02d\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			minutes,
			seconds)
	}

	// Penalty Minutes Leaders
	fmt.Println("\nPenalty Minutes Leaders")
	fmt.Printf("%-25s %-15s %3s\n", "Player", "Team", "PIM")
	fmt.Println(strings.Repeat("-", 45))
	for _, player := range leaders.PenaltyMins {
		fmt.Printf("%-25s %-15s %3.0f\n",
			fmt.Sprintf("%s %s", player.FirstName.Default, player.LastName.Default),
			player.TeamAbbrev,
			player.Value)
	}
}
