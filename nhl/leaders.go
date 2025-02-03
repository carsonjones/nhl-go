package nhl

import (
	"fmt"
	"go-nhl/internal/formatters"
)

// StatsLeaderPlayer represents a player in the stats leaders list
type StatsLeaderPlayer struct {
	ID            int           `json:"id"`
	FirstName     LanguageNames `json:"firstName"`
	LastName      LanguageNames `json:"lastName"`
	SweaterNumber int           `json:"sweaterNumber"`
	Headshot      string        `json:"headshot"`
	TeamAbbrev    string        `json:"teamAbbrev"`
	TeamName      LanguageNames `json:"teamName"`
	TeamLogo      string        `json:"teamLogo"`
	Position      string        `json:"position"`
	Value         float64       `json:"value"`
}

// StatsLeadersResponse represents the response from the stats leaders API
type StatsLeadersResponse struct {
	GoalsSh        []StatsLeaderPlayer `json:"goalsSh"`
	PlusMinus      []StatsLeaderPlayer `json:"plusMinus"`
	Assists        []StatsLeaderPlayer `json:"assists"`
	GoalsPp        []StatsLeaderPlayer `json:"goalsPp"`
	FaceoffLeaders []StatsLeaderPlayer `json:"faceoffLeaders"`
	PenaltyMins    []StatsLeaderPlayer `json:"penaltyMins"`
	Goals          []StatsLeaderPlayer `json:"goals"`
	Points         []StatsLeaderPlayer `json:"points"`
	TOI            []StatsLeaderPlayer `json:"toi"`
}

// GetStatsLeaders returns the NHL stats leaders for a given season
// If seasonID is 0, it returns stats for the current season
func (c *Client) GetStatsLeaders(seasonID int) (*StatsLeadersResponse, error) {
	// If no season provided, use current season
	if seasonID == 0 {
		seasonID = formatters.GetCurrentSeasonID()
	}

	url := fmt.Sprintf("%s/skater-stats-leaders/%d/2", c.baseURL, seasonID)
	var response StatsLeadersResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get stats leaders: %v", err)
	}
	return &response, nil
}
