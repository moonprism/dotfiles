export ZSH="$HOME/.oh-my-zsh"

ZSH_THEME='ys' #'agnoster'

# system

if   [[ $OSTYPE =~ ^darwin ]]; then
	BROWSER='open'
elif [[ $OSTYPE =~ ^linux ]];  then
	BROWSER='google-chrome'
fi

# alias

alias gs='git status'

alias vi='nvim'
alias vim="nvim -c ':call LoadSession()'"
alias q='exit'
alias p='ps aux | grep'
alias i='_lsofi(){lsof -i:$1};_lsofi'

alias -s gz='tar -xzvf'
alias -s tgz='tar -xzvf'
alias -s zip='unzip'
alias -s bz2='tar -xjvf'
alias -s xz='tar -xvJf'

alias -s go='go run'

alias simpleServer="$BROWSER http://localhost:8000; python -m SimpleHTTPServer"
alias phpServer="$BROWSER http://localhost:8087; php -S 0.0.0.0:8088"

alias tm0='tmux attach -t 0'

# plugin

plugins=(
	git
	extract 
	vi-mode
	z
	sudo
	fzf
)

source $ZSH/oh-my-zsh.sh

alias ls='exa'
alias cat='bat'
alias j='z'

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# load system zsh config

[ -f ~/.my.zsh ] && source ~/.my.zsh
