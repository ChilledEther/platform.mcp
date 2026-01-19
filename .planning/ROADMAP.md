# Roadmap: Platform MCP

## Overview

We will build a robust platform tooling solution starting with a pure Go core library for file generation. This foundation will power two consumers: a CLI for direct usage and an MCP server for AI agent integration. Finally, we'll containerize both applications and establish a fully automated CI/CD pipeline for releases and publishing.

## Domain Expertise

None

## Phases

- [ ] **Phase 1: Core Foundation** - Implement pure Go generation logic (`pkg/scaffold`) with embedded templates and TDD.
- [ ] **Phase 2: Platform CLI** - Build `platform-cli` with Cobra, wiring up core logic to disk I/O.
- [ ] **Phase 3: Platform MCP** - Build `platform-mcp` with `go-sdk`, exposing generation as MCP tools.
- [ ] **Phase 4: Docker Environment** - Create optimized multi-stage Alpine Dockerfiles for both binaries.
- [ ] **Phase 5: Deployment Pipeline** - Set up GitHub Actions, Release Please, and container publishing.

## Phase Details

### Phase 1: Core Foundation
**Goal**: Implement pure Go generation logic (`pkg/scaffold`) with embedded templates and TDD.
**Depends on**: Nothing
**Research**: Unlikely (standard Go, embed package, testing)
**Plans**: 3 plans

Plans:
- [ ] 01-01: Implement Core Logic (TDD) - Create types, config, and pure generation logic.
- [ ] 01-02: Embed Templates - Add Go `embed` support and template parsing.
- [ ] 01-03: Implement Specific Generators - Add logic for Actions, Docker, and FluxCD files.

### Phase 2: Platform CLI
**Goal**: Build `platform-cli` with Cobra, wiring up core logic to disk I/O.
**Depends on**: Phase 1
**Research**: Unlikely (Cobra is standard, I/O patterns established)
**Plans**: 2 plans

Plans:
- [ ] 02-01: CLI Skeleton & I/O - Set up Cobra and file writing utilities.
- [ ] 02-02: Wiring Commands - Connect CLI commands to core generation logic.

### Phase 3: Platform MCP
**Goal**: Build `platform-mcp` with `go-sdk`, exposing generation as MCP tools.
**Depends on**: Phase 1
**Research**: Likely (MCP protocol implementation details)
**Research topics**: `modelcontextprotocol/go-sdk` usage patterns, tool schema definition
**Plans**: 2 plans

Plans:
- [ ] 03-01: MCP Server Setup - Initialize server with stdio transport.
- [ ] 03-02: Tool Implementation - Expose core generation logic as MCP tools.

### Phase 4: Docker Environment
**Goal**: Create optimized multi-stage Alpine Dockerfiles for both binaries.
**Depends on**: Phase 2, Phase 3
**Research**: Unlikely (standard multi-stage Alpine builds)
**Plans**: 1 plan

Plans:
- [ ] 04-01: Containerization - Create Dockerfiles and build scripts for CLI and MCP.

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
| 1. Core Foundation | 0/3 | Not started | - |
| 2. Platform CLI | 0/2 | Not started | - |
| 3. Platform MCP | 0/2 | Not started | - |
| 4. Docker Environment | 0/1 | Not started | - |
| 5. Deployment Pipeline | 0/2 | Not started | - |
