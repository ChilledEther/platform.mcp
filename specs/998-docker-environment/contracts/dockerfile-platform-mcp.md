# Dockerfile Contract: Platform MCP

**File**: `build/package/platform-mcp/Dockerfile`  
**Purpose**: Multi-stage build for MCP server binary

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
    -o /platform-mcp \
    ./cmd/platform-mcp

# ============================================================================
# Stage 3: Runtime
# ============================================================================
FROM alpine:3.20

# OCI Labels
LABEL org.opencontainers.image.title="platform-mcp"
LABEL org.opencontainers.image.description="Platform MCP Server - AI agent tool provider"
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
COPY --from=build /platform-mcp /platform-mcp

# Create config directory for optional mounts
VOLUME ["/config"]

# Switch to non-root user
USER app

# MCP servers communicate via stdio
# Health check uses a simple flag
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/platform-mcp", "--health"]

ENTRYPOINT ["/platform-mcp"]
```

## Acceptance Criteria

- [ ] Image builds successfully with `docker build`
- [ ] Final image size < 50MB
- [ ] Binary executes with `--health` flag
- [ ] Non-root user (UID != 0)
- [ ] OCI labels present in image metadata
- [ ] MCP protocol communication works via stdio
- [ ] Health check passes within 30s

## Usage Examples

```bash
# Build
docker build -t platform-mcp:latest -f build/package/platform-mcp/Dockerfile .

# Run as stdio MCP server
docker run -i platform-mcp

# Check health
docker run platform-mcp --health

# With config mount
docker run -v ./config:/config platform-mcp

# Inspect labels
docker inspect platform-mcp:latest --format '{{json .Config.Labels}}'
```

## MCP Integration Notes

MCP servers communicate via stdio (stdin/stdout). When used with AI agents:

```json
{
  "mcpServers": {
    "platform": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "platform-mcp:latest"]
    }
  }
}
```
