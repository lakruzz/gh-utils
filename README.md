<!-- cspell:ignore gofmt golangci  -->

# gh-utils

A GitHub CLI extension providing utility commands for GitHub workflows and automation.

## Installation

Install as a GitHub CLI extension:

```bash
gh extension install lakruzz/gh-utils
```

## Usage

### `mkissue` - Create GitHub Issue from Markdown File

Create a GitHub issue from a markdown file with YAML frontmatter:

```bash
gh utils mkissue --file path/to/issue.md
# or using short form
gh utils mkissue -f path/to/issue.md
```

#### Reading from a Git Branch

You can read the issue file from a specific git branch without checking it out to your filesystem:

```bash
gh utils mkissue --file path/to/issue.md --branch secret
# or using short form
gh utils mkissue -f path/to/issue.md -b secret
```

This is useful for keeping issue templates in a separate orphan branch.

#### Issue File Format

The issue file must follow the format specified in [`exercises/template.issue.md`](exercises/template.issue.md). This template defines the contract for issue files:

- **Frontmatter** (YAML): Contains metadata like title, assignees, labels, milestone, and projects
- **Body** (Markdown): The issue description/content

**Example:**

```yaml
---
title: "My Issue Title"
assign:
  - "@me"
labels:
  - name: "bug"
  - name: "priority-high"
    color: "d73a4a"
    desc: "High priority issues"
milestone: "v1.0"
projects:
  - "Main Project"
---
## Issue Description

This is the issue body written in Markdown.

- Supports lists
- **Bold text**
- And all other Markdown features
```

See [`specs/template.issue.md`](specs/template.issue.md) for the complete specification and all supported fields.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for developer documentation and guidelines.

## License

This project is intended for use as specified in the repository.
