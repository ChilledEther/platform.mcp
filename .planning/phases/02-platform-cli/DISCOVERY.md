# Phase 2 Discovery: Platform CLI

## 1. Project Layout & Structure

We will adhere to the standard Go project layout, keeping the CLI entry point in `cmd/platform` and using `internal` for CLI-specific logic that shouldn't be exported.

```
.
├── cmd
│   └── platform
│       ├── main.go           # Entry point
│       ├── root.go           # Root command configuration
│       └── generate.go       # 'generate' command implementation
├── internal
│   └── cli
│       ├── logger            # UI/Logging helpers (color, spinner)
│       └── writer            # File I/O abstraction (Write, DryRun, Prompt)
└── pkg
    └── scaffold              # Existing core logic (Imported)
```

## 2. Command Hierarchy

The CLI will use `spf13/cobra`. The hierarchy will be:

*   **`platform`** (Root)
    *   Global Flags: `--verbose`
    *   **`generate`**
        *   Description: Scaffolds a new platform project or components.
        *   Flags:
            *   `--project-name` (string, required)
            *   `--with-actions` (bool)
            *   `--with-docker` (bool)
            *   `--with-flux` (bool)
            *   `--output` (string, default: ".")
            *   `--force` (bool, skip overwrite prompts)
            *   `--dry-run` (bool, print only)

*Note: Initially, `generate` will be a direct command. If we need `platform generate workflows` specifically, we can make `generate` a parent command, but based on Phase 1, the scaffold generates a mix of things. We will align with `pkg/scaffold` capabilities.*

## 3. Wiring Core Logic

The `pkg/scaffold` library is pure. The CLI is the "Consumer".

**Injection Strategy:**
1.  **Configuration Mapping:** The CLI `generate` command will parse flags into a `scaffold.Config` struct.
2.  **Execution:** It will call `scaffold.NewGenerator().Generate(config)`.
3.  **Output Handling:** The result `[]scaffold.File` will be passed to the `internal/cli/writer` service.

## 4. I/O Abstraction & Safety

To handle **FR-003 (Overwrite Prompt)** and **FR-004 (Dry Run)** robustly, we will implement a `FileWriter` in `internal/cli/writer`.

**Interface:**
```go
type Writer interface {
    Write(files []scaffold.File, opts WriteOptions) error
}

type WriteOptions struct {
    DryRun    bool
    Force     bool
    OutputDir string
}
```

**Behavior:**
*   **Dry Run:** Iterates files and prints `[DRY-RUN] Would write to <path>` with preview of content.
*   **Real Write:**
    *   Checks if file exists.
    *   If exists && !Force: Prompts user "Overwrite <path>? [y/N]".
    *   Creates parent directories (`os.MkdirAll`).
    *   Writes file.

## 5. Dependencies

We will add the following dependencies:
*   `github.com/spf13/cobra` (CLI Framework)
*   `github.com/spf13/pflag` (Flag parsing, transitive)

## 6. Testing Strategy

*   **Unit Tests (`internal/cli/writer`)**: Mock filesystem or use temporary directories to verify prompting and writing logic.
*   **Integration Tests (`cmd/platform`)**: Run the compiled binary against a test directory to verify end-to-end flow.

## 7. Action Items

1.  Initialize `go.mod` with Cobra.
2.  Create `cmd/platform/main.go` and `root.go`.
3.  Implement `internal/cli/writer`.
4.  Implement `cmd/platform/generate.go` wiring it all together.
