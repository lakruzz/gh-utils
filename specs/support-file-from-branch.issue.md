---
title: "Make mkissue support a different branch"
assign:
labels:
  - name: "spec"
    desc: "An issue that's designed to be a spec for and AI agent"
    color: "#881188"
  - name: "agentic ai"
    desc: "An issue that has been worked by an LLM in agentic mode"
    color: "#118811"
---

## Secret branches

I would like `utils mkissue` to support an optional extra switch: `-b, --branch <branch name to get the file from>`

The intent is that my repo with a "hidden" orphan branch where I have the issues files on.

Say I ran

```shell
gh utils mkissue --file exercises/sample.issue.md --branch secret
```

Kinda the logical equivalent to `git co secret -- exercises/sample.issue.md`

Only I don't actually want that file in my file system it should be read, but not actually retrieved.
