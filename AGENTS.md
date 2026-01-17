# mcp.github.agentic Development Guidelines

Auto-generated from all feature plans. Last updated: 2026-01-16

## Active Technologies

- TypeScript/Bun (implementations/typescript)
- Go 1.25+ (implementations/go)
- YAML (GitHub Actions Workflow)

## Project Structure

```text
mcp.github.agentic/
├── specs/                           # Shared specifications (source of truth)
│   ├── 001-infra-mcp-foundation/
│   ├── 002-azure-tools/
│   └── 999-deployment/
├── implementations/
│   ├── go/                          # Go implementation
│   │   ├── cmd/mcp-server/
│   │   ├── internal/
│   │   ├── Dockerfile
│   │   └── go.mod
│   └── typescript/                  # TypeScript/Bun implementation
│       ├── server/
│       ├── tools/
│       ├── utils/
│       ├── tests/
│       ├── Dockerfile
│       └── package.json
├── .github/workflows/               # CI/CD
├── .specify/                        # Spec toolkit
└── scripts/                         # PowerShell automation
```

## Commands

### TypeScript

```bash
cd implementations/typescript
bun install              # Install dependencies
bun run dev              # Run with hot-reload
bun test                 # Run tests
bun typecheck            # Type check
```

### Go

```bash
cd implementations/go
go build ./cmd/mcp-server   # Build
go test ./...               # Run tests
```

## Git Worktree Workflow

For parallel development on both implementations:

```powershell
# Create worktrees for parallel agent work
git worktree add ../mcp-go-worktree feature/new-tool
git worktree add ../mcp-ts-worktree feature/new-tool

# Agent 1 works on Go: ../mcp-go-worktree/implementations/go/
# Agent 2 works on TS: ../mcp-ts-worktree/implementations/typescript/
```

## Docker Images

| Implementation | Image Name |
|:---|:---|
| TypeScript | `ghcr.io/<owner>/mcp-github-ts` |
| Go | `ghcr.io/<owner>/mcp-github-go` |

## Code Style

- **TypeScript**: Strict mode, Zod for validation, `@modelcontextprotocol/sdk`
- **Go**: Standard library first, `github.com/modelcontextprotocol/go-sdk`
- **YAML**: Follow standard conventions for GitHub Actions

## Recent Changes

- 2026-01-16: Reorganized to dual-language structure with `implementations/` folder
- 999-github-deployment-workflow: GitHub Actions CI/CD

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
