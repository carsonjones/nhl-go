package examples

import (
	"fmt"

	"go-nhl/internal/display"
	"go-nhl/client"
)

// GetGameDetails demonstrates retrieving and displaying detailed game information
func GetGameDetails(client *nhl.Client, gameID int) error {
	// Get basic game details
	details, err := client.GetGameDetails(gameID)
	if err != nil {
		return fmt.Errorf("error getting game details: %v", err)
	}
	if details == nil {
		return fmt.Errorf("no game details found for ID: %d", gameID)
	}

	// Get boxscore
	boxscore, err := client.GetGameBoxscore(gameID)
	if err != nil {
		return fmt.Errorf("error getting game boxscore: %v", err)
	}
	if boxscore == nil {
		return fmt.Errorf("no boxscore found for ID: %d", gameID)
	}

	// Display game details with boxscore
	display.GameDetails(details, boxscore)

	// Display boxscore
	display.GameBoxscore(boxscore)

	// Get play-by-play
	pbp, err := client.GetGamePlayByPlay(gameID)
	if err != nil {
		return fmt.Errorf("error getting play-by-play: %v", err)
	}
	if pbp == nil {
		return fmt.Errorf("no play-by-play found for ID: %d", gameID)
	}
	display.GamePlayByPlay(pbp)

	return nil
}

// GetLiveGameUpdates demonstrates retrieving live game updates
func GetLiveGameUpdates(client *nhl.Client) error {
	updates, err := client.GetLiveGameUpdates()
	if err != nil {
		return fmt.Errorf("failed to get live game updates: %w", err)
	}

	display.LiveGameUpdates(updates)
	return nil
}
