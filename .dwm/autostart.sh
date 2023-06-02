#!/bin/bash
#

setxkbmap -option ctrl:swapcaps

exec fcitx5 &
