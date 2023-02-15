#!/bin/sh

version=$(curl --silent "https://api.github.com/repos/rockstaedt/commit-message-check/releases/latest" |
    grep '"tag_name":' |
    sed -E 's/.*"([^"]+)".*/\1/')

echo "-> Installing commit-message-check ${version}"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  os="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  os="darwin"
else
  echo "ERROR - Could not determine operating system. Please download binary manually."
  exit
fi

os="${os}-$(uname -m)"
binary="commit-message-check-${version}-${os}"
url="https://github.com/rockstaedt/commit-message-check/releases/latest/download/${binary}"

curl -o commit-message-check -L $url && chmod +x ./commit-message-check

if !(grep -Fxq "commit-message-check" .gitignore); then
  read -e -p "Do you want to add the binary to gitignore? (y/n) " choice
  [[ "$choice" == [Yy]* ]] && echo "\ncommit-message-check" >> .gitignore
fi

./commit-message-check setup

rm ./install.sh
