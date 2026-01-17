package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HealthCheckHandler handles the health_check tool execution.
func HealthCheckHandler(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
}

// RegisterHealthTool registers the health_check tool with the registry.
func RegisterHealthTool(r *Registry) {
	Register(r, &mcp.Tool{
		Name:        "health_check",
		Description: "Basic health check to verify server is operational",
	}, HealthCheckHandler)
}
