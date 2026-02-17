// Package cmd provides command-line interface commands for the gh-utils extension.
package cmd

import (
	"fmt"

	"github.com/lakruzz/gh-utils/cmd/mkissue"
	"github.com/spf13/cobra"
)

var (
	issueFile  string
	branchName string
	gistID     string
)

var mkissueCmd = &cobra.Command{
	Use:   "mkissue",
	Short: "Create a GitHub issue from a markdown file",
	Long: `Create a GitHub issue from a markdown file with frontmatter support.
The markdown file should contain YAML frontmatter with metadata and a markdown body.`,
	RunE: func(_ *cobra.Command, _ []string) error {
		// Validate that branch and gist are not both specified
		if branchName != "" && gistID != "" {
			return fmt.Errorf("cannot use both --branch and --gist flags together")
		}
		// Call the original mkissue logic with the file path, branch, and gist
		return mkissue.RunWithFile(issueFile, branchName, gistID)
	},
}

func init() {
	rootCmd.AddCommand(mkissueCmd)

	// Define flags for mkissue command
	mkissueCmd.Flags().StringVarP(&issueFile, "file", "f", "", "Path to the markdown file containing issue content (required)")
	mkissueCmd.Flags().StringVarP(&branchName, "branch", "b", "", "Branch name to get the file from (optional)")
	mkissueCmd.Flags().StringVarP(&gistID, "gist", "g", "", "Gist ID to get the file from (optional)")
	_ = mkissueCmd.MarkFlagRequired("file")
}
