# Quickstart: Azure Firewall Tool

## Overview
The `new_azure_firewall_rules` tool automates the management of your Azure firewall whitelist. It maintains a `azure-firewall-rules.yaml` file in the root of the repository.

## Usage

### Adding a new rule
Call the tool via your MCP client with the following parameters:

```json
{
  "name": "new_azure_firewall_rules",
  "arguments": {
    "team_name": "DevOps",
    "source_ip": "10.0.0.1",
    "destination_ip": "13.0.0.5"
  }
}
```

### Result
The file `azure-firewall-rules.yaml` will be created or updated:

```yaml
rules:
  - team: DevOps
    source: 10.0.0.1
    destination: 13.0.0.5
    port: 443
```

## Validation
- IPs must be valid (e.g., `192.168.1.1` or `10.0.0.0/24`).
- Duplicate rules (exact matches of all 4 fields) will be ignored to prevent bloat.
