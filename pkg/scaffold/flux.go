package scaffold

// FluxGenerator generates FluxCD manifests
type FluxGenerator struct{}

// Ensure FluxGenerator implements Generator
var _ Generator = (*FluxGenerator)(nil)

func (g *FluxGenerator) Generate(cfg Config) ([]File, error) {
	// Use the manifest-driven generator but ensure only Flux manifests are generated
	limitedCfg := cfg
	limitedCfg.WithFlux = true
	limitedCfg.WithActions = false
	limitedCfg.WithDocker = false

	return Generate(limitedCfg)
}
