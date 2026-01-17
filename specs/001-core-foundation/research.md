# Research: Core Foundation Library

## Go Version Compatibility
- **Decision**: Target Go 1.23+ for now, unless 1.25 provides critical features.
- **Rationale**: Go 1.25 is not yet the standard stable release in many environments. 1.23/1.24 provide necessary `embed` and library support.
- **Alternatives considered**: Using older versions, but `go:embed` requires 1.16+ and modern SDKs prefer 1.21+.

## MCP Go SDK Patterns
- **Decision**: Use `mcp.NewServer` and `mcp.Tool` for tool registration.
- **Rationale**: Follows the official SDK patterns for defining tool schemas and handlers.
- **Alternatives considered**: Manual JSON-RPC handling (rejected as too complex/brittle).

## Template Strategy
- **Decision**: Use `embed` package to bundle `.yaml.tmpl` files and `text/template` for rendering.
- **Rationale**: Standard Go practice for bundling assets. `text/template` is sufficient for YAML generation when paired with proper escaping if needed, though simple replacement is usually enough for workflows.
- **Alternatives considered**: Hardcoded strings (rejected as unmaintainable), External files (rejected by Constitution I/IV).

## Validation Logic
- **Decision**: Implement custom validation logic within the `Config` struct.
- **Rationale**: Keeps the core library dependencies minimal.
- **Alternatives considered**: Third-party validation libraries (rejected to keep `pkg/` lean).
