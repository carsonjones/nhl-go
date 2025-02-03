package display

import (
	"fmt"
	"go-nhl/internal/formatters"
	"go-nhl/nhl"
	"sort"
	"strings"
	"time"
)

// Games displays a list of games from the scoreboard
func Games(games *nhl.FilteredScoreboardResponse) {
	// Get the date from the first game if available
	displayDate := games.Date
	if displayDate == "" && len(games.Games) > 0 {
		if gameTime, err := time.Parse(time.RFC3339, games.Games[0].StartTimeUTC); err == nil {
			displayDate = gameTime.Format("2006-01-02")
		}
	}

	fmt.Printf("\nGames for %s:\n", displayDate)
	for _, game := range games.Games {
		gameTime, err := formatters.FormatGameTime(game.StartTimeUTC)
		if err != nil {
			fmt.Printf("Error formatting game time: %v\n", err)
			continue
		}

		fmt.Printf("%s at %s - %s\n",
			game.AwayTeam.Name.Default,
			game.HomeTeam.Name.Default,
			gameTime)

		if game.GameState == "LIVE" || game.GameState == "OFF" {
			fmt.Printf("Score: %s %d, %s %d\n",
				game.AwayTeam.Name.Default, game.AwayTeam.Score,
				game.HomeTeam.Name.Default, game.HomeTeam.Score)

			if game.GameState == "LIVE" {
				fmt.Printf("Period: %d (%s)\n",
					game.Period,
					game.PeriodDescriptor.PeriodType)
			}
		} else {
			fmt.Printf("Game Status: %s\n", game.GameState)
		}
		fmt.Println()
	}
}

// GameDetails displays detailed information about a specific game
func GameDetails(game *nhl.GameDetails, boxscore *nhl.BoxscoreResponse) {
	fmt.Printf("\nGame Details:\n")
	fmt.Printf("Date: %s\n", game.GameDate)
	fmt.Printf("Start Time (UTC): %s\n", game.StartTimeUTC)
	fmt.Printf("Venue: %s\n", game.Venue.Default)
	fmt.Printf("Status: %s\n", game.GameState)

	// Display team information
	fmt.Printf("\n%-20s %s\n", game.AwayTeam.PlaceName.Default, game.HomeTeam.PlaceName.Default)
	fmt.Printf("%-20s %s\n", game.AwayTeam.CommonName.Default, game.HomeTeam.CommonName.Default)
	fmt.Printf("%-20s %s\n", game.AwayTeam.Abbrev, game.HomeTeam.Abbrev)
	fmt.Printf("%-20d %d\n", game.AwayTeam.Score, game.HomeTeam.Score)

	// Display team stats
	fmt.Printf("\nTeam Stats:\n")
	fmt.Printf("%-6s %3s %3s %3s %5s %4s\n", "Team", "G", "SOG", "HIT", "FO%", "PIM")

	// Calculate team totals from boxscore
	var awayHits, awayPIM int
	var awayFOWins, awayFOTotal float64
	for _, player := range boxscore.PlayerByGameStats.AwayTeam.Forwards {
		awayHits += player.Hits
		awayPIM += player.PIM
		if player.FaceoffWinningPct > 0 {
			awayFOWins += player.FaceoffWinningPct
			awayFOTotal++
		}
	}
	for _, player := range boxscore.PlayerByGameStats.AwayTeam.Defense {
		awayHits += player.Hits
		awayPIM += player.PIM
	}

	var homeHits, homePIM int
	var homeFOWins, homeFOTotal float64
	for _, player := range boxscore.PlayerByGameStats.HomeTeam.Forwards {
		homeHits += player.Hits
		homePIM += player.PIM
		if player.FaceoffWinningPct > 0 {
			homeFOWins += player.FaceoffWinningPct
			homeFOTotal++
		}
	}
	for _, player := range boxscore.PlayerByGameStats.HomeTeam.Defense {
		homeHits += player.Hits
		homePIM += player.PIM
	}

	// Calculate average faceoff percentage
	awayFOPct := 0.0
	if awayFOTotal > 0 {
		awayFOPct = awayFOWins / awayFOTotal
	}
	homeFOPct := 0.0
	if homeFOTotal > 0 {
		homeFOPct = homeFOWins / homeFOTotal
	}

	fmt.Printf("%-6s %3d %3d %3d %4.1f%% %4d\n",
		game.AwayTeam.Abbrev,
		game.AwayTeam.Score,
		game.AwayTeam.ShotsOnGoal,
		awayHits,
		awayFOPct*100,
		awayPIM)
	fmt.Printf("%-6s %3d %3d %3d %4.1f%% %4d\n",
		game.HomeTeam.Abbrev,
		game.HomeTeam.Score,
		game.HomeTeam.ShotsOnGoal,
		homeHits,
		homeFOPct*100,
		homePIM)

	// Display scoring summary
	if len(game.Summary.Scoring) > 0 {
		fmt.Printf("\nScoring Summary:\n")
		for _, period := range game.Summary.Scoring {
			if len(period.Goals) > 0 {
				fmt.Printf("\nPeriod %d:\n", period.PeriodDescriptor.Number)
				for _, goal := range period.Goals {
					fmt.Printf("%s - %s (%s) %s\n",
						goal.TimeInPeriod,
						goal.Name.Default,
						goal.TeamAbbrev.Default,
						formatAssists(goal.Assists))
				}
			}
		}
	}

	// Display penalty summary
	if len(game.Summary.Penalties) > 0 {
		fmt.Printf("\nPenalty Summary:\n")
		for _, period := range game.Summary.Penalties {
			if len(period.Penalties) > 0 {
				fmt.Printf("\nPeriod %d:\n", period.PeriodDescriptor.Number)
				for _, penalty := range period.Penalties {
					fmt.Printf("%s - %s %s (%d min) drawn by %s\n",
						penalty.TimeInPeriod,
						penalty.CommittedByPlayer,
						penalty.DescKey,
						penalty.Duration,
						penalty.DrawnBy)
				}
			}
		}
	}

	// Display three stars
	if len(game.ThreeStars) > 0 {
		fmt.Printf("\nThree Stars:\n")
		for _, star := range game.ThreeStars {
			var stats string
			if star.Position == "G" {
				stats = fmt.Sprintf("Save %%: %.1f", star.SavePctg*100)
			} else {
				stats = fmt.Sprintf("G: %d, A: %d, P: %d", star.Goals, star.Assists, star.Points)
			}
			fmt.Printf("%d. %s (%s) - %s #%d - %s\n",
				star.Star,
				star.Name.Default,
				star.TeamAbbrev,
				star.Position,
				star.SweaterNo,
				stats)
		}
	}
}

// Helper function to format assists from the new model
func formatAssists(assists []nhl.AssistEvent) string {
	if len(assists) == 0 {
		return "Unassisted"
	}

	names := make([]string, len(assists))
	for i, assist := range assists {
		names[i] = assist.Name.Default
	}
	return strings.Join(names, ", ")
}

// GameBoxscore displays the boxscore for a game
func GameBoxscore(boxscore *nhl.BoxscoreResponse) {
	fmt.Printf("\nBoxscore Summary\n")
	fmt.Printf("%-6s %3s %3s %3s %5s %4s\n", "Team", "G", "SOG", "HIT", "FO%", "PIM")

	// Calculate team totals
	var awayHits, awayPIM, awaySOG int
	var awayFOWins, awayFOTotal float64
	for _, player := range boxscore.PlayerByGameStats.AwayTeam.Forwards {
		awayHits += player.Hits
		awayPIM += player.PIM
		awaySOG += player.SOG
		if player.FaceoffWinningPct > 0 {
			awayFOWins += player.FaceoffWinningPct
			awayFOTotal++
		}
	}
	for _, player := range boxscore.PlayerByGameStats.AwayTeam.Defense {
		awayHits += player.Hits
		awayPIM += player.PIM
		awaySOG += player.SOG
	}

	var homeHits, homePIM, homeSOG int
	var homeFOWins, homeFOTotal float64
	for _, player := range boxscore.PlayerByGameStats.HomeTeam.Forwards {
		homeHits += player.Hits
		homePIM += player.PIM
		homeSOG += player.SOG
		if player.FaceoffWinningPct > 0 {
			homeFOWins += player.FaceoffWinningPct
			homeFOTotal++
		}
	}
	for _, player := range boxscore.PlayerByGameStats.HomeTeam.Defense {
		homeHits += player.Hits
		homePIM += player.PIM
		homeSOG += player.SOG
	}

	// Calculate average faceoff percentage
	awayFOPct := 0.0
	if awayFOTotal > 0 {
		awayFOPct = awayFOWins / awayFOTotal
	}
	homeFOPct := 0.0
	if homeFOTotal > 0 {
		homeFOPct = homeFOWins / homeFOTotal
	}

	fmt.Printf("%-6s %3d %3d %3d %4.1f%% %4d\n",
		boxscore.AwayTeam.Abbrev,
		boxscore.AwayTeam.Score,
		awaySOG,
		awayHits,
		awayFOPct*100,
		awayPIM)
	fmt.Printf("%-6s %3d %3d %3d %4.1f%% %4d\n",
		boxscore.HomeTeam.Abbrev,
		boxscore.HomeTeam.Score,
		homeSOG,
		homeHits,
		homeFOPct*100,
		homePIM)

	// Display skater stats
	fmt.Printf("\nSkater Stats:\n")

	// Away team skaters
	fmt.Printf("\n%s Skaters:\n", boxscore.AwayTeam.Abbrev)
	fmt.Printf("%-20s %2s %2s %2s %3s %5s %3s %3s %3s\n",
		"Player", "G", "A", "P", "+/-", "TOI", "SOG", "HIT", "BLK")

	// Sort players by points, then goals
	awaySkaters := append(boxscore.PlayerByGameStats.AwayTeam.Forwards,
		boxscore.PlayerByGameStats.AwayTeam.Defense...)
	sort.Slice(awaySkaters, func(i, j int) bool {
		if awaySkaters[i].Points != awaySkaters[j].Points {
			return awaySkaters[i].Points > awaySkaters[j].Points
		}
		return awaySkaters[i].Goals > awaySkaters[j].Goals
	})

	for _, player := range awaySkaters {
		fmt.Printf("%-20s %2d %2d %2d %3d %5s %3d %3d %3d\n",
			player.Name.Default,
			player.Goals,
			player.Assists,
			player.Points,
			player.PlusMinus,
			player.TOI,
			player.SOG,
			player.Hits,
			player.BlockedShots)
	}

	// Home team skaters
	fmt.Printf("\n%s Skaters:\n", boxscore.HomeTeam.Abbrev)
	fmt.Printf("%-20s %2s %2s %2s %3s %5s %3s %3s %3s\n",
		"Player", "G", "A", "P", "+/-", "TOI", "SOG", "HIT", "BLK")

	// Sort players by points, then goals
	homeSkaters := append(boxscore.PlayerByGameStats.HomeTeam.Forwards,
		boxscore.PlayerByGameStats.HomeTeam.Defense...)
	sort.Slice(homeSkaters, func(i, j int) bool {
		if homeSkaters[i].Points != homeSkaters[j].Points {
			return homeSkaters[i].Points > homeSkaters[j].Points
		}
		return homeSkaters[i].Goals > homeSkaters[j].Goals
	})

	for _, player := range homeSkaters {
		fmt.Printf("%-20s %2d %2d %2d %3d %5s %3d %3d %3d\n",
			player.Name.Default,
			player.Goals,
			player.Assists,
			player.Points,
			player.PlusMinus,
			player.TOI,
			player.SOG,
			player.Hits,
			player.BlockedShots)
	}

	// Display goalie stats
	fmt.Printf("\nGoalie Stats:\n")
	fmt.Printf("%-20s %3s %3s %4s %5s %7s\n",
		"Player", "GA", "SV", "SV%", "TOI", "DEC")

	// Away team goalies
	for _, goalie := range boxscore.PlayerByGameStats.AwayTeam.Goalies {
		if goalie.TOI != "00:00" {
			fmt.Printf("%-20s %3d %3d %3.1f%% %5s %7s\n",
				goalie.Name.Default,
				goalie.GoalsAgainst,
				goalie.Saves,
				goalie.SavePctg*100,
				goalie.TOI,
				goalie.Decision)
		}
	}

	// Home team goalies
	for _, goalie := range boxscore.PlayerByGameStats.HomeTeam.Goalies {
		if goalie.TOI != "00:00" {
			fmt.Printf("%-20s %3d %3d %3.1f%% %5s %7s\n",
				goalie.Name.Default,
				goalie.GoalsAgainst,
				goalie.Saves,
				goalie.SavePctg*100,
				goalie.TOI,
				goalie.Decision)
		}
	}
}

// GamePlayByPlay displays play-by-play data for a game
func GamePlayByPlay(pbp *nhl.PlayByPlayResponse) {
	if len(pbp.Plays) == 0 {
		fmt.Println("\nNo play-by-play data available.")
		return
	}

	// Create a map of player IDs to names
	players := make(map[int]string)
	for _, player := range pbp.RosterSpots {
		players[player.PlayerID] = fmt.Sprintf("%s %s (%d)",
			player.FirstName.Default,
			player.LastName.Default,
			player.SweaterNumber)
	}

	fmt.Printf("\nPlay-by-Play:\n")
	fmt.Printf("%-6s %-8s %-8s %-50s\n", "Period", "Time", "Remain", "Event")
	fmt.Println(strings.Repeat("-", 80))

	for _, play := range pbp.Plays {
		// Skip non-significant events
		switch play.TypeDescKey {
		case "period-start", "period-end", "game-end", "stoppage",
			"giveaway", "takeaway", "delayed-penalty":
			continue
		}

		// Format event description based on type
		var description string
		switch play.TypeDescKey {
		case "shot-on-goal":
			shooter := players[play.Details.ShootingPlayerID]
			goalie := players[play.Details.GoalieInNetID]
			description = fmt.Sprintf("Shot by %s, saved by %s", shooter, goalie)
		case "goal":
			scorer := players[play.Details.ScoringPlayerID]
			description = fmt.Sprintf("GOAL! Scored by %s", scorer)
			if play.Details.Assist1PlayerID > 0 {
				assist1 := players[play.Details.Assist1PlayerID]
				if play.Details.Assist2PlayerID > 0 {
					assist2 := players[play.Details.Assist2PlayerID]
					description += fmt.Sprintf(" (Assists: %s, %s)", assist1, assist2)
				} else {
					description += fmt.Sprintf(" (Assist: %s)", assist1)
				}
			}
		case "blocked-shot":
			shooter := players[play.Details.ShootingPlayerID]
			blocker := players[play.Details.BlockingPlayerID]
			description = fmt.Sprintf("Shot by %s blocked by %s", shooter, blocker)
		case "missed-shot":
			shooter := players[play.Details.ShootingPlayerID]
			description = fmt.Sprintf("Shot by %s (%s)", shooter, play.Details.Reason)
		case "hit":
			hitter := players[play.Details.HittingPlayerID]
			hittee := players[play.Details.HitteePlayerID]
			description = fmt.Sprintf("%s hit %s", hitter, hittee)
		case "faceoff":
			winner := players[play.Details.WinningPlayerID]
			loser := players[play.Details.LosingPlayerID]
			description = fmt.Sprintf("Faceoff won by %s vs %s", winner, loser)
		case "penalty":
			offender := players[play.Details.CommittedByPlayerID]
			drawer := players[play.Details.DrawnByPlayerID]
			description = fmt.Sprintf("%s %s (%d min) drawn by %s",
				offender,
				play.Details.DescKey,
				play.Details.Duration,
				drawer)
		default:
			description = play.TypeDescKey
		}

		fmt.Printf("%-6d %-8s %-8s %-50s\n",
			play.PeriodDescriptor.Number,
			play.TimeInPeriod,
			play.TimeRemaining,
			description)
	}
}

// GameStory displays the game story/narrative
func GameStory(story *nhl.GameStoryResponse) {
	fmt.Println(story)
}

// FormatAssists formats a list of assists into a readable string
func FormatAssists(assists []nhl.PlayerBrief) string {
	if len(assists) == 0 {
		return "Unassisted"
	}

	names := make([]string, len(assists))
	for i, assist := range assists {
		names[i] = assist.Name.Default
	}
	return strings.Join(names, ", ")
}

// LiveGameUpdates displays live game information
func LiveGameUpdates(updates *nhl.ScoreboardResponse) {
	if updates == nil || len(updates.GamesByDate) == 0 {
		fmt.Println("No games found")
		return
	}

	now := time.Now()
	hasGames := false

	for _, gameDate := range updates.GamesByDate {
		var activeGames []nhl.Game
		for _, game := range gameDate.Games {
			gameTime, err := time.Parse("2006-01-02T15:04:05Z", game.StartTimeUTC)
			if err != nil {
				continue
			}

			// Include games that:
			// 1. Are currently live
			// 2. Have finished within the last hour
			// 3. Are starting within the next hour
			timeSinceEnd := now.Sub(gameTime)
			timeUntilStart := gameTime.Sub(now)

			if game.GameState == "LIVE" ||
				(game.GameState == "FINAL" && timeSinceEnd < time.Hour) ||
				(game.GameState == "OFF" && timeSinceEnd < time.Hour) ||
				(game.GameState == "PRE" && timeUntilStart < time.Hour) {
				activeGames = append(activeGames, game)
			}
		}

		if len(activeGames) > 0 {
			hasGames = true
			fmt.Printf("\nGames for %s:\n", gameDate.Date)
			fmt.Println(strings.Repeat("-", 80))

			for _, game := range activeGames {
				// Display game status
				var statusText string
				switch game.GameState {
				case "LIVE":
					statusText = fmt.Sprintf("LIVE - Period %d, %s", game.Period, game.Clock.TimeRemaining)
					if game.Clock.InIntermission {
						statusText += " (Intermission)"
					}
				case "FINAL", "OFF":
					statusText = "FINAL"
					if game.PeriodDescriptor.Number > 3 {
						if game.PeriodDescriptor.PeriodType == "OT" {
							statusText += " (OT)"
						} else if game.PeriodDescriptor.PeriodType == "SO" {
							statusText += " (SO)"
						}
					}
				case "PRE":
					localTime, _ := time.Parse("2006-01-02T15:04:05Z", game.StartTimeUTC)
					statusText = fmt.Sprintf("Starting at %s", localTime.Format("3:04 PM MST"))
				default:
					statusText = game.GameState
				}

				fmt.Printf("\n%s @ %s - %s\n", game.AwayTeam.Name.Default, game.HomeTeam.Name.Default, statusText)
				fmt.Printf("Score: %s %d, %s %d\n", game.AwayTeam.Abbrev, game.AwayTeam.Score, game.HomeTeam.Abbrev, game.HomeTeam.Score)

				if game.GameState == "LIVE" {
					fmt.Printf("Shots on Goal: %s %d, %s %d\n", game.AwayTeam.Abbrev, game.AwayTeam.ShotsOnGoal, game.HomeTeam.Abbrev, game.HomeTeam.ShotsOnGoal)

					// Display power play situation if applicable
					if game.Situation != nil && len(game.Situation.HomeTeam.SituationDescriptions) > 0 {
						fmt.Printf("Situation: %s %v\n", game.Situation.HomeTeam.Abbrev, game.Situation.HomeTeam.SituationDescriptions)
					}
				}

				fmt.Println(strings.Repeat("-", 40))
			}
		}
	}

	if !hasGames {
		fmt.Println("\nNo active games at the moment.")
		fmt.Println("Check back later for live game updates!")
	}
}
