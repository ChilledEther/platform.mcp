# UAT Issues: Phase 02 Plan 02

**Tested:** 2026-01-19
**Source:** .planning/phases/02-platform-cli/02-02-SUMMARY.md
**Tester:** User via /gsd-verify-work

## Open Issues

[None]

## Resolved Issues

### UAT-001: go.yaml workflow always generated regardless of flags

**Resolved:** 2026-01-19 - Fixed in 02-02-FIX.md

**Root Cause:** The `workflow_go` condition in `ShouldGenerate()` returned true when `WorkflowType == "go"` (the default), without checking if `WithActions` was set.

**Fix:** Modified `workflow_*` conditions in `pkg/scaffold/manifest.go` to require `WithActions` flag. Also updated:
- `internal/cli/cmd/generate.go` - workflows subcommand now sets `withActions=true`
- `internal/mcp/handler.go` - generate_workflows tool now sets `WithActions=true`

**Discovered:** 2026-01-19
**Phase/Plan:** 02-02
**Severity:** Minor
**Feature:** CLI generation (all flags)
**Description:** The CLI always generates `.github/workflows/go.yaml` regardless of which flags are used. This workflow should only appear when explicitly requested (e.g., `--with-actions` with go type, or a dedicated flag). Currently it appears with `--with-flux`, `--with-docker`, and other combinations where it wasn't requested.
**Expected:** go.yaml only generated when explicitly requested
**Actual:** go.yaml generated unconditionally with every generation command
**Repro:**
1. Run `go run cmd/platform/main.go generate --with-flux --project-name test`
2. Or run `go run cmd/platform/main.go generate --with-docker --project-name test`
3. Observe go.yaml appears in output alongside requested files

---

*Phase: 02-platform-cli*
*Plan: 02*
*Tested: 2026-01-19*
