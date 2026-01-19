# Roadmap: Platform MCP

## Overview

We will build a robust platform tooling solution starting with a pure Go core library for file generation. This foundation will power two consumers: a CLI for direct usage and an MCP server for AI agent integration. Finally, we'll containerize both applications and establish a fully automated CI/CD pipeline for releases and publishing.

## Domain Expertise

None

## Phases

- [x] **Phase 1: Core Foundation** - Implement pure Go generation logic (`pkg/scaffold`) with dynamic, manifest-driven templates.
- [x] **Phase 2: Platform CLI** - Build `platform-cli` with Cobra, wiring up core logic to disk I/O.
- [x] **Phase 3: Platform MCP** - Build `platform-mcp` with `go-sdk`, exposing generation as MCP tools.
- [x] **Phase 4: Docker Environment** - Create optimized multi-stage Alpine Dockerfile for the MCP server.
...
### Phase 4: Docker Environment
**Goal**: Create optimized multi-stage Alpine Dockerfile for the MCP server.
**Depends on**: Phase 2, Phase 3
**Research**: Unlikely (standard multi-stage Alpine builds)
**Plans**: 1 plan

Plans:
- [x] 04-01: Containerization - Create Dockerfile and build script for MCP server only.
...
| 4. Docker Environment | 1/1 | Complete | 2026-01-19 |
| 5. Deployment Pipeline | 0/2 | Not started | - |

