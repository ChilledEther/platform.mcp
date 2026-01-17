# Tasks: Infrastructure MCP Foundation

**Input**: Design documents from `/specs/001-infra-mcp-foundation/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md

**Status**: ✅ Complete (migrated to TypeScript/Bun)

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to

## Path Conventions

- **TypeScript Source**: `implementations/typescript/` (server/, tools/, utils/)
- **TypeScript Tests**: `implementations/typescript/tests/`
- **Go Source**: `implementations/go/` (cmd/, internal/)
- **Go Tests**: Alongside source files (`*_test.go`)

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure.

- [x] T001 Initialize TypeScript project with Bun (`package.json`, `tsconfig.json`)
- [x] T002 Install dependencies (`@modelcontextprotocol/sdk`, `zod`)
- [x] T003 [P] Configure TypeScript strict mode in `tsconfig.json`
- [x] T025 Initialize Go project (`go.mod`) with Go 1.25 requirement

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story.

- [x] T004 [P] Define `ToolHandler` type in `tools/registry.ts` (TS) and `internal/tools/registry.go` (Go)
- [x] T005 [P] Implement `Registry` class with tool registration in `tools/registry.ts` (TS) and `internal/tools/registry.go` (Go)
- [x] T006 [P] Implement `createConfig` function for flag/env parsing in `server/config.ts` (TS) and `internal/server/config.go` (Go)

**Checkpoint**: Foundation ready - user story implementation can begin.

---

## Phase 3: User Story 1 - Stdio Mode (Priority: P1) 🎯 MVP

**Goal**: Implement basic Stdio transport for local shell integration.

- [x] T007 [US1] Implement `MCPServer` class wrapping SDK server in `server/index.ts` (TS) and `internal/server/server.go` (Go)
- [x] T008 [US1] Implement `runStdio` method using `StdioServerTransport` in `server/index.ts` (TS) and `internal/server/server.go` (Go)
- [x] T009 [US1] Create entry point in `index.ts` (TS) and `cmd/mcp-server/main.go` (Go)
- [x] T010 [US1] Add graceful shutdown (SIGINT/SIGTERM handling)
- [x] T024 [US1] Implement logger utility to enforce stderr output for all logs to protect stdout transport


**Checkpoint**: User Story 1 functional - server runs in stdio mode.

---

## Phase 4: User Story 2 - HTTP Mode (Priority: P2)

**Goal**: Implement Streamable HTTP transport for containerized deployment.

- [x] T011 [US2] Implement `runHTTP` method using `StreamableHTTPServerTransport`
- [x] T012 [US2] Add `--transport` and `--addr` flag handling
- [x] T013 [US2] Create multi-stage `Dockerfile` using `bun build --compile` (TS) and static Go binary + Alpine to minimize size (<160MB)

**Checkpoint**: User Story 2 functional - server runs in HTTP mode.

---

## Phase 5: User Story 3 - Tool Registration API (Priority: P2)

**Goal**: Expose type-safe API for registering custom tools.

- [x] T014 [US3] Implement `register` method in `Registry` class
- [x] T015 [US3] Export `Registry` and types

**Checkpoint**: User Story 3 functional - tools can be registered.

---

## Phase 6: User Story 4 - Health Check Tool (Priority: P3)

**Goal**: Provide built-in health check tool for monitoring.

- [x] T016 [US4] Implement `registerHealthTool` function
- [x] T017 [US4] Register `health_check` tool in server initialization

**Checkpoint**: User Story 4 functional - health check available.

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Final cleanup and documentation.

- [x] T018 [P] Create unit tests for Registry
- [x] T019 [P] Create unit tests for health tool
- [x] T020 Run build and test suites for both implementations
- [x] T021 Update `README.md` with usage instructions
- [x] T022 Verify server startup time is < 100ms
- [x] T023 Limit container memory to 128MB in Dockerfile
- [x] T026 [P] Implement path-based CI/CD triggers to optimize resource usage
- [x] T027 [P] Update `Add-MCPToCatalog.ps1` script to support selecting between Go (`-Implementation go`) and TypeScript (`-Implementation typescript`) images while maintaining a single catalog entry name `mcp-github-agentic` for easy toggling.

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies
- **Foundational (Phase 2)**: Depends on Setup
- **User Stories (Phases 3-6)**: All depend on Foundational
  - US1 (Stdio) → US2 (HTTP) → US3 (Registration) → US4 (Health)
- **Polish (Phase 7)**: Depends on all user stories

### Parallel Opportunities

- T004, T005, T006 are independent (Foundational)
- T018 and T019 are independent (Tests)

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Setup + Foundational
2. Complete User Story 1 (Stdio)
3. **STOP and VALIDATE**: Test server in stdio mode
4. Deploy/demo if ready

---

## Notes

- All TypeScript paths reference `implementations/typescript/` per Constitution Principle I
- All Go paths reference `implementations/go/` per Constitution Principle Ib
- Tests use `bun:test` (TypeScript) and `go test` (Go) per Constitution Principle XI
- Both implementations must maintain feature parity
