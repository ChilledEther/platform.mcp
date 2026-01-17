package tools

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Registry is a wrapper around the MCP server to manage tool registration.
type Registry struct {
	server *mcp.Server
}

// NewRegistry creates a new Registry wrapper.
func NewRegistry(server *mcp.Server) *Registry {
	return &Registry{server: server}
}

// Register registers a tool with the server. It wraps the handler with UUID tracking.
func Register[In any](r *Registry, tool *mcp.Tool, handler func(context.Context, *mcp.CallToolRequest, In) (*mcp.CallToolResult, any, error)) {
	wrappedHandler := func(ctx context.Context, req *mcp.CallToolRequest, in In) (*mcp.CallToolResult, any, error) {
		executionID := uuid.New().String()
		fmt.Fprintf(os.Stderr, "[Tool:%s] [ID:%s] Starting execution...\n", tool.Name, executionID)

		result, meta, err := handler(ctx, req, in)

		if err != nil {
			fmt.Fprintf(os.Stderr, "[Tool:%s] [ID:%s] Execution failed: %v\n", tool.Name, executionID, err)
			return result, meta, err
		}

		if result == nil {
			result = &mcp.CallToolResult{}
		}

		// Ensure Meta is initialized and add the execution ID
		if result.Meta == nil {
			result.Meta = make(map[string]any)
		}
		result.Meta["execution_id"] = executionID

		fmt.Fprintf(os.Stderr, "[Tool:%s] [ID:%s] Execution completed successfully.\n", tool.Name, executionID)
		return result, meta, nil
	}

	mcp.AddTool(r.server, tool, wrappedHandler)
}
