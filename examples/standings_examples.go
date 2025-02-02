package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/nhl"
)

// GetCurrentStandings demonstrates retrieving and displaying current NHL standings
func GetCurrentStandings(client *nhl.Client) error {
	standings, err := client.GetStandings()
	if err != nil {
		return fmt.Errorf("error getting current standings: %v", err)
	}

	fmt.Println("\nCurrent NHL Standings:")
	display.Standings(standings)
	return nil
}

// GetStandingsByDate demonstrates retrieving and displaying NHL standings for a specific date
func GetStandingsByDate(client *nhl.Client, date string) error {
	standings, err := client.GetStandingsByDate(date)
	if err != nil {
		return fmt.Errorf("error getting standings for date %s: %v", date, err)
	}

	fmt.Printf("\nNHL Standings for %s:\n", date)
	display.Standings(standings)
	return nil
}
