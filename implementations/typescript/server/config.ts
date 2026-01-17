/**
 * Server configuration handling
 * Port of internal/server/config.go
 */

import { parseArgs } from 'util';

export type TransportMode = 'stdio' | 'http';

export interface ServerConfig {
  transport: TransportMode;
  addr: string;
}

/**
 * Parse command line arguments and environment variables to create config
 */
export function createConfig(): ServerConfig {
  // Parse CLI arguments
  const { values } = parseArgs({
    args: Bun.argv.slice(2),
    options: {
      transport: {
        type: 'string',
        default: 'stdio',
      },
      addr: {
        type: 'string',
        default: ':8080',
      },
    },
    strict: false,
    allowPositionals: true,
  });

  // Environment variables override CLI flags
  const transport = (process.env['MCP_TRANSPORT'] || String(values.transport ?? 'stdio')) as TransportMode;
  const addr = process.env['MCP_ADDR'] || String(values.addr ?? ':8080');

  // Validate transport mode
  if (transport !== 'stdio' && transport !== 'http') {
    throw new Error(`Invalid transport mode: ${transport}. Must be 'stdio' or 'http'`);
  }

  console.error(`Current working directory: ${process.cwd()}`);

  return {
    transport,
    addr,
  };
}
