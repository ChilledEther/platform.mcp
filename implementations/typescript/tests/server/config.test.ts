import { describe, expect, test, afterEach } from 'bun:test';
import { createConfig } from '../../server/config';

describe('createConfig', () => {
  const originalEnv = process.env;

  afterEach(() => {
    process.env = originalEnv;
  });

  test('should return default config when no args or env vars', () => {
    // We can't easily mock Bun.argv in the test process without side effects,
    // but we can ensure env vars are cleared.
    process.env = { ...originalEnv };
    delete process.env.MCP_TRANSPORT;
    delete process.env.MCP_ADDR;

    const config = createConfig();
    
    // Default values from src/server/config.ts
    expect(config.transport).toBe('stdio');
    expect(config.addr).toBe(':8080');
  });

  test('should prioritize environment variables over defaults', () => {
    process.env = {
      ...originalEnv,
      MCP_TRANSPORT: 'http',
      MCP_ADDR: ':9090',
    };

    const config = createConfig();
    
    expect(config.transport).toBe('http');
    expect(config.addr).toBe(':9090');
  });

  test('should throw error for invalid transport mode', () => {
    process.env = {
      ...originalEnv,
      MCP_TRANSPORT: 'invalid' as any,
    };

    expect(() => createConfig()).toThrow('Invalid transport mode: invalid');
  });
});
