package scaffold

import "fmt"

// ProjectGenerator orchestrates the generation of the entire project
type ProjectGenerator struct {
	Actions Generator
	Docker  Generator
	Flux    Generator
}

// NewProjectGenerator creates a new ProjectGenerator with default sub-generators
func NewProjectGenerator() *ProjectGenerator {
	return &ProjectGenerator{
		Actions: &ActionsGenerator{},
		Docker:  &DockerGenerator{},
		Flux:    &FluxGenerator{},
	}
}

// Ensure ProjectGenerator implements Generator
var _ Generator = (*ProjectGenerator)(nil)

func (g *ProjectGenerator) Generate(cfg Config) ([]File, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	var allFiles []File

	if cfg.WithActions {
		files, err := g.Actions.Generate(cfg)
		if err != nil {
			return nil, fmt.Errorf("actions generation failed: %w", err)
		}
		allFiles = append(allFiles, files...)
	}

	if cfg.WithDocker {
		files, err := g.Docker.Generate(cfg)
		if err != nil {
			return nil, fmt.Errorf("docker generation failed: %w", err)
		}
		allFiles = append(allFiles, files...)
	}

	if cfg.WithFlux {
		files, err := g.Flux.Generate(cfg)
		if err != nil {
			return nil, fmt.Errorf("flux generation failed: %w", err)
		}
		allFiles = append(allFiles, files...)
	}

	return allFiles, nil
}
