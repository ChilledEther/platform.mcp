# Infrastructure MCP Server 🛠️

An MCP server for infrastructure provisioning tools, built with **TypeScript** and **Bun**. This server bridges the gap between AI assistants and your infrastructure, allowing safe and structured interaction with configuration files like Azure Firewall rules.

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/ChilledEther/mcp.github.agentic/docker-build-publish.yaml?branch=main&label=Build%20%26%20Publish)

## ✨ Features

- **Standard MCP Protocol**: Full implementation of Model Context Protocol (Stdio & HTTP).
- **Dual Transport**: Run locally via stdio or deploy remotely via HTTP.
- **Type-Safe Tools**: Tools defined with rigorous Zod schemas for validation.
- **Resource Support**: Exposes schemas and configs as MCP resources.
- **Container Ready**: Optimized ~150MB Alpine-based Docker image.
- **Health Monitoring**: Built-in health check tool.

---

## 🚀 Quick Start

### Prerequisites

- [Bun](https://bun.sh/) 1.x or later
- Docker (optional, for container usage)

### Installation

```bash
bun install
```

### Running Locally

You can run the server in different modes depending on your needs:

```bash
# Stdio mode (Default - for local MCP clients)
bun run start

# HTTP mode (For remote connections)
bun run start -- --transport http --addr :8080

# Development mode (Watch for changes)
bun run dev
```

---

## 🐳 Docker Support

We provide a highly optimized, production-ready Docker image.

### Build Locally

Use our handy PowerShell script to build the image (requires Docker):

```powershell
./scripts/Invoke-DockerBuild.ps1
```

Or manually:

```bash
docker build -t mcp-github-agentic:latest .
```

### Run Container

```bash
# Run in Stdio mode
docker run -i --rm mcp-github-agentic:latest

# Run in HTTP mode
docker run -p 8080:8080 -e MCP_TRANSPORT=http mcp-github-agentic:latest
```

---

## 🛠️ Usage

This server exposes tools that your MCP client (Claude, Cursor, etc.) can discover and use.

### Available Tools

#### `health_check`
Verifies the server is operational.
- **Input**: `{}`
- **Output**: `"status": "healthy"`

#### `new_azure_firewall_rules`
Adds a new firewall whitelist rule to the YAML configuration.
- **Parameters**:
  - `team` (string, required): Team requesting the rule.
  - `source` (string, required): Valid IP/CIDR.
  - `destination` (string, required): Valid IP/CIDR.
  - `port` (number, default 443): Destination port.
  - `repo_path` (string, optional): Custom path for the rules file.

### Resources

#### `mcp://azure-firewall/schema`
Returns the JSON schema defining the required structure for firewall rules.

---

## 🔌 Client Configuration

Here is how to add this server to your favorite MCP client.

### Claude Desktop / Cursor

Add this to your configuration file (e.g., `claude_desktop_config.json` or `.cursor/mcp.json`):

**Local (Source)**:
```json
{
  "mcpServers": {
    "infra-mcp": {
      "command": "bun",
      "args": ["/absolute/path/to/mcp.github.agentic/src/index.ts"]
    }
  }
}
```

**Docker**:
```json
{
  "mcpServers": {
    "infra-mcp": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-v", "/path/to/your/configs:/data",
        "-e", "AZURE_FIREWALL_PATH=/data/rules.yaml",
        "mcp-github-agentic:latest"
      ]
    }
  }
}
```

---

## ⚙️ Configuration

Configure the server via environment variables or CLI flags:

| Env Variable | CLI Flag | Default | Description |
| :--- | :--- | :--- | :--- |
| `MCP_TRANSPORT` | `--transport` | `stdio` | Transport mode (`stdio` or `http`) |
| `MCP_ADDR` | `--addr` | `:8080` | HTTP server bind address |
| `AZURE_FIREWALL_PATH` | `--firewall-path` | `./azure-firewall-rules.yaml` | Path to the firewall rules file |

---

## 🏗️ Project Structure

```text
src/
├── index.ts              # Entry point
├── server/               # MCP Server core logic
├── tools/                # Tool definitions & registry
└── utils/                # Shared utilities
tests/                    # Bun test suite
legacy/                   # Historical Go implementation
.github/workflows/        # CI/CD pipelines
```

## 🧪 Development

```bash
# Run tests
bun test

# Type checking
bun run typecheck
```
