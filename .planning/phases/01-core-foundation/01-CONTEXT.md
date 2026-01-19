# Phase 01: Core Foundation (Extended) - Context

**Gathered:** 2026-01-19
**Status:** Ready for planning

<vision>
## How This Should Work

The Core Library (`pkg/scaffold`) should serve as a low-barrier engine for Platform Engineers to manage and distribute scaffolding templates. The vision is to move away from requiring Go knowledge to add new capabilities, making the platform truly "pluggable."

Platform Engineers should be able to simply "drop" new YAML templates into a designated directory. The Core logic then handles the discovery, parsing, and rendering of these files based on a central manifest, regardless of whether the consumer is the CLI or the MCP server.

</vision>

<essential>
## What Must Be Nailed

- **Zero-code template additions**: Platform engineers should never have to touch Go code to add a new template.
- **Unified Engine**: The same "drop file and configure" logic must power all consumer applications (CLI, MCP, etc.).
- **Dynamic Discovery**: The core library must be able to load and validate these external templates at runtime.

</essential>

<specifics>
## Specific Ideas

- **Drop-in YAML approach**: Templates are added as YAML files into a specific directory.
- **Central Manifest**: A configuration YAML (e.g., `templates.yaml`) that maps these template files to their intended destinations, roles, and variables.
- **Source of Truth**: The core library prioritizes or integrates these external templates alongside any embedded defaults.

</specifics>

<notes>
## Additional Context

This is a shift from purely hardcoded/embedded templates to a dynamic, configuration-driven engine. This impacts the architecture of `pkg/scaffold` significantly and must be resolved before finalizing the CLI or MCP consumers.

</notes>

---

*Phase: 01-core-foundation*
*Context gathered: 2026-01-19*
