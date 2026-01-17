<!-- SYNC IMPACT REPORT
Version: 2.2.1 -> 3.0.0
Changes:
- MAJOR: Dual-language support (TypeScript + Go) with shared specs
- MAJOR: New `implementations/` directory structure
- MAJOR: Git worktree workflow for parallel agent development
- MINOR: Removed legacy code principle (now active dual implementations)
- Templates:
  - .specify/templates/plan-template.md: ✅ Compatible (no changes required)
  - .specify/templates/spec-template.md: ✅ Compatible (no changes required)
  - .specify/templates/tasks-template.md: ✅ Compatible (no changes required)
Follow-up TODOs: Update CI/CD workflow for matrix builds
-->

# MCP GitHub Agentic Constitution

## Core Principles

### I. Dual-Language Implementation
The project maintains **parallel implementations** in both TypeScript and Go, serving as an experimentation platform for Spec-Driven Development (SDD).
- **TypeScript**: Located in `implementations/typescript/`, runs on Bun 1.x.
- **Go**: Located in `implementations/go/`, uses Go 1.25+.
- **Spec-First**: Both implementations MUST follow the same specifications in `specs/`.
- **Feature Parity**: New features SHOULD be implemented in both languages.
- **Reject Frontend**: Any UI, frontend frameworks, or client-side assets are explicitly forbidden.

### Ia. TypeScript-Specific Standards
- **Language**: TypeScript 5.x+ with strict mode enabled.
- **Runtime**: Bun 1.x (latest stable).
- **Tooling**: Use Bun toolchain (`bun run`, `bun test`, `bunx`).
- **SDK**: `@modelcontextprotocol/sdk` npm package.

### Ib. Go-Specific Standards
- **Language**: Go 1.25+ with standard library preferred.
- **SDK**: `github.com/modelcontextprotocol/go-sdk`.
- **Tooling**: Standard Go toolchain (`go build`, `go test`).

### II. Multi-Mode MCP Transport
The server MUST implement the official Model Context Protocol (MCP) JSON-RPC 2.0 specification.
- **SDK**: Both implementations MUST use their respective official MCP SDKs.
- **Modes**: MUST support both Standard I/O (stdio) for local shell integration AND Streamable HTTP for remote/containerized operation.
- **Default**: Stdio should be the default mode if no flags are provided.
- **Protocol**: JSON-RPC 2.0.

### III. Strict Type & Schema Safety
Security and robustness are paramount.
- **Type Safety**: Use TypeScript strict mode with no `any` types in production code.
- **Schema Generation**: Use Zod or similar for runtime validation and JSON schema generation.
- **Identifiers**: Use `crypto.randomUUID()` or a UUID library for all internal and external identifiers.
- **Validation**: Every input MUST be validated before execution using Zod schemas.

### IV. Concurrency & Timeout Management
- **Async Execution**: Use async/await patterns with proper Promise handling.
- **Context Awareness**: Use AbortController and AbortSignal for timeout propagation.
- **Graceful Shutdown**: Handle SIGINT/SIGTERM signals properly.

### X. Minimalist Dependency Philosophy
Prioritize a slim dependency tree to minimize security surface area and footprint.
- **Standard Library First**: Always favor Bun/Node built-in APIs over external packages.
- **Trusted Repositories**: Favor packages from official and trusted sources (e.g., `@modelcontextprotocol`, `zod`, well-maintained packages).
- **Careful Selection**: Avoid "utility" libraries that introduce large transitive dependency trees. Implement simple logic internally rather than adding a dependency.

### XI. Pragmatic Testing
Testing should provide high value and confidence without being overly burdensome.
- **Testing Framework**: Use `bun:test` for unit and integration testing.
- **Encourage TDD**: Test-driven development is encouraged where it helps drive design, but not strictly mandated.
- **Small Project Context**: Given the project's focused scope, prioritize meaningful coverage over absolute unit test parity.
- **Focus on Core & Tools**: Focus intensive testing on the core MCP server logic and the specific tools being exposed. Ensure tool schemas and tool logic are robustly verified.

### XII. Cloud-Native & Container Ready
The artifact MUST be deployable to cloud environments.
- **Docker**: Each implementation MUST maintain its own `Dockerfile` in its directory.
- **Images**: Separate Docker images per language: `mcp-github-go` and `mcp-github-ts`.
- **Compilation**: Production images MUST use compiled standalone binaries for minimal footprint.
- **Kubernetes**: Configuration for K8s deployment SHOULD be considered/provided.
- **Observability**: Logs MUST go to stderr/stdout.

### XIII. Git Worktree Workflow for Parallel Development
To enable parallel agent work on both implementations:
- **Feature Branches**: Each feature SHOULD have a single spec branch with implementation sub-branches.
- **Worktrees**: Use `git worktree` to check out feature branches in parallel directories.
- **Parallel Agents**: Multiple agents can work simultaneously on Go and TypeScript implementations.
- **Spec-First**: Always update `specs/` before implementing in either language.

Example workflow:
```bash
# Create feature branch from main
git checkout -b feature/new-tool main

# Create worktrees for parallel implementation
git worktree add ../mcp-go-worktree feature/new-tool
git worktree add ../mcp-ts-worktree feature/new-tool

# Agents work in parallel directories
# Agent 1: ../mcp-go-worktree/implementations/go/
# Agent 2: ../mcp-ts-worktree/implementations/typescript/
```

### XIV. Minimal Resource Footprint
The project aims for extreme efficiency and low resource consumption.
- **Memory Limit**: The runtime memory usage MUST be capped at **128 MiB** for the containerized server.
- **Compilation**: Use `bun build --compile` with `--minify` to produce a standalone executable.
- **Architecture**: Prefer stateless operations and streaming to avoid buffering large datasets in memory.

### XV. Structured Logging
Consistent and machine-readable logging is essential for observability.
- **Output Stream**: All logs MUST be written to `stderr` to avoid corrupting `stdout` (reserved for MCP JSON-RPC messages).
- **Format**: Logs SHOULD use a structured format (JSON) in production environments. Plain text with level prefixes (e.g., `[INFO]`) is acceptable for development and MVP stages.
- **Levels**: Use appropriate log levels (`DEBUG`, `INFO`, `WARN`, `ERROR`) to categorize log messages.
- **Context**: Include relevant context (e.g., execution ID, tool name) in log entries to facilitate debugging.

## Security & Safety

### V. Default Read-Only
- All operations are read-only by default.
- Write operations must be explicitly enabled and require safe-by-design workflows (e.g., branch-then-PR).

### VI. Secret Management
- **No Hardcoding**: Secrets (API tokens, app IDs) must NEVER be hardcoded.
- **Runtime Fetching**: Fetch secrets from environment variables or a dedicated secret store at runtime.

### VII. Input Sanitization
- All paths must be sanitized to prevent traversal.
- All commands/queries must be escaped/validated to prevent injection.

## Repository Structure & Tooling

### VIII. Specification Kit Alignment
Maintain structured documentation for all feature planning:
- `.specify/memory/`: Project constraints and constitutional documents.
- `specs/`: Feature specifications and implementation plans (at repository root for visibility).
- `.specify/templates/`: Templates for specs, plans, and tasks.

### IX. Scripting & Automation
Align with **Antigravity Global Standards**:
- **Scripts**: All automation MUST use **PowerShell** (`.ps1`). Bash scripts are not permitted.
- **Naming**: PowerShell scripts MUST use PascalCase Verb-Noun format (e.g., `Invoke-DockerBuild.ps1`).
- **Data**: Use **YAML** for configuration; avoid JSON where human readability is preferred.
- **Package Manager**: Use `bun` for package management and script execution.
- **Diagrams**: When creating diagrams in markdown files, MUST use **Mermaid** syntax instead of ASCII art. ASCII diagrams generated by agents are often malformed or misaligned; Mermaid provides consistent, renderable output.

## Governance

1. **Constitution Over All**: This document supersedes any local README or ad-hoc convention.
2. **Amendments**: Changes require a specific "Constitutional Amendment" PR with justified reasoning.
3. **Drafting Rule**: Before implementation, a Spec and Plan must be approved in the `specs/` directory.

**Version**: 3.0.0 | **Ratified**: 2026-01-16 | **Last Amended**: 2026-01-16
