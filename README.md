# Commit Message Checker

[![build](https://github.com/rockstaedt/commit-message-check/actions/workflows/build.yml/badge.svg)](https://github.com/rockstaedt/commit-message-check/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/rockstaedt/commit-message-check/branch/main/graph/badge.svg?token=VW245SMVP5)](https://codecov.io/gh/rockstaedt/commit-message-check)

**Attention: Still under development.**

This is a commit message git hook to ensure that commit messages are not too 
long.

## Usage

In your project in which a git repository is already initiated, you have to 
modify file `./.git/hooks/commit-msg.sample`. First, remove the 
prefix *sample* to enable the hook. Then, modify the executable as follows:

```shell
#!/bin/sh

./.git/hooks/commit-message-check $1
```

Afterwards, you have to put the Golang executable from this repository into 
the hook folder `./.git/hooks/`. This will run the Golang executable every time
a git commit is fired. For a commit message that is too long (>50 
characters), the commit process is 
aborted and a warning is shown. See the following as an example.

```shell
$ git commit -m "Veryyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy long"
2022/12/18 12:34:43  [INFO]     Validating commit message...
2022/12/18 12:34:43  [ERROR]    Abort commit. Subject line too long. Please fix.
```

See git documentation about hooks for further reference
[here](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks).