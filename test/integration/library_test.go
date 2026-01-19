package integration

import (
	"strings"
	"testing"

	"github.com/modelcontextprotocol/platform.mcp/pkg/scaffold"
)

func TestLibrary_FullFlow(t *testing.T) {
	cfg := scaffold.Config{
		ProjectName:  "integration-test",
		UseDocker:    true,
		WorkflowType: "go",
	}

	files, err := scaffold.Generate(cfg)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Verify Go Workflow
	foundGo := false
	for _, f := range files {
		if f.Path == ".github/workflows/go.yaml" {
			foundGo = true
			if !strings.Contains(f.Content, "Build integration-test") {
				t.Errorf("Go workflow content missing project name, got: %s", f.Content)
			}
		}
	}
	if !foundGo {
		t.Error("Go workflow file not generated")
	}

	// Verify Dockerfile
	foundDocker := false
	for _, f := range files {
		if f.Path == "Dockerfile" {
			foundDocker = true
			if !strings.Contains(f.Content, "cmd/integration-test/main.go") {
				t.Errorf("Dockerfile content missing project name, got: %s", f.Content)
			}
		}
	}
	if !foundDocker {
		t.Error("Dockerfile not generated")
	}
}
