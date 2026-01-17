# Quickstart: Infrastructure MCP Server

## Prerequisites

### TypeScript
- Bun 1.x (latest stable)

### Go
- Go 1.25+

## Install Dependencies

### TypeScript
```bash
cd implementations/typescript
bun install
```

### Go
```bash
cd implementations/go
go mod download
```

## Running

### Mode 1: Standard I/O (Default)
Best for local integration with MCP clients (Claude Desktop, IDEs).

**TypeScript:**
```bash
cd implementations/typescript
bun run dev
# OR explicitly
bun run start -- --transport stdio
```

**Go:**
```bash
cd implementations/go
go run ./cmd/mcp-server --transport stdio
```

### Mode 2: Streamable HTTP
Best for remote access or containerized environments.

**TypeScript:**
```bash
cd implementations/typescript
bun run start -- --transport http --addr :8080
```

**Go:**
```bash
cd implementations/go
go run ./cmd/mcp-server --transport http --addr :8080
```

### Mode 3: Docker Container

**TypeScript:**
```bash
cd implementations/typescript
docker build -t mcp-server-ts .
docker run -p 8080:8080 -e MCP_TRANSPORT=http mcp-server-ts
```

**Go:**
```bash
cd implementations/go
docker build -t mcp-server-go .
docker run -p 8080:8080 -e MCP_TRANSPORT=http mcp-server-go
```

## Verifying

### Health Check (HTTP Mode)

```bash
curl http://localhost:8080/health
# OK
```

### Tool List (Stdio Mode)
Paste this into the running process stdin:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/list",
  "params": {}
}
```

### Streamable HTTP Connection
Use an MCP-compliant client or a stream-capable HTTP client to connect to `http://localhost:8080/mcp`.

## Running Tests

### TypeScript
```bash
cd implementations/typescript
bun test
```

### Go
```bash
cd implementations/go
go test ./...
```

## Type Checking

### TypeScript
```bash
cd implementations/typescript
bun run typecheck
```

### Go
```bash
cd implementations/go
go build ./...
```
