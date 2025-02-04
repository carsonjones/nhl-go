package server

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Start() error {
	s := server.NewMCPServer(
		"NHL",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	slateTool := mcp.NewTool("slate",
		mcp.WithDescription("Get slate of games for a given date"),
		mcp.WithString("date",
			mcp.Required(),
			mcp.Description("Date (YYYY-MM-DD format)"),
		),
	)

	playerTool := mcp.NewTool("player",
		mcp.WithDescription("Get player info and stats"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Player name"),
		),
	)

	standingsTool := mcp.NewTool("standings",
		mcp.WithDescription("Get standings"),
		mcp.WithString("date",
			mcp.Description("Date (YYYY-MM-DD format)"),
		),
		mcp.WithString("type",
			mcp.Description("Standings type (conference, division, league)"),
			mcp.DefaultString("league"),
		),
	)

	rosterTool := mcp.NewTool("roster",
		mcp.WithDescription("Get team roster"),
		mcp.WithString("team",
			mcp.Required(),
			mcp.Description("Team abbreviation"),
		),
	)

	scheduleTool := mcp.NewTool("schedule",
		mcp.WithDescription("Get team schedule"),
		mcp.WithString("team",
			mcp.Required(),
			mcp.Description("Team abbreviation"),
		),
		mcp.WithNumber("seasonID",
			mcp.Description("Season ID (example: 20242025)"),
		),
	)

	s.AddTool(slateTool, SlateHandler)
	s.AddTool(playerTool, PlayerHandler)
	s.AddTool(standingsTool, StandingsHandler)
	s.AddTool(rosterTool, RosterHandler)
	s.AddTool(scheduleTool, ScheduleHandler)

	// Start the stdio server
	return server.ServeStdio(s)
}
