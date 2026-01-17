# Research: Docker Environment

**Feature**: 998-docker-environment  
**Date**: 2026-01-17  
**Status**: Complete

## Executive Summary

This feature creates Docker build infrastructure for the Platform MCP project, producing two minimal Alpine-based container images for the `platform` CLI and `platform-mcp` MCP server. The implementation follows Constitution V (Docker-First Deployment) and uses multi-stage builds to achieve images under 50MB.

## Technology Decisions

### Base Image: Alpine Linux

**Decision**: Use `golang:1.25-alpine` for build stage, `alpine:3.20` for runtime.

**Rationale**:
- Alpine produces smallest possible images (~5MB base)
- Go produces static binaries - no runtime dependencies needed
- Constitution mandates "minimal Alpine-based images"
- Industry standard for Go containerization

**Alternatives Rejected**:
- `scratch` - No shell for debugging, no ca-certificates
- `distroless` - Larger than Alpine, less flexibility
- `debian-slim` - 80MB+ base size

### Multi-Stage Build Pattern

**Decision**: Three-stage build: dependencies → build → runtime

```dockerfile
# Stage 1: Dependencies (cached)
FROM golang:1.25-alpine AS deps
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

# Stage 2: Build
FROM deps AS build
COPY . .
RUN go build -ldflags="-s -w" -o /app ./cmd/<binary>

# Stage 3: Runtime
FROM alpine:3.20
COPY --from=build /app /app
ENTRYPOINT ["/app"]
```

**Rationale**:
- Separate dependency stage enables Docker layer caching
- Build tools excluded from final image (FR-002, FR-004)
- `-ldflags="-s -w"` strips debug symbols (~30% size reduction)

### Build Arguments & Labels

**Decision**: Use OCI-standard labels and build-time arguments.

```dockerfile
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_DATE=unknown

LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.revision="${COMMIT}"
LABEL org.opencontainers.image.created="${BUILD_DATE}"
LABEL org.opencontainers.image.source="https://github.com/modelcontextprotocol/platform.mcp"
```

**Rationale**: OCI labels are the industry standard (FR-010), enable traceability.

## Directory Structure

Per Constitution directory structure:

```text
build/
└── package/
    ├── platform/
    │   └── Dockerfile
    └── platform-mcp/
        └── Dockerfile
```

This matches the existing PowerShell script expectation at `scripts/Invoke-DockerBuild.ps1:36`.

## Prerequisites Analysis

### Missing Components

The current codebase lacks `cmd/` entry points:
- `cmd/platform/main.go` - CLI entry point (depends on 001-core-foundation)
- `cmd/platform-mcp/main.go` - MCP server entry point (depends on 001-core-foundation)

**Impact**: Dockerfiles can be created but won't build successfully until entry points exist.

**Recommendation**: Create stub entry points that compile but have minimal functionality.

### Existing Components

- `pkg/scaffold/` - Core logic exists and is testable
- `scripts/Invoke-DockerBuild.ps1` - Build script ready, expects Dockerfile
- `go.mod` - Module initialized with Go 1.25.5

## Health Check Strategy

**Decision**: Use built-in command execution for health checks.

For CLI (`platform`):
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/platform", "--version"]
```

For MCP (`platform-mcp`):
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/platform-mcp", "--health"]
```

**Rationale**: Self-contained health checks, no external dependencies (FR-008).

## Volume Mount Strategy

For CLI operations that need file I/O:
```dockerfile
VOLUME ["/workspace"]
WORKDIR /workspace
```

**Usage**: `docker run -v $(pwd):/workspace platform generate`

## Testing Strategy

### Container Build Tests

1. **Build Success**: Both images build without errors
2. **Size Constraint**: Images under 50MB (SC-001)
3. **Binary Execution**: Container runs and responds to `--version`
4. **Label Verification**: OCI labels present and correct

### Test Script Location

`scripts/Test-Docker.ps1` - PowerShell script per global standards.

## Security Considerations

1. **Non-root user**: Create unprivileged user in container
2. **Read-only filesystem**: Where possible
3. **No secrets in image**: Build args only, no ENV secrets
4. **Minimal attack surface**: Only binary in final image

```dockerfile
RUN addgroup -S app && adduser -S app -G app
USER app
```

## Performance Targets

| Metric | Target | Measurement |
|--------|--------|-------------|
| Image size | < 50MB | `docker images` |
| Build time | < 5 min | Script timing |
| Startup time | < 3s | Health check |
| Layer count | < 5 | `docker history` |

## References

- [Docker Best Practices for Go](https://docs.docker.com/language/golang/build-images/)
- [OCI Image Spec - Annotations](https://github.com/opencontainers/image-spec/blob/main/annotations.md)
- [Alpine Docker Hub](https://hub.docker.com/_/alpine)
