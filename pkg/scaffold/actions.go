package scaffold

import (
	"embed"
	"fmt"

	"github.com/modelcontextprotocol/platform.mcp/internal/templates"
)

// ActionsGenerator generates GitHub Actions workflows
type ActionsGenerator struct {
	// FS is the file system containing templates
	// In production this would be the embed.FS
	FS *embed.FS
}

// Ensure ActionsGenerator implements Generator
var _ Generator = (*ActionsGenerator)(nil)

func (g *ActionsGenerator) Generate(cfg Config) ([]File, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	tmpl, err := templates.FindTemplate("actions-workflow")
	if err != nil {
		return nil, fmt.Errorf("failed to find template: %w", err)
	}

	tmplContent, err := templates.Load(tmpl.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to load template: %w", err)
	}

	content, err := templates.Render(tmplContent, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	return []File{
		{
			Path:    tmpl.Target,
			Content: content,
		},
	}, nil
}
