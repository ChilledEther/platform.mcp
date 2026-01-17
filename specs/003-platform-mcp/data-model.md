# Data Model: Platform MCP

## Tool Definitions

### `generate_workflows`
The primary tool exposed by the MCP server.

**Schema (JSON Schema)**:
```json
{
  "type": "object",
  "properties": {
    "project_name": {
      "type": "string",
      "description": "The name of the project for which to generate workflows."
    },
    "use_docker": {
      "type": "boolean",
      "description": "Whether to include Docker-related workflow steps.",
      "default": false
    },
    "workflow_type": {
      "type": "string",
      "description": "The type of workflow to generate (go, typescript, python).",
      "enum": ["go", "typescript", "python"],
      "default": "go"
    }
  },
  "required": ["project_name"]
}
```

## Internal Mapping

| MCP Parameter | Go Config Field | Type |
| :--- | :--- | :--- |
| `project_name` | `ProjectName` | `string` |
| `use_docker` | `UseDocker` | `bool` |
| `workflow_type` | `WorkflowType` | `string` |

## Response Model

The tool returns a list of files. Each file is represented as an MCP `TextContent` item.

**MCP Tool Result Structure**:
```go
mcp.CallToolResult{
    Content: []mcp.Content{
        mcp.TextContent{
            Type: "text",
            Text: "--- FILE: .github/workflows/go.yaml ---\n[CONTENT HERE]",
        },
        // ... more files
    },
}
```

Alternatively, we can return a single `TextContent` with all files formatted clearly, which is often easier for agents to process in one go.

**Decision**: Return individual `TextContent` items for each file. This allows the agent to distinguish between different files more easily if the client supports it, or it will just concatenate them.

## Error Model

| Error Case | MCP Error Code | Message Example |
| :--- | :--- | :--- |
| Missing `project_name` | `-32602` (Invalid Params) | "parameter 'project_name' is required" |
| Invalid `workflow_type` | `-32602` (Invalid Params) | "invalid workflow type: 'rust'" |
| Generation Failure | `-32000` (Internal Error) | "failed to generate scaffold: template not found" |
