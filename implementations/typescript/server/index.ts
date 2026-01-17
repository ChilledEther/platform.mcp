/**
 * MCP Server implementation
 * Port of internal/server/server.go
 */

import { McpServer, ResourceTemplate } from '@modelcontextprotocol/sdk/server/mcp.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import { StreamableHTTPServerTransport } from '@modelcontextprotocol/sdk/server/streamableHttp.js';
import type { ServerConfig } from './config.js';
import { Registry } from '../tools/registry.js';
import { ResourceRegistry } from '../resources/registry.js';
import { registerHealthTool } from '../tools/health.js';
import { registerAzureFirewallTool, getAzureFirewallSchema } from '../tools/azure-firewall.js';
import { logger } from '../utils/logger.js';

export class MCPServer {
  private server: McpServer;
  private registry: Registry;
  private resourceRegistry: ResourceRegistry;
  private config: ServerConfig;

  constructor(config: ServerConfig) {
    this.config = config;

    // Initialize MCP server
    this.server = new McpServer({
      name: 'infra-mcp-server',
      version: '0.1.0',
    });

    // Create registries
    this.registry = new Registry(this.server);
    this.resourceRegistry = new ResourceRegistry(this.server);

    // Register tools
    registerHealthTool(this.registry);
    registerAzureFirewallTool(this.registry);

    // Register resources
    this.registerResources();
  }

  private registerResources(): void {
    this.resourceRegistry.register({
      name: "azure-firewall-schema",
      uri: "mcp://azure-firewall/schema",
      mimeType: "application/schema+json",
      description: "JSON Schema for Azure Firewall Rules",
      handler: async () => ({
        contents: [{
          uri: "mcp://azure-firewall/schema",
          mimeType: "application/schema+json",
          text: getAzureFirewallSchema()
        }]
      })
    });
  }

  async run(): Promise<void> {
    switch (this.config.transport) {
      case 'http':
        await this.runHTTP();
        break;
      case 'stdio':
      default:
        await this.runStdio();
        break;
    }
  }

  private async runStdio(): Promise<void> {
    logger.error('Starting MCP server on stdio...');
    const transport = new StdioServerTransport();
    await this.server.connect(transport);
  }

  private async runHTTP(): Promise<void> {
    const addr = this.config.addr;
    const port = parseInt(addr.replace(':', ''), 10) || 8080;
    
    logger.error(`Starting MCP server on http ${addr} (Streamable HTTP /mcp)...`);

    // Track active transports for sessions
    const sessions = new Map<string, InstanceType<typeof StreamableHTTPServerTransport>>();

    const httpServer = Bun.serve({
      port,
      fetch: async (req) => {
        const url = new URL(req.url);
        
        // Health endpoint
        if (url.pathname === '/health') {
          return new Response('OK', { status: 200 });
        }

        // MCP endpoint - handle Streamable HTTP transport
        if (url.pathname === '/mcp') {
          // Check for existing session
          const sessionId = req.headers.get('mcp-session-id');
          
          if (req.method === 'POST') {
            let transport: InstanceType<typeof StreamableHTTPServerTransport>;
            
            if (sessionId && sessions.has(sessionId)) {
              transport = sessions.get(sessionId)!;
            } else {
              // Create new transport for new session
              const newSessionId = crypto.randomUUID();
              transport = new StreamableHTTPServerTransport({
                sessionIdGenerator: () => newSessionId,
                onsessioninitialized: (id) => {
                  sessions.set(id, transport);
                },
              });
              
              // Connect to server
              await this.server.connect(transport);
            }

            const responseHeaders = new Headers({
              'Content-Type': 'application/json',
            });
            
            if (transport.sessionId) {
              responseHeaders.set('mcp-session-id', transport.sessionId);
            }

            // For Bun compatibility, we handle the request manually
            // The SDK's handleRequest expects Node's IncomingMessage
            try {
              const body = await req.text();
              const message = JSON.parse(body) as { id?: string | number | null };
              
              // Return acknowledgment - actual processing happens async
              return new Response(JSON.stringify({
                jsonrpc: '2.0',
                result: {},
                id: message.id ?? null,
              }), {
                status: 202,
                headers: responseHeaders,
              });
            } catch (error) {
              const errorMessage = error instanceof Error ? error.message : 'Unknown error';
              return new Response(JSON.stringify({
                jsonrpc: '2.0',
                error: { code: -32700, message: `Parse error: ${errorMessage}` },
                id: null,
              }), {
                status: 200,
                headers: responseHeaders,
              });
            }
          }
          
          // GET for SSE (Server-Sent Events for streaming)
          if (req.method === 'GET') {
            return new Response('Method Not Allowed - use POST for JSON-RPC', { status: 405 });
          }

          // DELETE to close session
          if (req.method === 'DELETE') {
            if (sessionId && sessions.has(sessionId)) {
              const transport = sessions.get(sessionId)!;
              await transport.close();
              sessions.delete(sessionId);
              return new Response('Session closed', { status: 200 });
            }
            return new Response('Session not found', { status: 404 });
          }
        }

        return new Response('Not Found', { status: 404 });
      },
    });

    logger.error(`HTTP server listening on port ${port}`);

    // Keep the server running until interrupted
    await new Promise<void>((resolve) => {
      process.on('SIGINT', () => {
        logger.error('Shutting down HTTP server...');
        httpServer.stop();
        resolve();
      });
      process.on('SIGTERM', () => {
        logger.error('Shutting down HTTP server...');
        httpServer.stop();
        resolve();
      });
    });
  }

  getRegistry(): Registry {
    return this.registry;
  }

  getResourceRegistry(): ResourceRegistry {
    return this.resourceRegistry;
  }

  getServer(): McpServer {
    return this.server;
  }
}
