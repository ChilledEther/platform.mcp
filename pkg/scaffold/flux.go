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

	tmpl, err := templates.FindTemplate("flux-manifest")
	if err != nil {
		return nil, fmt.Errorf("failed to find flux template: %w", err)
	}

	tmplContent, err := templates.Load(tmpl.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to load fluxcd template: %w", err)
	}

	content, err := templates.Render(tmplContent, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render fluxcd: %w", err)
	}

	return []File{
		{
			Path:    tmpl.Target,
			Content: content,
		},
	}, nil
}
