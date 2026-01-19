package scaffold

import "fmt"

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
	if c.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	return nil
}
