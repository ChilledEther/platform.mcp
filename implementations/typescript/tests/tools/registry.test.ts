import { describe, expect, test, spyOn, beforeEach } from 'bun:test';
import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import { Registry } from '../../tools/registry';
import { z } from 'zod';

describe('Registry', () => {
  let server: McpServer;
  let registry: Registry;

  beforeEach(() => {
    server = new McpServer({
      name: 'test-server',
      version: '0.1.0',
    });
    registry = new Registry(server);
    spyOn(server, 'tool');
  });

  test('should register a tool with the MCP server', () => {
    registry.register({
      name: 'test_tool',
      description: 'A test tool',
      schema: z.object({
        input: z.string(),
      }),
      handler: async (input) => ({
        content: [{ type: 'text', text: `Hello ${input.input}` }],
      }),
    });

    expect(server.tool).toHaveBeenCalled();
    const callArgs = (server.tool as any).mock.calls[0];
    if (!callArgs) throw new Error('server.tool.mock.calls[0] is undefined');
    expect(callArgs[0]).toBe('test_tool');
    expect(callArgs[1]).toBe('A test tool');
  });

  test('should execute tool handler with validated input', async () => {
    let capturedInput: any = null;
    
    registry.register({
      name: 'test_tool',
      description: 'A test tool',
      schema: z.object({
        name: z.string(),
      }),
      handler: async (input) => {
        capturedInput = input;
        return {
          content: [{ type: 'text', text: `Hello ${input.name}` }],
        };
      },
    });

    // We need to trigger the tool execution. 
    // Since we can't easily trigger it through McpServer without full transport setup,
    // we'll rely on the fact that we're testing the Registry wrapper logic.
    // The registry.register calls server.tool(name, description, shape, callback)
    const toolCall = (server.tool as any).mock.calls[0][3];
    
    const result = await toolCall({ name: 'World' });
    
    expect(capturedInput).toEqual({ name: 'World' });
    expect(result.content[0].text).toBe('Hello World');
    expect(result.meta.execution_id).toBeDefined();
  });

  test('should return validation error for invalid input', async () => {
    registry.register({
      name: 'test_tool',
      description: 'A test tool',
      schema: z.object({
        age: z.number(),
      }),
      handler: async () => ({
        content: [{ type: 'text', text: 'success' }],
      }),
    });

    const toolCall = (server.tool as any).mock.calls[0][3];
    
    const result = await toolCall({ age: 'not a number' });
    
    expect(result.isError).toBe(true);
    expect(result.content[0].text).toContain('Validation error');
  });

  test('should handle handler exceptions gracefully', async () => {
    registry.register({
      name: 'fail_tool',
      description: 'A tool that fails',
      schema: z.object({}),
      handler: async () => {
        throw new Error('Planned failure');
      },
    });

    const toolCall = (server.tool as any).mock.calls[0][3];
    
    const result = await toolCall({});
    
    expect(result.isError).toBe(true);
    expect(result.content[0].text).toBe('Error: Planned failure');
  });
});
