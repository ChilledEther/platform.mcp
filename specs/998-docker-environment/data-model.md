# Data Model: Docker Environment

**Feature**: 998-docker-environment  
**Date**: 2026-01-17  
**Status**: Complete

## Overview

This feature is infrastructure-focused with minimal runtime data models. The primary "data" is the Docker build context and generated artifacts.

## Build Artifacts

### Container Images

| Image Name | Source | Entry Point | Purpose |
|------------|--------|-------------|---------|
| `platform:latest` | `build/package/platform/Dockerfile` | `/platform` | CLI tool for file generation |
| `platform-mcp:latest` | `build/package/platform-mcp/Dockerfile` | `/platform-mcp` | MCP server for AI agents |

### Build Context

```yaml
# Build context sent to Docker daemon
context:
  root: /                           # Repository root
  dockerfile: build/package/<name>/Dockerfile
  ignore:
    - .git/
    - .specify/
    - specs/
    - "*.md"
    - scripts/
```

## Configuration Data

### Build Arguments (runtime injection)

```go
// BuildArgs represents values passed at build time
type BuildArgs struct {
    Version   string // Semantic version (e.g., "0.1.0")
    Commit    string // Git SHA (e.g., "abc123f")
    BuildDate string // ISO 8601 timestamp
}
```

### OCI Labels (embedded in image)

```go
// ImageLabels represents OCI-compliant metadata
type ImageLabels struct {
    Version     string `label:"org.opencontainers.image.version"`
    Revision    string `label:"org.opencontainers.image.revision"`
    Created     string `label:"org.opencontainers.image.created"`
    Source      string `label:"org.opencontainers.image.source"`
    Title       string `label:"org.opencontainers.image.title"`
    Description string `label:"org.opencontainers.image.description"`
}
```

## File System Layout (Inside Container)

### CLI Container (`platform`)

```text
/
├── platform              # Binary (compiled, stripped)
└── workspace/            # Volume mount point
```

### MCP Container (`platform-mcp`)

```text
/
├── platform-mcp          # Binary (compiled, stripped)
└── config/               # Optional config mount point
```

## Environment Variables

### Runtime Environment

| Variable | Container | Purpose | Default |
|----------|-----------|---------|---------|
| `PLATFORM_LOG_LEVEL` | Both | Logging verbosity | `info` |
| `PLATFORM_OUTPUT_DIR` | CLI | Default output directory | `/workspace` |

### Build-time Only

| Variable | Purpose | Example |
|----------|---------|---------|
| `VERSION` | Image version label | `0.1.0` |
| `COMMIT` | Git commit SHA | `abc123f` |
| `BUILD_DATE` | Build timestamp | `2026-01-17T10:00:00Z` |

## Relationships

```text
┌─────────────────────────────────────────────────────────────┐
│                     Build Process                            │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │   go.mod    │────▶│  Dockerfile │────▶│   Image     │   │
│  │   go.sum    │     │ (multi-stage)│     │  (Alpine)   │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
│         │                   │                   │           │
│         ▼                   ▼                   ▼           │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │    pkg/     │────▶│   Binary    │────▶│  Container  │   │
│  │    cmd/     │     │ (stripped)  │     │  (running)  │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## Size Budget

| Component | Target Size | Notes |
|-----------|-------------|-------|
| Alpine base | ~5 MB | Minimal OS layer |
| Go binary | ~15-30 MB | Depends on dependencies |
| CA certificates | ~200 KB | For HTTPS |
| **Total** | **< 50 MB** | SC-001 requirement |
