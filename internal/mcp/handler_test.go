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
				// Expecting: ci.yaml (with_actions), go.yaml (workflow_go), Dockerfile, docker-build.yaml
				// Note: generate_workflows implicitly sets WithActions=true
				assert.Len(t, res.Content, 4)

				// Verify content exists, order might vary so we check existence in all items
				var text string
				for _, c := range res.Content {
					if tc, ok := c.(*mcp.TextContent); ok {
						text += tc.Text
					}
				}
				assert.Contains(t, text, ".github/workflows/go.yaml")
				assert.Contains(t, text, "Dockerfile")
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
