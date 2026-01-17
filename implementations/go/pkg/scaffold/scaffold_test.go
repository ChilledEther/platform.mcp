package scaffold

import (
	"os"
	"strings"
	"testing"
)

func TestGenerate_Basic(t *testing.T) {
	cfg := Config{
		ProjectName:  "test-project",
		WorkflowType: "go",
	}

	files, err := Generate(cfg)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if len(files) == 0 {
		t.Fatal("Expected files to be generated, got none")
	}

	foundWorkflow := false
	for _, f := range files {
		if f.Path == "" {
			t.Error("Generated file has empty path")
		}
		if f.Content == "" {
			t.Errorf("File %s has empty content", f.Path)
		}
		if strings.HasSuffix(f.Path, ".yaml") || strings.HasSuffix(f.Path, ".yml") {
			foundWorkflow = true
		}
	}

	if !foundWorkflow {
		t.Error("No workflow files generated")
	}
}

func TestGenerate_Table(t *testing.T) {
	tests := []struct {
		name         string
		cfg          Config
		expectedPath string
		expectExist  bool
	}{
		{
			name: "Go Workflow",
			cfg: Config{
				ProjectName:  "go-app",
				WorkflowType: "go",
			},
			expectedPath: ".github/workflows/go.yaml",
			expectExist:  true,
		},
		{
			name: "Docker Enabled",
			cfg: Config{
				ProjectName: "docker-app",
				UseDocker:   true,
			},
			expectedPath: "Dockerfile",
			expectExist:  true,
		},
		{
			name: "Docker Disabled",
			cfg: Config{
				ProjectName: "no-docker-app",
				UseDocker:   false,
			},
			expectedPath: "Dockerfile",
			expectExist:  false,
		},
		{
			name: "TypeScript Workflow",
			cfg: Config{
				ProjectName:  "ts-app",
				WorkflowType: "typescript",
			},
			expectedPath: ".github/workflows/typescript.yaml",
			expectExist:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := Generate(tt.cfg)
			if err != nil {
				t.Fatalf("Generate failed: %v", err)
			}

			found := false
			for _, f := range files {
				if f.Path == tt.expectedPath {
					found = true
					break
				}
			}

			if found != tt.expectExist {
				t.Errorf("Path %s existence check: got %v, want %v", tt.expectedPath, found, tt.expectExist)
			}
		})
	}
}

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{"Valid", Config{ProjectName: "valid-name", WorkflowType: "go"}, false},
		{"Empty Name", Config{ProjectName: "", WorkflowType: "go"}, true},
		{"Invalid Name", Config{ProjectName: "Invalid Name!", WorkflowType: "go"}, true},
		{"Unsupported Type", Config{ProjectName: "valid", WorkflowType: "ruby"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateConfig(tt.cfg); (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerate_NoSideEffects(t *testing.T) {
	// Since we are not using any mocking for FS, we just verify that no files are created in the current dir
	// In a real environment, we'd use a read-only filesystem check
	cfg := Config{
		ProjectName: "no-io",
	}

	_, err := Generate(cfg)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Simple check: .github directory should not be created by Generate
	_, err = os.Stat(".github")
	if !os.IsNotExist(err) {
		t.Error("Generate created .github directory (side effect detected)")
	}
}
