# Platform MCP Features 🚀

This document provides an overview of the features and capabilities of the Platform MCP project.

## 🧱 Core Foundation (001)
The shared library (`pkg/`) that powers the entire platform.
- **Scaffold Logic**: Pure functions to generate file structures.
- **Template Embedding**: Uses Go `embed` to bundle YAML templates directly into the binary.
- **TDD Driven**: 100% test coverage for all core generation logic.

## 💻 Platform CLI (002)
A human-friendly CLI tool for scaffolding projects.
- **Command Line Interface**: Built with `spf13/cobra`.
- **File I/O**: Writes generated workflows and Dockerfiles directly to the local filesystem.
- **User Friendly**: Clear error messages and usage instructions.

## 🤖 Platform MCP Server (003)
An MCP server that allows AI agents to perform scaffolding tasks.
- **MCP Protocol**: Implements the Model Context Protocol using `modelcontextprotocol/go-sdk`.
- **Tools**: Exposes `generate_workflows` and other tools to agents like Claude or Cursor.
- **Stateless**: Returns generated content as structured data rather than writing to disk directly.

## 🐳 Docker Environment (998)
Standardized containerization for all platform artifacts.
- **Multi-Stage Builds**: Optimized builds using Alpine Linux as the base.
- **Minimal Footprint**: Final images are typically under 50MB.
- **Reproducibility**: Guaranteed consistent behavior across development and production.

## 🔄 Deployment Pipeline (999)
Automated CI/CD workflows for reliable delivery.
- **Automated Testing**: Runs tests on every push and pull request.
- **Semantic Versioning**: Uses Google's `Release Please` for automated versioning and changelog generation.
- **Container Publishing**: Automatically builds and pushes images to GHCR on release.
- **Conventional Commits**: Enforces standardized commit messages for automation.
