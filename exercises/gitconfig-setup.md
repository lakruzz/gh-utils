# Set up git

## Add A Post create file

Create the file `.devcontainer/postCreateCommand.sh`

```bash

#!/usr/bin/env bash

set -e

PREFIX="🍰  "
echo "$PREFIX Running $(basename $0)"


if [ -n "$GH_TOKEN" ]; then
  echo "$PREFIX  \$GH_TOKEN" defined. It takes precedens over \$GITHUB_TOKEN ...looking fine so far
else
  if [[ "$CODESPACES" == "true" ]]; then
      echo "$PREFIX No \$GH_TOKEN defined - using the standard ghu_*** token injected by the codespace into \$GITHUB_TOKEN"
  else
      echo "$PREFIX ⚠️ No \$GH_TOKEN defined - skipping GitHub CLI login."
      echo "$PREFIX    1) Run 'gh auth login -s project' to login with OAuth and sufficient permissions"
  fi
fi

set +e
gh auth status >/dev/null 2>&1
AUTH_OK=$?
set -e
if [ $AUTH_OK -ne 0 ]; then
  echo "$PREFIX ⚠️ Not logged into GitHub CLI"
  echo "$PREFIX    This is not loogking good  — we want GitHub CLI to work!"
else
  echo "$PREFIX GitHub Authnetication is working smooth!
fi

echo "$PREFIX Setting up safe git repository to prevent dubious ownership errors"
git config --global --add safe.directory $(pwd)

echo "$PREFIX Setting up git configuration to support .gitconfig in repo-root"
git config --local --get include.path | grep -e ../.gitconfig >/dev/null 2>&1 || git config --local --add include.path ../.gitconfig


if [ -f "Gemfile.lock" ]; then
    echo "$PREFIX Installing ruby gems"
    echo "$PREFIX Freeze Gemfile.lock so it isn't modified during install"
    bundle config set frozen true
    bundle install
fi

if [ -f "package-lock.json" ]; then
    echo "$PREFIX Installing node modules"
    npm ci
fi


echo "$PREFIX ✅ SUCCESS"
exit 0

```

Make it an executable

```bash
chmod +x .devcontainer/postCreateComand.sh
```
