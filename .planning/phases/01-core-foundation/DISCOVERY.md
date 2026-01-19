# Phase 1 Context: Core Foundation

## Goal
Implement the pure Go core library (`pkg/scaffold`) that handles all file generation logic. This library must be stateless, side-effect free (no I/O), and distributable. It serves as the shared logic for both the CLI and MCP server.

## Existing State
- **Repo:** `.git` exists.
- **Code:** No existing code in `pkg/` or `internal/` found (clean slate).
- **Specs:** `.planning/PROJECT.md` defines requirements.
- **Constraints:**
  - Go 1.25+
  - `embed` package for templates
  - TDD (Red-Green-Refactor)

## Strategy
Since no legacy code exists in the target directories, we proceed with a **clean greenfield implementation**.

1. **Structure:** Create `pkg/scaffold` for public API and `internal/templates` for embedded assets.
2. **Data Model:** Define `Config` (inputs) and `File` (outputs) structs.
3. **Logic:** Implement generation functions using Go templates.
4. **Testing:** High coverage unit tests validating generation without writing to disk.

## Requirements Mapping
- **FR-001 (Pure Gen):** `Generate(cfg Config) ([]File, error)`
- **FR-002 (File Struct):** `type File struct { Path, Content, Mode }`
- **FR-005 (Embed):** `//go:embed *.tmpl` in `internal/templates`
- **FR-006 (Actions):** Generate `.github/workflows/*.yaml`

## Architecture
```go
// pkg/scaffold/types.go
type Config struct {
    ProjectName string
    Features    []string // "docker", "flux", etc.
}

type File struct {
    Path    string
    Content string
    Mode    os.FileMode
}

// pkg/scaffold/generator.go
func Generate(cfg Config) ([]File, error)
```
