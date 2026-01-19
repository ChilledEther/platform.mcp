# Phase 02: Platform CLI - Context

**Gathered:** 2026-01-19
**Status:** Ready for planning

<vision>
## How This Should Work

The Platform CLI should serve as a low-barrier tool for Platform Engineers to manage and distribute scaffolding templates. The vision is to move away from requiring Go knowledge to add new capabilities.

Platform Engineers should be able to simply "drop" new YAML templates into a designated directory. The CLI (consuming the Phase 01 core) then handles the logic of where these files should be placed and how they should be rendered based on a central configuration.

</vision>

<essential>
## What Must Be Nailed

- **Zero-code template additions**: Platform engineers should never have to touch Go code to add a new template.
- **Directory-based workflow**: A simple "drop file and configure" experience that feels like managing a library of assets rather than maintaining a codebase.

</essential>

<specifics>
## Specific Ideas

- **Drop-in YAML approach**: Templates are added as YAML files into a specific directory.
- **Central Manifest**: A configuration YAML in the root directory that maps these template files to their intended destinations and roles.
- **Managed Repository**: This repository is the "source of truth" managed by the platform team, designed for ease of contribution by non-developers.

</specifics>

<notes>
## Additional Context

The key differentiator here is making the platform "pluggable" for the engineers who own the standards (Platform Engineers), without them needing to understand the underlying implementation in Go.

</notes>

---

*Phase: 02-platform-cli*
*Context gathered: 2026-01-19*
