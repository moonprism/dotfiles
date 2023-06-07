#!/bin/bash
#

setxkbmap -option ctrl:swapcaps

exec fcitx5 &

exec ds &

amixer sset Master 50%
exec sxhkd -c $HOME/.sxhkdrc &
