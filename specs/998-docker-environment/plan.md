# Implementation Plan: Docker Environment

**Branch**: `998-docker-environment` | **Date**: 2026-01-17 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `/specs/998-docker-environment/spec.md`

## Summary

Create Docker build infrastructure for the Platform MCP project, producing two minimal Alpine-based container images (<50MB each) for the `platform` CLI and `platform-mcp` MCP server. Uses multi-stage builds with dependency caching, OCI-compliant labels, and non-root runtime users.

## Technical Context

**Language/Version**: Go 1.25.5  
**Primary Dependencies**: None (static binary, Alpine runtime)  
**Storage**: N/A (stateless containers)  
**Testing**: PowerShell test script, Docker build verification  
**Target Platform**: Linux containers (amd64), compatible with Docker Desktop on macOS/Windows  
**Project Type**: Infrastructure (Dockerfiles, scripts)  
**Performance Goals**: Image size <50MB, build time <5min, startup <3s  
**Constraints**: Constitution V mandates Alpine-based multi-stage builds  
**Scale/Scope**: 2 container images, 2 Dockerfiles, 1 test script

## Constitution Check

*GATE: All items passed*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Core-First Architecture | N/A | Infrastructure feature, no pkg/ changes |
| II. TDD | PASS | Test script validates all acceptance criteria |
| III. Modular Extensibility | N/A | No new modules |
| IV. Separate Artifacts | PASS | Separate images for CLI and MCP |
| V. Docker-First Deployment | PASS | This feature implements the mandate |
| VI. MCP Tool Naming | N/A | No new tools |

## Project Structure

### Documentation (this feature)

```text
specs/998-docker-environment/
├── plan.md              # This file
├── spec.md              # Feature specification
├── research.md          # Phase 0 - Technical research
├── data-model.md        # Phase 1 - Build artifacts model
├── quickstart.md        # Implementation guide
└── contracts/           # Phase 1 - Interface contracts
    ├── dockerfile-platform.md
    ├── dockerfile-platform-mcp.md
    └── test-docker-script.md
```

### Source Code (repository root)

```text
build/
└── package/
    ├── platform/
    │   └── Dockerfile       # CLI container (NEW)
    └── platform-mcp/
        └── Dockerfile       # MCP container (NEW)

cmd/
├── platform/
│   └── main.go              # CLI entry point (STUB if missing)
└── platform-mcp/
    └── main.go              # MCP entry point (STUB if missing)

scripts/
├── Invoke-DockerBuild.ps1   # Build script (EXISTS - UPDATE)
└── Test-Docker.ps1          # Test script (NEW)

.dockerignore                # Build context filter (NEW)
```

**Structure Decision**: Follows Constitution directory structure exactly. Dockerfiles in `build/package/<name>/` per golang-standards/project-layout.

## Implementation Tasks

### Phase 1: Directory Structure & Entry Points

| Task | Description | File(s) |
|------|-------------|---------|
| 1.1 | Create build/package directories | `build/package/platform/`, `build/package/platform-mcp/` |
| 1.2 | Create stub cmd/platform/main.go | `cmd/platform/main.go` |
| 1.3 | Create stub cmd/platform-mcp/main.go | `cmd/platform-mcp/main.go` |
| 1.4 | Create .dockerignore | `.dockerignore` |

### Phase 2: Dockerfiles

| Task | Description | Contract |
|------|-------------|----------|
| 2.1 | Create platform Dockerfile | `contracts/dockerfile-platform.md` |
| 2.2 | Create platform-mcp Dockerfile | `contracts/dockerfile-platform-mcp.md` |
| 2.3 | Verify builds succeed | `docker build` both images |

### Phase 3: Build Scripts

| Task | Description | Contract |
|------|-------------|----------|
| 3.1 | Create Test-Docker.ps1 | `contracts/test-docker-script.md` |
| 3.2 | Update Invoke-DockerBuild.ps1 | Support both images, build args |

### Phase 4: Validation

| Task | Description | Success Criteria |
|------|-------------|------------------|
| 4.1 | Run Test-Docker.ps1 | Exit code 0 |
| 4.2 | Verify image sizes | Both <50MB |
| 4.3 | Verify health checks | Both pass |
| 4.4 | Verify volume mounts | Files accessible |

## Dependencies

```text
998-docker-environment
    ├── [BLOCKED] cmd/platform/main.go (requires 001-core-foundation or stub)
    └── [BLOCKED] cmd/platform-mcp/main.go (requires 001-core-foundation or stub)
```

**Resolution**: Create minimal stub entry points that compile and support `--version` / `--health` flags. These will be replaced by full implementations from feature 001-core-foundation.

## Risk Assessment

| Risk | Impact | Mitigation |
|------|--------|------------|
| go.sum missing | Build fails | Run `go mod tidy` before Docker build |
| Image size exceeds limit | SC-001 fails | Verify .dockerignore, use `-ldflags="-s -w"` |
| Entry points don't exist | Build fails | Create stubs as part of this feature |

## Complexity Tracking

No Constitution violations. Implementation follows all mandates:

- Alpine-based images (V)
- Multi-stage builds (V)
- Separate artifacts (IV)
- PowerShell scripts (global standards)

## Success Metrics

| Metric | Target | Validation |
|--------|--------|------------|
| Image size | <50MB | `docker images --format '{{.Size}}'` |
| Build time | <5min | Script timing |
| Startup time | <3s | Health check interval |
| Test coverage | 100% criteria | Test-Docker.ps1 exit code |

## References

- [Research](./research.md) - Technology decisions and rationale
- [Data Model](./data-model.md) - Build artifacts and configuration
- [Quickstart](./quickstart.md) - Step-by-step implementation guide
- [Contracts](./contracts/) - Detailed file specifications
