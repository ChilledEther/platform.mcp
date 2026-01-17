# Tasks: Platform MCP Server

## Phase 1: Foundation & Contracts
- [x] Research MCP Go SDK and integration points <!-- id: 0 -->
- [x] Define `generate_workflows` tool schema and data model <!-- id: 1 -->
- [x] Create quickstart documentation <!-- id: 2 -->
- [x] Update agent context in `AGENTS.md` <!-- id: 3 -->

## Phase 2: Implementation (TDD)
- [ ] Create test harness for MCP tool execution in `internal/mcp` <!-- id: 4 -->
- [ ] Implement MCP server initialization in `cmd/platform-mcp/main.go` <!-- id: 5 -->
- [ ] Implement `generate_workflows` tool handler in `internal/mcp` <!-- id: 6 -->
- [ ] Verify tool execution with table-driven tests <!-- id: 7 -->

## Phase 3: Integration & Packaging
- [ ] Configure `stdio` transport for the server <!-- id: 8 -->
- [ ] Add `platform-mcp` Dockerfile in `build/package/platform-mcp/` <!-- id: 9 -->
- [ ] Verify Docker build and execution <!-- id: 10 -->
- [ ] Update project README with MCP server details <!-- id: 11 -->

## Phase 4: Final Validation
- [ ] Perform manual end-to-end test with Claude Desktop <!-- id: 12 -->
- [ ] Verify compliance with Constitution (especially naming and logic location) <!-- id: 13 -->
