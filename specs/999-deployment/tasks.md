# Executable Tasks: Deployment Pipeline

**Feature**: 999-deployment  
**Generated**: 2026-01-17  
**Status**: Ready for Implementation

## Task Summary

| Phase | Description | Tasks | Status |
|-------|-------------|-------|--------|
| 1 | Setup | 3 | Pending |
| 2 | Foundation | 2 | Pending |
| 3 | US1: Automated Testing | 4 | Pending |
| 4 | US2: Automated Releases | 4 | Pending |
| 5 | US3: Container Publishing | 5 | Pending |
| 6 | Polish | 3 | Pending |

**Total Tasks**: 21

---

## Phase 1: Setup

Initial directory structure and prerequisite verification.

- [ ] [T001] [P1] [Setup] Create directory `build/package/platform/` for CLI Dockerfile
- [ ] [T002] [P1] [Setup] Create directory `build/package/platform-mcp/` for MCP Dockerfile
- [ ] [T003] [P1] [Setup] Verify `.github/workflows/` directory exists (already present)

---

## Phase 2: Foundation

Core infrastructure dependencies that other phases require.

- [ ] [T004] [P1] [Foundation] Create `release-please-config.json` with multi-component Go configuration at `/release-please-config.json`
- [ ] [T005] [P1] [Foundation] Create `.release-please-manifest.json` with initial versions (0.0.1) at `/.release-please-manifest.json`

---

## Phase 3: US1 - Automated Testing on Push

**User Story**: As a developer, I want all tests to run automatically when I push code, so that I get immediate feedback on whether my changes break anything.

**Acceptance Criteria**:
- Tests run on push to any branch
- Tests run on pull requests
- Success/failure status reported to PR

- [ ] [T006] [P1] [US1] Create `.github/workflows/ci.yml` with push and pull_request triggers
- [ ] [T007] [P1] [US1] Add `test` job to `ci.yml`: checkout, setup-go@v5 with Go 1.25+, run `go test ./...`
- [ ] [T008] [P1] [US1] Add `build` job to `ci.yml`: depends on test, builds both `cmd/platform` and `cmd/platform-mcp`
- [ ] [T009] [P1] [US1] Configure Go module caching in `ci.yml` via `actions/setup-go` cache option

---

## Phase 4: US2 - Automated Releases

**User Story**: As a maintainer, I want releases to be created automatically based on commit history, so that versioning is consistent and changelogs are generated without manual effort.

**Acceptance Criteria**:
- Conventional Commits drive version bumps
- Release PRs created automatically
- Changelogs generated from commit messages

- [ ] [T010] [P2] [US2] Create `.github/workflows/release-please.yml` with push trigger on main branch
- [ ] [T011] [P2] [US2] Configure `googleapis/release-please-action@v4` with config-file and manifest-file options
- [ ] [T012] [P2] [US2] Set required permissions in `release-please.yml`: `contents: write`, `pull-requests: write`
- [ ] [T013] [P2] [US2] Export release outputs for downstream consumption (releases_created, tag_name per component)

---

## Phase 5: US3 - Publish Container Images

**User Story**: As a user, I want container images published automatically on release, so that I can pull the latest version without manual builds.

**Acceptance Criteria**:
- Images published on release event
- Multi-arch support (amd64, arm64)
- Tagged with version and `latest`

- [ ] [T014] [P3] [US3] Create `build/package/platform/Dockerfile` with multi-stage Alpine build for CLI
- [ ] [T015] [P3] [US3] Create `build/package/platform-mcp/Dockerfile` with multi-stage Alpine build for MCP server
- [ ] [T016] [P3] [US3] Create `.github/workflows/publish.yml` triggered on `release` event (type: published)
- [ ] [T017] [P3] [US3] Add `publish-platform` job with conditional trigger on `platform-v*` tags
- [ ] [T018] [P3] [US3] Add `publish-platform-mcp` job with conditional trigger on `platform-mcp-v*` tags

---

## Phase 6: Polish

Final verification, cleanup, and documentation updates.

- [ ] [T019] [P3] [Polish] Remove or refactor existing `.github/workflows/docker-build-publish.yaml` (superseded by publish.yml)
- [ ] [T020] [P3] [Polish] Verify all workflows pass lint checks (actionlint if available)
- [ ] [T021] [P3] [Polish] Update repository README with CI badge and container registry instructions

---

## Dependencies Graph

```text
Phase 1 (Setup)
    │
    ▼
Phase 2 (Foundation)
    │
    ├──────────────────┐
    ▼                  ▼
Phase 3 (US1)     Phase 4 (US2)
    │                  │
    └──────┬───────────┘
           ▼
     Phase 5 (US3)
           │
           ▼
     Phase 6 (Polish)
```

## File Manifest

| File | Phase | Description |
|------|-------|-------------|
| `build/package/platform/` | 1 | Directory for CLI Dockerfile |
| `build/package/platform-mcp/` | 1 | Directory for MCP Dockerfile |
| `release-please-config.json` | 2 | Release Please configuration |
| `.release-please-manifest.json` | 2 | Version tracking manifest |
| `.github/workflows/ci.yml` | 3 | CI workflow (test + build) |
| `.github/workflows/release-please.yml` | 4 | Automated release workflow |
| `build/package/platform/Dockerfile` | 5 | CLI container image |
| `build/package/platform-mcp/Dockerfile` | 5 | MCP container image |
| `.github/workflows/publish.yml` | 5 | Container publish workflow |

## Verification Checklist

After implementation, verify:

- [ ] Push to feature branch triggers CI workflow
- [ ] Tests pass and status reported to PR
- [ ] Merge to main creates release PR (if conventional commits present)
- [ ] Merging release PR creates GitHub Release with tag
- [ ] Release triggers publish workflow
- [ ] Container images available in GHCR with correct tags
