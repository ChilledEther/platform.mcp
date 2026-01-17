# Implementation Plan: Deployment Pipeline

**Branch**: `999-deployment` | **Date**: 2026-01-17 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/999-deployment/spec.md`

## Summary

Implement a complete CI/CD pipeline using GitHub Actions for automated testing, Release Please for versioning and changelogs, and GitHub Container Registry for publishing container images. The pipeline requires zero manual steps for standard releases while providing immediate feedback to developers on push.

## Technical Context

**Language/Version**: Go 1.25+ (per constitution)  
**Primary Dependencies**: GitHub Actions, Release Please, Docker Buildx  
**Storage**: N/A (GitHub-managed: Releases, Container Registry)  
**Testing**: `go test ./...` executed in CI  
**Target Platform**: GitHub Actions (ubuntu-latest runners)  
**Project Type**: CI/CD configuration files + Dockerfiles  
**Performance Goals**: CI feedback in <5 minutes, images published in <15 minutes  
**Constraints**: Must use Conventional Commits, Alpine-based images required  
**Scale/Scope**: Two binaries (platform, platform-mcp), multi-arch images (amd64, arm64)

## Constitution Check

*GATE: All checks pass*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Core-First Architecture | N/A | This feature is infrastructure, not business logic |
| II. TDD | PASS | CI workflow runs tests; no new Go code to test |
| III. Modular Extensibility | PASS | Workflows are independent, components configured separately |
| IV. Separate Artifacts | PASS | Pipeline builds/publishes both binaries independently |
| V. Docker-First Deployment | PASS | Multi-stage Alpine Dockerfiles in `build/package/` |
| VI. MCP Tool Naming | N/A | No new MCP tools |

## Project Structure

### Documentation (this feature)

```text
specs/999-deployment/
├── plan.md              # This file
├── research.md          # Technology decisions and rationale
├── data-model.md        # Configuration schemas and workflow outputs
├── quickstart.md        # Verification steps
├── contracts/           # Workflow specifications
│   └── workflows.md
└── tasks.md             # Phase 2 output (NOT created by /speckit.plan)
```

### Source Code (repository root)

```text
.github/
├── workflows/
│   ├── ci.yml               # Test and build on push/PR
│   ├── release-please.yml   # Automated versioning
│   └── publish.yml          # Container image publishing
├── dependabot.yml           # (existing)

build/
└── package/
    ├── platform/
    │   └── Dockerfile       # CLI image
    └── platform-mcp/
        └── Dockerfile       # MCP server image

release-please-config.json   # Release Please configuration
.release-please-manifest.json # Version tracking
```

**Structure Decision**: This feature adds CI/CD infrastructure files only. No changes to `pkg/`, `cmd/`, or `internal/` directories. Dockerfiles placed in `build/package/` per constitution.

## Complexity Tracking

No constitution violations. All requirements align with mandated standards.

## Implementation Phases

### Phase 1: CI Workflow
- Create `.github/workflows/ci.yml`
- Configure test job with Go caching
- Configure build job for both binaries
- Add status checks to PRs

### Phase 2: Release Please
- Create `release-please-config.json` for multi-component Go project
- Create `.release-please-manifest.json` with initial versions
- Create `.github/workflows/release-please.yml`
- Test with a conventional commit

### Phase 3: Dockerfiles
- Create `build/package/platform/Dockerfile`
- Create `build/package/platform-mcp/Dockerfile`
- Verify multi-stage builds produce minimal images

### Phase 4: Container Publishing
- Create `.github/workflows/publish.yml`
- Configure multi-arch builds (amd64, arm64)
- Configure GHCR authentication and tagging
- Test with a release

## Dependencies

External actions (pinned versions):
- `actions/checkout@v4`
- `actions/setup-go@v5`
- `googleapis/release-please-action@v4`
- `docker/setup-buildx-action@v3`
- `docker/login-action@v3`
- `docker/build-push-action@v6`
- `docker/metadata-action@v5`
