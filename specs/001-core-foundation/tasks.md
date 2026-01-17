# Tasks: 001-core-foundation

**Input**: Design documents from `/specs/001-core-foundation/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: TDD is mandatory per Constitution II. Tests are included for each user story and MUST be written first.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Initialize Go module in implementations/go/pkg/scaffold/go.mod
- [X] T002 Create initial directory structure for pkg/scaffold
- [X] T003 [P] Configure golangci-lint in implementations/go/.golangci.yml

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T004 Define File and Config structs in implementations/go/pkg/scaffold/types.go
- [X] T005 Define Generator interface in implementations/go/pkg/scaffold/generator.go
- [X] T006 Implement Config validation logic in implementations/go/pkg/scaffold/validate.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Generate File Content (Priority: P1) 🎯 MVP

**Goal**: Provide a Generate function that returns structured file data without I/O.

**Independent Test**: Call Generate with a valid Config and verify the returned File slice contains correct Paths and non-empty Content.

### Tests for User Story 1 (TDD Mandatory) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T007 [P] [US1] Create unit tests for basic generation in implementations/go/pkg/scaffold/scaffold_test.go
- [X] T008 [P] [US1] Create table-driven tests for workflow paths in implementations/go/pkg/scaffold/scaffold_test.go
- [X] T009 [US1] Implement Generate function skeleton in implementations/go/pkg/scaffold/scaffold.go
- [X] T010 [US1] Implement hardcoded workflow content generation in implementations/go/pkg/scaffold/scaffold.go
- [X] T011 [US1] Add side-effect verification (read-only filesystem check) in implementations/go/pkg/scaffold/scaffold_test.go

**Checkpoint**: User Story 1 functional and testable independently.

---

## Phase 4: User Story 2 - Configure Generation Options (Priority: P2)

**Goal**: Allow customization of generated files via Config (Docker, workflow types).

**Independent Test**: Verify that toggling UseDocker in Config adds/removes Dockerfile-related entries in the output.

### Tests for User Story 2 (TDD Mandatory) ⚠️

- [X] T012 [P] [US2] Add tests for UseDocker configuration in implementations/go/pkg/scaffold/scaffold_test.go
- [X] T013 [P] [US2] Add tests for WorkflowType selection in implementations/go/pkg/scaffold/scaffold_test.go
- [X] T014 [US2] Implement conditional Dockerfile generation in implementations/go/pkg/scaffold/scaffold.go
- [X] T015 [US2] Implement WorkflowType selection logic in implementations/go/pkg/scaffold/scaffold.go

**Checkpoint**: User Story 2 integrated and testable.

---

## Phase 5: User Story 3 - Embed Template Files (Priority: P3)

**Goal**: Bundle template files within the library using go:embed.

**Independent Test**: Remove hardcoded strings and verify templates are loaded from embedded FS and parsed correctly.

### Tests for User Story 3 (TDD Mandatory) ⚠️

- [X] T016 [P] [US3] Create tests for template embedding in implementations/go/pkg/scaffold/templates_test.go
- [X] T017 [P] [US3] Create tests for template parsing with placeholders in implementations/go/pkg/scaffold/templates_test.go
- [X] T018 [US3] Create template directory and initial workflow templates in implementations/go/internal/templates/
- [X] T019 [US3] Implement go:embed for templates in implementations/go/internal/templates/templates.go
- [X] T020 [US3] Refactor Generate to use text/template with embedded files in implementations/go/pkg/scaffold/scaffold.go

**Checkpoint**: All user stories functional.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T021 [P] Finalize API documentation in pkg/scaffold/doc.go
- [X] T022 [P] Add golangci-lint check to CI pipeline
- [X] T023 Verify 100% test coverage for public API
- [X] T024 Validate against quickstart.md instructions

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies.
- **Foundational (Phase 2)**: Depends on Setup.
- **User Stories (Phase 3+)**: Depend on Foundational completion.
- **Polish (Phase 6)**: Depends on all stories.

### Parallel Opportunities

- T003 (Linting) can run parallel with T001/T002.
- T007, T008 (Tests) can run parallel within US1.
- Once Foundation is done, US1 work starts. US2 and US3 have minor dependencies on US1 (shared scaffold.go file), but logic is mostly additive.

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Setup + Foundational.
2. Complete US1 (Basic generation).
3. Validate independent tests for US1.
