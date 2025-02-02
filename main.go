package main

import (
	"fmt"
	"go-nhl/examples"
	"go-nhl/nhl"
	"log"
)

func main() {
	client := nhl.NewClient()

	// Example flags
	exampleGetCurrentSchedule := false
	exampleGetScheduleByDate := false
	exampleGetRoster := false
	examplePlayerSearch := false
	exampleSkaterSearch := false
	exampleGoalieSearch := false
	exampleSeasonStats := false
	exampleTeamSchedule := false
	exampleCurrentStandings := false
	exampleStandingsByDate := false
	exampleLeagueStandings := false
	exampleConferenceStandings := false
	exampleDivisionStandings := false
	exampleGameDetails := true

	// Track if any example was run
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
		// Use a completed game (2024020750 - NYR vs CHI on Feb 9, 2024)
		if err := examples.GetGameDetails(client, 2024020750); err != nil {
			log.Fatal(err)
		}
	}

	if !examplesRun {
		fmt.Println("Available examples (set the corresponding flag to true to run):")
		fmt.Println("- exampleGetCurrentSchedule: Get today's NHL schedule")
		fmt.Println("- exampleGetScheduleByDate: Get schedule for a specific date")
		fmt.Println("- exampleGetRoster: Get team rosters")
		fmt.Println("- examplePlayerSearch: Search for any player")
		fmt.Println("- exampleSkaterSearch: Search for skaters with detailed stats")
		fmt.Println("- exampleGoalieSearch: Search for goalies with detailed stats")
		fmt.Println("- exampleSeasonStats: Get player stats across seasons")
		fmt.Println("- exampleTeamSchedule: Get a team's full schedule")
		fmt.Println("- exampleCurrentStandings: Get current NHL standings")
		fmt.Println("- exampleStandingsByDate: Get NHL standings for a specific date")
		fmt.Println("- exampleLeagueStandings: Get overall NHL standings")
		fmt.Println("- exampleConferenceStandings: Get standings by conference")
		fmt.Println("- exampleDivisionStandings: Get standings by division")
		fmt.Println("- exampleGameDetails: Get detailed game information")
	}
}
