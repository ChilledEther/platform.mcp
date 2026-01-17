#!/usr/bin/env bun
/**
 * MCP GitHub Agentic - Main Entry Point
 * Port of cmd/mcp-server/main.go
 */

import { createConfig } from './server/config.js';
import { MCPServer } from './server/index.js';
import { logger } from './utils/logger.js';

async function main(): Promise<void> {
  // Setup signal handlers
  let shutdownRequested = false;
  
  const handleSignal = (signal: string) => {
    if (shutdownRequested) return;
    shutdownRequested = true;
    logger.error(`Received ${signal}, shutting down...`);
    process.exit(0);
  };

  process.on('SIGINT', () => handleSignal('SIGINT'));
  process.on('SIGTERM', () => handleSignal('SIGTERM'));

  try {
    // Create configuration from CLI args and environment
    const config = createConfig();

    // Create and run the server
    const server = new MCPServer(config);
    await server.run();
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : String(error);
    logger.error(`Server exited with error: ${errorMessage}`);
    process.exit(1);
  }
}

// Run the main function
main().catch((error) => {
  logger.error('Fatal error:', error);
  process.exit(1);
});
