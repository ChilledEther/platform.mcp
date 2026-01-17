# Tasks: Azure Tools Group

**Input**: Design documents from `/specs/002-azure-tools/`
**Prerequisites**: plan.md (required), spec.md (required), data-model.md, contracts/

**Architecture**: Content-Return Pattern - Tool returns YAML content, agent writes to repository.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **TypeScript Source**: `implementations/typescript/` (server/, tools/, utils/)
- **TypeScript Tests**: `implementations/typescript/tests/`
- **Go Source**: `implementations/go/internal/` (server/, tools/, utils/)
- **Go Tests**: Alongside source files (`*_test.go`)

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Verify project structure and dependencies are in place.

- [x] T001 Verify TypeScript project structure exists per plan.md in `src/`
- [x] T002 Verify `yaml` package is installed in `package.json`
- [x] T003 Verify `zod` package is installed in `package.json`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core validation utilities that all tools depend on.

**⚠️ CRITICAL**: No user story work can begin until this phase is complete.

- [x] T004 [P] Implement `isValidIPOrCIDR` helper function in `src/utils/validation.ts`
- [x] T005 [P] Create unit tests for IP/CIDR validation in `tests/utils/validation.test.ts`

**Checkpoint**: Foundation ready - tool implementation can now begin.

---

## Phase 3: User Story 1 - Create Initial Firewall Rules (Priority: P1) 🎯 MVP

**Goal**: Implement the core tool that creates initial YAML content when no file exists.

**Independent Test**: Call tool with empty `existing_yaml` and verify returned YAML contains correct single entry with `action: created`.

### Implementation for User Story 1

- [x] T006 [P] [US1] Define `FirewallRuleSchema` with Zod validation in `src/tools/azure-firewall.ts`
- [x] T007 [P] [US1] Define `FirewallRule` and `FirewallConfig` interfaces in `src/tools/azure-firewall.ts`
- [x] T008 [US1] Implement `createAzureFirewallHandler` function (content-return pattern) in `src/tools/azure-firewall.ts`
- [x] T009 [US1] Implement `registerAzureFirewallTool` with agent workflow instructions in `src/tools/azure-firewall.ts`
- [x] T010 [US1] Wire up `registerAzureFirewallTool` in server initialization in `src/server/index.ts`
- [x] T011 [US1] Create unit tests for initial YAML creation in `tests/tools/azure-firewall.test.ts`

**Checkpoint**: User Story 1 functional - tool returns valid YAML for new rules.

---

## Phase 4: User Story 2 - Append Rules to Existing Configuration (Priority: P2)

**Goal**: Enable passing existing YAML content and returning merged result.

**Independent Test**: Pass existing YAML with 1 rule, add new rule, verify returned YAML has 2 rules.

### Implementation for User Story 2

- [x] T012 [US2] Add `existing_yaml` parameter handling in `src/tools/azure-firewall.ts`
- [x] T013 [US2] Implement YAML parsing of `existing_yaml` content in `src/tools/azure-firewall.ts`
- [x] T014 [US2] Implement rule append logic (merge new rule into existing array) in `src/tools/azure-firewall.ts`
- [x] T015 [US2] Create unit tests for rule appending in `tests/tools/azure-firewall.test.ts`

**Checkpoint**: User Story 2 functional - existing rules preserved when adding new ones.

---

## Phase 5: User Story 3 - Input Validation (Priority: P2)

**Goal**: Reject invalid IP addresses with clear error messages.

**Independent Test**: Call tool with malformed IPs (e.g., "999.999.999.999") and verify error response.

### Implementation for User Story 3

- [x] T016 [US3] Add regex `pattern` validation to schema for `source` and `destination` in `src/tools/azure-firewall.ts`
- [x] T017 [US3] Integrate `isValidIPOrCIDR` validation in handler (defense in depth) in `src/tools/azure-firewall.ts`
- [x] T018 [US3] Return descriptive error for invalid `existing_yaml` YAML content in `src/tools/azure-firewall.ts`
- [x] T019 [US3] Create unit tests for validation edge cases in `tests/tools/azure-firewall.test.ts`

**Checkpoint**: User Story 3 functional - invalid inputs rejected with clear messages.

---

## Phase 6: User Story 4 - Duplicate Rule Detection (Priority: P2)

**Goal**: Detect exact duplicate rules and return unchanged YAML with notification.

**Independent Test**: Pass existing YAML containing rule X, try to add rule X again, verify `action: duplicate_detected`.

### Implementation for User Story 4

- [x] T020 [US4] Implement duplicate detection logic (match team, source, destination, port) in `src/tools/azure-firewall.ts`
- [x] T021 [US4] Return `action: duplicate_detected` with unchanged YAML in `src/tools/azure-firewall.ts`
- [x] T022 [US4] Create unit tests for duplicate detection in `tests/tools/azure-firewall.test.ts`

**Checkpoint**: User Story 4 functional - duplicates prevented with clear notification.

---

## Phase 7: User Story 5 - Workspace Repository Selection (Priority: P1)

**Goal**: Tool description instructs agent to ask user for repository path when ambiguous.

**Independent Test**: Verify tool description contains multi-repo workspace instructions.

### Implementation for User Story 5

- [x] T023 [US5] Update tool description with repository path determination instructions in `src/tools/azure-firewall.ts`
- [x] T024 [US5] Include instruction to ask user if in multi-repo workspace in `src/tools/azure-firewall.ts`

**Checkpoint**: User Story 5 functional - agent knows to ask for clarification.

---

## Phase 8: Schema Resource (FR-FW-017)

**Goal**: Expose Azure Firewall Rule Schema as MCP Resource.

### Phase 8a: Foundation - Resource Capability (Missing from 001)

**Purpose**: Implement the generic `ResourceRegistry` and server capabilities required to expose resources.

- [ ] T025a [P] Define `ResourceHandler` and `Resource` interfaces in `src/resources/types.ts`
- [ ] T025b [P] Implement `ResourceRegistry` class in `src/resources/registry.ts`
- [ ] T025c Integrate `ResourceRegistry` into `MCPServer` (add `resources/list` and `resources/read` support) in `src/server/index.ts`

### Phase 8b: Feature - Schema Resource

- [ ] T025 Implement `getAzureFirewallSchema` function returning JSON Schema in `src/tools/azure-firewall.ts`
- [ ] T026 Register "Azure Firewall Rule Schema" resource at `mcp://azure-firewall/schema` in `src/server/index.ts`
- [ ] T027 Create test verifying schema resource returns valid JSON Schema in `tests/tools/azure-firewall.test.ts`

---

## Phase 9: Polish & Cross-Cutting Concerns

**Purpose**: Final cleanup and documentation.

- [x] T028 [P] Update `README.md` with usage examples for content-return pattern
- [x] T029 Run `bun run typecheck` to ensure no TypeScript errors
- [x] T030 Run `bun test` to verify all tests pass
- [x] T031 [P] Update `specs/002-azure-tools/contracts/azure_firewall.json` with current schema

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - verification only
- **Foundational (Phase 2)**: Depends on Setup - BLOCKS all user stories
- **User Stories (Phases 3-7)**: All depend on Foundational completion
  - US1 (Create) must complete before US2 (Append) - builds on same handler
  - US3 (Validation) and US4 (Duplicates) can run in parallel after US2
  - US5 (Workspace) can run in parallel - only changes tool description
- **Schema Resource (Phase 8)**: Can run after US1
- **Polish (Phase 9)**: Depends on all user stories complete

### Parallel Opportunities

- T004 and T005 are independent (Foundational phase)
- T006 and T007 are independent (US1 - interfaces)
- After Foundation, US3/US4/US5 can proceed in parallel
- All [P] tasks in same phase can run concurrently

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (verify)
2. Complete Phase 2: Foundational (validation utils)
3. Complete Phase 3: User Story 1 (create initial rules)
4. **STOP and VALIDATE**: Test tool returns valid YAML
5. Deploy/demo if ready

### Incremental Delivery

1. Foundation → US1 (Create) → MVP ready
2. Add US2 (Append) → Rules can accumulate
3. Add US3 (Validation) + US4 (Duplicates) → Robust input handling
4. Add US5 (Workspace) → Agent knows when to ask
5. Add Schema Resource → Clients can validate

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story
- Content-return pattern: Tool returns YAML, agent writes file
- TypeScript paths use `implementations/typescript/` per Constitution Principle Ia
- Go paths use `implementations/go/` per Constitution Principle Ib
- Tests use `bun:test` (TypeScript) and `go test` (Go)
- Both implementations must maintain feature parity
