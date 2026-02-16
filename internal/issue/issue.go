// Package issue provides functionality for creating GitHub issues from markdown files.
package issue

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CreateFromFile creates a GitHub issue from a markdown file.
// The filePath parameter is user-provided and validated before use.
func CreateFromFile(filePath string) error {
	// Read the file
	// #nosec G304 - filePath is a user-provided argument, validated by caller
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the markdown content
	title, body, err := parseMarkdown(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse markdown: %w", err)
	}

	// Create the issue using gh CLI
	return createIssue(title, body)
}

// parseMarkdown extracts title and body from markdown content
// Expects the first line to be the title (optionally starting with #)
func parseMarkdown(content string) (string, string, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	var title string
	var bodyLines []string
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			// First line is the title
			title = strings.TrimSpace(line)
			// Remove leading # if present
			title = strings.TrimPrefix(title, "#")
			title = strings.TrimSpace(title)
			firstLine = false
			continue
		}

		bodyLines = append(bodyLines, line)
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	if title == "" {
		return "", "", fmt.Errorf("no title found in markdown file")
	}

	body := strings.TrimSpace(strings.Join(bodyLines, "\n"))

	return title, body, nil
}

// createIssue creates a GitHub issue using gh CLI
func createIssue(title, body string) error {
	cmd := exec.Command("gh", "issue", "create", "--title", title, "--body", body)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gh command failed: %s: %w", stderr.String(), err)
	}

	return nil
}
