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

func TestLoad(t *testing.T) {
	// Test existing template
	content, err := Load("workflow.yaml.tmpl")
	if err != nil {
		t.Fatalf("Expected no error loading workflow.yaml.tmpl, got %v", err)
	}
	if len(content) == 0 {
		t.Error("Expected content, got empty string")
	}

	// Test missing template
	_, err = Load("nonexistent.tmpl")
	if err == nil {
		t.Error("Expected error for missing template, got nil")
	}
}

func TestRender(t *testing.T) {
	type TestData struct {
		ProjectName string
	}

	data := TestData{ProjectName: "MyProject"}
	tmpl := "Hello {{ .ProjectName }}"

	// Test success
	output, err := Render(tmpl, data)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	expected := "Hello MyProject"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	// Test invalid template
	_, err = Render("Hello {{ .Missing }", data)
	if err == nil {
		t.Error("Expected error for invalid template syntax, got nil")
	}
}
