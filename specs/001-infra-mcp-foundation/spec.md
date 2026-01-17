# Feature Specification: Infrastructure MCP Foundation

**Feature Branch**: `001-infra-mcp-foundation`  
**Created**: 2026-01-15  
**Status**: Draft  
**Input**: User description: "This is going to be a TypeScript/Bun mcp server that will expose tools for infrastructure provisioning. This is in the form of yaml files. I want to ensure we start with a strong foundation (initial spec) and later add the actual tools. We should use stdio for the mcp server. An idiomatic way to register tools and a strong tool interface. Include the basic health endpoints to prove this operates as expected. We'll create specs later for tool groups."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Core Server Connectivity (Priority: P1)

As an MCP Client (like Claude Desktop or an IDE), I want to connect to the server via standard input/output (stdio) so that I can discover its capabilities locally.

**Why this priority**: Without a running server and transport layer, no other functionality is possible.

**Independent Test**: Can be tested by running `bun run dev` (TypeScript) or `go run ./cmd/mcp-server` (Go) and sending a JSON-RPC initialization request via stdin, verifying the correct initialization response on stdout.

**Acceptance Scenarios**:

1. **Given** the server is started with the appropriate command for the implementation, **When** executed with no arguments (or explicit `--transport stdio`), **Then** it listens on stdin/stdout.
2. **Given** the server is running in stdio mode, **When** a valid JSON-RPC `initialize` request is sent, **Then** it returns the server capabilities including tool support.

---

### User Story 2 - Remote/Container Connectivity (Priority: P1)

As a Cloud Platform or Docker user, I want to connect to the server via **Streamable HTTP** so that I can access its tools remotely or from within a containerized environment.

**Why this priority**: Essential for the mandated "Cloud-Native" support and container usage.

**Independent Test**: Can be tested by running `curl` or an HTTP client against the `/mcp` endpoint.

**Acceptance Scenarios**:

1. **Given** the server is executed with `--transport http` (or `MCP_TRANSPORT=http`), **Then** it starts an HTTP server on the default or specified address.
2. **Given** the server is running in HTTP mode, **When** a client connects to the `/mcp` endpoint, **Then** the connection is established and events can be streamed.
3. **Given** the server is running, **When** the `/health` endpoint is requested, **Then** it returns a 200 OK status.

---

### User Story 3 - Tool Registration & Discovery (Priority: P2)

As a Developer, I want a structured way to register tools in the codebase, and as an MCP Client, I want to list these available tools so I know what operations can be performed.

**Why this priority**: Establish the architectural pattern for adding the future infrastructure tools (YAML generators) in a clean, idiomatic way.

**Independent Test**: Can be tested by implementing a dummy tool and verifying it appears in the `tools/list` response.

**Acceptance Scenarios**:

1. **Given** the server has registered tools, **When** a `tools/list` request is sent, **Then** the server returns a list of tools with names, descriptions, and input schemas.
2. **Given** a new tool definition is added to the code registry, **When** the server is restarted and queried, **Then** the new tool appears in the list.

---

### User Story 4 - Health Check Tool Execution (Priority: P3)

As an MCP Client, I want to execute a basic "health check" tool so that I can verify the server is operational and responding to tool execution requests correctly.

**Why this priority**: Proves the end-to-end execution flow (Client -> Server -> Tool Logic -> Server -> Client) works before complex logic is added.

**Independent Test**: Can be tested by sending a `tools/call` request for the "health_check" tool and asserting the result is positive.

**Acceptance Scenarios**:

1. **Given** the server is running, **When** a `tools/call` request is sent for `health_check`, **Then** it returns a success message confirming system status.
2. **Given** the server is running, **When** a `tools/call` request is sent for a non-existent tool, **Then** it returns a "tool not found" error.

### Edge Cases

- What happens when a tool panics during execution? (Should be recovered and return JSON-RPC error)
- How does the system handle concurrent tool requests? (Should handle safely without blocking other operations)
- What happens if the input schema for a tool call is invalid? (Should return validation error)
- What happens if the address/port is already in use (HTTP mode)? (Should exit with error)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST implement an MCP-compliant server in both TypeScript (Bun runtime) and Go, with feature parity between implementations.
- **FR-002**: The server MUST support **Dual Transport Modes**: Standard Input/Output (stdio) AND **Streamable HTTP**.
- **FR-003**: The server MUST default to `stdio` mode if no configuration is provided.
- **FR-004**: The server MUST allow configuration via CLI flags (`--transport`, `--addr`) AND Environment Variables (`MCP_TRANSPORT`, `MCP_ADDR`).
- **FR-005**: The server MUST implement the `tools/list` capability.
- **FR-006**: The server MUST implement the `tools/call` capability.
- **FR-007**: The system MUST provide a Type-Safe and Extensible interface for defining tools that prevents common registration errors.
- **FR-008**: The system MUST include a built-in `health_check` tool that requires no arguments and returns a healthy status.
- **FR-009**: The server MUST generate JSON schemas for tool inputs automatically or strictly enforce them.
- **FR-010**: The server MUST handle OS signals (SIGINT, SIGTERM) for graceful shutdown in both modes.
- **FR-011**: The system MUST be optimized for minimal memory usage, strictly staying within a **128 MiB** limit in containerized environments.
- **FR-012**: The system MUST explicitly direct all application logging (info, warn, debug) to **Standard Error (stderr)** to prevent corruption of the Standard Output (stdout) JSON-RPC stream.
- **FR-013**: The production container image size MUST be minimized (target < 160MB) by compiling to a single binary and stripping development dependencies.
- **FR-014**: The system MUST implement an optimized CI/CD pipeline that only triggers builds based on relevant file path changes to minimize resource consumption.
- **FR-015**: The system MUST provide a catalog registration script (`Add-MCPToCatalog.ps1`) that allows users to toggle the active MCP server between Go and TypeScript implementations within their local Docker desktop environment, reusing the same catalog entry name (`mcp-github-agentic`) for seamless switching.

### Key Entities

- **ToolRegistry**: A central component responsible for holding references to all available tools.
- **ToolDefinition**: A structure defining a tool's metadata (name, description) and input schema.
- **ToolHandler**: The function or method signature for executing a tool's logic.
- **ServerConfig**: Configuration struct holding transport mode, address, and other runtime settings.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Server executes successfully with `bun run dev` (TypeScript) or `go run ./cmd/mcp-server` (Go).
- **SC-002**: Server responds to `initialize` request within 100ms (Stdio).
- **SC-003**: Server successfully binds to HTTP port when configured and responds to **Streamable HTTP** connection requests on `/mcp`.
- **SC-004**: `tools/list` request returns valid JSON payload containing at least the `health_check` tool.
- **SC-005**: `health_check` tool executes and returns success in under 50ms.
- **SC-006**: 100% of defined tools have associated JSON schemas in the list response.
