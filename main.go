package main

import (
	"fmt"

	"github.com/RRonCChen/node-red-mcp-go/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"NodeRed MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// init mcp tools
	InitTools(s)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func InitTools(server *server.MCPServer) {
	server.AddTool(tools.NewGetFlowTool(), tools.GetFlowsHandlerFunc)
	server.AddTool(tools.NewGetFlowsStateTool(), tools.GetFlowsStateHandlerFunc)
	server.AddTool(tools.NewGetFlowByIdTool(), tools.GetFlowsByIdHandlerFunc)
}
