# Research: Infrastructure MCP Foundation

**Feature**: Infrastructure MCP Foundation (Dual Transport)
**Date**: 2026-01-15
**Updated**: 2026-01-16 (Migrated to TypeScript/Bun)
**Status**: Complete

## 1. Dual Transport Architecture

### Decision
We use a modular transport layer where `src/index.ts` parses configuration (flags/env vars) and initializes the appropriate transport (Stdio or Streamable HTTP) before passing it to the core MCP server instance.

### Rationale
- **Separation of Concerns**: The core server logic (tool handling, protocol processing) remains transport-agnostic.
- **Flexibility**: Allows adding other transports in the future without changing core logic.
- **SDK Compliance**: The `@modelcontextprotocol/sdk` provides `StdioServerTransport` and `StreamableHTTPServerTransport`.

### Implementation Detail
- **Stdio**: Uses `StdioServerTransport` from SDK.
- **Streamable HTTP**: Uses `StreamableHTTPServerTransport` with Bun's native HTTP server.
- **Flags**:
  - `--transport`: `stdio` (default) or `http`.
  - `--addr`: `:8080` (default for http).

## 2. Tool Registration Pattern

### Decision
We use a `Registry` class that wraps the MCP SDK server and provides a type-safe `register` method using Zod schemas.

### Rationale
- **Type Safety**: TypeScript's type system combined with Zod provides compile-time and runtime validation.
- **Schema Generation**: Zod schemas can be converted to JSON Schema for the `tools/list` response.
- **Simplicity**: Direct function handlers avoid complex reflection patterns.

### Alternatives Considered
- **Manual JSON Schema**: Writing raw JSON strings for schemas. *Rejected*: Error-prone and hard to maintain.
- **Code Generation**: Generating tool wrappers from a Spec file. *Rejected*: Too complex for the initial foundation.

## 3. Signal Handling

### Decision
Use `process.on('SIGINT')` and `process.on('SIGTERM')` for graceful shutdown.

### Rationale
- **Graceful Shutdown**: Essential for HTTP mode to close listeners and for Stdio mode to flush buffers.
- **Async Pattern**: Use Promise-based shutdown that resolves when cleanup is complete.
