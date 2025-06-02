package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"NodeRed MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
