package nhl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	BaseURLWeb = "https://api-web.nhle.com/v1"
)

// Client represents an NHL API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	teamCache  *TeamsResponse
	cacheMutex sync.RWMutex
}

// NewClient creates a new NHL API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		baseURL: BaseURLWeb,
	}
}

func (c *Client) get(url string, target interface{}) error {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d\nResponse: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// GetScheduleByDate retrieves the NHL schedule for a specific date
// date should be in YYYY-MM-DD format
// sortOrder determines the order of games (defaults to SortByDateAsc if not specified)
func (c *Client) GetScheduleByDate(date string, sortOrder ...GameSort) (*FilteredScoreboardResponse, error) {
	var scores ScoreboardResponse
	err := c.get(fmt.Sprintf("%s/scoreboard/%s", c.baseURL, date), &scores)
	if err != nil {
		return nil, fmt.Errorf("error getting schedule for date %s: %v", date, err)
	}

	// Filter games for the specific date
	filteredGames := []Game{}
	for _, gamesByDate := range scores.GamesByDate {
		if gamesByDate.Date == date {
			filteredGames = append(filteredGames, gamesByDate.Games...)
		}
	}

	// Use default sort order (SortByDateAsc) if none specified
	actualSortOrder := SortByDateAsc
	if len(sortOrder) > 0 {
		actualSortOrder = sortOrder[0]
	}

	// Sort games
	sort.Slice(filteredGames, func(i, j int) bool {
		timeI, errI := time.Parse(time.RFC3339, filteredGames[i].StartTimeUTC)
		timeJ, errJ := time.Parse(time.RFC3339, filteredGames[j].StartTimeUTC)

		// If we can't parse either time, maintain original order
		if errI != nil || errJ != nil {
			return i < j
		}

		if actualSortOrder == SortByDateAsc {
			return timeI.Before(timeJ)
		}
		return timeI.After(timeJ)
	})

	return &FilteredScoreboardResponse{
		Date:  date,
		Games: filteredGames,
	}, nil
}

// GetCurrentSchedule retrieves today's NHL schedule
// sortOrder determines the order of games (defaults to SortByDateAsc if not specified)
func (c *Client) GetCurrentSchedule(sortOrder ...GameSort) (*FilteredScoreboardResponse, error) {
	today := time.Now().Format("2006-01-02")
	return c.GetScheduleByDate(today, sortOrder...)
}

// GetTeams retrieves all NHL teams
func (c *Client) GetTeams() (*TeamsResponse, error) {
	// Check cache first
	c.cacheMutex.RLock()
	if c.teamCache != nil {
		defer c.cacheMutex.RUnlock()
		return c.teamCache, nil
	}
	c.cacheMutex.RUnlock()

	// Cache miss, fetch teams
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	// Double-check cache in case another goroutine populated it
	if c.teamCache != nil {
		return c.teamCache, nil
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

	c.teamCache = teams
	return teams, nil
}

// GetTeamByIdentifier finds a team by ID, abbreviation, or name
func (c *Client) GetTeamByIdentifier(identifier string) (*TeamInfo, error) {
	teams, err := c.GetTeams()
	if err != nil {
		return nil, err
	}

	// Try to parse identifier as team ID
	var teamID int
	if _, err := fmt.Sscanf(identifier, "%d", &teamID); err == nil {
		for _, team := range teams.Teams {
			if team.ID == teamID {
				return &team, nil
			}
		}
	}

	// Try to match by abbreviation or name
	identifierLower := strings.ToLower(identifier)
	for _, team := range teams.Teams {
		if strings.ToLower(team.Abbreviation) == identifierLower ||
			strings.ToLower(team.TriCode) == identifierLower ||
			strings.ToLower(team.Name.Default) == identifierLower {
			return &team, nil
		}
	}

	return nil, fmt.Errorf("team not found: %s", identifier)
}

// GetTeamRoster retrieves the roster for a team by ID, abbreviation, or name
func (c *Client) GetTeamRoster(identifier string) (*RosterResponse, error) {
	// First try to get the team info to validate and normalize the identifier
	team, err := c.GetTeamByIdentifier(identifier)
	if err != nil {
		return nil, fmt.Errorf("invalid team identifier %s: %v", identifier, err)
	}

	var roster RosterResponse
	err = c.get(fmt.Sprintf("%s/roster/%s/current", c.baseURL, team.Abbreviation), &roster)
	if err != nil {
		return nil, fmt.Errorf("error getting roster for team %s: %v", team.Name.Default, err)
	}

	return &roster, nil
}

// SearchPlayer searches for players by name (first, last, or full name)
func (c *Client) SearchPlayer(name string) ([]PlayerSearchResult, error) {
	// First try to find the player in the current rosters
	teams, err := c.GetTeams()
	if err != nil {
		return nil, fmt.Errorf("error getting teams: %v", err)
	}

	var results []PlayerSearchResult
	nameLower := strings.ToLower(name)

	// Search through all team rosters
	for _, team := range teams.Teams {
		roster, err := c.GetTeamRoster(team.Abbreviation)
		if err != nil {
			continue // Skip teams with errors
		}

		// Helper function to check if a player matches the search
		matchesSearch := func(player PlayerInfo) bool {
			firstLower := strings.ToLower(player.FirstName.Default)
			lastLower := strings.ToLower(player.LastName.Default)
			fullLower := firstLower + " " + lastLower

			return strings.Contains(firstLower, nameLower) ||
				strings.Contains(lastLower, nameLower) ||
				strings.Contains(fullLower, nameLower)
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

// GetPlayerStats retrieves statistics for a player
// reportType can be "summary" for regular season stats
func (c *Client) GetPlayerStats(playerID int, isGoalie bool, reportType string, filter *StatsFilter) (interface{}, error) {
	playerType := "skater"
	if isGoalie {
		playerType = "goalie"
	}

	// Build the cayenneExp query
	cayenneExp := fmt.Sprintf("playerId=%d", playerID)
	if filter != nil {
		if filter.GameType != 0 {
			cayenneExp += fmt.Sprintf(" and gameType=%d", filter.GameType)
		}
		if filter.SeasonID != 0 {
			cayenneExp += fmt.Sprintf(" and seasonId=%d", filter.SeasonID)
		}
	}

	// First get the generic response
	var genericResp StatsResponse
	url := fmt.Sprintf("https://api.nhle.com/stats/rest/en/%s/%s?cayenneExp=%s",
		playerType, reportType, url.QueryEscape(cayenneExp))
	err := c.get(url, &genericResp)
	if err != nil {
		return nil, fmt.Errorf("error getting stats for player %d: %v", playerID, err)
	}

	// Marshal the data back to JSON
	dataJSON, err := json.Marshal(genericResp)
	if err != nil {
		return nil, fmt.Errorf("error marshaling stats data: %v", err)
	}

	// Unmarshal into the appropriate type
	if isGoalie {
		var goalieResp GoalieStatsResponse
		if err := json.Unmarshal(dataJSON, &goalieResp); err != nil {
			return nil, fmt.Errorf("error unmarshaling goalie stats: %v", err)
		}
		return &goalieResp, nil
	}

	var skaterResp SkaterStatsResponse
	if err := json.Unmarshal(dataJSON, &skaterResp); err != nil {
		return nil, fmt.Errorf("error unmarshaling skater stats: %v", err)
	}
	return &skaterResp, nil
}

// GetPlayerSeasonStats retrieves a player's stats for all seasons and game types
func (c *Client) GetPlayerSeasonStats(playerID int) (*PlayerLandingResponse, error) {
	var landing PlayerLandingResponse
	err := c.get(fmt.Sprintf("%s/player/%d/landing", c.baseURL, playerID), &landing)
	if err != nil {
		return nil, fmt.Errorf("error getting player landing page: %v", err)
	}

	return &landing, nil
}

// GetFilteredPlayerStats retrieves a player's stats filtered by game type and/or season
func (c *Client) GetFilteredPlayerStats(playerID int, filter *StatsFilter) ([]SeasonTotal, error) {
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

// GetTeamSchedule retrieves a team's schedule for a given season
func (c *Client) GetTeamSchedule(team *TeamInfo, seasonID int) (*TeamScheduleResponse, error) {
	if team == nil {
		return nil, fmt.Errorf("team cannot be nil")
	}

	url := fmt.Sprintf("%s/club-schedule-season/%s/%d", c.baseURL, team.Abbreviation, seasonID)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule for %s: %w", team.Name.Default, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get schedule for %s: %s", team.Name.Default, resp.Status)
	}

	var schedule TeamScheduleResponse
	if err := json.NewDecoder(resp.Body).Decode(&schedule); err != nil {
		return nil, fmt.Errorf("failed to decode schedule response: %w", err)
	}

	return &schedule, nil
}
