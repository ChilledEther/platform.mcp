# Feature Specification: Platform CLI Tool

**Feature Branch**: `002-platform-cli`  
**Created**: 2026-01-17  
**Status**: Draft  
**Input**: User description: "CLI tool (platform) that generates GitHub Actions workflow YAML files to disk"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Generate Workflow Files to Disk (Priority: P1)

As a developer, I want to run a command that generates GitHub Actions workflow files in my project's `.github/workflows/` folder, so that I can quickly set up CI/CD without writing YAML manually.

**Why this priority**: This is the core value proposition — generating files to disk is the primary CLI function.

**Independent Test**: Can be fully tested by running the CLI in a test directory and verifying the expected files are created with correct content.

**Acceptance Scenarios**:

1. **Given** I am in a project directory, **When** I run `platform generate workflows`, **Then** workflow YAML files are created in `.github/workflows/`
2. **Given** the `.github/workflows/` directory doesn't exist, **When** I run the generate command, **Then** the directory is created automatically
3. **Given** files already exist, **When** I run the generate command, **Then** I am prompted before overwriting (unless --force is used)

---

### User Story 2 - Customize Generation via Flags (Priority: P2)

As a developer, I want to pass flags to customize what is generated, so that I can tailor the output to my project's needs.

**Why this priority**: Customization enables real-world adoption but depends on basic generation working first.

**Independent Test**: Can be tested by running the CLI with different flags and verifying output varies accordingly.

**Acceptance Scenarios**:

1. **Given** I pass `--project-name my-app`, **When** the command runs, **Then** generated files reference "my-app"
2. **Given** I pass `--docker`, **When** the command runs, **Then** Docker-related workflow steps are included
3. **Given** I pass `--output /custom/path`, **When** the command runs, **Then** files are written to the specified path

---

### User Story 3 - Preview Without Writing (Priority: P3)

As a developer, I want to preview what will be generated without writing files, so that I can review before committing changes.

**Why this priority**: Preview is a safety feature that enhances UX but is not required for core functionality.

**Independent Test**: Can be tested by running with `--dry-run` and verifying no files are created while output is displayed.

**Acceptance Scenarios**:

1. **Given** I pass `--dry-run`, **When** the command runs, **Then** file contents are displayed to stdout but not written to disk
2. **Given** I pass `--dry-run`, **When** the command runs, **Then** a summary shows what would be created/modified

---

### Edge Cases

- What happens when the target directory is not writable? Display clear error message with path.
- What happens when running outside a git repository? Warn but allow generation.
- What happens with invalid flag combinations? Display usage help with specific error.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: CLI MUST provide a `generate` command that writes files to disk
- **FR-002**: CLI MUST create parent directories if they don't exist
- **FR-003**: CLI MUST prompt before overwriting existing files (unless `--force` is specified)
- **FR-004**: CLI MUST support `--dry-run` to preview without writing
- **FR-005**: CLI MUST support `--output` to specify a custom output directory
- **FR-006**: CLI MUST support `--project-name` to set the project name in generated content
- **FR-007**: CLI MUST display progress feedback during file generation
- **FR-008**: The system MUST accept a `--workflow-type` flag to specify the CI/CD template (e.g., `go`, `node`, `python`)
- **FR-009**: CLI MUST return appropriate exit codes (0 for success, non-zero for errors)
- **FR-010**: CLI MUST import all generation logic from the shared core library (pkg/)
- **FR-011**: CLI MUST provide `--help` for all commands with clear usage examples

### Key Entities

- **Command**: A CLI action (e.g., generate, version, help)
- **Flags**: Options that modify command behavior (e.g., --force, --dry-run)
- **Output**: Files written to disk or displayed to stdout

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can generate a complete workflow setup in under 500ms (excluding user think time)
- **SC-002**: CLI provides clear error messages that help users resolve issues without external documentation
- **SC-003**: 100% of generated files are syntactically valid YAML
- **SC-004**: CLI works correctly on Linux, macOS, and Windows
- **SC-005**: Zero data loss — existing files are never overwritten without explicit user consent

## Assumptions

- The CLI binary will be named `platform`
- Default output directory is the current working directory
- The CLI uses the shared core library for all generation logic (no duplicated logic)
- Interactive prompts are skipped when stdin is not a terminal (CI environments)
