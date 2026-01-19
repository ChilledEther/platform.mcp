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
	dockerTmplInfo, err := templates.FindTemplate("dockerfile")
	if err != nil {
		return nil, fmt.Errorf("failed to find Dockerfile template: %w", err)
	}
	dockerTmpl, err := templates.Load(dockerTmplInfo.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to load Dockerfile template: %w", err)
	}
	dockerContent, err := templates.Render(dockerTmpl, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render Dockerfile: %w", err)
	}
	files = append(files, File{
		Path:    dockerTmplInfo.Target,
		Content: dockerContent,
	})

	// docker-build.yaml
	buildTmplInfo, err := templates.FindTemplate("docker-build")
	if err != nil {
		return nil, fmt.Errorf("failed to find docker-build template: %w", err)
	}
	buildTmpl, err := templates.Load(buildTmplInfo.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to load docker-build template: %w", err)
	}
	buildContent, err := templates.Render(buildTmpl, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to render docker-build: %w", err)
	}
	files = append(files, File{
		Path:    buildTmplInfo.Target,
		Content: buildContent,
	})

	return files, nil
}
