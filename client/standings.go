package nhl

import (
	"fmt"
)

// GetStandings returns the current NHL standings
func (c *Client) GetStandings() (*StandingsResponse, error) {
	url := fmt.Sprintf("%s/standings/now", c.baseURL)
	var response StandingsResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get standings: %v", err)
	}
	return &response, nil
}

// GetStandingsByDate returns the NHL standings for a specific date
// date should be in YYYY-MM-DD format
func (c *Client) GetStandingsByDate(date string) (*StandingsResponse, error) {
	url := fmt.Sprintf("%s/standings/%s", c.baseURL, date)
	var response StandingsResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get standings for date %s: %v", date, err)
	}
	return &response, nil
}
