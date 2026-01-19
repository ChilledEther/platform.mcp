package scaffold

import (
	"fmt"

	"github.com/modelcontextprotocol/platform.mcp/internal/templates"
)

// Generate returns a slice of File structs with the generated content.
func Generate(cfg Config) ([]File, error) {
	if err := ValidateConfig(cfg); err != nil {
		return nil, err
	}

	manifest, err := templates.GetManifest()
	if err != nil {
		return nil, fmt.Errorf("failed to get template manifest: %w", err)
	}

	mappings := FilterTemplates(manifest, cfg)

	var files []File
	for _, mapping := range mappings {
		tmplContent, err := templates.Load(mapping.Source)
		if err != nil {
			return nil, fmt.Errorf("failed to load template %s: %w", mapping.Source, err)
		}

		rendered, err := templates.Render(tmplContent, cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to render template %s: %w", mapping.Name, err)
		}

		files = append(files, File{
			Path:    mapping.Target,
			Content: rendered,
			Mode:    0644,
		})
	}

	return files, nil
}
