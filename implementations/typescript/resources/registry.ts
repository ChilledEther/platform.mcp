
import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import type { ResourceDefinition, ResourceContext } from './types.js';

/**
 * Registry wraps an MCP server and provides type-safe resource registration
 * with UUID tracking for each access
 */
export class ResourceRegistry {
  constructor(private server: McpServer) {}

  /**
   * Register a resource with the server
   */
  register(definition: ResourceDefinition): void {
    const { name, uri, mimeType, description, handler } = definition;

    this.server.resource(
      name,
      uri,
      { description, mimeType },
      async (url: URL) => {
        const executionId = crypto.randomUUID();
        console.error(`[Resource:${name}] [ID:${executionId}] Reading resource...`);

        try {
          const context: ResourceContext = { executionId };
          const result = await handler(url, context);
          
          console.error(`[Resource:${name}] [ID:${executionId}] Read completed successfully.`);
          
          return {
            contents: result.contents,
          };
        } catch (error) {
           const errorMessage = error instanceof Error ? error.message : String(error);
           console.error(`[Resource:${name}] [ID:${executionId}] Read failed: ${errorMessage}`);
           throw error;
        }
      }
    );
  }

  /**
   * Get the underlying MCP server
   */
  getServer(): McpServer {
    return this.server;
  }
}
