# Platform MCP 🛠️

A powerful Model Context Protocol (MCP) platform built with **Go**. This repository provides both a CLI tool and an MCP server, sharing a common core for consistent behavior.

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/ChilledEther/platform.mcp/docker-build-publish.yaml?branch=main&label=Build%20%26%20Publish)

## ✨ Features

Platform MCP is a comprehensive toolset for infrastructure scaffolding:

- **[Core Foundation](./docs/features.md#core-foundation-001)**: Shared logic for reliable file generation.
- **[Platform CLI](./docs/features.md#platform-cli-002)**: Cobra-based CLI for local development.
- **[MCP Server](./docs/features.md#platform-mcp-server-003)**: AI-native interface for agentic workflows.
- **[Dockerized](./docs/features.md#docker-environment-998)**: Minimal, multi-stage Alpine images.
- **[CI/CD Pipeline](./docs/features.md#deployment-pipeline-999)**: Automated tests, releases, and publishing.

---

## 📚 Documentation

- **[Architecture](./docs/architecture.md)**: Deep dive into the project structure and design principles.
- **[Feature List](./docs/features.md)**: Detailed breakdown of current capabilities.
- **[Development Guide](./docs/development.md)**: Workflow, scripts, and testing instructions.
- **[Planning](./.planning/)**: Project roadmap and execution plans.

---

## 🚀 Quick Start

### Prerequisites

- [Go](https://go.dev/) 1.25+
- Docker (optional, for container usage)

### Running the CLI

```bash
go run cmd/platform/main.go
```

### Running the MCP Server

The MCP server uses `stdio` transport. You can run it directly:

```bash
go run cmd/platform-mcp/main.go
```

To use it with Claude Desktop, add the following to your configuration:

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

## 🛠️ MCP Tools

### `generate_workflows`
Generates GitHub Actions workflows and Dockerfiles based on project parameters.

- **Parameters**:
  - `project_name` (string, required): The name of the project.
  - `workflow_type` (string, optional): One of `go`, `typescript`, `python`. Default is `go`.
  - `use_docker` (boolean, optional): Whether to generate a Dockerfile. Default is `false`.

---

## 🐳 Docker Support

### Build Locally

Use the PowerShell script to build the image:

```powershell
./scripts/Invoke-DockerBuild.ps1
```

Or manually:

```bash
docker build -t platform-mcp:latest -f build/package/platform-mcp/Dockerfile .
```

---

## 🏗️ Project Structure

Based on [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

```text
├── cmd/                        # Main applications (CLI & MCP)
├── pkg/                        # Shared logic and libraries
├── internal/                   # Private helpers and templates
├── .planning/                  # Project roadmap and plans
├── scripts/                    # Automation scripts
└── build/package/              # Dockerfiles
```

---

## 🧪 Development

```bash
# Run all tests
go test ./...

# Run linting
golangci-lint run
```
