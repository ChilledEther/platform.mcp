# Research: Platform MCP Server

## Goal
Implement an MCP server (`platform-mcp`) that provides the `generate_workflows` tool, using the shared logic in `pkg/scaffold`.

## MCP Go SDK Analysis

The `github.com/modelcontextprotocol/go-sdk` provides the necessary primitives for building MCP servers in Go.

### Server Lifecycle
1.  Initialize a new server: `server.NewServer("platform-mcp", "1.0.0")`.
2.  Register tools using `server.RegisterTool()`.
3.  Start the server with a transport (e.g., `stdio.NewServerTransport()`).

### Tool Registration
```go
server.RegisterTool(
    "generate_workflows",
    "Generates GitHub Actions workflow content for a project",
    func(ctx context.Context, args struct {
        ProjectName  string `json:"project_name" jsonschema:"description=The name of the project"`
        UseDocker    bool   `json:"use_docker" jsonschema:"description=Whether to include Docker-related workflows"`
        WorkflowType string `json:"workflow_type" jsonschema:"description=The type of workflow to generate (go, typescript, python),default=go"`
    }) (*mcp.CallToolResult, error) {
        // Logic here
    },
)
```

## Tool Contract: `generate_workflows`

### Input Parameters
| Parameter | Type | Description | Mapping to `scaffold.Config` |
| :--- | :--- | :--- | :--- |
| `project_name` | string | Name of the project | `ProjectName` |
| `use_docker` | boolean | Include Docker logic | `UseDocker` |
| `workflow_type` | string | Type of workflow | `WorkflowType` |

### Return Value
MCP tools return `CallToolResult`, which contains a slice of `Content` objects.
For file generation, we should return a list of files as a single text block or multiple text blocks.

**Proposal**: Return a JSON string containing the list of files, or a formatted text block. Since AI agents prefer structured data or clear text, we can return:
- A text block per file, or
- A single text block containing all files in a "pseudo-filesystem" format (e.g., `--- path/to/file ---`).

Actually, the spec says "response contains file paths and content as structured data".
We can use the `TextContent` type for each file.

## Integration Strategy

1.  **Transport**: Use `stdio` as the primary transport for local agent usage (Claude Desktop).
2.  **Logic**: Call `scaffold.Generate(cfg)` from the tool handler.
3.  **Error Handling**: Convert `pkg/scaffold` errors into MCP protocol errors.

## References
- [MCP Specification](https://modelcontextprotocol.io/)
- `pkg/scaffold/types.go`
- `pkg/scaffold/scaffold.go`
- `001-core-foundation` research and plan.
