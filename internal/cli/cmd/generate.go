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
	withDocker   bool
	withActions  bool
	workflowType string
	dryRun       bool
	force        bool
	outputDir    string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project scaffolds",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := scaffold.Config{
			ProjectName:  projectName,
			UseDocker:    withDocker || useDocker, // Support both for now
			WithDocker:   withDocker,
			WithActions:  withActions,
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

			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}

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

var workflowsCmd = &cobra.Command{
	Use:   "workflows",
	Short: "Generate GitHub Actions workflow files",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Delegate to generate command logic or keep independent?
		// For now, let's keep independent but using same flags to avoid breaking changes if any
		// But since we moved flags to Persistent on generate, they are available here.

		// Map legacy useDocker flag if needed
		return generateCmd.RunE(cmd, args)
	},
}

func isTerminal() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(workflowsCmd)

	// Shared flags (Persistent across generate and subcommands)
	generateCmd.PersistentFlags().StringVarP(&projectName, "project-name", "p", "", "Name of the project")
	generateCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", ".", "Target directory")
	generateCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Preview only")
	generateCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Overwrite existing files")
	generateCmd.PersistentFlags().StringVarP(&workflowType, "workflow-type", "t", "go", "Type of workflow (go, typescript, node, python)")

	// Generate specific flags
	generateCmd.Flags().BoolVar(&withDocker, "with-docker", false, "Include Dockerfile")
	generateCmd.Flags().BoolVar(&withActions, "with-actions", false, "Include GitHub Actions")

	// Legacy flags for workflows command (aliased or hidden if needed)
	// Since workflows inherits persistent flags, we don't need to re-add them.
	// But useDocker was specific to workflows.
	workflowsCmd.Flags().BoolVarP(&useDocker, "docker", "d", false, "Include Dockerfile (legacy)")
}
