# Development Guide 🛠️

This document outlines the development workflow and automation scripts for Platform MCP.

## 🚀 Development Workflow

1. **Specifications First**: Before any code is written, a specification must be created in `specs/###-feature-name/spec.md`.
2. **Implementation Plan**: Run `/speckit.plan` to generate `plan.md` and related design documents.
3. **TDD Implementation**: Write tests first in `pkg/` or `cmd/`, then implement the logic.
4. **Verification**: Run tests locally and in Docker.

## 📜 Automation Scripts (PowerShell)

We use PowerShell (`pwsh`) for all repository-level automation.

### `Invoke-DockerBuild.ps1`
Builds optimized, multi-stage Docker images for the project artifacts.
- **Usage**: `./scripts/Invoke-DockerBuild.ps1`
- **Output**: `platform-mcp:latest` image.

### `Add-MCPToCatalog.ps1`
Registers the built MCP server into the local Docker Desktop MCP catalog.
- **Usage**: `./scripts/Add-MCPToCatalog.ps1`
- **Prerequisites**: Requires `docker` with MCP support enabled.

## 🧪 Testing

```bash
# Run all Go tests
go test ./...

# Run integration tests specifically
go test -v ./test/integration/...
```

## 🧹 Linting

We use `golangci-lint` to ensure code quality.

```bash
golangci-lint run
```
