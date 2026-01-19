---
created: 2026-01-19T00:00
title: Externalize template configuration
area: tooling
files:
  - internal/templates
---

## Problem

The current template system requires manual updates to the core application package to include or support new template files. This makes it difficult for platform engineers to add, update, or remove templates without modifying the codebase.

## Solution

Implement a configuration-driven approach:
- Use a `.yaml` file in the repository root (or a config dir) to define available templates.
- Config should specify template metadata like name, source path, and destination path (for CLI generation).
- The core application should read this config to dynamically handle templates, decoupling the template assets from the core logic.
