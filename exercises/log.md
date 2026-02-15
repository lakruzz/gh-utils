# Worklog

## Create a workspace file in the root of your repo.

### How

Command Palette 
- Mac: ⬆⌘P (`shift+command+P`)
- Windows/Linux: ⬆^P (`shift+ctrl+P`)

`Workspaces:Save Workspace As...` → `default.code-workspace``

### Why

If you do not use an explicit workspace, VC Code wil implicitly create one for you.

But we don't want things to happen implicitly. That is an anti-pattern in a Configuration as Code approach. But it's also not just a _principle_. It comes with benefits. I can store VS Code settings in the workspace context, and I can add and remove additional folders, without resetting my CoPilot chat sessions.

Yep: If you _don't_ use a named workspace, but just go with VS Code's implied one — and you add a new folder to the workspace. Then VS Code will see that as a _new_ or _different_ workspace and as an undesired side effect all your current and historic chat sessions with Copilot are gone.

## Add a Dev Container

### How

Command Palette 
- Mac: ⬆⌘P (`shift+command+P`)
- Windows/Linux: ⬆^P (`shift+ctrl+P`)

`Dev Containers:Add Dev Container Configurations Files...`

→  Add configuration to workspace
→ Ubuntu
  → noble
→ Select additional features to install
  → GitHub CLI (devcontainers)
  → Common Utilities (devcontainers)
  → Keep Defaults

The steps above will create `.devcontainer/devcontainer.json`

## Reopen in container

### How

Command Palette 
- Mac: ⬆⌘P (`shift+command+P`)
- Windows/Linux: ⬆^P (`shift+ctrl+P`)

`Dev Containers:Reopen in Container`

Test:

```bash
lsb_release -a # Shows detailed distribution information (release, codename, etc.)
```

### Ohhhh why this may not work!

`Reopen in Container` requires that your host system (your PC!) can host a Docker container. In short **Docker must be installed and set up right"

**Prerequsite:** Install Docker Desktop
- [Mac](https://docs.docker.com/desktop/setup/install/mac-install/)
- [Windows](https://docs.docker.com/desktop/setup/install/windows-install/) (**IMPORTANT**: Use `WSL 2` as opposed to `Hyper-V`)

## Install extensions into the Dev Container

We need a few _extensions_ in addition to the _features_

- Git Graph (mhutchie)
- Better Git Line Blame (Mitchell Kember)
- GitHub Copilot Chat (GitHub)


## How
Open up the _Extensions_ panel

_View_ → _Extensions_

- Mac: ⬆⌘X (`shift+command+X`)
- Windows/Linux: ⬆^X (`shift+ctrl+X`)


Search for each extension one by one. When you find it:
- Right click and choose _Add to devcontainer.json_

👆 When you update the `devcontainer.json` file VS Code reminds you to rebuild the container, but don't do that — yet, add all three extensions first.

NOW, you should rebuild the container:

Command Palette 
- Mac: ⬆⌘P (`shift+command+P`)
- Windows/Linux: ⬆^P (`shift+ctrl+P`)

`Dev Containers:Rebuild Container`

## Log into GitHub...

We want to be able to talk to GitHub using the Command Line Interface (one fo the features we installed).

In the terminal run:

```bash
gh auth status
```

It replies:

```bash
You are not logged into any GitHub hosts. To log in, run: gh auth login
```

So, every time we start the container it needs to be authorized. That's boring!

Let's instead, once and for all authorize you host system (Your PC) and then let the container inherit that authorization.

At this point the terminal is running in the Dev Container, we need to get back in our local environment:


Command Palette 
- Mac: ⬆⌘P (`shift+command+P`)
- Windows/Linux: ⬆^P (`shift+ctrl+P`)

`Dev Containers:Reopen Folder Locally`

(VS Code is back on your PC, it sees that you have a devcontainer configured, and offers you to open it – do not do that yet)

In your terminal run:

```bash
gh auth status
```

1. Is the `gh` command even valid on your system?
2. Are you logged in?
3. If so, are you logged in with the `GH_TOKEN`?
4. Do you have the right scopes (we need `project` scope in addition to the default scopes)

Let's take them one by one (skip each step if it's not a concern)

### 1: Install GitGub CLI

- [Mac](https://github.com/cli/cli/blob/trunk/docs/install_macos.md)
- [Windows](https://github.com/cli/cli/blob/trunk/docs/install_windows.md)
- [Linux](https://github.com/cli/cli/blob/trunk/docs/install_linux.md)gh

### 2: Authenticate with additional `project` scope

```bash
gh auth login --hostname github.com --scopes project --git-protocol https --web
```

### 3 + 4 Right token and right scope?

running `gh auth status` should show something similar to:

```shell
github.com
  ✓ Logged in to github.com account lakruzz (GH_TOKEN)
  - Active account: true
  - Git operations protocol: https
  - Token: gho_************************************
  - Token scopes: 'gist', 'project', 'read:org', 'repo', 'workflow'
```

- You should use `GH_TOKEN`as login procedure
- Token should start with `gho`
- Token scopes should mention `'project'`


If that is not the case try:

```bash
gh auth logout
unset GITHUB_TOKEN
```

run  `gh auth status` again.

OK? If not run steps 1+2 again.

**GOOD TO KNOW**

What if you suspect, that someone stole your token, and you wanted to revoke or disable it?

- Go to https://github.com/settings/applications
- Find the OAuth application named "GitHub CLI"
- Click on the three dots and choose _Revoke_

Test it (come on ...its fun!): 

- Revoke the token
- Run `gh auth status` and see that you are now rejected access
- Set it up again

## Capture the token, encode it and store it to your profile

OK, so now your shell is logged in to GitHub. You want to capture this state, so that you do not have to do this every time you start your PC.

Let's go through how you do this in `bash` and `zsh` (If you use `fish` the steps are slightly different).

We do not want to store the token in bare text in any files, so we will encode it with `base64` first.

On you local system determine your shell:

```bash
echo $SHELL
```

If `bash` your profile is in `~/.bash_profile`
If `zsh`your profile is in  `~/.zprofile`

run

```bash
echo $GH_TOKEN 
echo $GH_TOKEN | base64
```

You'll see first the token (starting with `gho`) and then the base64 encoded token.

Find (or create) your profile file and add the following lines:

```bash
export _GH_TOKEN=<your-base64-encoded-token-here>
export GH_TOKEN=$(echo $_GH_TOKEN | base64 --decode)
```

At this point.  Start an new shell and run `gh auth status`

## Make the Dev Container utilize your hos settings

Open `.devcontainer/devcontainer.json` and add the following as a top-level key:

```json
    "remoteEnv": {
        "GH_TOKEN": "${localEnv:GH_TOKEN}"
    }
```

Note that `JSON` is really picky in regards of commas, so dependant on wether you put this key before, after or in between existing keys, you'll nee a comma after, before or both - respectively.

At this point your `devcontainer.json` looks like this

```json
// For format details, see https://aka.ms/devcontainer.json. For config options, see the
// README at: https://github.com/devcontainers/templates/tree/main/src/ubuntu
{
	"name": "Ubuntu",
	// Or use a Dockerfile or Docker Compose file. More info: https://containers.dev/guide/dockerfile
	"image": "mcr.microsoft.com/devcontainers/base:noble",
	"features": {
		"ghcr.io/devcontainers/features/common-utils:2": {},
		"ghcr.io/devcontainers/features/github-cli:1": {}
	},
	"customizations": {
		"vscode": {
			"extensions": [
				"mhutchie.git-graph",
				"mk12.better-git-line-blame",
				"GitHub.copilot-chat"
			]
		}
	},
    "remoteEnv": {
        "GH_TOKEN": "${localEnv:GH_TOKEN}"
    }

	// Features to add to the dev container. More info: https://containers.dev/features.
	// "features": {},

	// Use 'forwardPorts' to make a list of ports inside the container available locally.
	// "forwardPorts": [],

	// Use 'postCreateCommand' to run commands after the container is created.
	// "postCreateCommand": "uname -a",

	// Configure tool-specific properties.
	// "customizations": {},

	// Uncomment to connect as root instead. More info: https://aka.ms/dev-containers-non-root.
	// "remoteUser": "root"
}
```

Now rebuild in the dev container:

Command Palette 
- Mac: ⬆⌘P (`shift+command+P`)
- Windows/Linux: ⬆^P (`shift+ctrl+P`)

`Dev Containers:Rebuild and Reopen in Container`

run `gh auth status`


