# Data Model: Core Foundation

## Entities

### File
Represents a generated file.
- `Path`: string (e.g., ".github/workflows/ci.yaml")
- `Content`: string (The rendered YAML content)
- `Mode`: uint32 (File permissions, default 0644)

### Config
User-provided configuration for generation.
- `ProjectName`: string (Required, validated for length/characters)
- `UseDocker`: bool (Whether to include Docker-related workflows)
<<<<<<< HEAD
- `WorkflowType`: enum ["standard", "docker", "minimal"] (Default "standard")
=======
- `WorkflowType`: enum ["go", "typescript", "python"] (Default "go")
>>>>>>> 001-core-foundation

## Interfaces

### Generator
```go
type Generator interface {
    Generate(cfg Config) ([]File, error)
}
```

## Validation Rules
- `ProjectName` cannot be empty.
- `ProjectName` must be alphanumeric (hyphens allowed).
<<<<<<< HEAD
- `WorkflowType` must be one of the supported types (standard, docker, minimal).
=======
- `WorkflowType` must be one of the supported types.
>>>>>>> 001-core-foundation
