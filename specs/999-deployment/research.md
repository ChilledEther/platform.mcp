# Research: GitHub Actions Deployment Workflow

**Feature**: `999-deployment`
**Date**: 2026-01-16

## 1. Action Selection for Docker Build/Push

**Decision**: Use `docker/build-push-action` combined with `docker/login-action`.

**Rationale**:
- **Standard**: These are the official Docker actions recommended by GitHub and Docker.
- **Caching**: Built-in support for layer caching (`gha` cache) to speed up builds.
- **Security**: Handles credential management securely via `with: password: ${{ secrets.GITHUB_TOKEN }}`.
- **Multi-platform**: Supports multi-arch builds easily if needed in future (e.g., via QEMU).

**Alternatives Considered**:
- **Shell scripts (`docker build ...`)**: Harder to manage caching, layer pushing, and error handling manually. Less readable logs.
- **Other 3rd party actions**: Higher security risk, less maintenance.

## 2. Versioning Strategy

**Decision**: Tag images with both `latest` and `sha-{commit_short}`.

**Rationale**:
- **`latest`**: Provides a stable endpoint for "current production" references in development or simple deployments.
- **`sha-{commit_short}`**: Provides immutable, precise versioning for rollback and auditing. Allows k8s manifests to pin specific versions.

## 3. Registry Authentication

**Decision**: Use `GITHUB_TOKEN`.

**Rationale**:
- **Zero Config**: Available by default in GitHub Actions.
- **Secure**: Ephemeral, scoped to the repository.
- **Permission**: Default permissions usually allow package write, or can be easily adjusted in repository settings.

## 4. Workflow Triggers

**Decision**: `push` to `main`.

**Rationale**:
- Fits the requirement "whenever I push code to the main branch".
- Simple Continuous Delivery model.

**Alternatives Considered**:
- **`release` created**: Better for semantic versioning, but user specifically asked for push-based build.
