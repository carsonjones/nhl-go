package display

import (
	"fmt"
	"go-nhl/internal/formatters"
	"go-nhl/nhl"
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
