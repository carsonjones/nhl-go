package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	nhlserver "go-nhl/mcp/server"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8090"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost" + addr
	}

	// Create the MCP server
	s := server.NewMCPServer(
		"NHL",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Register all the NHL tools
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

	s.AddTool(slateTool, nhlserver.SlateHandler)
	s.AddTool(playerTool, nhlserver.PlayerHandler)
	s.AddTool(standingsTool, nhlserver.StandingsHandler)
	s.AddTool(rosterTool, nhlserver.RosterHandler)
	s.AddTool(scheduleTool, nhlserver.ScheduleHandler)
	s.AddTool(leadersTool, nhlserver.LeadersHandler)
	s.AddTool(gameTool, nhlserver.GameHandler)
	s.AddTool(liveTool, nhlserver.LiveHandler)
	s.AddTool(teamsTool, nhlserver.TeamsHandler)

	// Create SSE server
	sseServer := server.NewSSEServer(s,
		server.WithBaseURL(baseURL),
		server.WithSSEEndpoint("/sse"),
		server.WithMessageEndpoint("/message"),
	)

	// Create mux for multiple endpoints
	mux := http.NewServeMux()

	// SSE endpoints
	mux.Handle("/sse", sseServer)
	mux.Handle("/message", sseServer)

	// Also support /mcp as SSE endpoint (rewrite path)
	mux.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/sse"
		sseServer.ServeHTTP(w, r)
	})

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// CORS middleware
	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		mux.ServeHTTP(w, r)
	})

	log.Printf("NHL MCP SSE server listening on %s", addr)
	log.Printf("SSE endpoint: %s/sse", baseURL)
	log.Printf("Message endpoint: %s/message", baseURL)
	log.Fatal(http.ListenAndServe(addr, corsHandler))
}
