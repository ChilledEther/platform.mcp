# Tasks: Docker Environment

**Feature**: 998-docker-environment  
**Generated**: 2026-01-17  
**Source**: [plan.md](./plan.md) | [spec.md](./spec.md)

## Prerequisites Check

| Artifact | Status |
|----------|--------|
| research.md | Available |
| data-model.md | Available |
| contracts/ | Available |
| quickstart.md | Available |

## User Stories

| Story | Priority | Description |
|-------|----------|-------------|
| US1 | P1 | Build Container Images |
| US2 | P2 | Run Tests in Containers |
| US3 | P3 | Run Containers Locally |

---

## Phase 1: Setup

- [ ] [T1.1] [P0] [SETUP] Create directory `build/package/platform/`
- [ ] [T1.2] [P0] [SETUP] Create directory `build/package/platform-mcp/`
- [ ] [T1.3] [P0] [SETUP] Update `.dockerignore` with Go-specific exclusions per `quickstart.md`

---

## Phase 2: Foundation

- [ ] [T2.1] [P0] [FOUNDATION] Create stub `cmd/platform/main.go` with `--version` flag per `quickstart.md:99-117`
- [ ] [T2.2] [P0] [FOUNDATION] Create stub `cmd/platform-mcp/main.go` with `--version` and `--health` flags per `quickstart.md:120-145`
- [ ] [T2.3] [P0] [FOUNDATION] Run `go mod tidy` to ensure dependencies are resolved

---

## Phase 3: US1 - Build Container Images

**User Story**: As a developer, I want to build container images for both the CLI and MCP server, so that I can distribute and deploy them consistently.

### Dockerfiles

- [ ] [T3.1] [P1] [US1] Create `build/package/platform/Dockerfile` per `contracts/dockerfile-platform.md`
  - Multi-stage build: deps -> build -> runtime
  - Base: `golang:1.25-alpine` (build), `alpine:3.20` (runtime)
  - OCI labels for version, revision, created, source, title, description
  - Non-root user `app`
  - HEALTHCHECK with `--version`
  - VOLUME `/workspace`
- [ ] [T3.2] [P1] [US1] Create `build/package/platform-mcp/Dockerfile` per `contracts/dockerfile-platform-mcp.md`
  - Multi-stage build: deps -> build -> runtime
  - Base: `golang:1.25-alpine` (build), `alpine:3.20` (runtime)
  - OCI labels for version, revision, created, source, title, description
  - Non-root user `app`
  - HEALTHCHECK with `--health`
  - VOLUME `/config`

### Build Script

- [ ] [T3.3] [P1] [US1] Update `scripts/Invoke-DockerBuild.ps1` to support both images
  - Add `-ImageName` parameter with values: `platform`, `platform-mcp`, `all`
  - Add `-Version`, `-Commit`, `-BuildDate` build arguments
  - Support building single or both images
  - Pass build args to docker build command

### Validation

- [ ] [T3.4] [P1] [US1] Verify `docker build` succeeds for `platform` image
- [ ] [T3.5] [P1] [US1] Verify `docker build` succeeds for `platform-mcp` image
- [ ] [T3.6] [P1] [US1] Verify image sizes are under 50MB (SC-001)
- [ ] [T3.7] [P1] [US1] Verify OCI labels are present in both images (FR-010)

---

## Phase 4: US2 - Run Tests in Containers

**User Story**: As a developer, I want to run all tests inside containers, so that I can verify the software works in the same environment as production.

### Test Script

- [ ] [T4.1] [P2] [US2] Create `scripts/Test-Docker.ps1` per `contracts/test-docker-script.md`
  - Parameter: `-ImageName` with values: `platform`, `platform-mcp`, `all`
  - Test: Build both images
  - Test: Image size < 50MB
  - Test: Required OCI labels present
  - Test: Non-root user configured
  - Exit code 0 on success, 1 on failure

### Validation

- [ ] [T4.2] [P2] [US2] Run `Test-Docker.ps1` and verify exit code 0
- [ ] [T4.3] [P2] [US2] Verify all unit tests pass inside container (FR-005)
- [ ] [T4.4] [P2] [US2] Verify test output is visible on failure (AC: US2.3)

---

## Phase 5: US3 - Run Containers Locally

**User Story**: As a developer, I want to run containers locally for manual testing, so that I can verify behavior before deployment.

### CLI Container

- [ ] [T5.1] [P3] [US3] Verify `docker run platform --version` outputs version info
- [ ] [T5.2] [P3] [US3] Verify volume mount works: `docker run -v $(pwd):/workspace platform ls /workspace`
- [ ] [T5.3] [P3] [US3] Verify health check passes: `docker run platform --version` returns 0

### MCP Container

- [ ] [T5.4] [P3] [US3] Verify `docker run platform-mcp --health` outputs "ok"
- [ ] [T5.5] [P3] [US3] Verify container starts in under 3 seconds (SC-004)
- [ ] [T5.6] [P3] [US3] Verify MCP stdio mode works: `docker run -i platform-mcp`

---

## Phase 6: Polish

- [ ] [T6.1] [P4] [POLISH] Verify builds complete in under 5 minutes (SC-002)
- [ ] [T6.2] [P4] [POLISH] Verify images work on Linux (native Docker)
- [ ] [T6.3] [P4] [POLISH] Document usage in quickstart.md if needed
- [ ] [T6.4] [P4] [POLISH] Clean up test images: `docker rmi platform:test platform-mcp:test`

---

## Success Criteria Mapping

| Criteria | Task(s) | Target |
|----------|---------|--------|
| SC-001 | T3.6 | Image size < 50MB |
| SC-002 | T6.1 | Build time < 5 minutes |
| SC-003 | T4.3 | 100% tests pass in container |
| SC-004 | T5.5 | Container startup < 3 seconds |
| SC-005 | T6.2 | Works on Linux, macOS, Windows |

## Functional Requirements Mapping

| Requirement | Task(s) |
|-------------|---------|
| FR-001 | T3.1, T3.2 |
| FR-002 | T3.1, T3.2 |
| FR-003 | T3.1, T3.2 |
| FR-004 | T3.1, T3.2 |
| FR-005 | T4.3 |
| FR-006 | T5.1, T5.4 |
| FR-007 | T5.2 |
| FR-008 | T5.3, T5.4 |
| FR-009 | T3.4, T3.5 |
| FR-010 | T3.7 |

## Edge Cases

| Case | Mitigation | Task |
|------|------------|------|
| Missing dependencies | `go mod tidy` in T2.3 | T2.3 |
| Build fails | Clear error message in script | T4.1 |
| No container runtime | Guidance in Test-Docker.ps1 | T4.1 |
| Disk space insufficient | Check before build | T4.1 |

---

## Task Count Summary

| Phase | Tasks | Status |
|-------|-------|--------|
| Phase 1: Setup | 3 | Pending |
| Phase 2: Foundation | 3 | Pending |
| Phase 3: US1 | 7 | Pending |
| Phase 4: US2 | 4 | Pending |
| Phase 5: US3 | 6 | Pending |
| Phase 6: Polish | 4 | Pending |
| **Total** | **27** | - |
