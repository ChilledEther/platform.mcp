package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestGenerateCommand(t *testing.T) {
	root := &cobra.Command{Use: "platform"}
	root.AddCommand(generateCmd)

	cases := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "basic-generate",
			args:    []string{"generate", "workflows", "--project-name", "test-project"},
			wantErr: false,
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
		})
	}
}
