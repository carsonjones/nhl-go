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
	exampleGoalieSearch := true
	exampleSeasonStats := false
	exampleTeamSchedule := false

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
	}
}
