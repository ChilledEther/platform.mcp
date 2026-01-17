package scaffold

import (
	"strings"
	"testing"
)

func TestTemplates_Embedded(t *testing.T) {
	// This test will fail until templates are actually embedded and used
	cfg := Config{
		ProjectName:  "template-test",
		WorkflowType: "go",
	}

	files, err := Generate(cfg)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	for _, f := range files {
		if f.Path == ".github/workflows/go.yaml" {
			if !strings.Contains(f.Content, "Build template-test") {
				t.Errorf("Template rendering failed, content: %s", f.Content)
			}
		}
	}
}
