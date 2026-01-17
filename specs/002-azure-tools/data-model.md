# Data Model: Azure Tools Group

This document defines data structures for all tools in the Azure Tools group.

---

# Tool: `new_azure_firewall_rules`

## Architecture Pattern: Content-Return

## Entities

### FirewallRule

Represents a single whitelist entry in the Azure configuration.

| Field | Type | Description | Validation |
|-------|------|-------------|------------|
| `team` | `string` | Name of the team requesting the rule | Required, non-empty |
| `source` | `string` | Source IP or CIDR block | Valid IPv4/v6 or CIDR |
| `destination` | `string` | Destination IP or CIDR block | Valid IPv4/v6 or CIDR |
| `port` | `int` | Port number | Default: 443, Range: 1-65535 |

### FirewallConfig

The top-level structure of the `azure-firewall-rules.yaml` file.

| Field | Type | Description |
|-------|------|-------------|
| `rules` | `[]FirewallRule` | Array of firewall rules |

### ToolInput

Input parameters for the `new_azure_firewall_rules` tool.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `team` | `string` | Yes | Name of the team requesting the rule |
| `source` | `string` | Yes | Source IP or CIDR block |
| `destination` | `string` | Yes | Destination IP or CIDR block |
| `port` | `int` | No | Port number (defaults to 443) |
| `existing_yaml` | `string` | No | Current content of azure-firewall-rules.yaml (empty if file doesn't exist) |

### ToolResponse

Structured response returned by the tool.

| Field | Type | Description |
|-------|------|-------------|
| `yaml_content` | `string` | The complete updated YAML content |
| `filename` | `string` | Expected filename: `azure-firewall-rules.yaml` |
| `action` | `string` | One of: `created`, `updated`, `duplicate_detected` |
| `message` | `string` | Human-readable description of what occurred |

## Example YAML Output

```yaml
rules:
- team: Platform
  source: 10.0.0.1
  destination: 13.0.0.1
  port: 443
- team: Security
  source: 192.168.1.0/24
  destination: 10.20.30.40
  port: 443
```

## Example Tool Response

```json
{
  "yaml_content": "rules:\n- team: Platform\n  source: 10.0.0.1\n  destination: 13.0.0.1\n  port: 443\n",
  "filename": "azure-firewall-rules.yaml",
  "action": "created",
  "message": "Successfully created firewall rule for team Platform"
}
```

---

# Tool: *(Future Tool Placeholder)*

> Data models for additional Azure tools will be documented here as they are specified.

---

# Shared Data Structures

These structures are shared across multiple tools in the Azure Tools group.

## StandardToolResponse

Base response structure that all tools should follow.

| Field | Type | Description |
|-------|------|-------------|
| `success` | `boolean` | Whether the operation succeeded |
| `message` | `string` | Human-readable description |
| `data` | `object` | Tool-specific response data |

## ErrorResponse

Standard error response structure.

| Field | Type | Description |
|-------|------|-------------|
| `success` | `false` | Always false for errors |
| `error_code` | `string` | Machine-readable error identifier |
| `message` | `string` | Human-readable error description |
| `details` | `object?` | Optional additional error context |
