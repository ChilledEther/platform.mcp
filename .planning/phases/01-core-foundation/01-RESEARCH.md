# Phase 01: Core Foundation (Extended) - Research

**Researched:** 2026-01-19
**Domain:** Dynamic, Manifest-driven Scaffolding Engine in Go
**Confidence:** HIGH

<research_summary>
## Summary

Research focused on building a "pluggable" scaffolding engine using Go's `io/fs` and `text/template` packages. The standard approach for this in the Go ecosystem involves using a unified FileSystem abstraction (`fs.FS`) to merge embedded defaults with external overrides.

Key findings suggest using a central YAML manifest to map template sources to their target destinations. This manifest acts as the "contract" between platform engineers and the engine, allowing for zero-code additions. Security must be a priority when loading external templates to prevent path traversal.

**Primary recommendation:** Use `io/fs` as the primary interface. Merge an `embed.FS` (defaults) with an `os.DirFS` (external overrides) using a library like `mergefs`. Define a clear YAML schema for the `templates.yaml` manifest.

</research_summary>

<standard_stack>
## Standard Stack

The established libraries/tools for this domain:

### Core
| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| `io/fs` | stdlib | Filesystem abstraction | The standard interface for file access since Go 1.16 |
| `text/template` | stdlib | Text rendering | Powerful, safe, and built-in Go templating |
| `gopkg.in/yaml.v3` | v3.0.1 | YAML parsing | Supports line numbers and comments, standard in Go |

### Supporting
| Library | Version | Purpose | When to Use |
|---------|---------|---------|-------------|
| `github.com/laher/mergefs` | v0.1.1 | Chaining `fs.FS` | When you need to overlay external templates on top of embedded ones |
| `github.com/spf13/afero` | v1.15.0 | Advanced VFS | Use if you need a writable virtual filesystem for complex testing |

### Alternatives Considered
| Instead of | Could Use | Tradeoff |
|------------|-----------|----------|
| `mergefs` | Custom Loop | Simple to implement manually, but `mergefs` is battle-tested |
| `afero` | `io/fs` | `io/fs` is newer and idiomatic for read-only template access |

**Installation:**
```bash
go get gopkg.in/yaml.v3
go get github.com/laher/mergefs
```
</standard_stack>

<architecture_patterns>
## Architecture Patterns

### Recommended Project Structure
```
pkg/scaffold/
├── generator.go      # Main engine logic
├── manifest.go       # YAML parsing and validation
├── registry.go       # FS merging and discovery
├── templates/        # Embedded default templates
└── templates.yaml    # Default manifest
```

### Pattern 1: Manifest-Driven Discovery
**What:** Use a central YAML manifest to define "recipes."
**When to use:** To allow non-coders to add templates.
**Example:**
```yaml
templates:
  - name: github_action
    source: templates/action.yaml.tmpl
    target: .github/workflows/{{ .Name }}.yaml
    feature_flag: use_actions
```

### Pattern 2: FS Merging (Overlay)
**What:** Chain multiple `fs.FS` instances so external files take precedence.
**When to use:** To allow Platform Engineers to override embedded templates.
**Example:**
```go
import "github.com/laher/mergefs"

// internalFS is //go:embed templates/*
// externalFS is os.DirFS("/path/to/custom/templates")
combinedFS := mergefs.Merge(externalFS, internalFS)
```

### Anti-Patterns to Avoid
- **Hardcoded template paths:** Prevents externalization.
- **Implicit discovery:** Automatically loading every file in a directory without a manifest entry (makes it hard to control destination and features).
</architecture_patterns>

<dont_hand_roll>
## Don't Hand-Roll

Problems that look simple but have existing solutions:

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| FS Overlaying | Nested if/else for file checks | `mergefs` | Handles DirEntries and Globbing correctly across multiple FS |
| YAML Parsing | Regex or manual splits | `yaml.v3` | Handles complex structures and multi-document YAML safely |
| Path Joining | Manual string concatenation | `path/filepath` | Handles OS-specific separators and cleans `..` elements |

**Key insight:** Go's `io/fs` is extremely powerful but requires discipline. Using established patterns for `fs.FS` composition prevents "leaking" I/O logic into the core generator.
</dont_hand_roll>

<common_pitfalls>
## Common Pitfalls

### Pitfall 1: Path Traversal
**What goes wrong:** A malicious template manifest sets `target: ../../../etc/passwd`.
**Why it happens:** Directly using manifest `target` strings without sanitization.
**How to avoid:** Use `filepath.Join` and verify the resulting absolute path is within the target project root.

### Pitfall 2: Naming Collisions
**What goes wrong:** Engine finds `action.yaml` in both embedded and external FS.
**Why it happens:** Improper merge order.
**How to avoid:** Always put the `os.DirFS` (external) first in the merge chain.

### Pitfall 3: Template Context Drifts
**What goes wrong:** External template uses `{{ .NewField }}` but engine context only provides `{{ .Name }}`.
**Why it happens:** Decoupling templates from Go code.
**How to avoid:** Implement a strict validation step for external templates before rendering.
</common_pitfalls>

<code_examples>
## Code Examples

### Merging Embed and OS Filesystems
```go
// Source: https://pkg.go.dev/io/fs
package scaffold

import (
    "embed"
    "io/fs"
    "os"
    "github.com/laher/mergefs"
)

//go:embed internal/templates/*
var embeddedTemplates embed.FS

func LoadFS(externalPath string) fs.FS {
    if externalPath == "" {
        return embeddedTemplates
    }
    externalFS := os.DirFS(externalPath)
    return mergefs.Merge(externalFS, embeddedTemplates)
}
```

### Manifest Schema Design (Recommended)
```yaml
version: v1
overrides:
  target_dir: . # default root
templates:
  - id: dockerfile
    path: templates/Dockerfile.tmpl
    destination: Dockerfile
    conditions:
      - feature: docker
```
</code_examples>

<sota_updates>
## State of the Art (2024-2025)

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| `os.Open` | `fs.Open` | Go 1.16 | Code is now testable with `fstest.MapFS` |
| `ioutil.ReadFile` | `os.ReadFile` or `fs.ReadFile` | Go 1.16 | `ioutil` is deprecated |
| String templates | `embed` package | Go 1.16 | Templates are now compiled into binary |

**New tools/patterns to consider:**
- **`io/fs` Sub**: Use `fs.Sub` to focus the generator on a specific subdirectory within a merged FS.
- **Functional Options**: Use for configuring the Generator (e.g., `NewGenerator(WithExternalTemplates(path))`).
</sota_updates>

<sources>
## Sources

### Primary (HIGH confidence)
- `pkg.go.dev/io/fs` - Official FS interface documentation
- `github.com/tmrts/boilr` - Reference implementation for Go scaffolding
- `github.com/go-task/task` - Reference for YAML manifest patterns in Go

### Secondary (MEDIUM confidence)
- `github.com/laher/mergefs` - Verified functionality for merging FS

</sources>

<metadata>
## Metadata

**Research scope:**
- Core technology: Go `io/fs` and `text/template`
- Ecosystem: Scaffolding tools (boilr, copier)
- Patterns: Manifest-driven generation, FS merging
- Pitfalls: Security (traversal), validation

**Confidence breakdown:**
- Standard stack: HIGH - rely on stdlib and stable YAML lib
- Architecture: HIGH - manifest pattern is industry standard
- Pitfalls: HIGH - traversal is a well-known risk in scaffolding
- Code examples: HIGH - idiomatic Go patterns

**Research date:** 2026-01-19
**Valid until:** 2026-02-19
</metadata>

---

*Phase: 01-core-foundation*
*Research completed: 2026-01-19*
*Ready for planning: yes*
