package issue

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseMarkdown(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		wantTitle   string
		wantBody    string
		wantErr     bool
		errContains string
	}{
		{
			name:      "simple title and body",
			content:   "# Test Title\n\nThis is the body",
			wantTitle: "Test Title",
			wantBody:  "This is the body",
			wantErr:   false,
		},
		{
			name:      "title without hash",
			content:   "Test Title\n\nThis is the body",
			wantTitle: "Test Title",
			wantBody:  "This is the body",
			wantErr:   false,
		},
		{
			name:      "multiline body",
			content:   "# Test Title\n\nLine 1\nLine 2\nLine 3",
			wantTitle: "Test Title",
			wantBody:  "Line 1\nLine 2\nLine 3",
			wantErr:   false,
		},
		{
			name:        "empty content",
			content:     "",
			wantErr:     true,
			errContains: "no title found",
		},
		{
			name:      "title with extra whitespace",
			content:   "  # Test Title  \n\nBody content",
			wantTitle: "Test Title",
			wantBody:  "Body content",
			wantErr:   false,
		},
		{
			name:      "title only",
			content:   "# Test Title",
			wantTitle: "Test Title",
			wantBody:  "",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title, body, err := parseMarkdown(tt.content)

			if tt.wantErr {
				if err == nil {
					t.Errorf("parseMarkdown() expected error but got none")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("parseMarkdown() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("parseMarkdown() unexpected error = %v", err)
				return
			}

			if title != tt.wantTitle {
				t.Errorf("parseMarkdown() title = %v, want %v", title, tt.wantTitle)
			}

			if body != tt.wantBody {
				t.Errorf("parseMarkdown() body = %v, want %v", body, tt.wantBody)
			}
		})
	}
}

func TestCreateFromFile(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid markdown file",
			content: "# Test Issue\n\nThis is a test issue body",
			wantErr: false,
		},
		{
			name:        "empty file",
			content:     "",
			wantErr:     true,
			errContains: "no title found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "test.md")

			// #nosec G306 - test file permissions are intentionally relaxed
			if err := os.WriteFile(tmpFile, []byte(tt.content), 0644); err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}

			// Note: This will actually try to create a GitHub issue
			// In a real scenario, we'd mock the gh CLI call
			// For now, we're just testing the parsing part
			_, _, err := parseMarkdown(tt.content)

			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateFromFile() expected error but got none")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("CreateFromFile() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("CreateFromFile() unexpected error = %v", err)
			}
		})
	}
}

func TestCreateFromFileNotExists(t *testing.T) {
	err := CreateFromFile("/nonexistent/file.md")
	if err == nil {
		t.Error("CreateFromFile() expected error for nonexistent file but got none")
	}
	if !strings.Contains(err.Error(), "failed to read file") {
		t.Errorf("CreateFromFile() error = %v, want error containing 'failed to read file'", err)
	}
}
