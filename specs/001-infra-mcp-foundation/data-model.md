# Data Model: Infrastructure MCP Foundation

Since this is a stateless server, the "Data Model" refers to the internal TypeScript structures used to represent tools and configuration.

## Core Entities

### 1. ToolDefinition
Represents the metadata and schema of an exposed capability.

```typescript
interface ToolDefinition {
  name: string;
  description: string;
  schema: z.ZodSchema;  // Zod schema for input validation
}
```

### 2. ToolHandler
The executable logic for a tool.

```typescript
// ToolHandler takes validated input and context, returning a result
type ToolHandler<T> = (input: T, context: ToolContext) => Promise<ToolResult>;

interface ToolResult {
  content: Array<{ type: 'text'; text: string }>;
  isError?: boolean;
}

interface ToolContext {
  executionId: string;
}
```

### 3. RegisteredTool
Internal wrapper combining definition and implementation.

```typescript
interface RegisteredTool<T> {
  name: string;
  description: string;
  schema: z.ZodSchema<T>;
  handler: ToolHandler<T>;
}
```

### 4. ServerConfig
Runtime configuration.

```typescript
type TransportMode = 'stdio' | 'http';

interface ServerConfig {
  transport: TransportMode;
  addr: string;  // e.g. ":8080"
}
```

## Relationships

- `Registry` contains N `RegisteredTool`s.
- `MCPServer` uses 1 `Registry`.
- `MCPServer` runs on 1 `Transport` (configured by `ServerConfig`).
