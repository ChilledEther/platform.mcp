package scaffold

// FluxGenerator generates FluxCD manifests
type FluxGenerator struct{}

// Ensure FluxGenerator implements Generator
var _ Generator = (*FluxGenerator)(nil)

func (g *FluxGenerator) Generate(cfg Config) ([]File, error) {
	// TODO: Implement
	return nil, nil
}
