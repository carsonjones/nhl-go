package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/client"
	"strings"
)

// GetTeamRoster demonstrates retrieving and displaying a team's roster
func GetTeamRoster(client *nhl.Client) error {
	// Example: Get roster for teams using different identifier types
	identifiers := []string{
		"DAL",                // by abbreviation
		"Montreal Canadiens", // by full name
		"6",                  // by ID (Boston Bruins)
	}
	for _, identifier := range identifiers {
		roster, err := client.GetTeamRoster(identifier)
		if err != nil {
			fmt.Printf("Error getting roster for %s: %v\n", identifier, err)
			continue
		}
		display.Roster(roster, identifier)
		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}
	return nil
}
