package scaffold

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
}
