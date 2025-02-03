package display

import (
	"fmt"
	"go-nhl/client"
	"sort"
	"strings"
)

// formatRecord formats a team's record as "W-L-OTL"
func formatRecord(record nhl.Record) string {
	return fmt.Sprintf("%d-%d-%d", record.Wins, record.Losses, record.OtLosses)
}

// formatStreak formats a team's streak (e.g., "W3" for 3 wins in a row)
func formatStreak(code string, count int) string {
	if count == 0 {
		return "-"
	}
	return fmt.Sprintf("%s%d", code, count)
}

// SortTeams sorts teams by NHL standings rules:
// 1. Points (descending)
// 2. Regulation Wins (descending)
// 3. Goal Differential (descending)
func SortTeams(teams []nhl.StandingsTeam) {
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points != teams[j].Points {
			return teams[i].Points > teams[j].Points
		}
		if teams[i].RegulationWins != teams[j].RegulationWins {
			return teams[i].RegulationWins > teams[j].RegulationWins
		}
		return teams[i].GoalDifferential > teams[j].GoalDifferential
	})
}

// Standings displays the NHL standings
func Standings(standings *nhl.StandingsResponse) {
	// Group teams by conference and division
	conferences := make(map[string]map[string][]nhl.StandingsTeam)
	for _, team := range standings.Standings {
		conf := team.Conference
		div := team.Division
		if conferences[conf] == nil {
			conferences[conf] = make(map[string][]nhl.StandingsTeam)
		}
		conferences[conf][div] = append(conferences[conf][div], team)
	}

	// Sort conferences alphabetically
	confNames := make([]string, 0, len(conferences))
	for conf := range conferences {
		confNames = append(confNames, conf)
	}
	sort.Strings(confNames)

	// Display standings for each conference
	for _, conf := range confNames {
		fmt.Printf("\n%s Conference\n", conf)
		fmt.Println(strings.Repeat("=", len(conf)+11))

		// Sort divisions alphabetically
		divisions := make([]string, 0, len(conferences[conf]))
		for div := range conferences[conf] {
			divisions = append(divisions, div)
		}
		sort.Strings(divisions)

		// Display each division
		for _, div := range divisions {
			teams := conferences[conf][div]
			fmt.Printf("\n%s Division\n", div)
			fmt.Println(strings.Repeat("-", len(div)+9))

			// Sort teams using the shared sorting function
			SortTeams(teams)

			// Print header
			fmt.Printf("%-25s GP   W   L  OTL  PTS  REG  GF  GA DIFF  PTS%%  STRK  L10    HOME    AWAY\n", "Team")
			fmt.Println(strings.Repeat("-", 105))

			// Print each team
			for _, team := range teams {
				// Create records from individual stats
				homeRecord := nhl.Record{
					Wins:     team.HomeWins,
					Losses:   team.HomeLosses,
					OtLosses: team.HomeOtLosses,
				}
				l10Record := nhl.Record{
					Wins:     team.L10Wins,
					Losses:   team.L10Losses,
					OtLosses: team.L10OtLosses,
				}

				// Calculate points percentage
				ptsPercentage := float64(team.Points) / float64(team.GamesPlayed*2)

				fmt.Printf("%-25s %2d  %2d  %2d   %2d  %3d  %2d %3d %3d  %4d  .%03d  %4s  %5s  %7s  %7s\n",
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
					int(ptsPercentage*1000), // Convert to 3-digit format (e.g., .683)
					formatStreak(team.StreakCode, team.StreakCount),
					formatRecord(l10Record),
					formatRecord(homeRecord),
					formatRecord(nhl.Record{
						Wins:     team.Wins - team.HomeWins,
						Losses:   team.Losses - team.HomeLosses,
						OtLosses: team.OtLosses - team.HomeOtLosses,
					}))
			}
		}

		// Display wild card standings if in regular season
		if len(standings.Standings) > 0 && standings.Standings[0].WildCardSequence > 0 {
			fmt.Printf("\nWild Card\n")
			fmt.Println("---------")

			// Get all teams in conference
			var wcTeams []nhl.StandingsTeam
			for _, teams := range conferences[conf] {
				wcTeams = append(wcTeams, teams...)
			}

			// Sort by wild card sequence
			sort.Slice(wcTeams, func(i, j int) bool {
				return wcTeams[i].WildCardSequence < wcTeams[j].WildCardSequence
			})

			// Print wild card teams
			for _, team := range wcTeams {
				if team.WildCardSequence > 0 {
					fmt.Printf("%d. %-23s %3d pts (%d GP)\n",
						team.WildCardSequence,
						team.TeamName.Default,
						team.Points,
						team.GamesPlayed)
				}
			}
		}
	}
}
