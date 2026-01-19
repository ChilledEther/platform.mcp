package scaffold

// DockerGenerator generates Docker-related files
type DockerGenerator struct{}

// Ensure DockerGenerator implements Generator
var _ Generator = (*DockerGenerator)(nil)

func (g *DockerGenerator) Generate(cfg Config) ([]File, error) {
	// TODO: Implement
	return nil, nil
}
