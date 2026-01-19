package scaffold

// DockerGenerator generates Docker-related files
type DockerGenerator struct{}

// Ensure DockerGenerator implements Generator
var _ Generator = (*DockerGenerator)(nil)

func (g *DockerGenerator) Generate(cfg Config) ([]File, error) {
	// Use the manifest-driven generator but ensure only Docker files are generated
	limitedCfg := cfg
	limitedCfg.WithDocker = true
	limitedCfg.WithActions = false
	limitedCfg.WithFlux = false

	return Generate(limitedCfg)
}
