package nhl

import (
	"fmt"
	"strconv"
	"strings"
)

// GetTeams returns a list of all NHL teams
func (c *Client) GetTeams() (*TeamsResponse, error) {
	// Check cache first
	c.cacheMutex.RLock()
	if c.teams != nil {
		defer c.cacheMutex.RUnlock()
		return c.teams, nil
	}
	c.cacheMutex.RUnlock()

	// Cache miss, fetch teams
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	// Double-check cache in case another goroutine populated it
	if c.teams != nil {
		return c.teams, nil
	}

	// Known NHL teams with their IDs, abbreviations, and full names
	teamData := []struct {
		ID   int
		Abbr string
		Name string
		City string
	}{
		{1, "NJD", "New Jersey Devils", "New Jersey"},
		{2, "NYI", "New York Islanders", "New York"},
		{3, "NYR", "New York Rangers", "New York"},
		{4, "PHI", "Philadelphia Flyers", "Philadelphia"},
		{5, "PIT", "Pittsburgh Penguins", "Pittsburgh"},
		{6, "BOS", "Boston Bruins", "Boston"},
		{7, "BUF", "Buffalo Sabres", "Buffalo"},
		{8, "MTL", "Montreal Canadiens", "Montreal"},
		{9, "OTT", "Ottawa Senators", "Ottawa"},
		{10, "TOR", "Toronto Maple Leafs", "Toronto"},
		{12, "CAR", "Carolina Hurricanes", "Carolina"},
		{13, "FLA", "Florida Panthers", "Florida"},
		{14, "TBL", "Tampa Bay Lightning", "Tampa Bay"},
		{15, "WSH", "Washington Capitals", "Washington"},
		{16, "CHI", "Chicago Blackhawks", "Chicago"},
		{17, "DET", "Detroit Red Wings", "Detroit"},
		{18, "NSH", "Nashville Predators", "Nashville"},
		{19, "STL", "St. Louis Blues", "St. Louis"},
		{20, "CGY", "Calgary Flames", "Calgary"},
		{21, "COL", "Colorado Avalanche", "Colorado"},
		{22, "EDM", "Edmonton Oilers", "Edmonton"},
		{23, "VAN", "Vancouver Canucks", "Vancouver"},
		{24, "ANA", "Anaheim Ducks", "Anaheim"},
		{25, "DAL", "Dallas Stars", "Dallas"},
		{26, "LAK", "Los Angeles Kings", "Los Angeles"},
		{28, "SJS", "San Jose Sharks", "San Jose"},
		{29, "CBJ", "Columbus Blue Jackets", "Columbus"},
		{30, "MIN", "Minnesota Wild", "Minnesota"},
		{52, "WPG", "Winnipeg Jets", "Winnipeg"},
		{53, "ARI", "Arizona Coyotes", "Arizona"},
		{54, "VGK", "Vegas Golden Knights", "Vegas"},
		{55, "SEA", "Seattle Kraken", "Seattle"},
	}

	teams := &TeamsResponse{
		Teams: []TeamInfo{},
	}

	for _, td := range teamData {
		teams.Teams = append(teams.Teams, TeamInfo{
			ID:           td.ID,
			Abbreviation: td.Abbr,
			TriCode:      td.Abbr,
			Name: LanguageNames{
				Default: td.Name,
			},
			City: LanguageNames{
				Default: td.City,
			},
			Active: true,
		})
	}

	c.teams = teams
	return teams, nil
}

// GetTeamByIdentifier returns a team by its identifier (abbreviation, name, or ID)
func (c *Client) GetTeamByIdentifier(identifier string) (*TeamInfo, error) {
	teams, err := c.GetTeams()
	if err != nil {
		return nil, err
	}

	// Try to parse as team ID first
	if id, err := strconv.Atoi(identifier); err == nil {
		for _, team := range teams.Teams {
			if team.ID == id {
				return &team, nil
			}
		}
	}

	// Try to match by abbreviation or name
	identifier = strings.ToLower(identifier)
	for _, team := range teams.Teams {
		if strings.ToLower(team.Abbreviation) == identifier ||
			strings.ToLower(team.Name.Default) == identifier {
			return &team, nil
		}
	}

	return nil, fmt.Errorf("team not found: %s", identifier)
}

// GetTeamRoster returns the current roster for a team
func (c *Client) GetTeamRoster(identifier string) (*RosterResponse, error) {
	team, err := c.GetTeamByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/roster/%s/current", c.baseURL, team.Abbreviation)
	var response RosterResponse
	err = c.get(url, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get roster: %v", err)
	}

	return &response, nil
}
