package mcp

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
)

func TestHandleGenerate(t *testing.T) {
	tests := []struct {
		name    string
		input   GenerateInput
		wantErr bool
		check   func(*testing.T, *mcp.CallToolResult)
	}{
		{
			name: "valid request with all options",
			input: GenerateInput{
				ProjectName:  "test-project",
				WorkflowType: "go",
				UseDocker:    true,
				WithActions:  true,
				WithDocker:   true,
				WithFlux:     true,
			},
			wantErr: false,
			check: func(t *testing.T, res *mcp.CallToolResult) {
				assert.False(t, res.IsError)
				assert.NotEmpty(t, res.Content)
				// Basic check for content presence
				var text string
				for _, c := range res.Content {
					if tc, ok := c.(*mcp.TextContent); ok {
						text += tc.Text
					}
				}
				assert.Contains(t, text, "Dockerfile")
				assert.Contains(t, text, ".github/workflows")
				assert.Contains(t, text, "flux-system")
			},
		},
		{
			name: "missing project name",
			input: GenerateInput{
				ProjectName: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &mcp.CallToolRequest{}
			res, _, err := HandleGenerate(context.Background(), req, tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.check(t, res)
			}
		})
	}
}
