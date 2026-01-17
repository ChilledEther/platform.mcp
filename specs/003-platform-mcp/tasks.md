# Tasks: Platform MCP Server

## Phase 1: Foundation & Contracts
- [x] Research MCP Go SDK and integration points <!-- id: 0 -->
- [x] Define `generate_workflows` tool schema and data model <!-- id: 1 -->
- [x] Create quickstart documentation <!-- id: 2 -->
- [x] Update agent context in `AGENTS.md` <!-- id: 3 -->

## Phase 2: Implementation (TDD)
- [x] Create test harness for MCP tool execution in `internal/mcp` <!-- id: 4 -->
- [x] Implement MCP server initialization in `cmd/platform-mcp/main.go` <!-- id: 5 -->
- [x] Implement `generate_workflows` tool handler in `internal/mcp` <!-- id: 6 -->
- [x] Verify tool execution with table-driven tests <!-- id: 7 -->

## Phase 3: Integration & Packaging
- [x] Configure `stdio` transport for the server <!-- id: 8 -->
- [x] Add `platform-mcp` Dockerfile in `build/package/platform-mcp/` <!-- id: 9 -->
- [x] Verify Docker build and execution <!-- id: 10 -->
- [x] Update project README with MCP server details <!-- id: 11 -->

## Phase 4: Final Validation
- [x] Perform manual end-to-end test with Claude Desktop <!-- id: 12 -->
- [x] Verify compliance with Constitution (especially naming and logic location) <!-- id: 13 -->
