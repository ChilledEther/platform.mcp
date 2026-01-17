/**
 * Health Check Tool
 * Port of internal/tools/health.go
 */

import { z } from 'zod';
import type { Registry, ToolHandler, ToolResult } from './registry.js';

// Empty schema for health check - no input required
export const HealthCheckSchema = z.object({});

export type HealthCheckInput = z.infer<typeof HealthCheckSchema>;

/**
 * Health check handler - returns OK status
 */
export const healthCheckHandler: ToolHandler<HealthCheckInput> = async (
  _input,
  _context
): Promise<ToolResult> => {
  return {
    content: [{ type: 'text', text: 'OK' }],
  };
};

/**
 * Register the health_check tool with the registry
 */
export function registerHealthTool(registry: Registry): void {
  registry.register({
    name: 'health_check',
    description: 'Basic health check to verify server is operational',
    schema: HealthCheckSchema,
    handler: healthCheckHandler,
  });
}
