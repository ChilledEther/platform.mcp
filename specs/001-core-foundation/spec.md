# Feature Specification: Core Foundation Library

**Feature Branch**: `001-core-foundation`
**Created**: 2026-01-17
**Status**: Draft
**Input**: User description: "Core shared library (pkg/) that provides file generation logic for GitHub Actions workflows"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Generate File Content (Priority: P1)

As a developer building a CLI or MCP tool, I need to generate file content (path, content, permissions) without performing I/O, so that I can decide how to handle the output (write to disk, return to agent, etc.).

**Why this priority**: This is the foundational capability that enables both CLI and MCP consumers. Without pure data generation, neither consumer can function.

**Independent Test**: Can be fully tested by calling the Generate function with a configuration and verifying the returned file slice contains correct paths, content, and permissions.

**Acceptance Scenarios**:

1. **Given** a configuration with project name "my-project", **When** Generate is called, **Then** it returns a slice of File structs with non-empty Path and Content fields
2. **Given** a configuration requesting GitHub Actions workflows, **When** Generate is called, **Then** the returned files include `.github/workflows/*.yaml` paths
3. **Given** any configuration, **When** Generate is called, **Then** no files are written to disk (pure function)

---

### User Story 2 - Configure Generation Options (Priority: P2)

As a developer, I need to pass configuration options to control what files are generated, so that I can customize output for different project types.

**Why this priority**: Customization is essential for real-world use, but the basic generation must work first.

**Independent Test**: Can be tested by passing different Config values and verifying the generated files differ accordingly.

**Acceptance Scenarios**:

1. **Given** a Config with UseDocker=true, **When** Generate is called, **Then** Dockerfile-related content is included
2. **Given** a Config with UseDocker=false, **When** Generate is called, **Then** no Dockerfile content is generated
3. **Given** a Config with custom workflow options, **When** Generate is called, **Then** the workflow YAML reflects those options

---

### User Story 3 - Embed Template Files (Priority: P3)

As a developer, I need templates bundled within the library binary, so that consumers don't need to ship external template files.

**Why this priority**: Bundling simplifies distribution but is an enhancement over hardcoded templates.

**Independent Test**: Can be tested by verifying embedded templates are accessible and correctly parsed.

**Acceptance Scenarios**:

1. **Given** the library is compiled, **When** accessing embedded templates, **Then** template content is available without external files
2. **Given** a template with placeholders, **When** Generate processes it with Config, **Then** placeholders are replaced with actual values

---

### User Story 4 - Generate Standard Configuration Files (Priority: P1)

As a Platform Engineer, I want to generate Docker build and FluxCD configuration files using standardized templates so that my services follow best practices and consistent deployment patterns.

**Why this priority**: Essential for ensuring that all new services created with the platform comply with infrastructure standards (Docker and GitOps/Flux).

**Independent Test**: Can be tested by running the scaffold/generation command and verifying the output matches the expected structure defined in the templates.

**Acceptance Scenarios**:

1. **Given** the platform tool is installed, **When** I run the command to scaffold a Docker build file, **Then** a `docker-build.yaml` file is created in the target directory with content matching `docker-build.yaml.tmpl`.
2. **Given** the platform tool is installed, **When** I run the command to scaffold a FluxCD configuration, **Then** a `fluxcd.yaml` file is created in the target directory with content matching `fluxcd.yaml.tmpl`.

---

### User Story 5 - Template Consistency Verification (Priority: P1)

As a Platform Maintainer, I want the build system to verify that the required templates exist and are valid so that broken or missing templates do not reach the users.

**Why this priority**: Prevents regression and ensures reliability of the scaffolding tool.

**Independent Test**: Can be tested by removing a template from `internal/templates` and running the test suite, which should fail.

**Acceptance Scenarios**:

1. **Given** the source code, **When** I run the project test suite, **Then** it checks for the existence of `docker-build.yaml.tmpl` and `fluxcd.yaml.tmpl` in `internal/templates`.
2. **Given** a template file is missing from `internal/templates`, **When** I run the test suite, **Then** the test fails with a clear error message.

### Edge Cases

- What happens when an empty project name is provided? Return validation error.
- What happens when conflicting options are set? Return validation error with clear message.
- What happens when a template contains invalid syntax? Return parse error with template name and line.
- **Existing Output Files**: If the target output file (e.g., `docker-build.yaml`) already exists in the destination directory, the system should prompt for confirmation before overwriting or fail safely.
- **Corrupt Templates**: If a template file exists but is empty or unreadable, the generation process should fail with a descriptive error message.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Library MUST provide a Generate function that returns file data without performing I/O
- **FR-002**: Library MUST define a File struct with Path (string), Content (string), and Mode (permissions) fields
- **FR-003**: Library MUST define a Config struct for customizing generation behavior
- **FR-004**: Library MUST validate Config before generation and return descriptive errors for invalid input
- **FR-005**: Library MUST support embedding template files within the compiled binary
- **FR-006**: Library MUST generate valid GitHub Actions workflow YAML content
- **FR-007**: Library MUST be importable by external projects (public API in pkg/)
- **FR-008**: Library MUST NOT have side effects (no disk writes, no network calls, no stdout)
- **FR-009**: The system MUST store standard templates in `internal/templates`, including `docker-build.yaml.tmpl` and `fluxcd.yaml.tmpl`.
- **FR-010**: The system MUST include automated tests that verify the presence and accessibility of all required templates in `internal/templates`.
- **FR-011**: The template loading mechanism MUST support variable substitution (implied by `.tmpl` extension).

### Key Entities

- **File**: Represents a generated file with path, content, and permissions. Used by consumers to write to disk or return to agents.
- **Config**: Represents generation options (project name, features enabled, workflow type). Validated before use.
- **Generator**: Interface for feature modules to implement, enabling extensibility.
- **Template Registry**: The collection of file templates located in `internal/templates`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: All generated YAML files pass validation against GitHub Actions schema
- **SC-002**: Library functions complete in under 100ms for typical configurations
- **SC-003**: 100% of public API functions have corresponding test coverage
- **SC-004**: Zero side effects verified by running tests in read-only filesystem mode
- **SC-005**: External projects can import and use the library with a single `go get` command
- **SC-006**: 100% of the required templates (`docker-build.yaml.tmpl`, `fluxcd.yaml.tmpl`) are verified by the automated test suite.
- **SC-007**: Generated files are syntactically valid YAML.

## Assumptions

- Templates will initially be hardcoded, with embedded file support added in iteration
- Default permissions for generated files are 0644 (readable by all, writable by owner)
- GitHub Actions workflows target the latest stable workflow syntax version
- The library follows semantic versioning for API stability
