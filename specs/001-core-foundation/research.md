# Research: Core Foundation Library

## Go Version Compatibility
- **Decision**: Target Go 1.25+.
- **Rationale**: Mandated by Constitution. 1.25 is required for the project.
- **Alternatives considered**: None (Constitution mandate).

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
