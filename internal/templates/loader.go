package templates

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

//go:embed *.tmpl manifest.yaml
var FS embed.FS

// BaseDir is the directory to check for external templates.
var BaseDir string

// SetBaseDir sets the directory to check for external templates.
func SetBaseDir(path string) {
	BaseDir = path
}

// Load reads a template file from the external directory (if set) or the embedded filesystem.
func Load(name string) (string, error) {
	if BaseDir != "" {
		externalPath := filepath.Join(BaseDir, name)
		content, err := os.ReadFile(externalPath)
		if err == nil {
			return string(content), nil
		}
	}

	content, err := FS.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("failed to load template %s: %w", name, err)
	}
	return string(content), nil
}

// Manifest represents the template manifest structure.
type Manifest struct {
	Templates []TemplateMapping `yaml:"templates"`
}

// TemplateMapping defines a single template mapping.
type TemplateMapping struct {
	Name      string `yaml:"name" json:"name"`
	Source    string `yaml:"source" json:"source"`
	Target    string `yaml:"target" json:"target"`
	Condition string `yaml:"condition" json:"condition"`
}

// GetManifest parses the template manifest from the external directory or embedded filesystem.
func GetManifest() (*Manifest, error) {
	var content []byte
	var err error

	if BaseDir != "" {
		content, err = os.ReadFile(filepath.Join(BaseDir, "manifest.yaml"))
	}

	if content == nil {
		content, err = FS.ReadFile("manifest.yaml")
		if err != nil {
			return nil, fmt.Errorf("failed to read manifest.yaml: %w", err)
		}
	}

	var m Manifest
	if err := yaml.Unmarshal(content, &m); err != nil {
		return nil, fmt.Errorf("failed to unmarshal manifest: %w", err)
	}
	return &m, nil
}

// FindTemplate looks up a template mapping by name from the manifest.
func FindTemplate(name string) (*TemplateMapping, error) {
	m, err := GetManifest()
	if err != nil {
		return nil, err
	}

	for _, t := range m.Templates {
		if t.Name == name {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("template not found: %s", name)
}
