package nhl

import (
	"fmt"
	"time"
)

// GetCurrentSchedule returns the schedule for the current day
func (c *Client) GetCurrentSchedule() (*FilteredScoreboardResponse, error) {
	today := time.Now().Format("2006-01-02")
	return c.GetScheduleByDate(today, SortByDateAsc)
}

// GetScheduleByDate returns the schedule for a specific date
// date should be in YYYY-MM-DD format
// sortOrder can be either SortByDateAsc or SortByDateDesc
func (c *Client) GetScheduleByDate(date string, sortOrder SortOrder) (*FilteredScoreboardResponse, error) {
	url := fmt.Sprintf("%s/score/%s", c.baseURL, date)
	if sortOrder == SortByDateDesc {
		url += "?sort=desc"
	}

	var response FilteredScoreboardResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule: %v", err)
	}

	// Set the date in the response
	response.Date = date

	return &response, nil
}

// GetTeamSchedule returns the schedule for a specific team and season
func (c *Client) GetTeamSchedule(team *TeamInfo, seasonID int) (*TeamScheduleResponse, error) {
	if team == nil {
		return nil, fmt.Errorf("team cannot be nil")
	}

	if seasonID <= 0 {
		return nil, fmt.Errorf("invalid season ID: %d", seasonID)
	}

	url := fmt.Sprintf("%s/club-schedule-season/%s/%d", c.baseURL, team.Abbreviation, seasonID)

	var response TeamScheduleResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get team schedule: %v", err)
	}

	return &response, nil
}
