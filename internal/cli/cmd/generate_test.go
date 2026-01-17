package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestGenerateCommand(t *testing.T) {
	root := &cobra.Command{Use: "platform"}
	root.AddCommand(generateCmd)

	cases := []struct {
		name          string
		args          []string
		wantErr       bool
		expectedFiles []string
	}{
		{
			name:    "basic-generate",
			args:    []string{"generate", "workflows", "--project-name", "test-project"},
			wantErr: false,
			expectedFiles: []string{
				".github/workflows/go.yaml",
			},
		},
		{
			name:    "typescript-workflow",
			args:    []string{"generate", "workflows", "--project-name", "ts-project", "--workflow-type", "typescript"},
			wantErr: false,
			expectedFiles: []string{
				".github/workflows/typescript.yaml",
			},
		},
		{
			name:    "python-workflow",
			args:    []string{"generate", "workflows", "--project-name", "py-project", "--workflow-type", "python"},
			wantErr: false,
			expectedFiles: []string{
				".github/workflows/python.yaml",
			},
		},
		{
			name:    "node-alias-workflow",
			args:    []string{"generate", "workflows", "--project-name", "node-project", "--workflow-type", "node"},
			wantErr: false,
			expectedFiles: []string{
				".github/workflows/typescript.yaml",
			},
		},
		{
			name:    "invalid-type",
			args:    []string{"generate", "workflows", "--workflow-type", "invalid"},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "platform-test-*")
			if err != nil {
				t.Fatalf("failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			buf := new(bytes.Buffer)
			root.SetOut(buf)
			root.SetErr(buf)

			// Add output dir to args
			args := append(tc.args, "--output", tmpDir)
			root.SetArgs(args)

			err = root.Execute()
			if (err != nil) != tc.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr && len(tc.expectedFiles) > 0 {
				for _, f := range tc.expectedFiles {
					path := filepath.Join(tmpDir, f)
					if _, err := os.Stat(path); os.IsNotExist(err) {
						t.Errorf("Expected file %s to exist", path)
					}
				}
			}
		})
	}
}
