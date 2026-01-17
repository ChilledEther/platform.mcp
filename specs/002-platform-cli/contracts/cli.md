# CLI Contract: `platform`

## Command: `generate workflows`

Generates GitHub Actions workflow files.

### Arguments
None.

### Flags
| Flag | Short | Environment Variable | Default | Description |
|------|-------|----------------------|---------|-------------|
| `--project-name` | `-p` | `PLATFORM_PROJECT_NAME` | `current-dir` | Name of the project |
| `--docker` | `-d` | `PLATFORM_USE_DOCKER` | `false` | Include Dockerfile |
| `--workflow-type`| `-t` | `PLATFORM_WORKFLOW_TYPE`| `go` | Type of workflow |
| `--dry-run` | | | `false` | Preview only |
| `--force` | `-f` | | `false` | Overwrite existing files |
| `--output` | `-o` | `PLATFORM_OUTPUT_DIR` | `.` | Target directory |

### Exit Codes
| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | General error (invalid flags, I/O failure) |
| `130` | Interrupted by user (e.g., Ctrl+C during prompt) |

### Output Format (Success)
```text
✔ Created .github/workflows/go.yaml
✔ Created Dockerfile
Generation complete! 2 files created.
```

### Output Format (Dry Run)
```text
[DRY RUN] Would create .github/workflows/go.yaml
[DRY RUN] Would create Dockerfile
Dry run complete. No files were written.
```

### Output Format (Conflict)
```text
? File .github/workflows/go.yaml already exists. Overwrite? [y/N]
```
