#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE}")";

for file in `ls -a .`
do
  if [[ ! "$file" =~ ^(karabiner|bootstrap.sh|README.md|.git|.|..)$ ]]; then
    ln -s "${PWD}/$file" ~
  fi
done

# mac 系统使用karabiner映射按键
if [[ `uname` == 'Darwin' ]]; then
  cd karabiner
  for file in `ls .`
  do
    ln -s "${PWD}/$file" ~/.config/karabiner/assets/complex_modifications
  done
fi

# install zsh plugin 

# zsh-syntax-highlighting

# zsh-autosuggestions

# TODO install bat exa tmux-tmp ack rainbarf
