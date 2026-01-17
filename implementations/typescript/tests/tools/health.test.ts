import { describe, expect, test } from 'bun:test';
import { healthCheckHandler, HealthCheckSchema } from '../../tools/health';

describe('HealthCheckSchema', () => {
  test('accepts empty object', () => {
    const result = HealthCheckSchema.safeParse({});
    expect(result.success).toBe(true);
  });

  test('accepts object with extra properties (ignored)', () => {
    const result = HealthCheckSchema.safeParse({ extra: 'ignored' });
    expect(result.success).toBe(true);
  });
});

describe('healthCheckHandler', () => {
  test('returns OK response', async () => {
    const result = await healthCheckHandler({}, { executionId: 'test-123' });
    expect(result.content).toHaveLength(1);
    expect(result.content[0]).toEqual({ type: 'text', text: 'OK' });
    expect(result.isError).toBeUndefined();
  });
});
