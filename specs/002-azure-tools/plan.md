# Implementation Plan: Azure Tools Group

**Branch**: `002-azure-tools` | **Date**: 2026-01-15 | **Updated**: 2026-01-16 | **Spec**: [specs/002-azure-tools/spec.md](spec.md)

## Summary

This feature implements a **collection of Azure-related MCP tools**. Each tool may follow a different implementation pattern based on its requirements:

| Pattern | Description | Tools Using This Pattern |
|---------|-------------|-------------------------|
| **Content-Return** | Tool processes data, returns content for agent to persist | `new_azure_firewall_rules` |
| **REST Passthrough** | Tool makes REST calls to Azure APIs | *(future tools)* |
| **Hybrid** | Combines local processing with remote API calls | *(future tools)* |

This is a **living document** - new tools will be added as sections below.

---

# Tool: `new_azure_firewall_rules`

## Architecture: Content-Return Pattern

**Why this pattern?** The MCP server runs in a container with an isolated filesystem. The agent has direct access to the user's repository, so the tool returns content for the agent to write.

### Flow

1. Agent reads existing `azure-firewall-rules.yaml` from repository (if exists)
2. Agent calls tool with rule parameters + existing YAML content
3. Tool validates, checks duplicates, merges rule
4. Tool returns updated YAML content (does NOT write to filesystem)
5. Agent writes returned content to `azure-firewall-rules.yaml` at repository root

## Technical Context

**Languages/Versions**: 
- TypeScript (Bun runtime)
- Go 1.25+

**Primary Dependencies**: 
- TypeScript: `@modelcontextprotocol/sdk`, `yaml`, `zod`
- Go: `github.com/modelcontextprotocol/go-sdk`, `gopkg.in/yaml.v3`

**Storage**: YAML content is returned to the agent; agent handles file persistence  
**Performance Goals**: <100ms for YAML processing and response  
**Constraints**: Valid IPv4/IPv6/CIDR input, flat YAML structure, no duplicates

## Constitution Check

- [x] **V. Default Read-Only**: **COMPLIANT** - Tool returns content only, no filesystem writes.

## Source Code Location

### TypeScript
```text
implementations/typescript/tools/azure-firewall.ts   # Tool implementation
implementations/typescript/utils/validation.ts       # IP/CIDR validation
```

### Go
```text
implementations/go/internal/tools/azure_firewall.go  # Tool implementation
implementations/go/internal/utils/validation.go      # IP/CIDR validation
```

## Key Implementation Changes (from original)

### Remove Filesystem Operations

The tool handler must be refactored to:

1. **Remove** `readFile` and `writeFile` imports
2. **Remove** `existsSync` checks
3. **Add** `existing_yaml` input parameter
4. **Return** structured response with `yaml_content`, `filename`, `action`, `message`

### Update Tool Schema

```typescript
export const FirewallRuleSchema = z.object({
  team: z.string().describe('The name of the team requesting the rule.'),
  source: z.string().describe('The source IP address or CIDR block.'),
  destination: z.string().describe('The destination IP address or CIDR block.'),
  port: z.number().int().min(1).max(65535).optional().default(443),
  existing_yaml: z.string().optional().default('').describe(
    'Current content of azure-firewall-rules.yaml. Pass empty string if file does not exist.'
  ),
});
```

### Update Tool Description

The tool description must instruct the agent on:
- Repository root determination (ask user if ambiguous)
- Reading existing YAML file
- Passing content as `existing_yaml`
- Writing returned content to repository root

---

# Tool: *(Future Tool Placeholder)*

> Implementation details for additional Azure tools will be documented here.
> 
> Example future tools:
> - Azure Resource Group management (REST Passthrough pattern)
> - Azure Key Vault secret retrieval (REST Passthrough pattern)

---

# Group-Level Technical Context

**Testing**: `bun:test` (Unit & Integration)  
**Target Platform**: Linux/macOS/Windows (MCP Server in container, Agent local)  
**Scale/Scope**: Growing collection of Azure-related tools

## Project Structure

### Documentation

```text
specs/002-azure-tools/
├── plan.md              # This file
├── research.md          # Research decisions
├── data-model.md        # Data structures per tool
├── quickstart.md        # Usage examples
├── contracts/           # JSON schemas
│   └── azure_firewall.json
└── tasks.md             # Implementation tasks
```

### Source Code

```text
implementations/
├── typescript/
│   ├── server/
│   │   └── index.ts            # Resource and transport registration
│   ├── tools/
│   │   ├── azure-firewall.ts   # Firewall rules tool
│   │   ├── azure-*.ts          # Future Azure tools
│   │   └── registry.ts         # Tool registry
│   └── utils/
│       └── validation.ts       # Shared validation utilities
│
└── go/
    └── internal/
        ├── server/
        │   └── server.go       # Resource and transport registration
        ├── tools/
        │   ├── azure_firewall.go  # Firewall rules tool
        │   └── registry.go        # Tool registry
        └── utils/
            └── validation.go      # Shared validation utilities
```

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| None | Content-return pattern is simpler than filesystem writes | N/A |
