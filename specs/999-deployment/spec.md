# Feature Specification: GitHub Actions Deployment Workflow

**Feature Branch**: `999-deployment`
**Created**: 2026-01-16
**Status**: Draft
**Input**: User description: "Can we create a new specify (999 is the number) and this is for deployments. The idea is we specify how we want our deployments to go out. Our method is the following. GitHub workflows is what we use. The workflow must build the app (dockerfile) and then publish this to github registry. We do not need to worry about kubernetes yet, but perhaps make a note that it will reach kubernetes at some point."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Automated Build & Publish (Priority: P1)

As a Developer, I want my application container to be automatically built and published to the GitHub Container Registry (GHCR) whenever I push code to the main branch, so that the latest version is always available for deployment.

**Why this priority**: This is the core requirement. Without a built image in the registry, no downstream deployment can occur.

**Independent Test**: Can be tested by pushing a commit to the `main` branch and verifying that a new package version appears in the repository's package settings on GitHub.

**Acceptance Scenarios**:

1. **Given** valid Dockerfiles exist in both `implementations/typescript/` and `implementations/go/`, **When** code is pushed to the `main` branch, **Then** a GitHub Action triggers and successfully builds both Docker images using a matrix strategy.
2. **Given** the builds succeed, **When** the workflow attempts to publish, **Then** both images are pushed to `ghcr.io` under the repository namespace with appropriate suffixes (`-ts`, `-go`).
3. **Given** the images are pushed, **Then** each is tagged with `latest` and the short commit SHA.

---

### Edge Cases

- What happens if the Docker build fails? (Workflow should fail and notify)
- What happens if authentication to GHCR fails? (Workflow should fail)
- What happens if the `Dockerfile` is missing? (Workflow should fail early with clear error)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST implement a GitHub Actions workflow file (YAML).
- **FR-002**: The workflow MUST trigger automatically on `push` events to the `main` branch.
- **FR-003**: The workflow MUST authenticate with the GitHub Container Registry (GHCR) using the `GITHUB_TOKEN`.
- **FR-004**: The workflow MUST build Docker images using the Dockerfiles located in `implementations/typescript/` and `implementations/go/` using a matrix strategy.
- **FR-005**: The workflow MUST publish both built images to `ghcr.io` with appropriate suffixes (`-ts`, `-go`).
- **FR-006**: The published image MUST be tagged with `latest`.
- **FR-007**: The published image MUST also be tagged with the short commit SHA (7 characters) for version pinning.
- **FR-008**: The workflow MUST run on an `ubuntu-latest` runner.

### Assumptions

- The repository is hosted on GitHub.
- A `Dockerfile` exists in `implementations/typescript/` and `implementations/go/`.
- The repository permissions allow the `GITHUB_TOKEN` to push packages (default for most repos, but may need configuration).
- Kubernetes deployment is EXPLICITLY OUT OF SCOPE for this iteration, but the image format should be compatible with standard k8s runtimes.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A push to `main` results in a new image appearing in GHCR within 5 minutes (assuming typical build times).
- **SC-002**: The `latest` tag always points to the most recent successful build from `main`.
- **SC-003**: 100% of successful workflow runs produce a pullable Docker image.
