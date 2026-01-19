package scaffold

import (
	"github.com/modelcontextprotocol/platform.mcp/internal/templates"
)

// FilterTemplates returns the template mappings that should be generated based on the config.
func FilterTemplates(manifest *templates.Manifest, cfg Config) []templates.TemplateMapping {
	var filtered []templates.TemplateMapping
	for _, tmpl := range manifest.Templates {
		if ShouldGenerate(tmpl.Condition, cfg) {
			filtered = append(filtered, tmpl)
		}
	}
	return filtered
}

// ShouldGenerate evaluates a condition string against the config.
func ShouldGenerate(condition string, cfg Config) bool {
	switch condition {
	case "workflow_go":
		return cfg.WithActions && (cfg.WorkflowType == "go" || cfg.WorkflowType == "")
	case "workflow_typescript":
		return cfg.WithActions && (cfg.WorkflowType == "typescript" || cfg.WorkflowType == "node")
	case "workflow_python":
		return cfg.WithActions && cfg.WorkflowType == "python"
	case "use_docker":
		return cfg.UseDocker
	case "with_docker":
		return cfg.WithDocker || cfg.UseDocker
	case "with_flux":
		return cfg.WithFlux
	case "with_actions":
		return cfg.WithActions
	default:
		return false
	}
}
