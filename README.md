# Commit Message Check

[![build](https://github.com/rockstaedt/commit-message-check/actions/workflows/CI.yml/badge.svg)](https://github.com/rockstaedt/commit-message-check/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/rockstaedt/commit-message-check/branch/main/graph/badge.svg?token=VW245SMVP5)](https://codecov.io/gh/rockstaedt/commit-message-check)
[![Latest tag](https://img.shields.io/github/v/tag/rockstaedt/commit-message-check)](https://github.com/rockstaedt/commit-message-check/releases)

Commit-message-check is a CLI tool that you can use in your terminal to ensure that your commit messages aren't too
long. It is easily setup and can be used for any git project.

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

Everything is setup during the installation process. No further configuration is needed.

The tool is executed every time a commit is made. A commit message with more than 50 and less than 72 characters in the
subject line produces a warning. The exceeding characters are highlighted. See the following example:

![warning.png](docs%2Fwarning.png)

If the commit message is longer than 72 characters, the user will be asked if the commit should be aborted.

![abort.png](docs%2Fabort.png)

## Warning

The tool is not suitable for making commits from within an IDE. The commit will fail, showing an error message
indicating that the TTY device is not configured. I spent a lot of time researching if there is a possibility to cover
both use cases. Apparently, there is no way to distinguish between a commit made from within a terminal or an IDE.
If you have any ideas, please let me know! :) 

As for now, I recommend to use the tool only for commits made from within a terminal. If you want to make a commit
from within an IDE, you can temporarily disable the hook by executing the following command:

```shell
./commit-message-check uninstall
```

This removes the git hook, and you can make a commit from within your IDE. If you want to enable the hook again,
simply execute the following command:

```shell
./commit-message-check setup
```

## Updates

You can easily update the tool by executing the following command:

```shell
./commit-message-check update
```

This will check the repository for a new release and download the latest binary.

## Contributions

Contributions are welcome! Please feel free to open an issue or a pull request. Also, if you have any ideas for
improvements, please do not hesitate to contact me. If you like the tool, please give it a star. :)

## Background

The idea to make this tool initiated in a pairing session with [mschirmer](https://github.com/mschirmer1301). Using
Code With Me from Jetbrains, we found ourselves using git a lot from within the terminal because
the commit tool window is only visible for the host. We wanted to have a simple way to ensure that our commit messages
are not too long as we were used to by the highlighting from within the IDE. This is how commit-message-check was born.
