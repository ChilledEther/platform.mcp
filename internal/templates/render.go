package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

// Render applies the data to the template string
func Render(tmplStr string, data any) (string, error) {
	tmpl, err := template.New("base").Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
