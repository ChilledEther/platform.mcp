package scaffold

import (
	"strings"
	"testing"
)

func TestFluxGenerator_Generate(t *testing.T) {
	// Setup
	g := &FluxGenerator{}
	cfg := Config{
		ProjectName: "test-flux",
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
		if f.Path == "fluxcd.yaml" {
			found = true
			if !strings.Contains(f.Content, "test-flux") {
				t.Errorf("Expected content to contain project name 'test-flux', got: %s", f.Content)
			}
		}
	}

	if !found {
		t.Error("Expected fluxcd.yaml to be generated")
	}
}
