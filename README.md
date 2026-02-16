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

Create a GitHub issue from a markdown file:

```bash
gh utils mkissue --file path/to/issue.md
# or using short form
gh utils mkissue -f path/to/issue.md
```

The markdown file should have the following format:

```markdown
# Issue Title

Issue body goes here.
Multiple paragraphs are supported.
```

The first line (with or without `#`) becomes the issue title, and the rest becomes the issue body.

## Development

### Prerequisites

- Go 1.23 or higher
- Make
- golangci-lint (installed automatically via `make install-lint`)

### Building

```bash
# Build for current platform
make build

# Build for all platforms (Linux, Darwin, Windows on AMD64 and ARM64)
make build-all

# Clean build artifacts
make clean
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make coverage

# Run linter
make lint

# Format code
make fmt
```

### Development Workflow

1. Make your changes
2. Run `make test` to ensure tests pass
3. Run `make lint` to ensure code quality
4. Run `make build` to verify the build
5. Commit your changes (pre-commit hook will run all checks)

### Pre-commit Hooks

This repository uses pre-commit hooks to ensure code quality. The hooks run:

- Spell checking (cspell)
- Markdown linting
- Code formatting (prettier, gofmt)
- Go vet
- golangci-lint
- Go tests
- Build verification

Configure git to use the hooks:

```bash
git config core.hooksPath .githooks
```

## Project Structure

```txt
.
├── main.go              # Application entry point
├── cmd/                 # CLI commands
│   ├── root.go         # Root command
│   └── mkissue.go      # mkissue subcommand
├── internal/            # Internal packages
│   └── issue/          # Issue creation logic
├── Makefile            # Build automation
├── .golangci.yml       # Linter configuration
└── .githooks/          # Git hooks
```

## CI/CD

GitHub Actions workflow (`.github/workflows/copilot-setup-steps.yml`) automatically:

- Runs tests with race detector
- Checks code formatting
- Runs linters
- Builds the binary
- Verifies pre-commit hooks

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes following the coding standards
4. Ensure all tests and linters pass
5. Submit a pull request

### Coding Standards

- Follow Go community standards
- Use `gofmt` for formatting
- Write table-driven tests
- Use named flags over positional arguments
- Document exported functions

See `.github/copilot-instructions.md` and `.github/instructions/go-standards.instructions.md` for detailed guidelines.

## License

This project is intended for use as specified in the repository.
