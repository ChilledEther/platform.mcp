package scaffold

import (
	"strings"
	"testing"
)

func TestDockerGenerator_Generate(t *testing.T) {
	// Setup
	g := &DockerGenerator{}
	cfg := Config{
		ProjectName: "test-service",
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

	expectedFiles := map[string]bool{
		"Dockerfile":        false,
		"docker-build.yaml": false,
	}

	for _, f := range files {
		if _, ok := expectedFiles[f.Path]; ok {
			expectedFiles[f.Path] = true
			if !strings.Contains(f.Content, "test-service") {
				t.Errorf("Expected content of %s to contain project name, got: %s", f.Path, f.Content)
			}
		}
	}

	for path, found := range expectedFiles {
		if !found {
			t.Errorf("Expected %s to be generated", path)
		}
	}
}
