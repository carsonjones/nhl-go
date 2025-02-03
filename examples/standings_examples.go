package examples

import (
	"fmt"
	"go-nhl/internal/display"
	"go-nhl/client"
	"sort"
	"strings"
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

// GetLeagueStandings demonstrates retrieving and displaying overall NHL standings
func GetLeagueStandings(client *nhl.Client) error {
	standings, err := client.GetStandings()
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

// GetConferenceStandings demonstrates retrieving and displaying standings by conference
func GetConferenceStandings(client *nhl.Client) error {
	standings, err := client.GetStandings()
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

// GetDivisionStandings demonstrates retrieving and displaying standings by division
func GetDivisionStandings(client *nhl.Client) error {
	standings, err := client.GetStandings()
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

// Helper function to format streak
func formatStreak(code string, count int) string {
	if count == 0 {
		return "-"
	}
	return fmt.Sprintf("%s%d", code, count)
}
