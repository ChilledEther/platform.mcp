package mcp

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/modelcontextprotocol/platform.mcp/pkg/scaffold"
)

// GenerateInput defines the input for the generate tool.
type GenerateInput struct {
	ProjectName  string `json:"project_name" jsonschema:"description=The name of the project"`
	UseDocker    bool   `json:"use_docker,omitempty" jsonschema:"description=Whether to use Docker within the project templates"`
	WorkflowType string `json:"workflow_type,omitempty" jsonschema:"description=The type of workflow (go, typescript, python)"`
	WithActions  bool   `json:"with_actions,omitempty" jsonschema:"description=Whether to generate GitHub Actions workflows"`
	WithDocker   bool   `json:"with_docker,omitempty" jsonschema:"description=Whether to generate Dockerfiles"`
	WithFlux     bool   `json:"with_flux,omitempty" jsonschema:"description=Whether to generate Flux CD manifests"`
}

// HandleGenerate implements the generate MCP tool.
func HandleGenerate(ctx context.Context, request *mcp.CallToolRequest, input GenerateInput) (*mcp.CallToolResult, any, error) {
	cfg := scaffold.Config{
		ProjectName:  input.ProjectName,
		UseDocker:    input.UseDocker,
		WorkflowType: input.WorkflowType,
		WithActions:  input.WithActions,
		WithDocker:   input.WithDocker,
		WithFlux:     input.WithFlux,
	}

	generator := scaffold.NewProjectGenerator()
	files, err := generator.Generate(cfg)
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
