package scaffold

import (
	"strings"
	"testing"
)

func TestActionsGenerator_Generate(t *testing.T) {
	// Setup
	g := &ActionsGenerator{
		FS: nil, // Will need to mock or use real embed in implementation
	}
	cfg := Config{
		ProjectName: "test-project",
	}

	// Execute
	files, err := g.Generate(cfg)

	// Verify
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(files) == 0 {
		t.Error("Expected files to be generated, got 0")
	}

	found := false
	for _, f := range files {
		if f.Path == ".github/workflows/ci.yaml" {
			found = true
			if !strings.Contains(f.Content, "test-project") {
				t.Errorf("Expected content to contain project name 'test-project', got: %s", f.Content)
			}
		}
	}

	if !found {
		t.Error("Expected .github/workflows/ci.yaml to be generated")
	}
}
