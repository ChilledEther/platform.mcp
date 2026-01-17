# Implementation Plan: Platform CLI Tool

**Branch**: `002-platform-cli` | **Date**: 2026-01-17 | **Spec**: `/specs/002-platform-cli/spec.md`
**Input**: Feature specification from `/specs/002-platform-cli/spec.md`

## Summary

Implement a CLI tool named `platform` that generates GitHub Actions workflow YAML files and Dockerfiles to disk. The CLI serves as a consumer of the shared core library (`pkg/scaffold`), providing a user-friendly interface for project scaffolding with features like dry-run, force overwrite, and customization via flags.

## Technical Context

**Language/Version**: Go 1.25+
**Primary Dependencies**: `github.com/spf13/cobra`
**Storage**: Local filesystem (YAML files, Dockerfile)
**Testing**: `go test` (unit/integration), table-driven tests
**Target Platform**: Linux, macOS, Windows
**Project Type**: CLI Application (`cmd/platform`)
**Performance Goals**: < 100ms for file generation
**Constraints**: Zero data loss (prompt before overwrite), must use `pkg/scaffold` for logic

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

1. **I. Core-First Architecture**: CLI uses `pkg/scaffold` for all generation logic. (PASS)
2. **II. TDD**: Tests will be written for CLI command logic and I/O handlers. (PASS)
3. **III. Modular**: CLI commands are implemented as Cobra commands. (PASS)
4. **IV. Separate Artifacts**: CLI is a separate binary from MCP server. (PASS)
5. **V. Docker-First**: Build verified in `build/package/platform/Dockerfile`. (PASS)
6. **VI. MCP Naming**: N/A for CLI (applies to MCP tools). (PASS)

## Project Structure

### Documentation (this feature)

```text
specs/002-platform-cli/
├── plan.md              # This file
├── research.md          # Technical approach and alternatives
├── data-model.md        # CLI configuration and operation models
├── quickstart.md        # Usage examples and installation
├── contracts/           # CLI command and flag definitions
└── tasks.md             # Implementation tasks (Phase 2)
```

### Source Code (repository root)

```text
cmd/
└── platform/
    └── main.go         # Entry point

internal/
└── cli/
    ├── cmd/            # Cobra command definitions
    ├── io/             # File I/O and prompt handlers
    └── config/         # CLI configuration parsing
```

**Structure Decision**: Standard Go CLI layout using Cobra. Entry point in `cmd/platform`, logic in `internal/cli` to keep `main.go` minimal.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| None | | |
