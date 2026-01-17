# Tasks: GitHub Actions Deployment Workflow

**Input**: Design documents from `/specs/999-deployment/`
**Prerequisites**: plan.md (required), spec.md (required), data-model.md, contracts/

**Tests**: Manual verification of workflow execution (Push -> Build -> Publish)

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Workflows**: `.github/workflows/` at repository root
- **Docker**: `implementations/*/Dockerfile`
- **Implementations**: `implementations/typescript/` and `implementations/go/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Verify repository structure for CI/CD.

- [x] T001 Verify `Dockerfile` exists in `implementations/typescript/` and `implementations/go/`
- [x] T002 [P] Create `.github/workflows/` directory structure

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core configuration MUST be complete before workflow implementation.

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T003 [P] Define environment variables for registry and image name in `.github/workflows/docker-build-publish.yml`

**Checkpoint**: Foundation ready - user story implementation can now begin.

---

## Phase 3: User Story 1 - Automated Build & Publish (Priority: P1) 🎯 MVP

**Goal**: Implement the core workflow that builds and pushes the Docker image.

**Independent Test**: Push to `main` branch triggers workflow, which succeeds and publishes image to GHCR.

### Implementation for User Story 1

- [x] T004 [US1] Create workflow file `.github/workflows/docker-build-publish.yml` with basic triggers (push to main)
- [x] T005 [US1] Add `checkout` step using `actions/checkout@v6` in `.github/workflows/docker-build-publish.yml`
- [x] T006 [US1] Configure matrix strategy for `typescript` and `go` implementations
- [x] T007 [US1] Add `docker/login-action@v3` step for GHCR authentication
- [x] T008 [US1] Add `docker/metadata-action@v5` step for generating tags with matrix suffixes (`-ts`, `-go`)
- [x] T009 [US1] Add `docker/build-push-action@v6` step using matrix context
- [x] T010 [US1] Configure workflow permissions (`contents: read`, `packages: write`)

**Checkpoint**: At this point, User Story 1 should be fully functional.

---

## Phase 4: Polish & Cross-Cutting Concerns

**Purpose**: Final cleanup and documentation.

- [x] T011 [P] Validate YAML syntax of `.github/workflows/docker-build-publish.yml`
- [x] T012 [P] Update `README.md` with badge or instructions for the new workflow
- [x] T013 Run manual validation: Trigger workflow and verify GHCR packages (`-ts` and `-go`)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion
- **User Stories (Phase 3)**: Depends on Foundational phase completion
- **Polish (Final Phase)**: Depends on user story completion

### User Story Dependencies

- **User Story 1 (P1)**: Core workflow implementation.

### Parallel Opportunities

- T001 and T002 can run in parallel.
- T010 and T011 can run in parallel.

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Trigger workflow via push to main
5. Verify package in GitHub UI

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
