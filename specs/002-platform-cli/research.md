# Research: Platform CLI Tool

## Technical Approach

### CLI Framework: Cobra
We will use `github.com/spf13/cobra` to implement the CLI. Cobra is the industry standard for Go CLI applications (used by Kubernetes, Hugo, etc.).

**Key Commands:**
- `platform generate workflows`: The main command to generate files.
- `platform version`: Shows version information.

**Key Flags for `generate`:**
- `--project-name`: Maps to `scaffold.Config.ProjectName`.
- `--docker`: Maps to `scaffold.Config.UseDocker`.
- `--workflow-type`: Maps to `scaffold.Config.WorkflowType`.
- `--dry-run`: Preview changes without writing to disk.
- `--force`: Overwrite existing files without prompting.
- `--output`: Base directory for file generation (default: `.`).

### Core Integration
The CLI will import `github.com/modelcontextprotocol/platform.mcp/pkg/scaffold`.
It will:
1. Parse CLI flags into a `scaffold.Config`.
2. Call `scaffold.Generate(cfg)` which returns `[]scaffold.File`.
3. Iterate over the files and perform I/O operations.

### Safe File I/O
To ensure "Zero data loss" (SC-005):
- Check if file exists before writing.
- If it exists and `--force` is not set, prompt the user for confirmation.
- Use `os.MkdirAll` to create parent directories.
- Use `os.WriteFile` with the mode provided in `scaffold.File`.

### User Interaction
- Use `fmt.Scanln` or a lightweight library for prompts.
- Respect `stdin` connectivity for CI environments (auto-skip prompts).
- Use `stdout` for progress and `stderr` for errors.

## Alternatives Considered

### Alternative 1: Hand-written YAML templates in CLI
- **Pros**: Simple for a small number of files.
- **Cons**: Violates **Constitution I (Core-First)** and **IV (Shared Core)**. Harder to maintain as the platform grows.
- **Decision**: Rejected in favor of `pkg/scaffold` integration.

### Alternative 2: Using `survey` library for prompts
- **Pros**: Richer UI, better UX.
- **Cons**: Adds a large dependency for a simple "yes/no" prompt.
- **Decision**: Use standard `fmt` for now to keep dependencies minimal, reconsider if complex interaction is needed.

## Performance
- Go's file I/O is extremely fast. Generating < 10 files should take < 100ms, well within the 5s goal (SC-001).
