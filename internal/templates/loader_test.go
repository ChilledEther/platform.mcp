package templates

import (
	"os"
	"path/filepath"
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

func TestGetManifest(t *testing.T) {
	manifest, err := GetManifest()
	if err != nil {
		t.Fatalf("Expected no error getting manifest, got %v", err)
	}

	if len(manifest.Templates) == 0 {
		t.Error("Expected templates in manifest, got none")
	}

	foundGo := false
	for _, tmpl := range manifest.Templates {
		if tmpl.Name == "go-workflow" {
			foundGo = true
			break
		}
	}

	if !foundGo {
		t.Error("Expected go-workflow in manifest")
	}
}

func TestExternalLoading(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "templates-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	testTmpl := "external.tmpl"
	content := "external content"
	if err := os.WriteFile(filepath.Join(tmpDir, testTmpl), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// Set external base dir
	origBaseDir := BaseDir
	SetBaseDir(tmpDir)
	defer SetBaseDir(origBaseDir)

	// Test loading external
	loaded, err := Load(testTmpl)
	if err != nil {
		t.Fatalf("Failed to load external template: %v", err)
	}
	if loaded != content {
		t.Errorf("Expected %q, got %q", content, loaded)
	}

	// Test fallback to embedded
	embedded, err := Load("go.yaml.tmpl")
	if err != nil {
		t.Fatalf("Failed to load embedded template with BaseDir set: %v", err)
	}
	if len(embedded) == 0 {
		t.Error("Expected embedded content, got empty")
	}
}
