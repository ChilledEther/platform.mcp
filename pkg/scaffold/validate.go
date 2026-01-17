package scaffold

import (
	"errors"
	"regexp"
)

var projectNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)

// ValidateConfig checks if the configuration is valid.
func ValidateConfig(cfg Config) error {
	if cfg.ProjectName == "" {
		return errors.New("project name cannot be empty")
	}

	if !projectNameRegex.MatchString(cfg.ProjectName) {
		return errors.New("project name must be alphanumeric (hyphens allowed)")
	}

	switch cfg.WorkflowType {
	case "go", "typescript", "python", "":
		// Valid
	default:
		return errors.New("unsupported workflow type")
	}

	return nil
}
