# Implementation Plan: Infrastructure MCP Foundation

**Branch**: `001-infra-mcp-foundation` | **Date**: 2026-01-15 | **Updated**: 2026-01-16 | **Spec**: [specs/001-infra-mcp-foundation/spec.md](../spec.md)

## Summary

This feature implements the core Model Context Protocol (MCP) server in TypeScript/Bun, serving as the foundation for future infrastructure provisioning tools. It establishes the server lifecycle, supporting both **Standard I/O (stdio)** for local usage and **Streamable HTTP** for remote/containerized environments. It includes a type-safe, extensible architecture for tool registration using the official `@modelcontextprotocol/sdk` and a built-in `health_check` tool.

## Technical Context

**Languages/Versions**: 
- TypeScript 5.x+ (Bun runtime)
- Go 1.25+

**Primary Dependencies**: 
- TypeScript: `@modelcontextprotocol/sdk`, `zod`
- Go: `github.com/modelcontextprotocol/go-sdk`, `gopkg.in/yaml.v3`

**Storage**: N/A (Stateless server)
**Testing**: `bun:test` (TypeScript), `go test` (Go)
**Target Platform**: Linux/macOS/Windows (Container via Bun or Go binary), Docker/K8s
**Project Type**: Dual-language CLI/Server binary
**Performance Goals**: <100ms startup, low memory footprint (<128MB)
**Constraints**: Stdio & Streamable HTTP transports, no external database, strict schema validation, 128MB Memory Limit
**Scale/Scope**: Foundation for 10-50 future tools

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] **I. Dual-Language Implementation**: Both TypeScript/Bun and Go implementations maintained in parallel.
- [x] **Ia. TypeScript-Specific**: Uses TypeScript 5.x+, Bun runtime, `@modelcontextprotocol/sdk`.
- [x] **Ib. Go-Specific**: Uses Go 1.25+, `github.com/modelcontextprotocol/go-sdk`.
- [x] **II. Multi-Mode MCP Transport**: Implements both Stdio and Streamable HTTP using the official SDKs.
- [x] **III. Strict Type & Schema Safety**: Uses Zod schemas (TS) and struct validation (Go).
- [x] **IV. Concurrency & Timeout Management**: Async/await patterns (TS), context.Context (Go).
- [x] **V. Default Read-Only**: `health_check` is read-only. Future write tools will need safeguards.
- [x] **VI. Secret Management**: No secrets required for foundation (env vars supported if needed).
- [x] **VII. Input Sanitization**: Zod validation (TS), struct validation (Go) applied to all inputs.
- [x] **VIII. Specification Kit Alignment**: Follows `.specify` structure.
- [x] **IX. Scripting & Automation**: Uses Bun (TS) and Go for execution.
- [x] **XII. Cloud-Native & Container Ready**: Separate Dockerfiles maintained, logging to stderr.
- [x] **XIII. Git Worktree Workflow**: Enables parallel agent work on both implementations.
- [x] **XIV. Minimal Resource Footprint**: Uses `bun build --compile` (TS) and static Go binary + Alpine to minimize image size (<160MB).
- [x] **XV. Optimized CI/CD Workflow**: Implements path-based triggers to ensure builds only run when relevant source or configuration files change.

## Project Structure

### Documentation (this feature)

```text
specs/001-infra-mcp-foundation/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
└── tasks.md             # Phase 2 output
```

### Source Code (implementations/)

```text
implementations/
├── typescript/                    # TypeScript/Bun implementation
│   ├── Dockerfile                 # Container image (Bun base)
│   ├── index.ts                   # Entry point
│   ├── package.json
│   ├── tsconfig.json
│   ├── server/
│   │   ├── index.ts               # Server lifecycle & transport selection
│   │   └── config.ts              # Configuration parsing
│   ├── tools/
│   │   ├── registry.ts            # Tool registration logic
│   │   └── health.ts              # Health check implementation
│   ├── resources/
│   │   ├── registry.ts            # Resource registration logic
│   │   └── types.ts               # Resource interfaces
│   ├── utils/
│   │   └── validation.ts          # Schema validation helpers
│   └── tests/                     # TypeScript tests
│
└── go/                            # Go implementation
    ├── Dockerfile                 # Container image (Go + Alpine)
    ├── go.mod
    ├── go.sum
    ├── cmd/mcp-server/
    │   └── main.go                # Entry point
    └── internal/
        ├── server/
        │   ├── server.go          # Server lifecycle & transport selection
        │   └── config.go          # Configuration parsing
        ├── tools/
        │   ├── registry.go        # Tool registration logic
        │   └── health.go          # Health check implementation
        └── utils/
            └── validation.go      # Schema validation helpers

### Automation (scripts/)

```text
scripts/
└── Add-MCPToCatalog.ps1           # FR-015: Catalog registration & implementation toggling
```
```

**Structure Decision**: Dual-language project layout with `implementations/typescript/` and `implementations/go/` for parallel development.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A | | |
