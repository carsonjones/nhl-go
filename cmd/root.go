package cmd

import (
	"flag"
	"fmt"
	"go-nhl/client"
	"log"
	"time"
)

type Config struct {
	// Command flags
	TodaysSchedule      bool
	Slate               bool
	Roster              bool
	PlayerSearch        bool
	SkaterSearch        bool
	GoalieSearch        bool
	Stats               bool
	Schedule            bool
	Standings           bool
	StandingsByDate     bool
	LeagueStandings     bool
	ConferenceStandings bool
	DivisionStandings   bool
	GameDetails         bool
	LiveUpdates         bool
	Leaders             bool

	// Parameters
	Date           string
	Name           string
	GameID         int
	UpdateInterval int

	// NHL Client
	Client *nhl.Client
}

func NewConfig() *Config {
	return &Config{
		Client: nhl.NewClient(),
	}
}

func (c *Config) ParseFlags() {
	// Command flags
	flag.BoolVar(&c.TodaysSchedule, "today", false, "Get today's NHL schedule")
	flag.BoolVar(&c.Slate, "slate", false, "Get schedule for a specific date")
	flag.BoolVar(&c.Roster, "roster", false, "Get team rosters")
	flag.BoolVar(&c.PlayerSearch, "player", false, "Search for any player")
	flag.BoolVar(&c.SkaterSearch, "skater", false, "Search for skaters with detailed stats")
	flag.BoolVar(&c.GoalieSearch, "goalie", false, "Search for goalies with detailed stats")
	flag.BoolVar(&c.Stats, "stats", false, "Get player stats across seasons")
	flag.BoolVar(&c.Schedule, "schedule", false, "Get a team's full schedule")
	flag.BoolVar(&c.Standings, "standings", false, "Get current NHL standings")
	flag.BoolVar(&c.StandingsByDate, "standings-by-date", false, "Get NHL standings for a specific date")
	flag.BoolVar(&c.LeagueStandings, "league-standings", false, "Get overall NHL standings")
	flag.BoolVar(&c.ConferenceStandings, "conference", false, "Get standings by conference")
	flag.BoolVar(&c.DivisionStandings, "division", false, "Get standings by division")
	flag.BoolVar(&c.GameDetails, "game", false, "Get detailed game information")
	flag.BoolVar(&c.Leaders, "leaders", false, "Get NHL league leaders")
	flag.BoolVar(&c.LiveUpdates, "live", false, "Show live game updates")

	// Parameters
	flag.IntVar(&c.GameID, "game-id", 2024020750, "Game ID for game details (default: NYR vs CHI on Feb 9, 2024)")
	flag.IntVar(&c.UpdateInterval, "interval", 60, "Update interval in seconds for live updates")
	flag.StringVar(&c.Date, "date", "", "Date to get schedule for (format: YYYY-MM-DD)")
	flag.StringVar(&c.Name, "name", "", "Team name for roster, schedule, and standings")

	flag.Parse()
}

func (c *Config) Execute() error {
	commandsRun := false

	if c.TodaysSchedule {
		commandsRun = true
		if err := c.RunTodaysSchedule(); err != nil {
			return err
		}
	}

	if c.Slate {
		commandsRun = true
		slateDate := c.Date
		if slateDate == "" {
			slateDate = time.Now().Format("2006-01-02")
		}
		if err := c.RunScheduleByDate(slateDate); err != nil {
			return err
		}
	}

	if c.Roster {
		commandsRun = true
		if err := c.RunTeamRoster(); err != nil {
			return err
		}
	}

	if c.PlayerSearch {
		commandsRun = true
		playerName := c.Name
		if playerName == "" {
			playerName = "Robertson"
		}
		if err := c.RunPlayerSearch(playerName); err != nil {
			return err
		}
	}

	if c.SkaterSearch {
		commandsRun = true
		skaterName := c.Name
		if skaterName == "" {
			skaterName = "Hintz"
		}
		if err := c.RunSkaterSearch(skaterName); err != nil {
			return err
		}
	}

	if c.GoalieSearch {
		commandsRun = true
		goalieName := c.Name
		if goalieName == "" {
			goalieName = "Oettinger"
		}
		if err := c.RunGoalieSearch(goalieName); err != nil {
			return err
		}
	}

	if c.Stats {
		commandsRun = true
		playerName := c.Name
		if playerName == "" {
			playerName = "Johnston"
		}
		if err := c.RunSeasonStats(playerName); err != nil {
			return err
		}
	}

	if c.Schedule {
		commandsRun = true
		teamName := c.Name
		if teamName == "" {
			teamName = "DAL"
		}
		if err := c.RunTeamSchedule(teamName); err != nil {
			return err
		}
	}

	if c.Standings {
		commandsRun = true
		if err := c.RunCurrentStandings(); err != nil {
			return err
		}
	}

	if c.StandingsByDate {
		commandsRun = true
		standingsDate := c.Date
		if standingsDate == "" {
			standingsDate = time.Now().Format("2006-01-02")
		}
		if err := c.RunStandingsByDate(standingsDate); err != nil {
			return err
		}
	}

	if c.LeagueStandings {
		commandsRun = true
		if err := c.RunLeagueStandings(); err != nil {
			return err
		}
	}

	if c.ConferenceStandings {
		commandsRun = true
		if err := c.RunConferenceStandings(); err != nil {
			return err
		}
	}

	if c.DivisionStandings {
		commandsRun = true
		if err := c.RunDivisionStandings(); err != nil {
			return err
		}
	}

	if c.GameDetails {
		commandsRun = true
		if err := c.RunGameDetails(); err != nil {
			return err
		}
	}

	if c.Leaders {
		commandsRun = true
		if err := c.RunLeagueLeaders(); err != nil {
			return err
		}
	}

	if c.LiveUpdates {
		commandsRun = true
		fmt.Printf("Starting live game updates (refreshing every %d seconds). Press Ctrl+C to stop.\n", c.UpdateInterval)
		for {
			if err := c.RunLiveGameUpdates(); err != nil {
				log.Printf("Error getting live updates: %v", err)
			}
			time.Sleep(time.Duration(c.UpdateInterval) * time.Second)
		}
	}

	if !commandsRun {
		c.PrintUsage()
	}

	return nil
}

func (c *Config) PrintUsage() {
	fmt.Println("Available commands (use -h flag to see all options):")
	fmt.Println("- today: Get today's NHL schedule")
	fmt.Println("- slate: Get schedule for a specific date")
	fmt.Println("- roster: Get team rosters")
	fmt.Println("- player: Search for any player")
	fmt.Println("- skater: Search for skaters with detailed stats")
	fmt.Println("- goalie: Search for goalies with detailed stats")
	fmt.Println("- stats: Get player stats across seasons")
	fmt.Println("- schedule: Get a team's full schedule")
	fmt.Println("- standings: Get current NHL standings")
	fmt.Println("- standings-by-date: Get NHL standings for a specific date")
	fmt.Println("- league-standings: Get overall NHL standings")
	fmt.Println("- conference: Get standings by conference")
	fmt.Println("- division: Get standings by division")
	fmt.Println("- game: Get detailed game information")
	fmt.Println("- live: Show live game updates")
	fmt.Println("- leaders: Get NHL league leaders")
}
