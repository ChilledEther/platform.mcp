# Quickstart: Deployment Pipeline

**Date**: 2026-01-17  
**Feature**: 999-deployment

## Prerequisites

- Git repository hosted on GitHub
- Go 1.25+ project with `cmd/platform` and `cmd/platform-mcp` binaries
- Conventional Commits format enforced

## Files to Create

### 1. GitHub Workflows

```bash
mkdir -p .github/workflows
```

Create these workflow files:
- `.github/workflows/ci.yml` - Continuous integration
- `.github/workflows/release-please.yml` - Automated releases
- `.github/workflows/publish.yml` - Container publishing

### 2. Release Please Configuration

Create `release-please-config.json`:
```json
{
  "$schema": "https://raw.githubusercontent.com/googleapis/release-please/main/schemas/config.json",
  "release-type": "go",
  "packages": {
    "cmd/platform": { "component": "platform" },
    "cmd/platform-mcp": { "component": "platform-mcp" }
  }
}
```

Create `.release-please-manifest.json`:
```json
{
  "cmd/platform": "0.0.1",
  "cmd/platform-mcp": "0.0.1"
}
```

### 3. Dockerfiles

```bash
mkdir -p build/package/platform build/package/platform-mcp
```

Each Dockerfile should:
1. Use multi-stage build
2. Build with `golang:1.25-alpine`
3. Run on `alpine:3.21`
4. Produce minimal static binary

## Verification Steps

### Test CI Pipeline
```bash
git push origin feature-branch
# Verify GitHub Actions "CI" workflow runs and passes
```

### Test Release Please
```bash
git checkout main
git merge --no-ff feature-branch -m "feat: add new feature"
git push origin main
# Verify Release Please creates a PR
```

### Test Container Publishing
```bash
# Merge the Release Please PR
# Verify containers appear in GHCR:
docker pull ghcr.io/<owner>/platform:latest
docker pull ghcr.io/<owner>/platform-mcp:latest
```

## Common Issues

| Issue | Solution |
|-------|----------|
| "No releases found" | Ensure commits follow Conventional Commits |
| "Package write permission denied" | Check workflow `permissions` block |
| "Docker build fails" | Verify Dockerfile path in workflow |

## Success Criteria

- [ ] CI runs on every push
- [ ] Tests block merges when failing
- [ ] Release PRs created automatically
- [ ] Changelogs generated from commits
- [ ] Container images published on release
- [ ] Images tagged with version and `latest`
