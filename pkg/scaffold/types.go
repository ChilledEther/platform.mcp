package scaffold

import "fmt"

// File represents a generated file.
type File struct {
	Path    string
	Content string
	Mode    uint32
}

// Config represents the generation options.
type Config struct {
	ProjectName  string
	UseDocker    bool
	WorkflowType string // "go", "typescript", "python"
	WithActions  bool
	WithDocker   bool
	WithFlux     bool
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	return nil
}
