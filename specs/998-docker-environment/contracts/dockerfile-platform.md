# Dockerfile Contract: Platform CLI

**File**: `build/package/platform/Dockerfile`  
**Purpose**: Multi-stage build for CLI binary

## Contract

```dockerfile
# syntax=docker/dockerfile:1

# ============================================================================
# Stage 1: Dependencies
# ============================================================================
FROM golang:1.25-alpine AS deps

WORKDIR /src

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Download dependencies (cached layer)
COPY go.mod go.sum ./
RUN go mod download

# ============================================================================
# Stage 2: Build
# ============================================================================
FROM deps AS build

# Build arguments for metadata
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_DATE=unknown

# Copy source code
COPY . .

# Run tests
RUN go test ./...

# Build binary with size optimization
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \
    -o /platform \
    ./cmd/platform

# ============================================================================
# Stage 3: Runtime
# ============================================================================
FROM alpine:3.20

# OCI Labels
LABEL org.opencontainers.image.title="platform"
LABEL org.opencontainers.image.description="Platform CLI - scaffold generator for DevOps files"
LABEL org.opencontainers.image.source="https://github.com/modelcontextprotocol/platform.mcp"
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_DATE=unknown
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.revision="${COMMIT}"
LABEL org.opencontainers.image.created="${BUILD_DATE}"

# Security: non-root user
RUN addgroup -S app && adduser -S app -G app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy binary from build stage
COPY --from=build /platform /platform

# Create workspace for file I/O
VOLUME ["/workspace"]
WORKDIR /workspace

# Switch to non-root user
USER app

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/platform", "--version"]

ENTRYPOINT ["/platform"]
```

## Acceptance Criteria

- [ ] Image builds successfully with `docker build`
- [ ] Final image size < 50MB
- [ ] Binary executes with `--version` flag
- [ ] Non-root user (UID != 0)
- [ ] OCI labels present in image metadata
- [ ] Volume mount works for file output
- [ ] Health check passes within 30s

## Usage Examples

```bash
# Build
docker build -t platform:latest -f build/package/platform/Dockerfile .

# Run with volume mount
docker run -v $(pwd):/workspace platform generate --output /workspace

# Check version
docker run platform --version

# Inspect labels
docker inspect platform:latest --format '{{json .Config.Labels}}'
```
