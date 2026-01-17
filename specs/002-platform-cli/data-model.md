# Data Model: Platform CLI Tool

## CLI Configuration (`CliConfig`)

This model represents the parsed command-line arguments and environment variables.

| Field | Type | Description | Mapping to `scaffold.Config` |
|-------|------|-------------|-----------------------------|
| `ProjectName` | `string` | Name of the project | `ProjectName` |
| `UseDocker` | `bool` | Whether to include Dockerfile | `UseDocker` |
| `WorkflowType`| `string` | Type of workflow (go, ts, python) | `WorkflowType` |
| `DryRun` | `bool` | Preview only, no I/O | N/A |
| `Force` | `bool` | Overwrite without prompting | N/A |
| `OutputDir` | `string` | Target directory | N/A |

## File Operation Result (`OpResult`)

Used to track the status of file operations during generation.

| Field | Type | Description |
|-------|------|-------------|
| `Path` | `string` | Relative path of the file |
| `Status` | `string` | `Created`, `Overwritten`, `Skipped`, `Error` |
| `Message` | `string` | Details (e.g., error message) |

## User Confirmation (`Confirmation`)

Internal model for handling interactive prompts.

| Field | Type | Description |
|-------|------|-------------|
| `FilePath` | `string` | File that already exists |
| `UserResponse`| `bool` | Whether the user agreed to overwrite |
