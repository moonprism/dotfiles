# export ZSH="/home/aqua/.oh-my-zsh"
export ZSH="$HOME/.oh-my-zsh"

ZSH_THEME="agnoster"

set -o vi

alias q="exit"

alias gs="git status"

alias -s gz='tar -xzvf'
alias -s tgz='tar -xzvf'
alias -s zip='unzip'
alias -s bz2='tar -xjvf'
alias -s xz='tar -xvJf'
# ??
alias vi='nvim'

alias -s go='go run'

# plugin

plugins=(
	git
	extract 
	vi-mode
	z
	sudo
	fzf
	zsh-syntax-highlighting
	zsh-autosuggestions
)

# auto start tmux
ZSH_TMUX_AUTOSTART=true

source $ZSH/oh-my-zsh.sh

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# local zsh config

[ -f ~/.my.zsh ] && source ~/.my.zsh

# source <(curl -sSL git.io/forgit)

alias ls='exa'
alias cat='bat'
alias j='z'
