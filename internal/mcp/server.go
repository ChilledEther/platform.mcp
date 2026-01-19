package mcp

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// NewServer creates and initializes a new MCP server with the specified configuration.
func NewServer(version string) *mcp.Server {
	return mcp.NewServer(
		&mcp.Implementation{
			Name:    "platform-mcp",
			Version: version,
		},
		nil, // Default options
	)
}

// RegisterTools adds all available tools to the server instance.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "generate_workflows",
		Description: "Generate GitHub Actions workflows for a project",
	}, HandleGenerateWorkflows)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "generate",
		Description: "Generate project scaffolding including Actions, Docker, and Flux",
	}, HandleGenerate)
}
