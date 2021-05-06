# export ZSH="/home/aqua/.oh-my-zsh"
export ZSH="$HOME/.oh-my-zsh"

ZSH_THEME="ys" #"agnoster"

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
	tmux
)

# ZSH_TMUX_AUTOSTART=true

source $ZSH/oh-my-zsh.sh

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# load my system zsh config

[ -f ~/.my.zsh ] && source ~/.my.zsh

# source <(curl -sSL git.io/forgit)

alias ls='exa'
alias cat='bat'
alias j='z'
export HOMEBREW_NO_AUTO_UPDATE=true

export PATH="/usr/local/opt/php@7.2/bin:/usr/local/opt/php@7.2/sbin:$PATH"
export PATH="/usr/local/opt/node@12/bin:$PATH"
source /Users/yuanguoan/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
source $(brew --prefix)/share/zsh-autosuggestions/zsh-autosuggestions.zsh
