package scaffold

import (
	"embed"
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
	// Use the manifest-driven generator but ensure only Actions are generated
	limitedCfg := cfg
	limitedCfg.WithActions = true
	limitedCfg.WithDocker = false
	limitedCfg.WithFlux = false

	return Generate(limitedCfg)
}
