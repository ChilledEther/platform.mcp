# Research: Azure Tools Group - Firewall Rules

## Decision: IP Validation
**Choice**: Use custom `isValidIPOrCIDR` helper with regex patterns for IPv4 and CIDR blocks.
**Rationale**: Bun/Node's `net` module provides `isIP()` but not CIDR validation. A focused helper keeps dependencies minimal per Constitution Principle X.
**Alternatives considered**: External validation libraries (unnecessary dependency), relying only on Zod regex (less defense in depth).

## Decision: YAML Manipulation
**Choice**: `yaml` npm package
**Rationale**: Well-maintained, lightweight YAML parser/serializer with good TypeScript support.
**Alternatives considered**: `js-yaml` (larger), `@std/yaml` (Deno-focused).

## Decision: Content-Return Strategy
**Choice**: Tool receives `existing_yaml` content, returns updated YAML content. Agent handles file I/O.
**Rationale**: MCP server runs in container with isolated filesystem. Agent has direct repository access. This keeps tool stateless and side-effect free per Constitution Principle V (Default Read-Only).
**Alternatives considered**: Volume mounts (requires Docker config), filesystem writes (violates read-only principle).

## Decision: Schema Validation
**Choice**: Zod schemas with regex patterns for IP validation.
**Rationale**: Per Constitution Principle III, every input MUST be validated using Zod. Regex patterns in schema enable client-side validation before tool execution.
**Alternatives considered**: Runtime-only validation (misses client-side opportunity).

## Decision: Tool Registration
**Choice**: `registerAzureFirewallTool(registry: Registry)` in `src/tools/azure-firewall.ts`.
**Rationale**: Follows pattern established by `registerHealthTool` in Feature 001.

## Decision: Structured Response
**Choice**: Return JSON object with `yaml_content`, `filename`, `action`, `message`.
**Rationale**: Enables agent to understand outcome and take appropriate action (write file, notify user, etc.).
