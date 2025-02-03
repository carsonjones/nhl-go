package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/internal/formatters"
	"go-nhl/nhl"
	"time"
)

// GetTodaysSchedule demonstrates retrieving and displaying the current day's schedule
func GetTodaysSchedule(client *nhl.Client) error {
	// Get today's schedule with default sort (ascending - earliest games first)
	scores, err := client.GetCurrentSchedule()
	if err != nil {
		return fmt.Errorf("error getting current schedule: %v", err)
	}
	fmt.Println("Games sorted by start time (earliest first - default):")
	display.Games(scores)
	return nil
}

// GetScheduleByDate demonstrates retrieving and displaying a schedule for a specific date
func GetScheduleByDate(client *nhl.Client, date string) error {
	// Example: Get schedule for a specific date with explicit descending sort
	scores, err := client.GetScheduleByDate(date, nhl.SortByDateDesc)
	if err != nil {
		return fmt.Errorf("error getting schedule for date %s: %v", date, err)
	}
	fmt.Println("\nGames sorted by start time (latest first):")
	display.Games(scores)
	return nil
}

// GetTeamSchedule demonstrates retrieving and displaying a team's schedule
func GetTeamSchedule(client *nhl.Client, teamIdentifier string) error {
	team, err := client.GetTeamByIdentifier(teamIdentifier)
	if err != nil {
		return fmt.Errorf("failed to get team: %v", err)
	}

	seasonID := formatters.GetCurrentSeasonID()
	schedule, err := client.GetTeamSchedule(team, seasonID)
	if err != nil {
		return fmt.Errorf("failed to get schedule: %v", err)
	}

	// Print schedule
	fmt.Printf("Schedule for %s (%d-%d):\n", team.Name.Default, seasonID/10000, (seasonID/10000)+1)
	for _, game := range schedule.Games {
		gameTime, err := time.Parse(time.RFC3339, game.StartTimeUTC)
		if err != nil {
			fmt.Printf("Error parsing game time: %v\n", err)
			continue
		}

		var opponentAbbrev string
		var location string
		if game.HomeTeam.Abbreviation == team.Abbreviation {
			opponentAbbrev = game.AwayTeam.Abbreviation
			location = "vs"
		} else {
			opponentAbbrev = game.HomeTeam.Abbreviation
			location = "@"
		}

		fmt.Printf("%s: %s %s\n", gameTime.Format("2006-01-02"), location, opponentAbbrev)
	}
	return nil
}
