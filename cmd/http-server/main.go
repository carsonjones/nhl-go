package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

type MCPBridge struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout *bufio.Reader
	mu     sync.Mutex
}

func NewMCPBridge(mcpPath string) (*MCPBridge, error) {
	cmd := exec.Command(mcpPath)
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start MCP: %w", err)
	}

	return &MCPBridge{
		cmd:    cmd,
		stdin:  stdin,
		stdout: bufio.NewReader(stdout),
	}, nil
}

func (b *MCPBridge) Call(request []byte) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Write request
	if _, err := b.stdin.Write(append(request, '\n')); err != nil {
		return nil, fmt.Errorf("failed to write to MCP: %w", err)
	}

	// Read response
	line, err := b.stdout.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read from MCP: %w", err)
	}

	return line, nil
}

func (b *MCPBridge) Close() error {
	b.stdin.Close()
	return b.cmd.Wait()
}

func main() {
	mcpPath := os.Getenv("MCP_PATH")
	if mcpPath == "" {
		mcpPath = "/home/exedev/src/nhl-go/nhl-mcp"
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8090"
	}

	bridge, err := NewMCPBridge(mcpPath)
	if err != nil {
		log.Fatalf("failed to start MCP bridge: %v", err)
	}
	defer bridge.Close()

	// Initialize MCP
	initReq := `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"http-bridge","version":"1.0"}}}`
	if _, err := bridge.Call([]byte(initReq)); err != nil {
		log.Fatalf("failed to initialize MCP: %v", err)
	}
	log.Println("MCP initialized")

	http.HandleFunc("/tools", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		req := `{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`
		resp, err := bridge.Call([]byte(req))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)
	})

	http.HandleFunc("/call/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return
		}

		toolName := r.URL.Path[len("/call/"):]
		if toolName == "" {
			http.Error(w, "tool name required", http.StatusBadRequest)
			return
		}

		var args map[string]interface{}
		if r.Method == "POST" && r.Body != nil {
			if err := json.NewDecoder(r.Body).Decode(&args); err != nil && err != io.EOF {
				http.Error(w, fmt.Sprintf("invalid JSON: %v", err), http.StatusBadRequest)
				return
			}
		}
		if args == nil {
			args = make(map[string]interface{})
		}

		// Parse query params as fallback
		for k, v := range r.URL.Query() {
			if len(v) > 0 && args[k] == nil {
				args[k] = v[0]
			}
		}

		req := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      3,
			"method":  "tools/call",
			"params": map[string]interface{}{
				"name":      toolName,
				"arguments": args,
			},
		}

		reqBytes, _ := json.Marshal(req)
		resp, err := bridge.Call(reqBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse and extract just the result content
		var jsonResp struct {
			Result struct {
				Content []struct {
					Text string `json:"text"`
				} `json:"content"`
			} `json:"result"`
			Error *struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.Unmarshal(resp, &jsonResp); err == nil {
			if jsonResp.Error != nil {
				http.Error(w, jsonResp.Error.Message, http.StatusBadRequest)
				return
			}
			if len(jsonResp.Result.Content) > 0 {
				w.Write([]byte(jsonResp.Result.Content[0].Text))
				return
			}
		}
		w.Write(resp)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("HTTP-MCP bridge listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
