package scaffold

import (
	"fmt"

	"github.com/modelcontextprotocol/platform.mcp/internal/templates"
)

// FluxGenerator generates FluxCD manifests
type FluxGenerator struct{}

// Ensure FluxGenerator implements Generator
var _ Generator = (*FluxGenerator)(nil)

func (g *FluxGenerator) Generate(cfg Config) ([]File, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	tmplContent, err := templates.Load("fluxcd.yaml.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to load fluxcd template: %w", err)
	}

	content, err := templates.Render(tmplContent, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render fluxcd: %w", err)
	}

	return []File{
		{
			Path:    "fluxcd.yaml",
			Content: content,
		},
	}, nil
}
