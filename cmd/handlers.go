package cmd

import (
	"fmt"
	nhl "go-nhl/client"
	"go-nhl/internal/display"
	"go-nhl/internal/formatters"
	"sort"
	"strings"
	"time"
)

// Schedule Commands
func (c *Config) RunTodaysSchedule() error {
	scores, err := c.Client.GetCurrentSchedule()
	if err != nil {
		return fmt.Errorf("error getting current schedule: %v", err)
	}
	fmt.Println("Games sorted by start time (earliest first - default):")
	display.Games(scores)
	return nil
}

func (c *Config) RunScheduleByDate(date string) error {
	scores, err := c.Client.GetScheduleByDate(date, nhl.SortByDateDesc)
	if err != nil {
		return fmt.Errorf("error getting schedule for date %s: %v", date, err)
	}
	fmt.Println("\nGames sorted by start time (latest first):")
	display.Games(scores)
	return nil
}

func (c *Config) RunTeamSchedule(teamIdentifier string) error {
	team, err := c.Client.GetTeamByIdentifier(teamIdentifier)
	if err != nil {
		return fmt.Errorf("failed to get team: %v", err)
	}

	seasonID := formatters.GetCurrentSeasonID()
	schedule, err := c.Client.GetTeamSchedule(team, seasonID)
	if err != nil {
		return fmt.Errorf("failed to get schedule: %v", err)
	}

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

// Player Commands
func (c *Config) RunPlayerSearch(searchName string) error {
	players, err := c.Client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for player %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	fmt.Printf("\nFound %d players matching '%s':\n", len(players), searchName)
	for i, player := range players {
		fmt.Printf("%d. %s %s (#%d) - %s %s\n",
			i+1,
			player.FirstName.Default,
			player.LastName.Default,
			player.JerseyNumber,
			player.TeamAbbrev,
			player.Position)
	}

	// Get stats for the first player found
	player := players[0]
	stats, err := c.Client.GetFilteredPlayerStats(player.PlayerID, nil)
	if err != nil {
		return fmt.Errorf("error getting stats for player %d: %v", player.PlayerID, err)
	}

	display.SeasonStats(stats, nhl.GameTypeRegularSeason)
	return nil
}

func (c *Config) RunSkaterSearch(searchName string) error {
	players, err := c.Client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for skater %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	// Filter for skaters only
	var skaters []nhl.PlayerSearchResult
	for _, player := range players {
		if player.Position != "G" {
			skaters = append(skaters, player)
		}
	}

	if len(skaters) == 0 {
		fmt.Printf("No skaters found matching '%s'\n", searchName)
		return nil
	}

	fmt.Printf("\nFound %d skaters matching '%s':\n", len(skaters), searchName)
	for i, player := range skaters {
		fmt.Printf("%d. %s %s (#%d) - %s %s\n",
			i+1,
			player.FirstName.Default,
			player.LastName.Default,
			player.JerseyNumber,
			player.TeamAbbrev,
			player.Position)
	}

	// Get stats for the first skater found
	player := skaters[0]

	// Get regular season stats
	fmt.Println("\nRegular Season Stats:")
	stats, err := c.Client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypeRegularSeason,
	})
	if err != nil {
		return fmt.Errorf("error getting regular season stats for skater %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypeRegularSeason)

	// Get playoff stats
	fmt.Println("\nPlayoff Stats:")
	stats, err = c.Client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypePlayoffs,
	})
	if err != nil {
		return fmt.Errorf("error getting playoff stats for skater %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypePlayoffs)
	return nil
}

func (c *Config) RunGoalieSearch(searchName string) error {
	players, err := c.Client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for goalie %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	// Filter for goalies only
	var goalies []nhl.PlayerSearchResult
	for _, player := range players {
		if player.Position == "G" {
			goalies = append(goalies, player)
		}
	}

	if len(goalies) == 0 {
		fmt.Printf("No goalies found matching '%s'\n", searchName)
		return nil
	}

	fmt.Printf("\nFound %d goalies matching '%s':\n", len(goalies), searchName)
	for i, player := range goalies {
		fmt.Printf("%d. %s %s (#%d) - %s %s\n",
			i+1,
			player.FirstName.Default,
			player.LastName.Default,
			player.JerseyNumber,
			player.TeamAbbrev,
			player.Position)
	}

	// Get stats for the first goalie found
	player := goalies[0]

	// Get regular season stats
	fmt.Println("\nRegular Season Stats:")
	stats, err := c.Client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypeRegularSeason,
	})
	if err != nil {
		return fmt.Errorf("error getting regular season stats for goalie %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypeRegularSeason)

	// Get playoff stats
	fmt.Println("\nPlayoff Stats:")
	stats, err = c.Client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
		GameType: nhl.GameTypePlayoffs,
	})
	if err != nil {
		return fmt.Errorf("error getting playoff stats for goalie %d: %v", player.PlayerID, err)
	}
	display.SeasonStats(stats, nhl.GameTypePlayoffs)
	return nil
}

func (c *Config) RunSeasonStats(searchName string) error {
	players, err := c.Client.SearchPlayer(searchName)
	if err != nil {
		return fmt.Errorf("error searching for player %s: %v", searchName, err)
	}

	if len(players) == 0 {
		fmt.Printf("No players found matching '%s'\n", searchName)
		return nil
	}

	// Get the first matching player
	player := players[0]
	fmt.Printf("\nShowing stats for %s %s:\n",
		player.FirstName.Default,
		player.LastName.Default)

	// Get all seasons to show what's available
	allStats, err := c.Client.GetFilteredPlayerStats(player.PlayerID, nil)
	if err != nil {
		return fmt.Errorf("error getting player stats: %v", err)
	}

	fmt.Println("\nAvailable NHL Seasons:")
	for _, season := range allStats {
		fmt.Printf("- %d-%d (%s): %d games played, %d goals, %d points\n",
			season.Season/10000,
			(season.Season/10000)+1,
			display.GetGameTypeName(nhl.GameType(season.GameTypeID)),
			season.GamesPlayed,
			season.Goals,
			season.Points)
	}

	// Calculate current and previous season IDs
	currentSeasonID := formatters.GetCurrentSeasonID()
	previousSeasonID := currentSeasonID - 10000

	// Example seasons to show
	seasons := []struct {
		seasonID int
		gameType nhl.GameType
	}{
		{currentSeasonID, nhl.GameTypeRegularSeason},  // Current season
		{previousSeasonID, nhl.GameTypeRegularSeason}, // Previous season
		{previousSeasonID, nhl.GameTypePlayoffs},      // Previous season playoffs
	}

	// Show stats for each season
	for _, s := range seasons {
		fmt.Printf("\nStats for %d-%d %s:\n",
			s.seasonID/10000,
			(s.seasonID/10000)+1,
			display.GetGameTypeName(s.gameType))

		stats, err := c.Client.GetFilteredPlayerStats(player.PlayerID, &nhl.StatsFilter{
			GameType: s.gameType,
			SeasonID: s.seasonID,
		})
		if err != nil {
			fmt.Printf("Error getting stats: %v\n", err)
			continue
		}

		if len(stats) == 0 {
			fmt.Println("No stats available")
			continue
		}

		display.SeasonStats(stats, s.gameType)
	}

	return nil
}

// Team Commands
func (c *Config) RunTeamRoster() error {
	// Example: Get roster for teams using different identifier types
	identifiers := []string{
		"DAL",                // by abbreviation
		"Montreal Canadiens", // by full name
		"6",                  // by ID (Boston Bruins)
	}
	for _, identifier := range identifiers {
		roster, err := c.Client.GetTeamRoster(identifier)
		if err != nil {
			fmt.Printf("Error getting roster for %s: %v\n", identifier, err)
			continue
		}
		display.Roster(roster, identifier)
		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}
	return nil
}

// Standings Commands
func (c *Config) RunCurrentStandings() error {
	standings, err := c.Client.GetStandings()
	if err != nil {
		return fmt.Errorf("error getting current standings: %v", err)
	}

	fmt.Println("\nCurrent NHL Standings:")
	display.Standings(standings)
	return nil
}

func (c *Config) RunStandingsByDate(date string) error {
	standings, err := c.Client.GetStandingsByDate(date)
	if err != nil {
		return fmt.Errorf("error getting standings for date %s: %v", date, err)
	}

	fmt.Printf("\nNHL Standings for %s:\n", date)
	display.Standings(standings)
	return nil
}

func (c *Config) RunLeagueStandings() error {
	standings, err := c.Client.GetStandings()
	if err != nil {
		return fmt.Errorf("error getting standings: %v", err)
	}

	// Sort all teams by points, regulation wins, goal differential
	teams := standings.Standings
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points != teams[j].Points {
			return teams[i].Points > teams[j].Points
		}
		if teams[i].RegulationWins != teams[j].RegulationWins {
			return teams[i].RegulationWins > teams[j].RegulationWins
		}
		return teams[i].GoalDifferential > teams[j].GoalDifferential
	})

	fmt.Println("\nOverall NHL Standings:")
	fmt.Printf("%-25s GP   W   L  OTL  PTS  REG  GF  GA DIFF  PTS%%  STRK  L10\n", "Team")
	fmt.Println(strings.Repeat("-", 90))

	for i, team := range teams {
		ptsPercentage := float64(team.Points) / float64(team.GamesPlayed*2)
		l10Record := fmt.Sprintf("%d-%d-%d", team.L10Wins, team.L10Losses, team.L10OtLosses)

		fmt.Printf("%2d. %-22s %2d  %2d  %2d   %2d  %3d  %2d %3d %3d  %4d  .%03d  %4s  %5s\n",
			i+1,
			team.TeamName.Default,
			team.GamesPlayed,
			team.Wins,
			team.Losses,
			team.OtLosses,
			team.Points,
			team.RegulationWins,
			team.GoalsFor,
			team.GoalsAgainst,
			team.GoalDifferential,
			int(ptsPercentage*1000),
			formatStreak(team.StreakCode, team.StreakCount),
			l10Record)
	}
	return nil
}

func (c *Config) RunConferenceStandings() error {
	standings, err := c.Client.GetStandings()
	if err != nil {
		return fmt.Errorf("error getting standings: %v", err)
	}

	// Group teams by conference
	conferences := make(map[string][]nhl.StandingsTeam)
	for _, team := range standings.Standings {
		conferences[team.Conference] = append(conferences[team.Conference], team)
	}

	// Sort conferences alphabetically
	confNames := make([]string, 0, len(conferences))
	for conf := range conferences {
		confNames = append(confNames, conf)
	}
	sort.Strings(confNames)

	// Display each conference
	for _, conf := range confNames {
		teams := conferences[conf]
		// Sort teams by points, regulation wins, goal differential
		sort.Slice(teams, func(i, j int) bool {
			if teams[i].Points != teams[j].Points {
				return teams[i].Points > teams[j].Points
			}
			if teams[i].RegulationWins != teams[j].RegulationWins {
				return teams[i].RegulationWins > teams[j].RegulationWins
			}
			return teams[i].GoalDifferential > teams[j].GoalDifferential
		})

		fmt.Printf("\n%s Conference Standings:\n", conf)
		fmt.Printf("%-25s GP   W   L  OTL  PTS  REG  GF  GA DIFF  PTS%%  STRK  L10\n", "Team")
		fmt.Println(strings.Repeat("-", 90))

		for i, team := range teams {
			ptsPercentage := float64(team.Points) / float64(team.GamesPlayed*2)
			l10Record := fmt.Sprintf("%d-%d-%d", team.L10Wins, team.L10Losses, team.L10OtLosses)

			fmt.Printf("%2d. %-22s %2d  %2d  %2d   %2d  %3d  %2d %3d %3d  %4d  .%03d  %4s  %5s\n",
				i+1,
				team.TeamName.Default,
				team.GamesPlayed,
				team.Wins,
				team.Losses,
				team.OtLosses,
				team.Points,
				team.RegulationWins,
				team.GoalsFor,
				team.GoalsAgainst,
				team.GoalDifferential,
				int(ptsPercentage*1000),
				formatStreak(team.StreakCode, team.StreakCount),
				l10Record)
		}
	}
	return nil
}

func (c *Config) RunDivisionStandings() error {
	standings, err := c.Client.GetStandings()
	if err != nil {
		return fmt.Errorf("error getting standings: %v", err)
	}

	// Group teams by division
	divisions := make(map[string][]nhl.StandingsTeam)
	for _, team := range standings.Standings {
		divisions[team.Division] = append(divisions[team.Division], team)
	}

	// Sort divisions alphabetically
	divNames := make([]string, 0, len(divisions))
	for div := range divisions {
		divNames = append(divNames, div)
	}
	sort.Strings(divNames)

	// Display each division
	for _, div := range divNames {
		teams := divisions[div]
		// Sort teams by points, regulation wins, goal differential
		sort.Slice(teams, func(i, j int) bool {
			if teams[i].Points != teams[j].Points {
				return teams[i].Points > teams[j].Points
			}
			if teams[i].RegulationWins != teams[j].RegulationWins {
				return teams[i].RegulationWins > teams[j].RegulationWins
			}
			return teams[i].GoalDifferential > teams[j].GoalDifferential
		})

		fmt.Printf("\n%s Division Standings:\n", div)
		fmt.Printf("%-25s GP   W   L  OTL  PTS  REG  GF  GA DIFF  PTS%%  STRK  L10\n", "Team")
		fmt.Println(strings.Repeat("-", 90))

		for i, team := range teams {
			ptsPercentage := float64(team.Points) / float64(team.GamesPlayed*2)
			l10Record := fmt.Sprintf("%d-%d-%d", team.L10Wins, team.L10Losses, team.L10OtLosses)

			fmt.Printf("%2d. %-22s %2d  %2d  %2d   %2d  %3d  %2d %3d %3d  %4d  .%03d  %4s  %5s\n",
				i+1,
				team.TeamName.Default,
				team.GamesPlayed,
				team.Wins,
				team.Losses,
				team.OtLosses,
				team.Points,
				team.RegulationWins,
				team.GoalsFor,
				team.GoalsAgainst,
				team.GoalDifferential,
				int(ptsPercentage*1000),
				formatStreak(team.StreakCode, team.StreakCount),
				l10Record)
		}
	}
	return nil
}

// Game Commands
func (c *Config) RunGameDetails() error {
	// Get basic game details
	details, err := c.Client.GetGameDetails(c.GameID)
	if err != nil {
		return fmt.Errorf("error getting game details: %v", err)
	}
	if details == nil {
		return fmt.Errorf("no game details found for ID: %d", c.GameID)
	}

	// Get boxscore
	boxscore, err := c.Client.GetGameBoxscore(c.GameID)
	if err != nil {
		return fmt.Errorf("error getting game boxscore: %v", err)
	}
	if boxscore == nil {
		return fmt.Errorf("no boxscore found for ID: %d", c.GameID)
	}

	// Display game details with boxscore
	display.GameDetails(details, boxscore)

	// Display boxscore
	display.GameBoxscore(boxscore)

	// Get play-by-play
	pbp, err := c.Client.GetGamePlayByPlay(c.GameID)
	if err != nil {
		return fmt.Errorf("error getting play-by-play: %v", err)
	}
	if pbp == nil {
		return fmt.Errorf("no play-by-play found for ID: %d", c.GameID)
	}
	display.GamePlayByPlay(pbp)

	return nil
}

func (c *Config) RunLiveGameUpdates() error {
	updates, err := c.Client.GetLiveGameUpdates()
	if err != nil {
		return fmt.Errorf("failed to get live game updates: %w", err)
	}

	display.LiveGameUpdates(updates)
	return nil
}

func (c *Config) RunLeagueLeaders() error {
	// Get current season leaders
	leaders, err := c.Client.GetStatsLeaders(0)
	if err != nil {
		return fmt.Errorf("error getting stats leaders: %v", err)
	}

	display.StatsLeaders(leaders, 0)

	// Get previous season leaders
	prevSeasonLeaders, err := c.Client.GetStatsLeaders(20222023)
	if err != nil {
		return fmt.Errorf("error getting previous season stats leaders: %v", err)
	}

	display.StatsLeaders(prevSeasonLeaders, 20222023)
	return nil
}

// Helper function to format streak
func formatStreak(code string, count int) string {
	if count == 0 {
		return "-"
	}
	return fmt.Sprintf("%s%d", code, count)
}
