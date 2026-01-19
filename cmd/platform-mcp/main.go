package main

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	internalmcp "github.com/modelcontextprotocol/platform.mcp/internal/mcp"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// 1. Initialize MCP server
	server := internalmcp.NewServer("0.1.0")

	// 2. Register tools
	internalmcp.RegisterTools(server)

	// 3. Start server with stdio transport
	fmt.Fprintf(os.Stderr, "platform-mcp server starting...\n")
	return server.Run(ctx, &mcp.StdioTransport{})
}
