# Implementation Plan: Platform MCP Server

**Branch**: `003-platform-mcp` | **Date**: 2026-01-17 | **Spec**: `/specs/003-platform-mcp/spec.md`
**Input**: Feature specification from `/specs/003-platform-mcp/spec.md`

## Summary

Implement the `platform-mcp` server that exposes the `generate_workflows` tool to AI agents. The server will use the `github.com/modelcontextprotocol/go-sdk` and integrate with the shared generation logic in `pkg/scaffold`. It will support the `stdio` transport for seamless integration with local agent clients.

## Technical Context

**Language/Version**: Go 1.25+
**Primary Dependencies**: `github.com/modelcontextprotocol/go-sdk`, `pkg/scaffold`
**Storage**: N/A (Stateless server)
**Testing**: `go test` with MCP client simulator
**Target Platform**: Linux, Docker (Alpine-based)
**Project Type**: MCP Server (`cmd/platform-mcp`)
**Performance Goals**: < 2s for tool calls
**Constraints**: Must NOT perform disk I/O, must follow `verb_noun` naming
**Scale/Scope**: Initial MCP server with one tool

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

1.  **I. Core-First**: Logic resides in `pkg/scaffold`. (PASS)
2.  **II. TDD**: Tests will be written for tool registration and execution. (PASS)
3.  **III. Modular**: New tools can be registered without changing server core. (PASS)
4.  **IV. Separate Artifacts**: `platform-mcp` is a distinct binary in `cmd/`. (PASS)
5.  **V. Docker-First**: Dockerfile to be created in `build/package/platform-mcp/`. (PASS)
6.  **VI. MCP Naming**: Tool named `generate_workflows` (verb_noun). (PASS)

## Project Structure

### Documentation (this feature)

```text
specs/003-platform-mcp/
├── plan.md              # This file
├── research.md          # MCP protocol and integration research
├── data-model.md        # Tool schema and response format
├── quickstart.md        # Usage instructions
├── contracts/           # Tool JSON schemas
│   └── tool-generate-workflows.json
└── tasks.md             # Implementation tasks
```

### Source Code (repository root)

```text
cmd/
└── platform-mcp/
    └── main.go          # MCP server entry point

internal/
└── mcp/
    ├── server.go        # MCP server initialization
    └── tools.go         # Tool registration logic

pkg/
└── scaffold/            # Shared generation logic (from 001-core-foundation)

build/package/
└── platform-mcp/
    └── Dockerfile       # MCP server Dockerfile
```

**Structure Decision**: Standard Go project layout with separate `cmd` for the binary and `internal/mcp` for protocol-specific logic.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| None      | -          | -                                   |
