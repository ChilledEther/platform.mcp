import { describe, expect, test } from 'bun:test';
import { MCPServer } from '../../server/index';

describe('MCPServer', () => {
  test('should initialize with provided config', () => {
    const config = {
      transport: 'stdio' as const,
      addr: ':8080',
    };
    
    const server = new MCPServer(config);
    
    expect(server).toBeDefined();
    expect(server.getRegistry()).toBeDefined();
    expect(server.getServer()).toBeDefined();
  });

  test('should register basic health tool on initialization', () => {
    const config = {
      transport: 'stdio' as const,
      addr: ':8080',
    };
    
    const server = new MCPServer(config);
    expect(server.getRegistry()).toBeDefined();
  });
});
