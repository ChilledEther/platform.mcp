package scaffold

// File represents a file to be generated
type File struct {
	Path    string
	Content string
}

// Config holds the configuration for the scaffold generation
type Config struct {
	ProjectName string
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	return nil
}
