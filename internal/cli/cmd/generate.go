package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/modelcontextprotocol/platform.mcp/internal/cli/io"
	"github.com/modelcontextprotocol/platform.mcp/pkg/scaffold"
	"github.com/spf13/cobra"
)

var (
	projectName  string
	useDocker    bool
	workflowType string
	dryRun       bool
	force        bool
	outputDir    string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project scaffolds",
}

var workflowsCmd = &cobra.Command{
	Use:   "workflows",
	Short: "Generate GitHub Actions workflow files",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := scaffold.Config{
			ProjectName:  projectName,
			UseDocker:    useDocker,
			WorkflowType: workflowType,
		}

		if cfg.ProjectName == "" {
			dir, _ := os.Getwd()
			cfg.ProjectName = filepath.Base(dir)
		}

		files, err := scaffold.Generate(cfg)
		if err != nil {
			return fmt.Errorf("failed to generate scaffold: %w", err)
		}

		for _, file := range files {
			targetPath := filepath.Join(outputDir, file.Path)
			if dryRun {
				fmt.Printf("[DRY RUN] Would create %s\n", targetPath)
				continue
			}

			// Task 4: Implement file writing logic with directory creation
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}

			// Task 6: Implement overwrite protection
			if _, err := os.Stat(targetPath); err == nil && !force {
				if !isTerminal() {
					return fmt.Errorf("file %s already exists, use --force to overwrite (non-interactive mode)", targetPath)
				}
				if !io.Confirm(fmt.Sprintf("? File %s already exists. Overwrite?", targetPath)) {
					fmt.Printf("Skipped %s\n", targetPath)
					continue
				}
			}

			if err := os.WriteFile(targetPath, []byte(file.Content), os.FileMode(file.Mode)); err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}
			fmt.Printf("✔ Created %s\n", targetPath)
		}

		if dryRun {
			fmt.Println("Dry run complete. No files were written.")
		} else {
			fmt.Printf("Generation complete! %d files created.\n", len(files))
		}

		return nil
	},
}

func isTerminal() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(workflowsCmd)

	workflowsCmd.Flags().StringVarP(&projectName, "project-name", "p", "", "Name of the project")
	workflowsCmd.Flags().BoolVarP(&useDocker, "docker", "d", false, "Include Dockerfile")
	workflowsCmd.Flags().StringVarP(&workflowType, "workflow-type", "t", "go", "Type of workflow (go, typescript, python)")
	workflowsCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview only")
	workflowsCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing files")
	workflowsCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Target directory")
}
