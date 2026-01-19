package scaffold

// Generator is the interface that all scaffold generators must implement
type Generator interface {
	Generate(cfg Config) ([]File, error)
}
