# Antigravity Agent Instructions: Platform MCP

You are Antigravity, an agentic coding assistant. This file defines the operational standards and development workflows for the **Platform MCP** project.

## Core Mandates

- **Architecture**: Shared Core (`pkg/`) + Consumer Applications (`cmd/`).
  - `pkg/`: Pure logic, return data, no side effects (I/O).
  - `cmd/`: Entry points for applications (CLI, MCP, etc.).
- **Standards**: All development MUST follow the **Constitution** at `.specify/memory/constitution.md`.
- **TDD**: Red-Green-Refactor is non-negotiable. Write failing tests before implementation.

## Development Workflow

### Build & Run
- **Build Docker**: `./scripts/Invoke-DockerBuild.ps1` (PowerShell)
- **Add to Catalog**: `./scripts/Add-MCPToCatalog.ps1` (PowerShell)
- **Run CLI**: `go run cmd/platform/main.go`
- **Run MCP**: `go run cmd/platform-mcp/main.go`

### Testing
- **Run All Tests**: `go test ./...`
- **Single Package**: `go test ./pkg/scaffold/...`
- **Single Test**: `go test -v -run TestSpecificName ./pkg/scaffold`
- **Docker Tests**: `./scripts/Test-Docker.ps1`

### Linting & Formatting
- **Format**: `go fmt ./...`
- **Lint**: `golangci-lint run` (if available)

## Code Style & Guidelines

### MCP Tool Naming (CONSTITUTION VI)
- **Convention**: `snake_case` using `verb_noun`.
- **Verbs**: `get`, `create`, `update`, `delete`, `list`, `search`, `analyze`, `generate`, `validate`.
- **Examples**: `get_firewall_rules`, `list_resources`.

### Go Implementation
- **Project Layout**: Follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout).
- **Naming**: 
  - Packages: short, lowercase, single word.
  - Exported members: PascalCase.
  - Local variables: camelCase.
- **Error Handling**: 
  - Return errors as the last value.
  - Wrap errors with context: `fmt.Errorf("failed to generate scaffold: %w", err)`.
  - Use `errors.Is` and `errors.As` for checking.
- **Imports**: Standard library first, then third-party, then internal. Grouped by blank lines.

### TypeScript/Bun (Historical/Future)
- **Runtime**: Always use **Bun**.
- **Naming**: `kebab-case` for files, `PascalCase` for classes/types, `camelCase` for functions/variables.
- **Validation**: Use **Zod** for all I/O and MCP tool arguments.

## Repository Structure

- `/cmd`: Entry points for CLI and MCP binaries.
- `/pkg`: Shared core logic (importable library).
- `/internal`: Private code (CLI helpers, MCP registry, templates).
- `/specs`: Markdown specifications (Source of Truth).
- `/scripts`: PowerShell (`.ps1`) automation scripts.
- `/build/package`: Multi-stage Alpine-based Dockerfiles.

## AI Alignment

- **Context**: Read `specs/` before implementation.
- **Plan**: Use `/speckit.plan` (via Task tool) to generate design before writing code.
- **Validation**: Ensure all tests pass in Docker.

**Version**: 2.2.0 | **Updated**: 2026-01-17

## Active Technologies
- Go 1.25+ + `github.com/spf13/cobra` (CLI), `github.com/modelcontextprotocol/go-sdk` (MCP), Go `embed` package (001-core-foundation)
- N/A (Pure functions) (001-core-foundation)

## Recent Changes
- 001-core-foundation: Added Go 1.25+ + `github.com/spf13/cobra` (CLI), `github.com/modelcontextprotocol/go-sdk` (MCP), Go `embed` package
