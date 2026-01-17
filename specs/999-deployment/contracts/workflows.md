# GitHub Workflow Contracts

**Date**: 2026-01-17  
**Feature**: 999-deployment

## ci.yml Contract

**Location**: `.github/workflows/ci.yml`

### Triggers
- `push` to any branch
- `pull_request` to any branch

### Jobs

#### test
- **Runs on**: `ubuntu-latest`
- **Steps**:
  1. Checkout code
  2. Setup Go 1.25+ with module caching
  3. Run `go test ./...`
- **Outputs**: Test results

#### build
- **Runs on**: `ubuntu-latest`
- **Needs**: `test` (sequential)
- **Steps**:
  1. Checkout code
  2. Setup Go 1.25+
  3. Build `cmd/platform`
  4. Build `cmd/platform-mcp`
- **Outputs**: Binary artifacts (not persisted)

### Required Secrets
- None (uses default `GITHUB_TOKEN`)

### Required Permissions
- `contents: read`

---

## release-please.yml Contract

**Location**: `.github/workflows/release-please.yml`

### Triggers
- `push` to `main` branch only

### Jobs

#### release-please
- **Runs on**: `ubuntu-latest`
- **Steps**:
  1. Run `googleapis/release-please-action@v4`
- **Outputs**:
  - `releases_created`: boolean
  - `platform--release_created`: boolean
  - `platform--tag_name`: string (e.g., `platform-v1.0.0`)
  - `platform-mcp--release_created`: boolean
  - `platform-mcp--tag_name`: string

### Required Secrets
- None (uses default `GITHUB_TOKEN`)

### Required Permissions
- `contents: write`
- `pull-requests: write`

---

## publish.yml Contract

**Location**: `.github/workflows/publish.yml`

### Triggers
- `release` event, type `published`

### Jobs

#### publish-platform
- **Runs on**: `ubuntu-latest`
- **Condition**: Release tag starts with `platform-v`
- **Steps**:
  1. Checkout code
  2. Setup Docker Buildx
  3. Login to GHCR
  4. Extract metadata (tags, labels)
  5. Build multi-arch image (amd64, arm64)
  6. Push to `ghcr.io/${{ github.repository_owner }}/platform`

#### publish-platform-mcp
- **Runs on**: `ubuntu-latest`
- **Condition**: Release tag starts with `platform-mcp-v`
- **Steps**: Same as publish-platform but for MCP binary

### Image Tags Generated
- `ghcr.io/<owner>/platform:v1.0.0`
- `ghcr.io/<owner>/platform:latest`

### Required Secrets
- None (uses default `GITHUB_TOKEN`)

### Required Permissions
- `contents: read`
- `packages: write`

---

## Configuration File Contracts

### release-please-config.json

**Location**: Repository root

```yaml
required_fields:
  - $schema
  - release-type: "go"
  - packages: object with at least one component

packages_structure:
  "cmd/platform":
    component: "platform"
  "cmd/platform-mcp":
    component: "platform-mcp"
```

### .release-please-manifest.json

**Location**: Repository root

```yaml
required_structure:
  "cmd/platform": "semver string"
  "cmd/platform-mcp": "semver string"

initial_values:
  "cmd/platform": "0.0.1"
  "cmd/platform-mcp": "0.0.1"
```

---

## Dockerfile Contracts

### build/package/platform/Dockerfile

**Requirements**:
- Multi-stage build
- Build stage: `golang:1.25-alpine`
- Runtime stage: `alpine:3.21`
- Binary: `/platform`
- Entrypoint: `["/platform"]`

### build/package/platform-mcp/Dockerfile

**Requirements**:
- Multi-stage build
- Build stage: `golang:1.25-alpine`
- Runtime stage: `alpine:3.21`
- Binary: `/platform-mcp`
- Entrypoint: `["/platform-mcp"]`
