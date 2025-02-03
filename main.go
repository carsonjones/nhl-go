package main

import (
	"flag"
	"fmt"
	"go-nhl/examples"
	"go-nhl/nhl"
	"log"
	"time"
)

func main() {
	// Command line flags
	var (
		exampleGetCurrentSchedule  bool
		exampleGetScheduleByDate   bool
		exampleGetRoster           bool
		examplePlayerSearch        bool
		exampleSkaterSearch        bool
		exampleGoalieSearch        bool
		exampleSeasonStats         bool
		exampleTeamSchedule        bool
		exampleCurrentStandings    bool
		exampleStandingsByDate     bool
		exampleLeagueStandings     bool
		exampleConferenceStandings bool
		exampleDivisionStandings   bool
		exampleGameDetails         bool
		exampleLiveUpdates         bool
		gameID                     int
		updateInterval             int
	)

	flag.BoolVar(&exampleGetCurrentSchedule, "current-schedule", false, "Get today's NHL schedule")
	flag.BoolVar(&exampleGetScheduleByDate, "schedule-by-date", false, "Get schedule for a specific date")
	flag.BoolVar(&exampleGetRoster, "get-roster", false, "Get team rosters")
	flag.BoolVar(&examplePlayerSearch, "player-search", false, "Search for any player")
	flag.BoolVar(&exampleSkaterSearch, "skater-search", false, "Search for skaters with detailed stats")
	flag.BoolVar(&exampleGoalieSearch, "goalie-search", false, "Search for goalies with detailed stats")
	flag.BoolVar(&exampleSeasonStats, "season-stats", false, "Get player stats across seasons")
	flag.BoolVar(&exampleTeamSchedule, "team-schedule", false, "Get a team's full schedule")
	flag.BoolVar(&exampleCurrentStandings, "current-standings", false, "Get current NHL standings")
	flag.BoolVar(&exampleStandingsByDate, "standings-by-date", false, "Get NHL standings for a specific date")
	flag.BoolVar(&exampleLeagueStandings, "league-standings", false, "Get overall NHL standings")
	flag.BoolVar(&exampleConferenceStandings, "conference", false, "Get standings by conference")
	flag.BoolVar(&exampleDivisionStandings, "division", false, "Get standings by division")
	flag.BoolVar(&exampleGameDetails, "game", false, "Get detailed game information")
	flag.IntVar(&gameID, "game-id", 2024020750, "Game ID for game details (default: NYR vs CHI on Feb 9, 2024)")
	flag.BoolVar(&exampleLiveUpdates, "live", false, "Show live game updates")
	flag.IntVar(&updateInterval, "interval", 60, "Update interval in seconds for live updates")
	flag.Parse()

	// Create NHL client
	client := nhl.NewClient()

	// Track if any examples were run
	examplesRun := false

	if exampleGetCurrentSchedule {
		examplesRun = true
		if err := examples.GetCurrentSchedule(client); err != nil {
			log.Fatal(err)
		}
	}

	if exampleGetScheduleByDate {
		examplesRun = true
		if err := examples.GetScheduleByDate(client, "2025-02-01"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleGetRoster {
		examplesRun = true
		if err := examples.GetTeamRoster(client); err != nil {
			log.Fatal(err)
		}
	}

	if examplePlayerSearch {
		examplesRun = true
		if err := examples.SearchPlayer(client, "Robertson"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleSkaterSearch {
		examplesRun = true
		if err := examples.SearchSkater(client, "Hintz"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleGoalieSearch {
		examplesRun = true
		if err := examples.SearchGoalie(client, "Oettinger"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleSeasonStats {
		examplesRun = true
		if err := examples.GetSeasonStats(client, "Johnston"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleTeamSchedule {
		examplesRun = true
		if err := examples.GetTeamSchedule(client, "DAL"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleCurrentStandings {
		examplesRun = true
		if err := examples.GetCurrentStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if exampleStandingsByDate {
		examplesRun = true
		if err := examples.GetStandingsByDate(client, "2024-02-01"); err != nil {
			log.Fatal(err)
		}
	}

	if exampleLeagueStandings {
		examplesRun = true
		if err := examples.GetLeagueStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if exampleConferenceStandings {
		examplesRun = true
		if err := examples.GetConferenceStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if exampleDivisionStandings {
		examplesRun = true
		if err := examples.GetDivisionStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if exampleGameDetails {
		examplesRun = true
		if err := examples.GetGameDetails(client, gameID); err != nil {
			log.Fatal(err)
		}
	}

	if exampleLiveUpdates {
		examplesRun = true
		fmt.Printf("Starting live game updates (refreshing every %d seconds). Press Ctrl+C to stop.\n", updateInterval)
		for {
			if err := examples.GetLiveGameUpdates(client); err != nil {
				log.Printf("Error getting live updates: %v", err)
			}
			time.Sleep(time.Duration(updateInterval) * time.Second)
		}
	}

	if !examplesRun {
		fmt.Println("Available examples (use -h flag to see all options):")
		fmt.Println("- current-schedule: Get today's NHL schedule")
		fmt.Println("- schedule-by-date: Get schedule for a specific date")
		fmt.Println("- get-roster: Get team rosters")
		fmt.Println("- player-search: Search for any player")
		fmt.Println("- skater-search: Search for skaters with detailed stats")
		fmt.Println("- goalie-search: Search for goalies with detailed stats")
		fmt.Println("- season-stats: Get player stats across seasons")
		fmt.Println("- team-schedule: Get a team's full schedule")
		fmt.Println("- current-standings: Get current NHL standings")
		fmt.Println("- standings-by-date: Get NHL standings for a specific date")
		fmt.Println("- league-standings: Get overall NHL standings")
		fmt.Println("- conference: Get standings by conference")
		fmt.Println("- division: Get standings by division")
		fmt.Println("- game: Get detailed game information")
		fmt.Println("- live: Show live game updates")
	}
}
