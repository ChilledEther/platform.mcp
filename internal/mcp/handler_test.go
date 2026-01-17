package mcp

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
)

func TestHandleGenerateWorkflows(t *testing.T) {
	tests := []struct {
		name    string
		input   GenerateWorkflowsInput
		wantErr bool
		check   func(*testing.T, *mcp.CallToolResult)
	}{
		{
			name: "valid request",
			input: GenerateWorkflowsInput{
				ProjectName:  "test-project",
				WorkflowType: "go",
				UseDocker:    true,
			},
			wantErr: false,
			check: func(t *testing.T, res *mcp.CallToolResult) {
				assert.False(t, res.IsError)
				assert.Len(t, res.Content, 2)
				assert.Contains(t, res.Content[0].(*mcp.TextContent).Text, ".github/workflows/go.yaml")
				assert.Contains(t, res.Content[1].(*mcp.TextContent).Text, "Dockerfile")
			},
		},
		{
			name: "invalid workflow_type (handled by scaffold.Generate)",
			input: GenerateWorkflowsInput{
				ProjectName:  "test-project",
				WorkflowType: "rust",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &mcp.CallToolRequest{}
			res, _, err := HandleGenerateWorkflows(context.Background(), req, tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.check(t, res)
			}
		})
	}
}
