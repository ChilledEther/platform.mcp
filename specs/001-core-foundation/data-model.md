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
- `WorkflowType`: enum ["standard", "docker", "minimal"] (Default "standard")

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
- `WorkflowType` must be one of the supported types (standard, docker, minimal).
