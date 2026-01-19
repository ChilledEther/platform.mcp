package templates

import (
	"os"
	"testing"
)

func TestTemplatesExist(t *testing.T) {
	files := []string{
		"workflow.yaml.tmpl",
		"docker-build.yaml.tmpl",
	}

	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			t.Errorf("Template file %s not found: %v", f, err)
		}
	}
}
