package scaffold

import (
	"fmt"

	"github.com/modelcontextprotocol/platform.mcp/internal/templates"
)

// DockerGenerator generates Docker-related files
type DockerGenerator struct{}

// Ensure DockerGenerator implements Generator
var _ Generator = (*DockerGenerator)(nil)

func (g *DockerGenerator) Generate(cfg Config) ([]File, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	files := []File{}

	// Dockerfile
	dockerTmpl, err := templates.Load("Dockerfile.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to load Dockerfile template: %w", err)
	}
	dockerContent, err := templates.Render(dockerTmpl, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render Dockerfile: %w", err)
	}
	files = append(files, File{
		Path:    "Dockerfile",
		Content: dockerContent,
	})

	// docker-build.yaml
	buildTmpl, err := templates.Load("docker-build.yaml.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to load docker-build template: %w", err)
	}
	buildContent, err := templates.Render(buildTmpl, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render docker-build: %w", err)
	}
	files = append(files, File{
		Path:    "docker-build.yaml",
		Content: buildContent,
	})

	return files, nil
}
