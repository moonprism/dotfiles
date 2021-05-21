#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE}")";

for file in `ls -a .`
do
  if [[ ! "$file" =~ ^(joplin|karabiner|bootstrap.sh|README.md|.git|.|..)$ ]]; then
    ln -s "${PWD}/$file" ~
  fi
done

# mac 系统使用karabiner映射按键
if [[ `uname` == 'Darwin' ]]; then
  for file in `ls karabiner`
  do
    ln -s "${PWD}/karabiner/$file" ~/.config/karabiner/assets/complex_modifications
  done
fi

for file in `ls joplin`
do
  ln -s "${PWD}/joplin/$file" ~/.config/joplin
done

# install zsh plugin 

# zsh-syntax-highlighting

# zsh-autosuggestions

# TODO install bat exa tmux-tmp ack rainbarf
