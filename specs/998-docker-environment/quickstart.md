# Quickstart: Docker Environment

**Feature**: 998-docker-environment  
**Date**: 2026-01-17  
**Estimated Effort**: 2-3 hours

## Prerequisites

- Docker 20.10+ installed and running
- PowerShell 7+ (for build scripts)
- Go 1.25+ (for local development)

## Implementation Order

```text
┌─────────────────────────────────────────────────────────────┐
│  1. Create directory structure                              │
│     └── build/package/{platform,platform-mcp}/              │
├─────────────────────────────────────────────────────────────┤
│  2. Create stub entry points (if not exist)                 │
│     └── cmd/{platform,platform-mcp}/main.go                 │
├─────────────────────────────────────────────────────────────┤
│  3. Create Dockerfiles                                      │
│     └── Use contracts as reference                          │
├─────────────────────────────────────────────────────────────┤
│  4. Create Test-Docker.ps1 script                           │
│     └── Validates size, labels, user                        │
├─────────────────────────────────────────────────────────────┤
│  5. Create .dockerignore                                    │
│     └── Exclude unnecessary files from context              │
├─────────────────────────────────────────────────────────────┤
│  6. Update Invoke-DockerBuild.ps1                           │
│     └── Support both images, add build args                 │
├─────────────────────────────────────────────────────────────┤
│  7. Verify all tests pass                                   │
│     └── ./scripts/Test-Docker.ps1                           │
└─────────────────────────────────────────────────────────────┘
```

## Quick Commands

```bash
# Create directory structure
mkdir -p build/package/platform build/package/platform-mcp

# Build platform CLI image
docker build -t platform:latest -f build/package/platform/Dockerfile .

# Build platform-mcp image  
docker build -t platform-mcp:latest -f build/package/platform-mcp/Dockerfile .

# Run tests
pwsh scripts/Test-Docker.ps1

# Check image size
docker images platform --format "table {{.Repository}}\t{{.Size}}"
```

## Files to Create

| File | Purpose | Contract |
|------|---------|----------|
| `build/package/platform/Dockerfile` | CLI container | `contracts/dockerfile-platform.md` |
| `build/package/platform-mcp/Dockerfile` | MCP container | `contracts/dockerfile-platform-mcp.md` |
| `scripts/Test-Docker.ps1` | Validation script | `contracts/test-docker-script.md` |
| `.dockerignore` | Build context filter | See below |
| `cmd/platform/main.go` | CLI entry (stub) | Minimal compilable |
| `cmd/platform-mcp/main.go` | MCP entry (stub) | Minimal compilable |

## .dockerignore Template

```text
# Git
.git/
.gitignore

# Specs and docs
.specify/
specs/
*.md

# Scripts (not needed in container)
scripts/

# IDE
.vscode/
.idea/

# Build artifacts
*.exe
*.test
*.out
```

## Stub Entry Points

If `cmd/` doesn't exist, create minimal stubs:

**cmd/platform/main.go**:
```go
package main

import (
    "fmt"
    "os"
)

var version = "dev"
var commit = "unknown"

func main() {
    if len(os.Args) > 1 && os.Args[1] == "--version" {
        fmt.Printf("platform %s (%s)\n", version, commit)
        return
    }
    fmt.Println("platform: scaffold generator")
}
```

**cmd/platform-mcp/main.go**:
```go
package main

import (
    "fmt"
    "os"
)

var version = "dev"
var commit = "unknown"

func main() {
    if len(os.Args) > 1 {
        switch os.Args[1] {
        case "--version":
            fmt.Printf("platform-mcp %s (%s)\n", version, commit)
            return
        case "--health":
            fmt.Println("ok")
            return
        }
    }
    fmt.Println("platform-mcp: MCP server ready")
}
```

## Verification Checklist

- [ ] Both images build without errors
- [ ] Image sizes < 50MB each
- [ ] `docker run platform --version` works
- [ ] `docker run platform-mcp --health` returns "ok"
- [ ] OCI labels present: `docker inspect platform:latest --format '{{json .Config.Labels}}'`
- [ ] Non-root user: `docker inspect platform:latest --format '{{.Config.User}}'`
- [ ] Volume mounts work: `docker run -v $(pwd):/workspace platform ls /workspace`

## Common Issues

| Issue | Solution |
|-------|----------|
| "go.sum missing" | Run `go mod tidy` before build |
| "undefined: Config" | Ensure all Go files compile locally first |
| Image too large | Check for unnecessary files in context, verify .dockerignore |
| Permission denied | Ensure volume mount permissions match container user |
