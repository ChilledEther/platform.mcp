# Quickstart: Platform MCP

## Prerequisites
- Go 1.25+
- An MCP client (e.g., [Claude Desktop](https://claude.ai/download), [Cursor](https://cursor.com/))

## Running the Server Locally

```bash
go run cmd/platform-mcp/main.go
```
Note: The server uses stdio transport, so it will wait for JSON-RPC input on stdin.

## Configuring Claude Desktop

Add the following to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "platform": {
      "command": "go",
      "args": ["run", "/path/to/platform.mcp/cmd/platform-mcp/main.go"]
    }
  }
}
```

## Example Tool Call

If you are using an agent, you can ask:
> "Generate a Go GitHub Actions workflow for a project named 'my-cool-app' using the platform tool."

The agent will call `generate_workflows(project_name="my-cool-app", workflow_type="go")`.

## Development

### Running Tests
```bash
go test ./internal/mcp/...
```

### Building Docker Image
```bash
./scripts/Invoke-DockerBuild.ps1 -Artifact platform-mcp
```
