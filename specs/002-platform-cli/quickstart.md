# Quickstart: Platform CLI Tool

## Installation

```bash
# Clone the repository
git clone https://github.com/modelcontextprotocol/platform.mcp.git
cd platform.mcp

# Build the CLI
go build -o platform cmd/platform/main.go

# Move to your path
mv platform /usr/local/bin/
```

## Basic Usage

### Generate Default Go Workflow
```bash
platform generate workflows
```

### Generate Dockerfile and TypeScript Workflow
```bash
platform generate workflows --docker --workflow-type typescript
```

### Preview Changes (Dry Run)
```bash
platform generate workflows --dry-run
```

### Force Overwrite
```bash
platform generate workflows --force
```

## Troubleshooting

### "File already exists" error
The CLI will prompt you before overwriting files. Use `-f` or `--force` to bypass this.

### Missing templates
Ensure the CLI was built correctly with `go:embed` support (standard in Go 1.16+).
