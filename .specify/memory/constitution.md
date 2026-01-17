# <!--

# SYNC IMPACT REPORT

Version Change: 2.1.1 → 2.2.0 (MINOR - added MCP tool naming standards)

Modified Principles: None

Added Principles:
- VI. MCP Tool Naming Standards

Added Sections: None

Removed Sections: None

Templates Requiring Updates:
- .specify/templates/plan-template.md ✅ (compatible)
- .specify/templates/spec-template.md ✅ (compatible)
- .specify/templates/tasks-template.md ✅ (compatible)
- .specify/templates/checklist-template.md ✅ (compatible)
- .specify/templates/agent-file-template.md ✅ (compatible)

# Follow-up TODOs: None

-->

# Platform MCP Constitution

## Core Principles

### I. Core-First Architecture

All functionality MUST be implemented in shared core packages (`pkg/`) before any consumer (CLI or MCP) is built.

- Core packages are responsible for **logic only**; they return data, never perform I/O directly
- CLI handles disk I/O (writing files, reading user input)
- MCP handles network I/O (returning structured responses to agents)
- Both consumers import the same core—no duplicated logic

**Rationale**: Ensures testability, prevents code duplication, and enables consistent behavior across both artifacts.

### II. Test-Driven Development (TDD) — NON-NEGOTIABLE

Every feature MUST follow the Red-Green-Refactor cycle. No exceptions.

1. **Write tests first**: Tests MUST be written before implementation
2. **User approval**: Test plan MUST be reviewed before implementation begins
3. **Tests MUST fail**: Verify tests fail before writing implementation
4. **Implement to pass**: Write minimal code to make tests pass
5. **Refactor**: Clean up while keeping tests green

**Rationale**: TDD catches design flaws early, produces self-documenting code, and ensures all functionality is verified.

### III. Modular Extensibility

The architecture MUST support adding new features without modifying existing code.

- Each feature (e.g., GitHub Actions generator, Dockerfile generator) is a self-contained module
- Modules register themselves via a common interface (`Generator`, `Tool`, etc.)
- Adding a new module MUST NOT require changes to the CLI or MCP entry points
- Feature flags or configuration control which modules are active

**Rationale**: Enables the platform to grow into a general-purpose MCP while keeping each feature isolated and maintainable.

### IV. Separate Artifacts, Shared Core

This repository produces **two distinct binaries** that share the same core packages:

| Artifact   | Binary Name    | Purpose                                         |
| :--------- | :------------- | :---------------------------------------------- |
| CLI Tool   | `platform`     | For human operators; writes files to disk       |
| MCP Server | `platform-mcp` | For AI agents; returns content via MCP protocol |

**Requirements**:

- Both binaries MUST import logic exclusively from `pkg/`
- The CLI (`platform`) handles disk I/O and user interaction
- The MCP server (`platform-mcp`) handles MCP protocol and agent communication
- No shared entry point—each has its own `main.go` in `cmd/`

**Rationale**: Separate binaries allow independent versioning, smaller container images, and clearer separation of concerns while maintaining a single source of truth for business logic.

### V. Docker-First Deployment

All development, testing, and deployment MUST be containerized.

- Dockerfiles MUST reside in `/build/package/`
- All tests MUST pass inside the container (same environment as production)
- Multi-stage builds MUST produce minimal Alpine-based images
- Local development MAY use native Go, but CI/CD runs in Docker

**Rationale**: Eliminates "works on my machine" issues and ensures consistent behavior across environments.

### VI. MCP Tool Naming Standards

Tools exposed via MCP MUST follow a consistent naming convention to ensure predictability for AI agents.

- Tools MUST be named using `snake_case`
- Names MUST follow a `verb_noun` pattern (e.g., `get_user`, `list_repos`)
- Standard verbs MUST be used where applicable: `get`, `create`, `update`, `delete`, `list`, `search`, `analyze`, `generate`, `validate`
- Nouns MUST be singular unless representing a collection (e.g., `list_items` is acceptable if it returns a list)

**Rationale**: Consistent naming allows AI agents to more easily discover and understand the purpose of tools, leading to higher reliability and better performance.

## Technology Stack

The following technologies are mandated for the Go implementation:

| Component          | Technology                                                     | Notes                                         |
| :----------------- | :------------------------------------------------------------- | :-------------------------------------------- |
| Language           | Go 1.25+                                                       | Use latest stable features                    |
| MCP SDK            | `github.com/modelcontextprotocol/go-sdk`                       | Official MCP Go SDK                           |
| CLI Framework      | `github.com/spf13/cobra`                                       | For `platform` CLI                            |
| Templates          | Go `embed` package                                             | Bundle templates in binaries                  |
| Testing            | `go test` with table-driven tests                              | TDD mandatory                                 |
| Containerization   | Docker (Alpine-based)                                          | Multi-stage builds                            |
| Config Format      | YAML                                                           | All generated files (`.github/` workflows)    |
| Scripts            | Bash (`.sh`)                                                   | Standard                                      |
| Release Automation | [Release Please](https://github.com/googleapis/release-please) | Automated versioning via Conventional Commits |

### Release Please Configuration

This repository uses Google's **Release Please** for automated release management:

- **Conventional Commits**: All commits MUST follow [Conventional Commits](https://www.conventionalcommits.org/) format
- **Automated CHANGELOG**: Release Please generates `CHANGELOG.md` from commit history
- **Version Bumping**: Semantic versioning determined automatically from commit types:
  - `feat:` → MINOR bump
  - `fix:` → PATCH bump
  - `feat!:` or `BREAKING CHANGE:` → MAJOR bump
- **Release PRs**: Release Please creates PRs to prepare releases
- **Multi-component**: Configured for monorepo with separate versioning per artifact

**Commit Format**:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

**Examples**:

```
feat(github): add workflow generation for Go projects
fix(scaffold): correct file permissions on generated files
feat(mcp)!: redesign tool registration API
docs: update constitution to v2.1.1
```

## Development Workflow

### Adding a New Feature

1. **Spec First**: Create specification in `/specs/###-feature-name/`
2. **Core Package**: Implement logic in `pkg/<feature>/` with full test coverage
3. **CLI Integration**: Add command in `cmd/platform/` that calls core
4. **MCP Integration**: Register tool in `cmd/platform-mcp/` that calls core
5. **Docker Verification**: Ensure all tests pass in container for both artifacts

### Directory Structure (Go Implementation)

Based on [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

```text
implementations/go/
├── cmd/                        # Main applications
│   ├── platform/               # CLI entry point (main.go)
│   └── platform-mcp/           # MCP server entry point (main.go)
│
├── pkg/                        # Public library code (importable by external projects)
│   ├── github/                 # GitHub Actions generation
│   └── scaffold/               # File scaffolding utilities
│
├── internal/                   # Private code (compiler-enforced)
│   ├── cli/                    # CLI-specific helpers
│   ├── mcp/                    # MCP server and tool registration
│   └── templates/              # Embedded template files (via go:embed)
│
├── test/                       # Additional test apps and test data
│   ├── testdata/               # Test fixtures and golden files
│   └── integration/            # Integration test harnesses
│
├── build/                      # Packaging and CI
│   └── package/                # Dockerfiles
│       ├── platform/           # Dockerfile for CLI
│       │   └── Dockerfile
│       └── platform-mcp/       # Dockerfile for MCP server
│           └── Dockerfile
│
├── scripts/                    # Build automation (PowerShell)
│   ├── Build-All.ps1
│   └── Test-Docker.ps1
│
├── configs/                    # Configuration templates
│   └── default.yaml
│
├── api/                        # API definitions (future: OpenAPI, protobuf)
│
├── .release-please-manifest.json  # Release Please version tracking
├── release-please-config.json     # Release Please configuration
└── go.mod
```

### Directory Purpose Reference

| Directory        | Purpose                                    | golang-standards alignment |
| :--------------- | :----------------------------------------- | :------------------------- |
| `/cmd`           | Main application entry points              | Standard                   |
| `/pkg`           | Public library code, importable externally | Standard                   |
| `/internal`      | Private code, compiler-enforced isolation  | Standard                   |
| `/test`          | External test apps and test data           | Standard                   |
| `/build/package` | Dockerfiles and container configs          | Standard                   |
| `/scripts`       | Build, install, analysis scripts           | Standard                   |
| `/configs`       | Configuration file templates               | Standard                   |
| `/api`           | OpenAPI specs, protocol definitions        | Standard (future)          |

**Note**: Do NOT use `/src` — this is a Java pattern, not idiomatic Go.

## Governance

This constitution supersedes all other development practices for Platform MCP.

**Amendment Process**:

1. Proposed changes MUST be documented with rationale
2. Changes MUST be reviewed and approved before implementation
3. Version number MUST be updated according to semantic versioning:
   - MAJOR: Backward-incompatible principle changes
   - MINOR: New principles or sections added
   - PATCH: Clarifications and wording improvements

**Compliance**:

- All PRs MUST verify compliance with these principles
- Deviations MUST be justified in the Complexity Tracking section of the plan
- Runtime development guidance lives in `AGENTS.md`

**Version**: 2.2.0 | **Ratified**: 2026-01-17 | **Last Amended**: 2026-01-17
