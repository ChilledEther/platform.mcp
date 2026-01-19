# Antigravity Agent Instructions: Platform MCP

You are Antigravity, an agentic coding assistant. This file defines the operational standards and development workflows for the **Platform MCP** project.

## Core Mandates

- **Architecture**: Core-First. Shared logic in `pkg/`, consumers in `cmd/`.
  - `pkg/`: Pure logic, return data, no side effects (I/O).
  - `cmd/`: Application entry points (CLI, MCP server).
- **TDD**: Red-Green-Refactor is non-negotiable. Write failing tests before implementation.
- **Dependencies**: Minimal footprint. Use Go standard library over third-party packages unless they provide significant value (e.g., `cobra`, `mcp-sdk`).
- **Standards**: All generated files must be valid, standards-compliant YAML.

## Development Workflow

### Build & Run
- **Build Docker**: `./scripts/Invoke-DockerBuild.ps1`
- **Run CLI**: `go run cmd/platform/main.go`
- **Run MCP**: `go run cmd/platform-mcp/main.go`

### Testing
- **Run All Tests**: `go test ./...`
- **Package Tests**: `go test ./pkg/scaffold/...`
- **Docker Tests**: `./scripts/Test-Docker.ps1`

### Linting
- **Format**: `go fmt ./...`
- **Lint**: `golangci-lint run`

## Code Style & Guidelines

### MCP Tool Naming
- **Convention**: `snake_case` using `verb_noun`.
- **Verbs**: `get`, `create`, `update`, `delete`, `list`, `search`, `analyze`, `generate`, `validate`.

### Go Implementation
- **Layout**: Follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout).
- **Error Handling**: Return errors as the last value. Wrap with context: `fmt.Errorf("failed to do X: %w", err)`.
- **Templates**: Use Go `embed` package to bundle assets into the binary.

## Repository Structure

- `cmd/`: Entry points for `platform` (CLI) and `platform-mcp` (MCP).
- `pkg/`: Shared core logic (scaffolding, generators).
- `internal/`: Private code (templates, registry helpers).
- `.planning/`: Execution plans, roadmap, and state tracking.
- `scripts/`: PowerShell automation for local builds/testing.
- `build/package/`: Dockerfiles for production artifacts.

## AI Alignment

- **Context**: Read `.planning/` before starting any phase.
- **Plan**: Follow `.planning/phases/` sequentially.
- **Validation**: Ensure all tests pass in Docker before declaring completion.

**Version**: 3.0.0 | **Updated**: 2026-01-19
