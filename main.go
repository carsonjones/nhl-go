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
		// Flags
		todaysSchedule      bool
		slate               bool
		roster              bool
		playerSearch        bool
		skaterSearch        bool
		goalieSearch        bool
		stats               bool
		schedule            bool
		standings           bool
		standingsByDate     bool
		leagueStandings     bool
		conferenceStandings bool
		divisionStandings   bool
		gameDetails         bool
		liveUpdates         bool
		leaders             bool
		// Parameters
		date           string
		name           string
		gameID         int
		updateInterval int
	)

	flag.BoolVar(&todaysSchedule, "today", false, "Get today's NHL schedule")
	flag.BoolVar(&slate, "slate", false, "Get schedule for a specific date")
	flag.BoolVar(&roster, "roster", false, "Get team rosters")
	flag.BoolVar(&playerSearch, "player", false, "Search for any player")
	flag.BoolVar(&skaterSearch, "skater", false, "Search for skaters with detailed stats")
	flag.BoolVar(&goalieSearch, "goalie", false, "Search for goalies with detailed stats")
	flag.BoolVar(&stats, "stats", false, "Get player stats across seasons")
	flag.BoolVar(&schedule, "schedule", false, "Get a team's full schedule")
	flag.BoolVar(&standings, "standings", false, "Get current NHL standings")
	flag.BoolVar(&standingsByDate, "standings-by-date", false, "Get NHL standings for a specific date")
	flag.BoolVar(&leagueStandings, "league-standings", false, "Get overall NHL standings")
	flag.BoolVar(&conferenceStandings, "conference", false, "Get standings by conference")
	flag.BoolVar(&divisionStandings, "division", false, "Get standings by division")
	flag.BoolVar(&gameDetails, "game", false, "Get detailed game information")
	flag.BoolVar(&leaders, "leaders", false, "Get NHL league leaders")
	flag.BoolVar(&liveUpdates, "live", false, "Show live game updates")

	// Parameters
	flag.IntVar(&gameID, "game-id", 2024020750, "Game ID for game details (default: NYR vs CHI on Feb 9, 2024)")
	flag.IntVar(&updateInterval, "interval", 60, "Update interval in seconds for live updates")
	flag.StringVar(&date, "date", "", "Date to get schedule for (format: YYYY-MM-DD)")
	flag.StringVar(&name, "name", "", "Team name for roster, schedule, and standings")

	flag.Parse()

	// Create NHL client
	client := nhl.NewClient()

	// Track if any examples were run
	examplesRun := false

	if todaysSchedule {
		examplesRun = true
		if err := examples.GetTodaysSchedule(client); err != nil {
			log.Fatal(err)
		}
	}

	if slate {
		examplesRun = true
		slateDate := date
		if slateDate == "" {
			slateDate = time.Now().Format("2006-01-02")
		}
		if err := examples.GetScheduleByDate(client, slateDate); err != nil {
			log.Fatal(err)
		}
	}

	if roster {
		examplesRun = true
		if err := examples.GetTeamRoster(client); err != nil {
			log.Fatal(err)
		}
	}

	if playerSearch {
		examplesRun = true
		playerName := name
		if playerName == "" {
			playerName = "Robertson"
		}
		if err := examples.SearchPlayer(client, playerName); err != nil {
			log.Fatal(err)
		}
	}

	if skaterSearch {
		examplesRun = true
		skaterName := name
		if skaterName == "" {
			skaterName = "Hintz"
		}
		if err := examples.SearchSkater(client, skaterName); err != nil {
			log.Fatal(err)
		}
	}

	if goalieSearch {
		examplesRun = true
		goalieName := name
		if goalieName == "" {
			goalieName = "Oettinger"
		}
		if err := examples.SearchGoalie(client, goalieName); err != nil {
			log.Fatal(err)
		}
	}

	if stats {
		examplesRun = true
		playerName := name
		if playerName == "" {
			playerName = "Johnston"
		}
		if err := examples.GetSeasonStats(client, playerName); err != nil {
			log.Fatal(err)
		}
	}

	if schedule {
		examplesRun = true
		teamName := name
		if teamName == "" {
			teamName = "DAL"
		}
		if err := examples.GetTeamSchedule(client, teamName); err != nil {
			log.Fatal(err)
		}
	}

	if standings {
		examplesRun = true
		if err := examples.GetCurrentStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if standingsByDate {
		examplesRun = true
		standingsDate := date
		if standingsDate == "" {
			standingsDate = time.Now().Format("2006-01-02")
		}
		if err := examples.GetStandingsByDate(client, standingsDate); err != nil {
			log.Fatal(err)
		}
	}

	if leagueStandings {
		examplesRun = true
		if err := examples.GetLeagueStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if conferenceStandings {
		examplesRun = true
		if err := examples.GetConferenceStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if divisionStandings {
		examplesRun = true
		if err := examples.GetDivisionStandings(client); err != nil {
			log.Fatal(err)
		}
	}

	if gameDetails {
		examplesRun = true
		if err := examples.GetGameDetails(client, gameID); err != nil {
			log.Fatal(err)
		}
	}

	if leaders {
		examplesRun = true
		if err := examples.GetLeagueLeaders(client); err != nil {
			log.Fatal(err)
		}
	}

	if liveUpdates {
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
		fmt.Println("- leaders: Get NHL league leaders")
	}
}
