// Package cmd provides command-line interface commands for the gh-utils extension.
package cmd

import (
	"fmt"
	"os"

	"github.com/lakruzz/gh-utils/internal/issue"
	"github.com/spf13/cobra"
)

var (
	issueFile string
)

var mkissueCmd = &cobra.Command{
	Use:   "mkissue",
	Short: "Create a GitHub issue from a markdown file",
	Long: `Create a GitHub issue from a markdown file.
The markdown file should contain the issue title and body.`,
	RunE: func(_ *cobra.Command, _ []string) error {
		if issueFile == "" {
			return fmt.Errorf("file flag is required")
		}

		// Check if file exists
		if _, err := os.Stat(issueFile); os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", issueFile)
		}

		// Create the issue
		if err := issue.CreateFromFile(issueFile); err != nil {
			return fmt.Errorf("failed to create issue: %w", err)
		}

		fmt.Printf("✅ Issue created successfully from %s\n", issueFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mkissueCmd)

	// Define flags for mkissue command
	mkissueCmd.Flags().StringVarP(&issueFile, "file", "f", "", "Path to the markdown file containing issue content (required)")
	_ = mkissueCmd.MarkFlagRequired("file")
}
