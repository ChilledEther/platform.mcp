package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCLIIntegration(t *testing.T) {
	// Build the CLI
	tmpDir, err := os.MkdirTemp("", "platform-build-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	binaryName := "platform"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	binaryPath := filepath.Join(tmpDir, binaryName)

	buildCmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/platform")
	buildCmd.Dir = "../../" // Run from worktree root
	if output, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build CLI: %v\nOutput: %s", err, string(output))
	}

	// Test generation
	workDir, err := os.MkdirTemp("", "platform-work-*")
	if err != nil {
		t.Fatalf("failed to create work dir: %v", err)
	}
	defer os.RemoveAll(workDir)

	genCmd := exec.Command(binaryPath, "generate", "workflows", "--project-name", "integration-test", "--docker", "--output", workDir)
	if output, err := genCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to run CLI: %v\nOutput: %s", err, string(output))
	}

	// Verify files
	expectedFiles := []string{
		filepath.Join(workDir, ".github/workflows/go.yaml"),
		filepath.Join(workDir, "Dockerfile"),
	}

	for _, f := range expectedFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected file %s was not created", f)
		}
	}
}
