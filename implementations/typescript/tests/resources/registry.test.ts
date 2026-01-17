
import { describe, expect, test, mock } from 'bun:test';
import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js';
import { ResourceRegistry } from '../../resources/registry.js';
import type { ResourceDefinition } from '../../resources/types.js';

describe('ResourceRegistry', () => {
  test('registers a resource with the server', () => {
    // Mock McpServer
    const mockResource = mock((_name: string, _uri: string, _metadata: any, _handler: any) => ({}));
    const mockServer = {
      resource: mockResource,
    } as unknown as McpServer;

    const registry = new ResourceRegistry(mockServer);

    const handler = async () => ({ contents: [] });

    const definition: ResourceDefinition = {
      name: 'test-resource',
      uri: 'test://resource',
      mimeType: 'text/plain',
      description: 'Test Resource',
      handler,
    };

    registry.register(definition);

    expect(mockResource).toHaveBeenCalled();
    expect(mockResource).toHaveBeenCalledWith(
      'test-resource',
      'test://resource',
      { description: 'Test Resource', mimeType: 'text/plain' },
      expect.any(Function)
    );
  });
});
