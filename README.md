# Commit Message Checker

[![build](https://github.com/rockstaedt/commit-message-check/actions/workflows/CI.yml/badge.svg)](https://github.com/rockstaedt/commit-message-check/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/rockstaedt/commit-message-check/branch/main/graph/badge.svg?token=VW245SMVP5)](https://codecov.io/gh/rockstaedt/commit-message-check)
[![Latest tag](https://img.shields.io/github/v/tag/rockstaedt/commit-message-check)](https://github.com/rockstaedt/commit-message-check/releases)

**Attention: Still under development!**

This is a commit message git hook to ensure that commit messages are not too 
long.

## Installation

Execute the command within the root folder of your git 
project. This will download the installation script that determines the OS and
downloads the latest binary. Afterwards, the initialization is started and the 
corresponding hook files are prepared. See git documentation about hooks for 
further [reference](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks).

```shell
curl -o install.sh  -L https://raw.githubusercontent.com/rockstaedt/commit-message-check/main/install.sh && chmod +x ./install.sh && ./install.sh
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
