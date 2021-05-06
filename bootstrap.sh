#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE}")";

for file in `ls -a .`
do
  if [[ ! "$file" =~ ^(bootstrap.sh|README.md|.git|.|..)$ ]]; then
    ln -s "${PWD}/$file" ~
  fi
done

# TODO install bat exa tmux-tmp ack rainbarf
