# Platform MCP

## What This Is

A Go-based platform tooling solution that generates standardized configuration files (GitHub Actions workflows, Docker builds, FluxCD configs) for platform engineers. Provides both a CLI for direct file generation and an MCP server for AI agent integration.

## Core Value

The core library must generate valid, standards-compliant YAML without side effects — everything else builds on this foundation.

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] Core library (`pkg/scaffold`) generates file content without I/O
- [ ] File struct with Path, Content, and Mode fields
- [ ] Config struct for customizing generation behavior
- [ ] Embedded templates via Go `embed` package
- [ ] GitHub Actions workflow YAML generation
- [ ] Docker build configuration generation
- [ ] FluxCD configuration generation
- [ ] CLI (`platform-cli`) writes generated files to disk
- [ ] CLI supports `--dry-run`, `--force`, `--output`, `--project-name` flags
- [ ] MCP server (`platform-mcp`) returns content to AI agents via stdio
- [ ] MCP tool discovery with parameter schemas
- [ ] Multi-stage Alpine Docker builds (<50MB images)
- [ ] GitHub Actions CI/CD pipeline
- [ ] Release Please automated versioning
- [ ] Container registry publishing

### Out of Scope

- Multi-architecture builds (amd64/arm64) — future enhancement after v1
- HTTP transport for MCP — stdio is primary, HTTP can be added later
- Additional template types beyond GitHub Actions, Docker, FluxCD — expand after core is solid

## Context

**Architecture**: Shared Core (`pkg/`) + Consumer Applications (`cmd/`)

- `pkg/`: Pure logic, returns data, no side effects
- `cmd/platform-cli/`: CLI entry point using Cobra
- `cmd/platform-mcp/`: MCP server entry point using go-sdk
- `internal/templates/`: Embedded template files

**Tech Stack** (locked):

- Go 1.25+
- `github.com/spf13/cobra` for CLI
- `github.com/modelcontextprotocol/go-sdk` for MCP
- Go `embed` package for templates

**Existing State**:

- Specs completed via SpecKit (5 feature specs in `specs/`)
- Empty `cmd/` directory structure
- README.md contains stale content (references non-existent TypeScript implementation) — needs cleanup

## Constraints

- **Tech Stack**: Go 1.25+, Cobra, go-sdk, embed — per specs
- **TDD**: Red-Green-Refactor workflow required per AGENTS.md
- **Scripting**: PowerShell for all automation scripts per global standards
- **Image Size**: Final Docker images must be under 50MB each

## Key Decisions

| Decision                       | Rationale                                                                      | Outcome   |
| ------------------------------ | ------------------------------------------------------------------------------ | --------- |
| Shared Core + Consumer pattern | Enables both CLI and MCP to use identical generation logic without duplication | — Pending |
| Embedded templates via `embed` | Single binary distribution, no external template files needed                  | — Pending |
| Stdio transport first for MCP  | Primary use case is local MCP clients (Claude Desktop, Cursor)                 | — Pending |

---

_Last updated: 2026-01-19 after initialization_
