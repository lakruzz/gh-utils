---
title: "Make the utility a standed gh extension"
assign:
  - "@me        
labels:           
  - name: "spec"
    desc: "An issue that's desinged to be a spec for and AI agent"
    color: "#881188"
  - name: "agentic ai"
    desc: "An issue that has been worked by an LLM in agentic mode"
    color: "#118811"
---

## Turn this repo into a standard GH CLI extension

We hacked `utils` and the subcommand `mkissue` as a mvp. It serves the purpose descibed in `exercises/template.issue.md`

We need to professionalize this setup as I imagine `utils` will grow in the future.

## put source in `src``

I'd like to keep all code in a dedicated `src` directory. A repo contains all kinds of cofigurations, documentations, etc so as a principle I want the separated from all this and put into `src`.

## Rename this repo to `lakruzz/gh-utiils`

The contract for a GH CLI extension is that it's name starts with `gh-` and than the root of the repo contains an executable what has the same name as part of the repo name that follows `gh-`

So this repos should be named `lakruzz/gh-utiils` and the go project should compile to `./utils` (as it does already).

That would enable this feature to be installed using the GH CLI built-in package manager and called as a gh extension.
Like this:

```bash
gh ext install lakruzz/gh-utils
gh utils mkissue ...
```

## When designing the CLI always use named switches over anonymous arguments

In the current state `utils mkissue` takes an anonymous argument like this:

```shell
Usage: utils <subcommand> [args]

Available subcommands:
  mkissue <file.issue.md>  Create a GitHub issue from a markdown file
  help                     Show this help message
```

Id like all arguments to belong to a switch like:

```shell
Usage: utils <subcommand> [args] [flags]

Available subcommands:
  mkissue -f, --file <file.issue.md>  Create a GitHub issue from a markdown file
  help                                Show this help message
```

I will not rule out that _some_ subcommands can occasionally be optimized to take an anonymous argument, but the general rule and principle is that we use named flags.

And that flags (usually) come in both a long version (`--[a-z+]{3:18}`). and a short version (`-[a-z]`). Although there might be _some_ rarely used or vers specialized flags tah only come in the long version.

## Setup build according to community standards

I'm new to Go, but I'd like the set up the go project according to community standards.

- Does go have some thing smilar to Ruby's `rake` file og Node's `npm run` then we should use that!
- Initialize and setup a community standard Go unit test/mock frame work with coverage and add unittests for the code that is already written.
- A build feature should be provided that builds automatically on FileSave and which builds to both ARM and AMD architecture on both Windows, Darwin and Linux.
- A linter should be set up, preferably a power full one (like Ruff for Python) preferably one that can also check cyclomatic or McCabe complexity.

## Store what ever is needed in RAG files

Grab what you must from this instructions and store it for permanent RAG instructions in either.

- `.github/copilot-instructions.md`
- `.github/instructions/*.instructions.md`

...as you see fit (it's for your own good)

Also create:

- `.github/workflows/copilot-setup-steps.yml`

Set it up so that you can work in a similar setup as the one described in:

- `.devcontainers/devcontainer.json`
- `.devcontainer/postCreateCommand.sh`

A specific note to the `copilot-setup-steps-yml`:

You should setup to be dependant on the `pre-commit` hook. and you should run it manually. to verify that the workflow works. So the last two steps in the workflow should be along the lines of

```yaml
# Configure git hooks path
- name: Configure Git hooks
  run: git config core.hooksPath .githooks

# Verify the setup by running the pre-commit hook
- name: Test pre-commit hook
  run: .githooks/pre-commit
```

## Finalize the `pre-commit` hook to serve the new setup

Consider if any more steps could be added to the `pre-commit` hook like

- run go unit tests
- Run Go linter check
- Run a go build
- Any other static analysis that we could benefit from
