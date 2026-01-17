# Technical Research: Deployment Pipeline

**Date**: 2026-01-17  
**Feature**: 999-deployment  
**Status**: Complete

## Summary

This research documents the technical approach for implementing a CI/CD pipeline using GitHub Actions, Release Please for automated versioning, and GitHub Container Registry (GHCR) for container image publishing.

## Technology Decisions

### 1. CI/CD Platform: GitHub Actions

**Why**: Native integration with GitHub repository, free tier for public repos, extensive marketplace of reusable actions.

**Key workflows needed**:
- `ci.yml` - Run tests on push/PR (all branches)
- `release-please.yml` - Create release PRs on merge to main
- `publish.yml` - Build and push container images on release

### 2. Release Automation: Release Please

**Why**: Mandated by constitution (Section V). Generates changelogs automatically from Conventional Commits.

**Configuration files**:
- `.release-please-manifest.json` - Tracks current versions
- `release-please-config.json` - Configures behavior per component

**Multi-component setup** (per constitution):
- `platform` - CLI binary
- `platform-mcp` - MCP server binary

### 3. Container Registry: GitHub Container Registry (GHCR)

**Why**: Native GitHub integration, free for public packages, supports multi-architecture images.

**Image naming**:
- `ghcr.io/modelcontextprotocol/platform:v1.0.0`
- `ghcr.io/modelcontextprotocol/platform-mcp:v1.0.0`

**Tags per release**:
- Semantic version (`v1.0.0`)
- `latest` (updated on each release)

### 4. Docker Build: Multi-stage Alpine

**Why**: Mandated by constitution (Section V). Produces minimal images.

**Location** (per constitution):
- `build/package/platform/Dockerfile`
- `build/package/platform-mcp/Dockerfile`

## Workflow Specifications

### CI Workflow (`ci.yml`)

**Triggers**: `push`, `pull_request`
**Jobs**:
1. `test` - Run `go test ./...`
2. `lint` - Run `golangci-lint` (optional)
3. `build` - Verify binaries compile

**Caching**: Go modules cached via `actions/setup-go`

### Release Please Workflow (`release-please.yml`)

**Triggers**: `push` to `main`
**Behavior**:
- Creates/updates release PRs with version bumps
- On PR merge, creates GitHub Release with tag
- Outputs release metadata for downstream workflows

### Publish Workflow (`publish.yml`)

**Triggers**: `release` event (type: `published`)
**Jobs**:
1. Build multi-arch images (amd64, arm64)
2. Push to GHCR with version + latest tags
3. Update container registry metadata

## Dependencies (GitHub Actions)

| Action | Version | Purpose |
|--------|---------|---------|
| `actions/checkout` | v4 | Clone repository |
| `actions/setup-go` | v5 | Install Go with caching |
| `googleapis/release-please-action` | v4 | Automated releases |
| `docker/setup-buildx-action` | v3 | Multi-arch builds |
| `docker/login-action` | v3 | GHCR authentication |
| `docker/build-push-action` | v6 | Build and push images |
| `docker/metadata-action` | v5 | Generate tags/labels |

## Directory Structure Impact

New files to create:
```text
.github/
├── workflows/
│   ├── ci.yml
│   ├── release-please.yml
│   └── publish.yml

build/
└── package/
    ├── platform/
    │   └── Dockerfile
    └── platform-mcp/
        └── Dockerfile

release-please-config.json
.release-please-manifest.json
```

## Risk Assessment

| Risk | Mitigation |
|------|------------|
| Container registry unavailable | Retry with exponential backoff in workflow |
| Release Please version conflict | Pin to specific action version |
| Docker build cache invalidation | Use GitHub Actions cache mounts |
| Multi-arch build slow | Build architectures in parallel |

## Constraints

- Go 1.25+ required (per constitution)
- Alpine-based images required (per constitution)
- Conventional Commits required for Release Please
- GHCR authentication requires `GITHUB_TOKEN` permissions

## References

- [Release Please](https://github.com/googleapis/release-please)
- [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
- [Docker Build Push Action](https://github.com/docker/build-push-action)
- [Conventional Commits](https://www.conventionalcommits.org/)
