# Commit Message Checker

[![build](https://github.com/rockstaedt/commit-message-check/actions/workflows/build.yml/badge.svg)](https://github.com/rockstaedt/commit-message-check/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/rockstaedt/commit-message-check/branch/main/graph/badge.svg?token=VW245SMVP5)](https://codecov.io/gh/rockstaedt/commit-message-check)

**Attention: Still under development.**

This is a commit message git hook to ensure that commit messages are not too 
long.

## Installation

Execute the command for your platform within the root folder of your git 
project. This will download the latest release, puts it in the `./.git/hooks/`
folder and makes it executable. See git documentation about hooks for 
further [reference](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks).

### MacOS

```shell
curl -o .git/hooks/commit-msg  -L https://github.com/rockstaedt/commit-message-check/releases/download/v0.0.4/commit-message-check-v0.0.4-darwin-arm64 && chmod +x .git/hooks/commit-msg
```

### Linux

```shell
curl -o .git/hooks/commit-msg -L https://github.com/rockstaedt/commit-message-check/releases/download/v0.0.4/commit-message-check-v0.0.4-linux-amd64 && chmod +x .git/hooks/commit-msg
```

## Usage

Every time a git commit is made, the corresponding hook is fired. For a commit 
message that is too long (>50 characters), the commit process is 
aborted and a warning is shown. See the following as an example.

```shell
$ git commit -m "Veryyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy long"
2022/12/18 12:34:43  [INFO]     Validating commit message...
2022/12/18 12:34:43  [ERROR]    Abort commit. Subject line too long. Please fix.
```
