package mcp

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/modelcontextprotocol/platform.mcp/pkg/scaffold"
)

// GenerateWorkflowsInput defines the input for the generate_workflows tool.
type GenerateWorkflowsInput struct {
	ProjectName  string `json:"project_name" jsonschema:"The name of the project for which to generate workflows."`
	UseDocker    bool   `json:"use_docker" jsonschema:"Whether to include Docker-related workflow steps."`
	WorkflowType string `json:"workflow_type" jsonschema:"The type of workflow to generate (go, typescript, python)."`
}

// HandleGenerateWorkflows implements the generate_workflows MCP tool.
func HandleGenerateWorkflows(ctx context.Context, request *mcp.CallToolRequest, input GenerateWorkflowsInput) (*mcp.CallToolResult, any, error) {
	cfg := scaffold.Config{
		ProjectName:  input.ProjectName,
		UseDocker:    input.UseDocker,
		WorkflowType: input.WorkflowType,
	}

	files, err := scaffold.Generate(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("generation failed: %w", err)
	}

	var content []mcp.Content
	for _, f := range files {
		content = append(content, &mcp.TextContent{
			Text: fmt.Sprintf("--- FILE: %s ---\n%s", f.Path, f.Content),
		})
	}

	return &mcp.CallToolResult{
		Content: content,
	}, nil, nil
}
