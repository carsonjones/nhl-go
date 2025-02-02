package nhl

import (
	"fmt"
)

// GetGameDetails returns detailed information about a specific game
func (c *Client) GetGameDetails(gameID int) (*GameDetails, error) {
	url := fmt.Sprintf("%s/gamecenter/%d/landing", c.baseURL, gameID)
	var response GameDetails
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get game details: %v", err)
	}
	return &response, nil
}

// GetGameBoxscore returns the boxscore for a specific game
func (c *Client) GetGameBoxscore(gameID int) (*BoxscoreResponse, error) {
	url := fmt.Sprintf("%s/gamecenter/%d/boxscore", c.baseURL, gameID)
	var response BoxscoreResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get game boxscore: %v", err)
	}
	return &response, nil
}

// GetGamePlayByPlay returns the play-by-play data for a specific game
func (c *Client) GetGamePlayByPlay(gameID int) (*PlayByPlayResponse, error) {
	url := fmt.Sprintf("%s/gamecenter/%d/play-by-play", c.baseURL, gameID)
	var response PlayByPlayResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get game play-by-play: %v", err)
	}
	return &response, nil
}

// GetGameStory returns the game story/narrative for a specific game
func (c *Client) GetGameStory(gameID int) (*GameStoryResponse, error) {
	url := fmt.Sprintf("%s/wsc/game-story/%d", c.baseURL, gameID)
	var response GameStoryResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get game story: %v", err)
	}
	return &response, nil
}
