# Data Model: Deployment Pipeline

**Date**: 2026-01-17  
**Feature**: 999-deployment  
**Status**: Complete

## Overview

This feature primarily creates GitHub Actions workflows and configuration files. The "data model" consists of YAML schemas for workflows and Release Please configuration.

## Configuration Schemas

### Release Please Manifest (`.release-please-manifest.json`)

Tracks current versions for each component:

```json
{
  "cmd/platform": "0.0.1",
  "cmd/platform-mcp": "0.0.1"
}
```

### Release Please Config (`release-please-config.json`)

Controls release behavior:

```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json",
  "release-type": "go",
  "packages": {
    "cmd/platform": {
      "component": "platform",
      "changelog-path": "CHANGELOG.md"
    },
    "cmd/platform-mcp": {
      "component": "platform-mcp",
      "changelog-path": "CHANGELOG.md"
    }
  },
  "separate-pull-requests": false,
  "bump-minor-pre-major": true,
  "bump-patch-for-minor-pre-major": true
}
```

## Workflow Data Structures

### CI Workflow Outputs

| Output | Type | Description |
|--------|------|-------------|
| `test_result` | `success` \| `failure` | Test execution status |
| `build_result` | `success` \| `failure` | Binary compilation status |

### Release Please Outputs

| Output | Type | Description |
|--------|------|-------------|
| `releases_created` | boolean | Whether any releases were created |
| `tag_name` | string | Git tag for the release (e.g., `platform-v1.0.0`) |
| `release_created` | boolean | Whether a release was created |

### Publish Workflow Inputs

| Input | Type | Required | Description |
|-------|------|----------|-------------|
| `tag_name` | string | From release | Version tag to publish |
| `component` | string | From release | Which binary to publish |

## Container Image Metadata

Images are tagged with OCI annotations:

| Label | Value |
|-------|-------|
| `org.opencontainers.image.source` | Repository URL |
| `org.opencontainers.image.version` | Semantic version |
| `org.opencontainers.image.created` | Build timestamp |
| `org.opencontainers.image.title` | `platform` or `platform-mcp` |
| `org.opencontainers.image.description` | Binary description |

## Dockerfile Build Args

| Arg | Default | Description |
|-----|---------|-------------|
| `GO_VERSION` | `1.25` | Go version for build stage |
| `ALPINE_VERSION` | `3.21` | Alpine version for runtime |

## State Management

No persistent state is required. All state is managed by:
- GitHub Actions (workflow runs)
- GitHub Releases (version history)
- GitHub Container Registry (image storage)
- Git tags (version references)

## Relationships

```text
push to main
    │
    ▼
┌─────────────────┐
│ release-please  │──creates──▶ Release PR
└────────┬────────┘
         │ merge
         ▼
┌─────────────────┐
│ GitHub Release  │──triggers──▶ publish.yml
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Container Image │──tagged──▶ GHCR
└─────────────────┘
```
