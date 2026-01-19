package scaffold

import (
	"testing"
)

func TestProjectGenerator_Generate(t *testing.T) {
	// Setup
	g := NewProjectGenerator()

	tests := []struct {
		name            string
		cfg             Config
		expectedFiles   []string
		unexpectedFiles []string
	}{
		{
			name: "All Enabled",
			cfg: Config{
				ProjectName: "full-project",
				WithActions: true,
				WithDocker:  true,
				WithFlux:    true,
			},
			expectedFiles: []string{
				".github/workflows/ci.yaml",
				"Dockerfile",
				"docker-build.yaml",
				"fluxcd.yaml",
			},
		},
		{
			name: "Only Actions",
			cfg: Config{
				ProjectName: "actions-project",
				WithActions: true,
				WithDocker:  false,
				WithFlux:    false,
			},
			expectedFiles: []string{
				".github/workflows/ci.yaml",
			},
			unexpectedFiles: []string{
				"Dockerfile",
				"fluxcd.yaml",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := g.Generate(tt.cfg)
			if err != nil {
				t.Fatalf("Generate failed: %v", err)
			}

			// Check expected files
			for _, exp := range tt.expectedFiles {
				found := false
				for _, f := range files {
					if f.Path == exp {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected file %s not found", exp)
				}
			}

			// Check unexpected files
			for _, unexp := range tt.unexpectedFiles {
				for _, f := range files {
					if f.Path == unexp {
						t.Errorf("Unexpected file %s found", unexp)
					}
				}
			}
		})
	}
}
