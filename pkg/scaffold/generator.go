package scaffold

// Generator is the interface for feature modules to implement.
type Generator interface {
	Generate(cfg Config) ([]File, error)
}
