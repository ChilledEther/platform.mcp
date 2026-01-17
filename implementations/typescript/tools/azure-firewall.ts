/**
 * Azure Firewall Rules Tool
 * Implements content-return pattern per spec 002-azure-tools
 */

import { z } from 'zod';
import * as yaml from 'yaml';
import { isValidIPOrCIDR } from '../utils/validation.js';
import type { Registry, ToolHandler, ToolResult } from './registry.js';

// IPv4/CIDR regex pattern for schema validation (FR-FW-007)
const IP_CIDR_PATTERN = '^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))?$';

// Firewall rule schema with validation (FR-FW-002, FR-FW-003, FR-FW-007)
export const FirewallRuleSchema = z.object({
  team: z.string().min(1).describe('The name of the team requesting the rule.'),
  source: z.string()
    .regex(new RegExp(IP_CIDR_PATTERN), 'Must be a valid IPv4 address or CIDR block')
    .describe('The source IP address or CIDR block.'),
  destination: z.string()
    .regex(new RegExp(IP_CIDR_PATTERN), 'Must be a valid IPv4 address or CIDR block')
    .describe('The destination IP address or CIDR block.'),
  port: z.number().int().min(1).max(65535).optional().default(443)
    .describe('The destination port (defaults to 443).'),
  existing_yaml: z.string().optional().default('')
    .describe('Current content of azure-firewall-rules.yaml. Pass empty string if file does not exist. The agent MUST read this file from the repository root before calling this tool.'),
});

export type FirewallRuleInput = z.infer<typeof FirewallRuleSchema>;

interface FirewallRule {
  team: string;
  source: string;
  destination: string;
  port: number;
}

interface FirewallConfig {
  rules: FirewallRule[];
}

// Response action types (FR-FW-005)
type ToolAction = 'created' | 'updated' | 'duplicate_detected';

// Structured response interface (FR-FW-005)
interface FirewallToolResponse {
  yaml_content: string;
  filename: string;
  action: ToolAction;
  message: string;
}

/**
 * Create handler for azure firewall rules using content-return pattern
 * Tool does NOT write to filesystem - returns YAML for agent to write (FR-FW-004)
 */
export function createAzureFirewallHandler(): ToolHandler<FirewallRuleInput> {
  return async (input, _context): Promise<ToolResult> => {
    const filename = 'azure-firewall-rules.yaml';
    
    // Validate IPs (FR-FW-006) - double validation for security
    if (!isValidIPOrCIDR(input.source)) {
      return {
        content: [{ type: 'text', text: `Invalid source IP or CIDR: ${input.source}` }],
        isError: true,
      };
    }
    if (!isValidIPOrCIDR(input.destination)) {
      return {
        content: [{ type: 'text', text: `Invalid destination IP or CIDR: ${input.destination}` }],
        isError: true,
      };
    }

    // Parse existing YAML or start fresh (FR-FW-008)
    let config: FirewallConfig = { rules: [] };
    const existingYaml = input.existing_yaml?.trim() || '';
    
    if (existingYaml) {
      try {
        const parsed = yaml.parse(existingYaml);
        if (parsed && typeof parsed === 'object') {
          if (Array.isArray(parsed.rules)) {
            config = parsed as FirewallConfig;
          } else {
            // YAML exists but no rules key - initialize it
            config = { rules: [] };
          }
        }
      } catch (error) {
        // FR-FW-008: Return descriptive parse error for invalid YAML
        const errorMessage = error instanceof Error ? error.message : String(error);
        return {
          content: [{ type: 'text', text: `Failed to parse existing_yaml: ${errorMessage}. Please ensure the content is valid YAML.` }],
          isError: true,
        };
      }
    }

    // Set default port (FR-FW-011)
    const port = input.port ?? 443;

    // Check for duplicates (FR-FW-012, FR-FW-013)
    for (const rule of config.rules) {
      if (
        rule.team === input.team &&
        rule.source === input.source &&
        rule.destination === input.destination &&
        rule.port === port
      ) {
        // FR-FW-014: Return existing YAML unchanged with duplicate_detected action
        const existingYamlOutput = formatYaml(config);
        const response: FirewallToolResponse = {
          yaml_content: existingYamlOutput,
          filename,
          action: 'duplicate_detected',
          message: `Rule already exists for team "${input.team}" (source: ${input.source}, destination: ${input.destination}, port: ${port}). No changes made.`,
        };
        return {
          content: [{ type: 'text', text: JSON.stringify(response, null, 2) }],
        };
      }
    }

    // Determine action based on whether rules existed
    const action: ToolAction = config.rules.length === 0 ? 'created' : 'updated';

    // Append new rule (FR-FW-010)
    config.rules.push({
      team: input.team,
      source: input.source,
      destination: input.destination,
      port,
    });

    // Serialize to YAML (FR-FW-009)
    const yamlOutput = formatYaml(config);

    // Return structured response (FR-FW-005)
    const response: FirewallToolResponse = {
      yaml_content: yamlOutput,
      filename,
      action,
      message: `Successfully ${action === 'created' ? 'created' : 'added'} firewall rule for team "${input.team}". Write the yaml_content to ${filename} at your repository root.`,
    };

    return {
      content: [{ type: 'text', text: JSON.stringify(response, null, 2) }],
    };
  };
}

/**
 * Format FirewallConfig to YAML with consistent styling
 */
function formatYaml(config: FirewallConfig): string {
  let yamlOutput = yaml.stringify(config, {
    indent: 2,
    lineWidth: 0,
  });

  // Adjust indentation to match expected flat format
  yamlOutput = yamlOutput.replace(/\n  - /g, '\n- ');
  yamlOutput = yamlOutput.replace(/\n    /g, '\n  ');

  return yamlOutput;
}

/**
 * Register the new_azure_firewall_rules tool with the registry (FR-FW-001, FR-FW-015, FR-FW-016)
 */
export function registerAzureFirewallTool(registry: Registry): void {
  registry.register({
    name: 'new_azure_firewall_rules',
    description: `Adds a new Azure firewall whitelist rule and returns updated YAML content.

IMPORTANT - Content-Return Pattern:
This tool does NOT write files directly. You (the agent) MUST:
1. FIRST, determine the repository root. If working in a multi-repo workspace, ASK the user which repository to use.
2. Read the existing 'azure-firewall-rules.yaml' file from the repository root (if it exists).
3. Pass the file content as the 'existing_yaml' parameter (or empty string if file doesn't exist).
4. After receiving the response, write the 'yaml_content' from the response to 'azure-firewall-rules.yaml' at the repository root.

NEVER guess or infer data from the project structure. If the user's request is missing required parameters (team, source, destination), you MUST ASK the user for clarification.

Use the 'Azure Firewall Rule Schema' resource (mcp://azure-firewall/schema) for the expected YAML format.`,
    schema: FirewallRuleSchema,
    handler: createAzureFirewallHandler(),
  });
}

/**
 * Get the Azure Firewall schema as JSON for the MCP resource (FR-FW-017)
 */
export function getAzureFirewallSchema(): string {
  return JSON.stringify({
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "title": "Azure Firewall Rules Configuration",
    "description": "Schema for azure-firewall-rules.yaml file",
    "type": "object",
    "properties": {
      "rules": {
        "type": "array",
        "description": "Array of firewall whitelist rules",
        "items": {
          "type": "object",
          "properties": {
            "team": { 
              "type": "string",
              "description": "Name of the team requesting the rule"
            },
            "source": {
              "type": "string",
              "description": "Source IP address or CIDR block",
              "pattern": IP_CIDR_PATTERN
            },
            "destination": {
              "type": "string",
              "description": "Destination IP address or CIDR block",
              "pattern": IP_CIDR_PATTERN
            },
            "port": {
              "type": "integer",
              "description": "Destination port number",
              "minimum": 1,
              "maximum": 65535,
              "default": 443
            }
          },
          "required": ["team", "source", "destination"]
        }
      }
    },
    "required": ["rules"]
  }, null, 2);
}
