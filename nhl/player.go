package nhl

import (
	"fmt"
	"sort"
	"strings"
)

// SearchPlayer searches for players by name
func (c *Client) SearchPlayer(name string) ([]PlayerSearchResult, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	// Get all teams to search through rosters
	teams, err := c.GetTeams()
	if err != nil {
		return nil, fmt.Errorf("failed to get teams: %v", err)
	}

	var results []PlayerSearchResult
	name = strings.ToLower(name)

	// Search through each team's roster
	for _, team := range teams.Teams {
		roster, err := c.GetTeamRoster(team.Abbreviation)
		if err != nil {
			continue // Skip teams with errors
		}

		// Helper function to check if a player matches the search
		matchesSearch := func(player PlayerInfo) bool {
			return strings.Contains(strings.ToLower(player.FirstName.Default), name) ||
				strings.Contains(strings.ToLower(player.LastName.Default), name)
		}

		// Check forwards
		for _, player := range roster.Forwards {
			if matchesSearch(player) {
				results = append(results, PlayerSearchResult{
					FirstName:    player.FirstName,
					LastName:     player.LastName,
					Position:     player.Position,
					JerseyNumber: player.JerseyNumber,
					TeamID:       team.ID,
					TeamAbbrev:   team.Abbreviation,
					PlayerID:     player.ID,
				})
			}
		}

		// Check defensemen
		for _, player := range roster.Defensemen {
			if matchesSearch(player) {
				results = append(results, PlayerSearchResult{
					FirstName:    player.FirstName,
					LastName:     player.LastName,
					Position:     player.Position,
					JerseyNumber: player.JerseyNumber,
					TeamID:       team.ID,
					TeamAbbrev:   team.Abbreviation,
					PlayerID:     player.ID,
				})
			}
		}

		// Check goalies
		for _, player := range roster.Goalies {
			if matchesSearch(player) {
				results = append(results, PlayerSearchResult{
					FirstName:    player.FirstName,
					LastName:     player.LastName,
					Position:     player.Position,
					JerseyNumber: player.JerseyNumber,
					TeamID:       team.ID,
					TeamAbbrev:   team.Abbreviation,
					PlayerID:     player.ID,
				})
			}
		}
	}

	return results, nil
}

// GetPlayerStats returns stats for a player
func (c *Client) GetPlayerStats(playerID int, isGoalie bool, reportType string, filter *StatsFilter) (interface{}, error) {
	if playerID <= 0 {
		return nil, fmt.Errorf("invalid player ID: %d", playerID)
	}

	// Validate report type
	validReportTypes := map[string]bool{
		"regularSeason": true,
		"playoffs":      true,
	}
	if !validReportTypes[reportType] {
		return nil, fmt.Errorf("invalid report type: %s", reportType)
	}

	url := fmt.Sprintf("%s/player/%d/landing", c.baseURL, playerID)
	if filter != nil {
		if filter.SeasonID > 0 {
			url = fmt.Sprintf("%s/player/%d/stats/%d", c.baseURL, playerID, filter.SeasonID)
		}
		if filter.GameType != 0 {
			url = fmt.Sprintf("%s?cayenneExp=gameTypeId=%d", url, filter.GameType)
		}
	}

	if isGoalie {
		var response GoalieStatsResponse
		err := c.get(url, &response)
		if err != nil {
			return nil, fmt.Errorf("failed to get goalie stats: %v", err)
		}
		return &response, nil
	}

	var response SkaterStatsResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get skater stats: %v", err)
	}
	return &response, nil
}

// GetPlayerSeasonStats returns a player's stats for all seasons
func (c *Client) GetPlayerSeasonStats(playerID int) (*PlayerLandingResponse, error) {
	if playerID <= 0 {
		return nil, fmt.Errorf("invalid player ID: %d", playerID)
	}

	url := fmt.Sprintf("%s/player/%d/landing", c.baseURL, playerID)
	var response PlayerLandingResponse
	err := c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get player season stats: %v", err)
	}

	return &response, nil
}

// GetFilteredPlayerStats returns filtered stats for a player
func (c *Client) GetFilteredPlayerStats(playerID int, filter *StatsFilter) ([]SeasonTotal, error) {
	if playerID <= 0 {
		return nil, fmt.Errorf("invalid player ID: %d", playerID)
	}

	landing, err := c.GetPlayerSeasonStats(playerID)
	if err != nil {
		return nil, err
	}

	var filtered []SeasonTotal
	for _, season := range landing.SeasonTotals {
		// Skip non-NHL seasons
		if season.LeagueAbbrev != "NHL" {
			continue
		}

		// Apply filters
		if filter != nil {
			if filter.GameType != 0 && season.GameTypeID != int(filter.GameType) {
				continue
			}
			if filter.SeasonID != 0 && season.Season != filter.SeasonID {
				continue
			}
		}

		filtered = append(filtered, season)
	}

	// Sort by season in descending order (most recent first)
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Season > filtered[j].Season
	})

	return filtered, nil
}
