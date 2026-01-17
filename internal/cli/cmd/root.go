package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "platform",
	Short: "Platform CLI tool for generating project scaffolds",
	Long: `A CLI tool that generates GitHub Actions workflow YAML files and Dockerfiles 
to disk, using the shared core library.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Root flags if any
}
