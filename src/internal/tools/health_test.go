package tools

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestHealthCheckHandler(t *testing.T) {
	result, _, err := HealthCheckHandler(context.Background(), nil, struct{}{})
	if err != nil {
		t.Fatalf("HealthCheckHandler returned error: %v", err)
	}

	if len(result.Content) != 1 {
		t.Errorf("Expected 1 content item, got %d", len(result.Content))
	}

	textContent, ok := result.Content[0].(*mcp.TextContent)
	if !ok {
		t.Errorf("Expected TextContent, got %T", result.Content[0])
	}

	if textContent.Text != "OK" {
		t.Errorf("Expected 'OK', got '%s'", textContent.Text)
	}
}
