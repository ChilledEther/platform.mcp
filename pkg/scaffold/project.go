package scaffold

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
	// TODO: Implement wiring
	return nil, nil
}
