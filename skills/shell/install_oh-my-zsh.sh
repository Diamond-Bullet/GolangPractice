#!/bin/bash
# installation instruction https://github.com/ohmyzsh/ohmyzsh/wiki
# nowadays, you can simply ask ChatGPT how to achieve it.
# you can also install it from source, which is quite easy as well.

cd ~

function install_pkg() {
  if [ $# -eq 1 ]; then # argument count is 1.
    _=$(yum install "${1}")
    _=$(apt install "${1}")
  fi
}

# install zsh
if [[ "$(which zsh)" == *"not found" ]]; then
  install_pkg zsh
fi
if [[ "$SHELL" != *"zsh" ]]; then
  chsh -s "$(which zsh)"
fi

# install curl
if [[ "$(which curl)" == *"not found" ]]; then
  install_pkg curl
fi

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" || exit 1

source ~/.zshrc
