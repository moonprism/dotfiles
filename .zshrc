export ZSH="$HOME/.oh-my-zsh"

ZSH_THEME="ys" #"agnoster"

# alias

alias gs="git status"

alias vi="nvim"
alias vim="nvim -c \":call LoadSession()\""
alias q="exit"
alias p="ps aux | grep"

alias -s gz="tar -xzvf"
alias -s tgz="tar -xzvf"
alias -s zip="unzip"
alias -s bz2="tar -xjvf"
alias -s xz="tar -xvJf"

alias -s go="go run"

alias simpleServer="open http://localhost:8000; python -m SimpleHTTPServer"
alias phpServer="open http://localhost:8087; php -S 0.0.0.0:8087"

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

alias ls="exa"
alias cat="bat"
alias j="z"

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# load system zsh config

[ -f ~/.my.zsh ] && source ~/.my.zsh
