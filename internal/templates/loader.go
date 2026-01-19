package templates

import (
	"embed"
	"fmt"
)

//go:embed *.tmpl
var fs embed.FS

// Load reads a template file from the embedded filesystem
func Load(name string) (string, error) {
	content, err := fs.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("failed to load template %s: %w", name, err)
	}
	return string(content), nil
}
