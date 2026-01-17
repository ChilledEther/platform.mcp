/**
 * Tool Registry - Type-safe wrapper for MCP tool registration
 * Port of internal/tools/registry.go
 */

import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import type { z } from 'zod';

export interface ToolContext {
  executionId: string;
}

export interface ToolResult {
  content: Array<{ type: 'text'; text: string }>;
  isError?: boolean;
  meta?: Record<string, unknown>;
}

export type ToolHandler<TInput> = (
  input: TInput,
  context: ToolContext
) => Promise<ToolResult>;

export interface ToolDefinition<TInput> {
  name: string;
  description: string;
  schema: z.ZodType<TInput>;
  handler: ToolHandler<TInput>;
}

/**
 * Registry wraps an MCP server and provides type-safe tool registration
 * with UUID tracking for each execution
 */
export class Registry {
  constructor(private server: McpServer) {}

  /**
   * Register a tool with the server, wrapping the handler with UUID tracking
   */
  register<TInput>(definition: ToolDefinition<TInput>): void {
    const { name, description, schema, handler } = definition;

    this.server.tool(
      name,
      description,
      // Convert Zod schema to JSON schema shape for MCP
      this.zodToMcpShape(schema),
      async (args: Record<string, unknown>) => {
        const executionId = crypto.randomUUID();
        console.error(`[Tool:${name}] [ID:${executionId}] Starting execution...`);

        try {
          // Validate and parse input using Zod
          const parseResult = schema.safeParse(args);
          if (!parseResult.success) {
            console.error(`[Tool:${name}] [ID:${executionId}] Validation failed: ${parseResult.error.message}`);
            return {
              content: [{ type: 'text' as const, text: `Validation error: ${parseResult.error.message}` }],
              isError: true,
            };
          }

          const input = parseResult.data;
          const context: ToolContext = { executionId };
          const result = await handler(input, context);

          // Add execution ID to result metadata
          const meta = { ...result.meta, execution_id: executionId };

          console.error(`[Tool:${name}] [ID:${executionId}] Execution completed successfully.`);
          return {
            ...result,
            meta,
          };
        } catch (error) {
          const errorMessage = error instanceof Error ? error.message : String(error);
          console.error(`[Tool:${name}] [ID:${executionId}] Execution failed: ${errorMessage}`);
          return {
            content: [{ type: 'text' as const, text: `Error: ${errorMessage}` }],
            isError: true,
          };
        }
      }
    );
  }

  /**
   * Convert a Zod schema to MCP-compatible shape object
   * MCP SDK expects a plain object shape, not a JSON Schema
   */
  private zodToMcpShape(schema: z.ZodType<unknown>): Record<string, unknown> {
    // For object schemas, extract the shape
    if ('shape' in schema && typeof schema.shape === 'object') {
      return schema.shape as Record<string, unknown>;
    }
    // Return empty shape for non-object schemas
    return {};
  }

  /**
   * Get the underlying MCP server
   */
  getServer(): McpServer {
    return this.server;
  }
}
