import { describe, expect, test } from 'bun:test';
import { 
  FirewallRuleSchema, 
  createAzureFirewallHandler,
  getAzureFirewallSchema 
} from '../../tools/azure-firewall';

describe('FirewallRuleSchema', () => {
  test('accepts valid input with all fields', () => {
    const result = FirewallRuleSchema.safeParse({
      team: 'platform',
      source: '10.0.0.0/8',
      destination: '192.168.1.100',
      port: 443,
      existing_yaml: '',
    });
    expect(result.success).toBe(true);
    if (result.success) {
      expect(result.data.team).toBe('platform');
      expect(result.data.port).toBe(443);
    }
  });

  test('accepts input without port (uses default)', () => {
    const result = FirewallRuleSchema.safeParse({
      team: 'platform',
      source: '10.0.0.1',
      destination: '192.168.1.100',
      existing_yaml: '',
    });
    expect(result.success).toBe(true);
    if (result.success) {
      expect(result.data.port).toBe(443);
    }
  });

  test('rejects invalid port', () => {
    const result = FirewallRuleSchema.safeParse({
      team: 'platform',
      source: '10.0.0.1',
      destination: '192.168.1.100',
      port: 70000,
      existing_yaml: '',
    });
    expect(result.success).toBe(false);
  });

  test('rejects missing required fields', () => {
    const result = FirewallRuleSchema.safeParse({
      team: 'platform',
    });
    expect(result.success).toBe(false);
  });

  test('rejects invalid IP format via regex pattern', () => {
    const result = FirewallRuleSchema.safeParse({
      team: 'platform',
      source: 'not-an-ip',
      destination: '192.168.1.100',
      existing_yaml: '',
    });
    expect(result.success).toBe(false);
  });

  test('accepts CIDR notation', () => {
    const result = FirewallRuleSchema.safeParse({
      team: 'platform',
      source: '10.0.0.0/24',
      destination: '192.168.0.0/16',
      existing_yaml: '',
    });
    expect(result.success).toBe(true);
  });
});

describe('createAzureFirewallHandler - Content Return Pattern', () => {
  test('creates new YAML content when existing_yaml is empty', async () => {
    const handler = createAzureFirewallHandler();
    const result = await handler(
      {
        team: 'platform',
        source: '10.0.0.1',
        destination: '192.168.1.1',
        port: 443,
        existing_yaml: '',
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBeUndefined();
    
    const response = JSON.parse(result.content[0]?.text || '{}');
    expect(response.action).toBe('created');
    expect(response.filename).toBe('azure-firewall-rules.yaml');
    expect(response.yaml_content).toContain('team: platform');
    expect(response.yaml_content).toContain('source: 10.0.0.1');
    expect(response.message).toContain('created');
  });

  test('updates existing YAML content', async () => {
    const handler = createAzureFirewallHandler();
    
    const existingYaml = `rules:
- team: existing-team
  source: 1.1.1.1
  destination: 2.2.2.2
  port: 80
`;
    
    const result = await handler(
      {
        team: 'new-team',
        source: '10.0.0.1',
        destination: '192.168.1.1',
        port: 443,
        existing_yaml: existingYaml,
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBeUndefined();
    
    const response = JSON.parse(result.content[0]?.text || '{}');
    expect(response.action).toBe('updated');
    expect(response.yaml_content).toContain('existing-team');
    expect(response.yaml_content).toContain('new-team');
  });

  test('detects duplicate rules and returns duplicate_detected action', async () => {
    const handler = createAzureFirewallHandler();
    
    const existingYaml = `rules:
- team: platform
  source: 10.0.0.1
  destination: 192.168.1.1
  port: 443
`;
    
    const result = await handler(
      {
        team: 'platform',
        source: '10.0.0.1',
        destination: '192.168.1.1',
        port: 443,
        existing_yaml: existingYaml,
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBeUndefined();
    
    const response = JSON.parse(result.content[0]?.text || '{}');
    expect(response.action).toBe('duplicate_detected');
    expect(response.message).toContain('already exists');
  });

  test('rejects invalid source IP', async () => {
    const handler = createAzureFirewallHandler();
    const result = await handler(
      {
        team: 'platform',
        source: 'invalid-ip',
        destination: '192.168.1.1',
        port: 443,
        existing_yaml: '',
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBe(true);
    expect(result.content[0]?.text).toContain('Invalid source IP or CIDR');
  });

  test('rejects invalid destination IP', async () => {
    const handler = createAzureFirewallHandler();
    const result = await handler(
      {
        team: 'platform',
        source: '10.0.0.1',
        destination: 'not-an-ip',
        port: 443,
        existing_yaml: '',
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBe(true);
    expect(result.content[0]?.text).toContain('Invalid destination IP or CIDR');
  });

  test('returns parse error for invalid existing_yaml', async () => {
    const handler = createAzureFirewallHandler();
    const result = await handler(
      {
        team: 'platform',
        source: '10.0.0.1',
        destination: '192.168.1.1',
        port: 443,
        existing_yaml: 'this: is: not: valid: yaml: [[[',
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBe(true);
    expect(result.content[0]?.text).toContain('Failed to parse existing_yaml');
  });

  test('handles existing YAML without rules key', async () => {
    const handler = createAzureFirewallHandler();
    const result = await handler(
      {
        team: 'platform',
        source: '10.0.0.1',
        destination: '192.168.1.1',
        port: 443,
        existing_yaml: 'some_other_key: value',
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBeUndefined();
    
    const response = JSON.parse(result.content[0]?.text || '{}');
    expect(response.action).toBe('created');
    expect(response.yaml_content).toContain('rules:');
  });

  test('uses default port 443 when not specified', async () => {
    const handler = createAzureFirewallHandler();
    // Port defaults via schema, but handler receives the defaulted value
    const result = await handler(
      {
        team: 'platform',
        source: '10.0.0.1',
        destination: '192.168.1.1',
        port: 443, // Schema default applied before handler
        existing_yaml: '',
      },
      { executionId: 'test-123' }
    );

    expect(result.isError).toBeUndefined();
    
    const response = JSON.parse(result.content[0]?.text || '{}');
    expect(response.yaml_content).toContain('port: 443');
  });
});

describe('getAzureFirewallSchema', () => {
  test('returns valid JSON schema', () => {
    const schema = getAzureFirewallSchema();
    const parsed = JSON.parse(schema);
    
    expect(parsed.$schema).toBe('https://json-schema.org/draft/2020-12/schema');
    expect(parsed.type).toBe('object');
    expect(parsed.properties.rules).toBeDefined();
    expect(parsed.required).toContain('rules');
  });

  test('includes pattern validation for IP addresses', () => {
    const schema = getAzureFirewallSchema();
    const parsed = JSON.parse(schema);
    
    expect(parsed.properties.rules.items.properties.source.pattern).toBeDefined();
    expect(parsed.properties.rules.items.properties.destination.pattern).toBeDefined();
  });
});
