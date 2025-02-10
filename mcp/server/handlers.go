package server

import (
	"context"
	"encoding/json"
	"fmt"
	nhl "go-nhl/client"
	"go-nhl/internal/formatters"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var (
	SlateHandler server.ToolHandlerFunc = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := nhl.NewClient()

		var date string
		if dateArg, ok := request.Params.Arguments["date"]; ok && dateArg != nil {
			date, ok = dateArg.(string)
			if !ok {
				return nil, fmt.Errorf("if provided, date must be a string in YYYY-MM-DD format")
			}
		}
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}

		result, err := client.GetScheduleByDate(date, nhl.SortByDateDesc)
		if err != nil {
			return nil, err
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshaling response: %v", err)
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}

	PlayerHandler server.ToolHandlerFunc = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := nhl.NewClient()

		nameArg, ok := request.Params.Arguments["name"]
		if !ok || nameArg == nil {
			return nil, fmt.Errorf("name parameter is required")
		}

		searchName, ok := nameArg.(string)
		if !ok {
			return nil, fmt.Errorf("name must be a string")
		}

		players, err := client.SearchPlayer(searchName)
		if err != nil {
			return nil, fmt.Errorf("error searching for player %s: %v", searchName, err)
		}

		if len(players) == 0 {
			return nil, fmt.Errorf("could not find any players matching '%s'", searchName)
		}

		// Get stats for the first player found
		player := players[0]
		result, err := client.GetFilteredPlayerStats(player.PlayerID, nil)
		if err != nil {
			return nil, fmt.Errorf("error getting stats for player %d: %v", player.PlayerID, err)
		}

		response := struct {
			Players []nhl.PlayerSearchResult `json:"players"`
			Stats   []nhl.SeasonTotal        `json:"stats"`
		}{
			Players: players,
			Stats:   result,
		}

		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshaling response: %v", err)
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}

	StandingsHandler server.ToolHandlerFunc = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := nhl.NewClient()

		var date string
		if dateArg, ok := request.Params.Arguments["date"]; ok && dateArg != nil {
			date, ok = dateArg.(string)
			if !ok {
				return nil, fmt.Errorf("if provided, date must be a string in YYYY-MM-DD format")
			}
		}

		var typ string
		if typeArg, ok := request.Params.Arguments["type"]; ok && typeArg != nil {
			typ, ok = typeArg.(string)
			if !ok {
				return nil, fmt.Errorf("if provided, type must be one of: conference, division, league")
			}
		}

		// Get standings first
		var standings *nhl.StandingsResponse
		var err error
		if date != "" {
			standings, err = client.GetStandingsByDate(date)
		} else {
			standings, err = client.GetStandings()
		}
		if err != nil {
			return nil, err
		}

		// Filter by type if specified
		var result interface{}
		switch typ {
		case "conference":
			// Group teams by conference
			conferences := make(map[string][]nhl.StandingsTeam)
			for _, team := range standings.Standings {
				conferences[team.Conference] = append(conferences[team.Conference], team)
			}
			result = conferences
		case "division":
			// Group teams by division
			divisions := make(map[string][]nhl.StandingsTeam)
			for _, team := range standings.Standings {
				divisions[team.Division] = append(divisions[team.Division], team)
			}
			result = divisions
		default:
			// Return league standings (all teams)
			result = standings
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshaling response: %v", err)
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}

	RosterHandler server.ToolHandlerFunc = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := nhl.NewClient()

		teamArg, ok := request.Params.Arguments["team"]
		if !ok || teamArg == nil {
			return nil, fmt.Errorf("team parameter is required")
		}

		team, ok := teamArg.(string)
		if !ok {
			return nil, fmt.Errorf("team must be a string")
		}

		result, err := client.GetTeamRoster(team)
		if err != nil {
			return nil, fmt.Errorf("error getting team roster for %s: %v", team, err)
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshaling response: %v", err)
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}

	ScheduleHandler server.ToolHandlerFunc = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := nhl.NewClient()

		teamArg, ok := request.Params.Arguments["team"]
		if !ok || teamArg == nil {
			return nil, fmt.Errorf("team parameter is required")
		}

		team, ok := teamArg.(string)
		if !ok {
			return nil, fmt.Errorf("team must be a string")
		}

		var seasonID int
		if seasonIDArg, ok := request.Params.Arguments["seasonID"]; ok && seasonIDArg != nil {
			seasonID, ok = seasonIDArg.(int)
			if !ok {
				return nil, fmt.Errorf("if provided, seasonID must be an integer")
			}
		}

		if seasonID == 0 {
			seasonID = formatters.GetCurrentSeasonID()
		}

		teamInfo, err := client.GetTeamByIdentifier(team)
		if err != nil {
			return nil, fmt.Errorf("error getting team: %v", err)
		}

		result, err := client.GetTeamSchedule(teamInfo, seasonID)
		if err != nil {
			return nil, fmt.Errorf("error getting schedule for season %d: %v", seasonID, err)
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshaling response: %v", err)
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}

	LeadersHandler server.ToolHandlerFunc = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := nhl.NewClient()

		var seasonID int
		if seasonIDArg, ok := request.Params.Arguments["seasonID"]; ok && seasonIDArg != nil {
			seasonID, ok = seasonIDArg.(int)
			if !ok {
				return nil, fmt.Errorf("if provided, seasonID must be an integer")
			}
		}

		if seasonID == 0 {
			seasonID = formatters.GetCurrentSeasonID()
		}

		result, err := client.GetStatsLeaders(seasonID)
		if err != nil {
			return nil, fmt.Errorf("error getting leaders: %v", err)
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshaling response: %v", err)
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}
)
