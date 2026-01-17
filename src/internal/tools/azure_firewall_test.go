package tools

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestNewAzureFirewallRules_ContentReturnPattern_CreateNew(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "10.0.0.1",
		Destination:  "192.168.1.1",
		Port:         443,
		ExistingYAML: "",
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.IsError {
		t.Fatalf("Expected no result error, got true")
	}

	// Parse the response JSON
	var response FirewallToolResponse
	textContent := result.Content[0].(*mcp.TextContent)
	if err := json.Unmarshal([]byte(textContent.Text), &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if response.Action != ActionCreated {
		t.Errorf("Expected action 'created', got %s", response.Action)
	}

	if response.Filename != "azure-firewall-rules.yaml" {
		t.Errorf("Expected filename 'azure-firewall-rules.yaml', got %s", response.Filename)
	}

	if !strings.Contains(response.YAMLContent, "team: platform") {
		t.Errorf("Expected yaml_content to contain 'team: platform', got %s", response.YAMLContent)
	}

	if !strings.Contains(response.YAMLContent, "source: 10.0.0.1") {
		t.Errorf("Expected yaml_content to contain 'source: 10.0.0.1', got %s", response.YAMLContent)
	}

	if !strings.Contains(response.Message, "created") {
		t.Errorf("Expected message to contain 'created', got %s", response.Message)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_UpdateExisting(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	existingYAML := `rules:
- team: existing-team
  source: 1.1.1.1
  destination: 2.2.2.2
  port: 80
`

	input := FirewallRuleInput{
		Team:         "new-team",
		Source:       "10.0.0.1",
		Destination:  "192.168.1.1",
		Port:         443,
		ExistingYAML: existingYAML,
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.IsError {
		t.Fatalf("Expected no result error, got true")
	}

	var response FirewallToolResponse
	textContent := result.Content[0].(*mcp.TextContent)
	if err := json.Unmarshal([]byte(textContent.Text), &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if response.Action != ActionUpdated {
		t.Errorf("Expected action 'updated', got %s", response.Action)
	}

	if !strings.Contains(response.YAMLContent, "existing-team") {
		t.Errorf("Expected yaml_content to contain 'existing-team', got %s", response.YAMLContent)
	}

	if !strings.Contains(response.YAMLContent, "new-team") {
		t.Errorf("Expected yaml_content to contain 'new-team', got %s", response.YAMLContent)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_DuplicateDetected(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	existingYAML := `rules:
- team: platform
  source: 10.0.0.1
  destination: 192.168.1.1
  port: 443
`

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "10.0.0.1",
		Destination:  "192.168.1.1",
		Port:         443,
		ExistingYAML: existingYAML,
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.IsError {
		t.Fatalf("Expected no result error for duplicate detection, got true")
	}

	var response FirewallToolResponse
	textContent := result.Content[0].(*mcp.TextContent)
	if err := json.Unmarshal([]byte(textContent.Text), &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if response.Action != ActionDuplicateDetected {
		t.Errorf("Expected action 'duplicate_detected', got %s", response.Action)
	}

	if !strings.Contains(response.Message, "already exists") {
		t.Errorf("Expected message to contain 'already exists', got %s", response.Message)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_InvalidSourceIP(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "invalid-ip",
		Destination:  "192.168.1.1",
		Port:         443,
		ExistingYAML: "",
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !result.IsError {
		t.Fatalf("Expected result.IsError to be true for invalid source IP")
	}

	textContent := result.Content[0].(*mcp.TextContent)
	if !strings.Contains(textContent.Text, "Invalid source IP or CIDR") {
		t.Errorf("Expected error message about invalid source IP, got %s", textContent.Text)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_InvalidDestinationIP(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "10.0.0.1",
		Destination:  "not-an-ip",
		Port:         443,
		ExistingYAML: "",
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !result.IsError {
		t.Fatalf("Expected result.IsError to be true for invalid destination IP")
	}

	textContent := result.Content[0].(*mcp.TextContent)
	if !strings.Contains(textContent.Text, "Invalid destination IP or CIDR") {
		t.Errorf("Expected error message about invalid destination IP, got %s", textContent.Text)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_InvalidYAML(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "10.0.0.1",
		Destination:  "192.168.1.1",
		Port:         443,
		ExistingYAML: "this: is: not: valid: yaml: [[[",
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !result.IsError {
		t.Fatalf("Expected result.IsError to be true for invalid YAML")
	}

	textContent := result.Content[0].(*mcp.TextContent)
	if !strings.Contains(textContent.Text, "Failed to parse existing_yaml") {
		t.Errorf("Expected error message about parsing YAML, got %s", textContent.Text)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_ExistingYAMLWithoutRulesKey(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "10.0.0.1",
		Destination:  "192.168.1.1",
		Port:         443,
		ExistingYAML: "some_other_key: value",
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.IsError {
		t.Fatalf("Expected no result error, got true")
	}

	var response FirewallToolResponse
	textContent := result.Content[0].(*mcp.TextContent)
	if err := json.Unmarshal([]byte(textContent.Text), &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if response.Action != ActionCreated {
		t.Errorf("Expected action 'created' for YAML without rules key, got %s", response.Action)
	}

	if !strings.Contains(response.YAMLContent, "rules:") {
		t.Errorf("Expected yaml_content to contain 'rules:', got %s", response.YAMLContent)
	}
}

func TestNewAzureFirewallRules_ContentReturnPattern_DefaultPort(t *testing.T) {
	ctx := context.Background()
	req := &mcp.CallToolRequest{}

	input := FirewallRuleInput{
		Team:         "platform",
		Source:       "10.0.0.1",
		Destination:  "192.168.1.1",
		Port:         0, // 0 means not provided, should default to 443
		ExistingYAML: "",
	}

	handler := NewAzureFirewallRulesHandler()
	result, _, err := handler(ctx, req, input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.IsError {
		t.Fatalf("Expected no result error, got true")
	}

	var response FirewallToolResponse
	textContent := result.Content[0].(*mcp.TextContent)
	if err := json.Unmarshal([]byte(textContent.Text), &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if !strings.Contains(response.YAMLContent, "port: 443") {
		t.Errorf("Expected yaml_content to contain 'port: 443', got %s", response.YAMLContent)
	}
}

func TestGetAzureFirewallSchema_ValidJSONSchema(t *testing.T) {
	schemaStr := GetAzureFirewallSchema()

	var schema map[string]any
	if err := json.Unmarshal([]byte(schemaStr), &schema); err != nil {
		t.Fatalf("Failed to parse schema JSON: %v", err)
	}

	if schema["$schema"] != "https://json-schema.org/draft/2020-12/schema" {
		t.Errorf("Expected $schema to be draft/2020-12, got %v", schema["$schema"])
	}

	if schema["type"] != "object" {
		t.Errorf("Expected type to be 'object', got %v", schema["type"])
	}

	properties, ok := schema["properties"].(map[string]any)
	if !ok {
		t.Fatalf("Expected properties to be a map")
	}

	if _, exists := properties["rules"]; !exists {
		t.Errorf("Expected properties to contain 'rules'")
	}

	required, ok := schema["required"].([]any)
	if !ok {
		t.Fatalf("Expected required to be an array")
	}

	hasRules := false
	for _, r := range required {
		if r == "rules" {
			hasRules = true
			break
		}
	}
	if !hasRules {
		t.Errorf("Expected required to contain 'rules'")
	}
}

func TestGetAzureFirewallSchema_IncludesPatternValidation(t *testing.T) {
	schemaStr := GetAzureFirewallSchema()

	var schema map[string]any
	if err := json.Unmarshal([]byte(schemaStr), &schema); err != nil {
		t.Fatalf("Failed to parse schema JSON: %v", err)
	}

	properties := schema["properties"].(map[string]any)
	rules := properties["rules"].(map[string]any)
	items := rules["items"].(map[string]any)
	itemProps := items["properties"].(map[string]any)

	source := itemProps["source"].(map[string]any)
	if _, exists := source["pattern"]; !exists {
		t.Errorf("Expected source to have pattern validation")
	}

	destination := itemProps["destination"].(map[string]any)
	if _, exists := destination["pattern"]; !exists {
		t.Errorf("Expected destination to have pattern validation")
	}
}
