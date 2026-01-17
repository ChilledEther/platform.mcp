# Architecture 🏛️

Platform MCP follows a **Core-First Architecture** to ensure consistency across different interfaces.

## 🏗️ Design Principles

1. **Shared Core (`pkg/`)**: All business logic, data models, and generation logic reside in shared packages. These packages are pure logic and return data, avoiding direct side effects like disk or network I/O.
2. **Consumer Applications (`cmd/`)**:
   - **CLI (`platform`)**: Handles user interaction and local file I/O.
   - **MCP Server (`platform-mcp`)**: Handles MCP protocol communication and returns structured content to agents.
3. **Internal Packages (`internal/`)**: Private code such as templates and helpers that are not intended for external use.

## 🗺️ Project Structure

```text
├── cmd/                        # Main applications
│   ├── platform/               # CLI entry point
│   └── platform-mcp/           # MCP server entry point
│
├── pkg/                        # Public library code
│   ├── github/                 # GitHub Actions logic
│   └── scaffold/               # Scaffolding utilities
│
├── internal/                   # Private code
│   └── templates/              # Embedded templates
│
├── specs/                      # Feature specifications (Source of Truth)
├── docs/                       # Human-oriented documentation
├── scripts/                    # Automation and build scripts
└── build/package/              # Dockerfiles
```

## 🔄 Data Flow

1. **User/Agent** provides input parameters.
2. **Consumer App** validates parameters and calls **Core Library**.
3. **Core Library** processes logic and returns **Structured Data**.
4. **Consumer App** performs I/O (Write to disk or Send over network).
