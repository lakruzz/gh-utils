# Setup linting

## Install

### VS Code install

Add the following extensions to the `devcontainer.json`

- Prettier - Code formatter (Prettier)
- Markdownlint (David Anson)
- Code Cspell Checker (Street Side Software)
- Danish - Code Spell Checker (Street Side Software)

```json
	"customizations": {
		"vscode": {
			"extensions": [
				...
				"esbenp.prettier-vscode",
				"DavidAnson.vscode-markdownlint",
				"streetsidesoftware.code-spell-checker",
				"streetsidesoftware.code-spell-checker-danish"
			]
		}
	}
```

## Install the same tools as features, directly in the devcontainer

```json
		"ghcr.io/devcontainers-extra/features/npm-packages:1": {
			"packages": ["prettier", "markdownlint-cli2", "cspell", "@cspell/dict-da-dk"]
		}
```

From the terminal run

```bash
cspell
markdownlint-cli2
prettier
```

They should all work fine.

## Configure linters to run automatically

Add `.vscode/settings.json``

```json
{
  "files.autoSave": "afterDelay",
  "files.autoSaveDelay": 2000,

  "editor.codeActionsOnSave": {
    "source.fixAll.markdownlint": "always",
    "source.fixAll.cspell": "explicit",
    "source.fixAll.prettier": "always"
  },
  "editor.formatOnSave": true,
  "editor.formatOnType": true,
  "editor.formatOnSaveMode": "file",

  "markdownlint.configFile": ".markdownlint-cli2.jsonc", // Force extension to use the same config file as the CLI tool

  "[markdown]": {
    "editor.defaultFormatter": "DavidAnson.vscode-markdownlint"
  },

  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },

  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },

  "[yaml]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  }
}
```

### Set up the `markdownlilnt-cli2` configuration

Create `.markdownlint-cli2.jsonc` and add:

```json
// Configuration for markdownlint-cli2

// This file is used by markdownlint-cli2 as opposed to markdownlint which uses .markdownlint.json(c).
// The format between the two differ, markdownlint-cli2 supports "globs" which is nice in order to avoid noise.

// VSCode setting will look for .markdownlint.json(c) so to get the same result when using the VS code extenison
// and the markdownlint-cli2 you need to tell VS Code in '.vscode/settings.json' to use this file by
// explicitly adding:
//     "markdownlint.configFile": ".markdownlint-cli2.jsonc",
{
  "globs": [
    "**/*.md",
    "**/*.markdown",
    "!node_modules/**", // Exclude node_modules (JavaScript)
    "!temp/**", // Exclude temp directory
    "!.bundle/**", // Exclude .bundle (ruby)
    "!vendor/**" // Exclude vendor (ruby)
  ],
  "config": {
    "MD033": false, // Allow inline HTML
    "MD004": { "style": "dash" }, // Enforce consistent unordered list style (dash)
    "MD013": false, // Disable line length rule
    "MD036": false, // Disable Emphasis used instead of a header
    "MD041": { "allow_preamble": true, "level": 1 }, // Allow preamble, require first header to be level 1
    "MD026": { "punctuation": ".," }, // Allow punctuation at the end of a list item, except for "." and ","
    "MD024": { "siblings_only": true }, // Rule only apply to siblings (same parent)
    "MD010": { "code_blocks": false, "spaces_per_tab": 2 } // Allow tabs in code blocks and set spaces per tab to 2
  }
}
```
