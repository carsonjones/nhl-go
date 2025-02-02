package display

import (
	"fmt"
	"go-nhl/nhl"
)

// Roster displays a team's roster
func Roster(roster *nhl.RosterResponse, teamAbbr string) {
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
