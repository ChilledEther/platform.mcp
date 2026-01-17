package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jjr/mcp.github.agentic/internal/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"gopkg.in/yaml.v3"
)

// IPv4/CIDR regex pattern for schema validation (FR-FW-007)
const IPCIDRPattern = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))?$`

// FirewallRuleInput represents the input parameters for the tool.
// Matches TypeScript FirewallRuleSchema.
type FirewallRuleInput struct {
	Team         string `json:"team"`
	Source       string `json:"source"`
	Destination  string `json:"destination"`
	Port         int    `json:"port,omitempty"`
	ExistingYAML string `json:"existing_yaml,omitempty"`
}

// FirewallRule represents a single whitelist entry in the Azure configuration.
// Used internally for YAML serialization.
type FirewallRule struct {
	Team        string `yaml:"team" json:"team"`
	Source      string `yaml:"source" json:"source"`
	Destination string `yaml:"destination" json:"destination"`
	Port        int    `yaml:"port" json:"port"`
}

// FirewallConfig represents the top-level structure of the azure-firewall-rules.yaml file.
type FirewallConfig struct {
	Rules []FirewallRule `yaml:"rules"`
}

// ToolAction represents the action taken by the tool (FR-FW-005)
type ToolAction string

const (
	ActionCreated           ToolAction = "created"
	ActionUpdated           ToolAction = "updated"
	ActionDuplicateDetected ToolAction = "duplicate_detected"
)

// FirewallToolResponse is the structured response interface (FR-FW-005).
// Matches TypeScript FirewallToolResponse.
type FirewallToolResponse struct {
	YAMLContent string     `json:"yaml_content"`
	Filename    string     `json:"filename"`
	Action      ToolAction `json:"action"`
	Message     string     `json:"message"`
}

// NewAzureFirewallRulesHandler returns a handler implementing the content-return pattern.
// Tool does NOT write to filesystem - returns YAML for agent to write (FR-FW-004).
func NewAzureFirewallRulesHandler() func(context.Context, *mcp.CallToolRequest, FirewallRuleInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input FirewallRuleInput) (*mcp.CallToolResult, any, error) {
		const filename = "azure-firewall-rules.yaml"

		// Validate IPs (FR-FW-006) - double validation for security
		if !utils.IsValidIPOrCIDR(input.Source) {
			return &mcp.CallToolResult{
				IsError: true,
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Invalid source IP or CIDR: %s", input.Source)},
				},
			}, nil, nil
		}
		if !utils.IsValidIPOrCIDR(input.Destination) {
			return &mcp.CallToolResult{
				IsError: true,
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Invalid destination IP or CIDR: %s", input.Destination)},
				},
			}, nil, nil
		}

		// Parse existing YAML or start fresh (FR-FW-008)
		config := FirewallConfig{Rules: []FirewallRule{}}
		existingYAML := strings.TrimSpace(input.ExistingYAML)

		if existingYAML != "" {
			if err := yaml.Unmarshal([]byte(existingYAML), &config); err != nil {
				// FR-FW-008: Return descriptive parse error for invalid YAML
				return &mcp.CallToolResult{
					IsError: true,
					Content: []mcp.Content{
						&mcp.TextContent{Text: fmt.Sprintf("Failed to parse existing_yaml: %v. Please ensure the content is valid YAML.", err)},
					},
				}, nil, nil
			}
			// If parsed but rules is nil, initialize it
			if config.Rules == nil {
				config.Rules = []FirewallRule{}
			}
		}

		// Set default port (FR-FW-011)
		port := input.Port
		if port == 0 {
			port = 443
		}

		// Check for duplicates (FR-FW-012, FR-FW-013)
		for _, rule := range config.Rules {
			if rule.Team == input.Team &&
				rule.Source == input.Source &&
				rule.Destination == input.Destination &&
				rule.Port == port {
				// FR-FW-014: Return existing YAML unchanged with duplicate_detected action
				yamlOutput, err := formatYAML(config)
				if err != nil {
					return &mcp.CallToolResult{
						IsError: true,
						Content: []mcp.Content{
							&mcp.TextContent{Text: fmt.Sprintf("Failed to encode YAML: %v", err)},
						},
					}, nil, nil
				}

				response := FirewallToolResponse{
					YAMLContent: yamlOutput,
					Filename:    filename,
					Action:      ActionDuplicateDetected,
					Message:     fmt.Sprintf("Rule already exists for team %q (source: %s, destination: %s, port: %d). No changes made.", input.Team, input.Source, input.Destination, port),
				}

				responseJSON, _ := json.MarshalIndent(response, "", "  ")
				return &mcp.CallToolResult{
					Content: []mcp.Content{
						&mcp.TextContent{Text: string(responseJSON)},
					},
				}, nil, nil
			}
		}

		// Determine action based on whether rules existed
		action := ActionUpdated
		if len(config.Rules) == 0 {
			action = ActionCreated
		}

		// Append new rule (FR-FW-010)
		config.Rules = append(config.Rules, FirewallRule{
			Team:        input.Team,
			Source:      input.Source,
			Destination: input.Destination,
			Port:        port,
		})

		// Serialize to YAML (FR-FW-009)
		yamlOutput, err := formatYAML(config)
		if err != nil {
			return &mcp.CallToolResult{
				IsError: true,
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Failed to encode YAML: %v", err)},
				},
			}, nil, nil
		}

		// Construct message matching TypeScript format
		actionVerb := "added"
		if action == ActionCreated {
			actionVerb = "created"
		}

		// Return structured response (FR-FW-005)
		response := FirewallToolResponse{
			YAMLContent: yamlOutput,
			Filename:    filename,
			Action:      action,
			Message:     fmt.Sprintf("Successfully %s firewall rule for team %q. Write the yaml_content to %s at your repository root.", actionVerb, input.Team, filename),
		}

		responseJSON, _ := json.MarshalIndent(response, "", "  ")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: string(responseJSON)},
			},
		}, nil, nil
	}
}

// formatYAML formats FirewallConfig to YAML with consistent styling matching TypeScript.
func formatYAML(config FirewallConfig) (string, error) {
	var buf strings.Builder
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(config); err != nil {
		return "", err
	}

	// Adjust indentation to match expected flat format (matching TypeScript)
	yamlOutput := buf.String()
	yamlOutput = strings.ReplaceAll(yamlOutput, "\n  - ", "\n- ")
	yamlOutput = strings.ReplaceAll(yamlOutput, "\n    ", "\n  ")

	return yamlOutput, nil
}

// RegisterAzureFirewallTool registers the new_azure_firewall_rules tool with the registry.
// Matches TypeScript registerAzureFirewallTool (FR-FW-001, FR-FW-015, FR-FW-016).
func RegisterAzureFirewallTool(r *Registry) {
	Register(r, &mcp.Tool{
		Name: "new_azure_firewall_rules",
		Description: `Adds a new Azure firewall whitelist rule and returns updated YAML content.

IMPORTANT - Content-Return Pattern:
This tool does NOT write files directly. You (the agent) MUST:
1. FIRST, determine the repository root. If working in a multi-repo workspace, ASK the user which repository to use.
2. Read the existing 'azure-firewall-rules.yaml' file from the repository root (if it exists).
3. Pass the file content as the 'existing_yaml' parameter (or empty string if file doesn't exist).
4. After receiving the response, write the 'yaml_content' from the response to 'azure-firewall-rules.yaml' at the repository root.

NEVER guess or infer data from the project structure. If the user's request is missing required parameters (team, source, destination), you MUST ASK the user for clarification.

Use the 'Azure Firewall Rule Schema' resource (mcp://azure-firewall/schema) for the expected YAML format.`,
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"team": map[string]any{
					"type":        "string",
					"minLength":   1,
					"description": "The name of the team requesting the rule.",
				},
				"source": map[string]any{
					"type":        "string",
					"description": "The source IP address or CIDR block.",
					"pattern":     IPCIDRPattern,
				},
				"destination": map[string]any{
					"type":        "string",
					"description": "The destination IP address or CIDR block.",
					"pattern":     IPCIDRPattern,
				},
				"port": map[string]any{
					"type":        "integer",
					"description": "The destination port (defaults to 443).",
					"minimum":     1,
					"maximum":     65535,
					"default":     443,
				},
				"existing_yaml": map[string]any{
					"type":        "string",
					"description": "Current content of azure-firewall-rules.yaml. Pass empty string if file does not exist. The agent MUST read this file from the repository root before calling this tool.",
					"default":     "",
				},
			},
			"required": []string{"team", "source", "destination"},
		},
	}, NewAzureFirewallRulesHandler())
}

// GetAzureFirewallSchema returns the Azure Firewall schema as JSON for the MCP resource (FR-FW-017).
// Matches TypeScript getAzureFirewallSchema.
func GetAzureFirewallSchema() string {
	schema := map[string]any{
		"$schema":     "https://json-schema.org/draft/2020-12/schema",
		"title":       "Azure Firewall Rules Configuration",
		"description": "Schema for azure-firewall-rules.yaml file",
		"type":        "object",
		"properties": map[string]any{
			"rules": map[string]any{
				"type":        "array",
				"description": "Array of firewall whitelist rules",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"team": map[string]any{
							"type":        "string",
							"description": "Name of the team requesting the rule",
						},
						"source": map[string]any{
							"type":        "string",
							"description": "Source IP address or CIDR block",
							"pattern":     IPCIDRPattern,
						},
						"destination": map[string]any{
							"type":        "string",
							"description": "Destination IP address or CIDR block",
							"pattern":     IPCIDRPattern,
						},
						"port": map[string]any{
							"type":        "integer",
							"description": "Destination port number",
							"minimum":     1,
							"maximum":     65535,
							"default":     443,
						},
					},
					"required": []string{"team", "source", "destination"},
				},
			},
		},
		"required": []string{"rules"},
	}

	jsonBytes, _ := json.MarshalIndent(schema, "", "  ")
	return string(jsonBytes)
}
