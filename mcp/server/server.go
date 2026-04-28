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

	slateTool := mcp.NewTool("nhl-slate",
		mcp.WithDescription("Get slate of games for a given date"),
		mcp.WithString("date",
			mcp.Required(),
			mcp.Description("Date (YYYY-MM-DD format)"),
		),
	)

	playerTool := mcp.NewTool("nhl-player",
		mcp.WithDescription("Get player info and stats"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Player name"),
		),
	)

	standingsTool := mcp.NewTool("nhl-standings",
		mcp.WithDescription("Get standings"),
		mcp.WithString("date",
			mcp.Description("Date (YYYY-MM-DD format)"),
		),
		mcp.WithString("type",
			mcp.Description("Standings type (conference, division, league)"),
			mcp.DefaultString("league"),
		),
	)

	rosterTool := mcp.NewTool("nhl-roster",
		mcp.WithDescription("Get team roster"),
		mcp.WithString("team",
			mcp.Required(),
			mcp.Description("Team abbreviation"),
		),
	)

	scheduleTool := mcp.NewTool("nhl-schedule",
		mcp.WithDescription("Get team schedule"),
		mcp.WithString("team",
			mcp.Required(),
			mcp.Description("Team abbreviation"),
		),
		mcp.WithNumber("seasonID",
			mcp.Description("Season ID (example: 20242025)"),
		),
	)

	leadersTool := mcp.NewTool("nhl-leaders",
		mcp.WithDescription("Get leaders"),
		// mcp.WithString("type",
		// 	mcp.Description("Leader type (points, goals, assists, etc.)"),
		// ),
		mcp.WithString("seasonID",
			mcp.Description("Season ID (example: 20242025)"),
		),
	)

	gameTool := mcp.NewTool("nhl-game",
		mcp.WithDescription("Get detailed game information including boxscore, play-by-play, and game story"),
		mcp.WithNumber("gameId",
			mcp.Required(),
			mcp.Description("Game ID"),
		),
		mcp.WithString("include",
			mcp.Description("What to include: details, boxscore, plays, story, or all (default: details)"),
			mcp.DefaultString("details"),
		),
	)

	liveTool := mcp.NewTool("nhl-live",
		mcp.WithDescription("Get live game updates and current scoreboard"),
	)

	teamsTool := mcp.NewTool("nhl-teams",
		mcp.WithDescription("Get list of all NHL teams"),
	)

	s.AddTool(slateTool, SlateHandler)
	s.AddTool(playerTool, PlayerHandler)
	s.AddTool(standingsTool, StandingsHandler)
	s.AddTool(rosterTool, RosterHandler)
	s.AddTool(scheduleTool, ScheduleHandler)
	s.AddTool(leadersTool, LeadersHandler)
	s.AddTool(gameTool, GameHandler)
	s.AddTool(liveTool, LiveHandler)
	s.AddTool(teamsTool, TeamsHandler)

	// Start the stdio server
	return server.ServeStdio(s)
}
