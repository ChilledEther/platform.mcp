package mcp

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestHandleGenerateWorkflows_Concurrency(t *testing.T) {
	// Task 10: Implement concurrent load test for generate_workflows
	// We want to verify that the handler can handle 10+ parallel requests.
	// We'll use 50 to be safe and robust.
	concurrency := 50
	var wg sync.WaitGroup
	wg.Add(concurrency)

	errChan := make(chan error, concurrency)

	for i := 0; i < concurrency; i++ {
		go func(id int) {
			defer wg.Done()

			input := GenerateWorkflowsInput{
				ProjectName:  fmt.Sprintf("project-%d", id),
				WorkflowType: "go",
				UseDocker:    true,
			}

			// Call the handler
			res, _, err := HandleGenerateWorkflows(context.Background(), &mcp.CallToolRequest{}, input)
			if err != nil {
				errChan <- fmt.Errorf("worker %d failed: %w", id, err)
				return
			}

			if res.IsError {
				errChan <- fmt.Errorf("worker %d returned error result", id)
				return
			}

			// Check if we got the expected content
			if len(res.Content) < 2 {
				errChan <- fmt.Errorf("worker %d returned insufficient content: %d items", id, len(res.Content))
				return
			}

			// Basic content verification
			hasWorkflow := false
			hasDockerfile := false
			for _, c := range res.Content {
				textContent, ok := c.(*mcp.TextContent)
				if !ok {
					continue
				}
				if strings.Contains(textContent.Text, ".github/workflows/go.yaml") {
					hasWorkflow = true
				}
				if strings.Contains(textContent.Text, "Dockerfile") {
					hasDockerfile = true
				}
			}

			if !hasWorkflow {
				errChan <- fmt.Errorf("worker %d missing workflow file", id)
			}
			if !hasDockerfile {
				errChan <- fmt.Errorf("worker %d missing Dockerfile", id)
			}

		}(i)
	}

	wg.Wait()
	close(errChan)

	// Collect errors
	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		t.Errorf("Concurrency test failed with %d errors:", len(errors))
		for _, err := range errors {
			t.Errorf("  %v", err)
		}
	}
}
