package scaffold

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/modelcontextprotocol/platform.mcp/internal/templates"
)

// Generate returns a slice of File structs with the generated content.
func Generate(cfg Config) ([]File, error) {
	if err := ValidateConfig(cfg); err != nil {
		return nil, err
	}

	var files []File

	// Workflow generation
	wtype := cfg.WorkflowType
	if wtype == "" {
		wtype = "go"
	}

	tmplName := fmt.Sprintf("%s.yaml.tmpl", wtype)
	tmplContent, err := templates.FS.ReadFile(tmplName)
	if err == nil {
		tmpl, err := template.New(wtype).Parse(string(tmplContent))
		if err != nil {
			return nil, fmt.Errorf("failed to parse template: %w", err)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, cfg); err != nil {
			return nil, fmt.Errorf("failed to execute template: %w", err)
		}

		files = append(files, File{
			Path:    fmt.Sprintf(".github/workflows/%s.yaml", wtype),
			Content: buf.String(),
			Mode:    0644,
		})
	}

	// Dockerfile generation
	if cfg.UseDocker {
		files = append(files, File{
			Path:    "Dockerfile",
			Content: fmt.Sprintf("FROM alpine:latest\nLABEL project=%s\nCMD [\"echo\", \"Hello from %s\"]", cfg.ProjectName, cfg.ProjectName),
			Mode:    0644,
		})
	}

	return files, nil
}
