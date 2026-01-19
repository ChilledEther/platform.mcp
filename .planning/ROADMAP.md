# Roadmap: Platform MCP

## Overview

We will build a robust platform tooling solution starting with a pure Go core library for file generation. This foundation will power two consumers: a CLI for direct usage and an MCP server for AI agent integration. Finally, we'll containerize both applications and establish a fully automated CI/CD pipeline for releases and publishing.

## Domain Expertise

None

## Phases

- [x] **Phase 1: Core Foundation** - Implement pure Go generation logic (`pkg/scaffold`) with dynamic, manifest-driven templates.
- [ ] **Phase 2: Platform CLI** - Build `platform-cli` with Cobra, wiring up core logic to disk I/O. (In progress)
- [x] **Phase 3: Platform MCP** - Build `platform-mcp` with `go-sdk`, exposing generation as MCP tools.
- [ ] **Phase 4: Docker Environment** - Create optimized multi-stage Alpine Dockerfile for the MCP server.
- [ ] **Phase 5: Deployment Pipeline** - Set up GitHub Actions, Release Please, and container publishing.

## Phase Details

### Phase 1: Core Foundation
**Goal**: Implement pure Go generation logic (`pkg/scaffold`) with dynamic, manifest-driven templates.
**Depends on**: Nothing
**Research**: Unlikely (standard Go, YAML parsing, manifest pattern)
**Plans**: 4 plans

Plans:
- [x] 01-01: Implement Core Logic (TDD) - Create types, config, and pure generation logic.
- [x] 01-02: Embed Templates - Add Go `embed` support and template parsing.
- [x] 01-03: Implement Specific Generators - Add logic for Actions, Docker, and FluxCD files.
- [x] 01-04: Dynamic Manifest & External Templates - Implement "drop-in" YAML support and central manifest mapping.

### Phase 2: Platform CLI
**Goal**: Build `platform-cli` with Cobra, wiring up core logic to disk I/O.
**Depends on**: Phase 1
**Research**: Unlikely (Cobra is standard, I/O patterns established)
**Plans**: 2 plans

Plans:
- [x] 02-01: CLI Skeleton & I/O - Set up Cobra and file writing utilities.
- [ ] 02-02: Wiring Commands - Connect CLI commands to core generation logic. (In progress)

### Phase 3: Platform MCP
**Goal**: Build `platform-mcp` with `go-sdk`, exposing generation as MCP tools.
**Depends on**: Phase 1
**Research**: Likely (MCP protocol implementation details)
**Research topics**: `modelcontextprotocol/go-sdk` usage patterns, tool schema definition
**Plans**: 2 plans

Plans:
- [x] 03-01: MCP Server Setup - Initialize server with stdio transport.
- [x] 03-02: Tool Implementation - Expose core generation logic as MCP tools.

### Phase 4: Docker Environment
**Goal**: Create optimized multi-stage Alpine Dockerfile for the MCP server.
**Depends on**: Phase 2, Phase 3
**Research**: Unlikely (standard multi-stage Alpine builds)
**Plans**: 1 plan

Plans:
- [ ] 04-01: Containerization - Create Dockerfile and build script for MCP server only.

### Phase 5: Deployment Pipeline
**Goal**: Set up GitHub Actions, Release Please, and container publishing.
**Depends on**: Phase 4
**Research**: Unlikely (standard GitHub Actions and Release Please)
**Plans**: 2 plans

Plans:
- [ ] 05-01: CI Workflows - Add linting, testing, and build verification.
- [ ] 05-02: Release & Publish - Configure Release Please and container registry publishing.

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2 → 3 → 4 → 5

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Core Foundation | 4/4 | Complete | 2026-01-19 |
| 2. Platform CLI | 1/2 | In progress | - |
| 3. Platform MCP | 2/2 | Complete | 2026-01-19 |
| 4. Docker Environment | 0/1 | Not started | - |
| 5. Deployment Pipeline | 0/2 | Not started | - |
