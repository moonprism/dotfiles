## init
if   [[ $OSTYPE =~ ^darwin ]]; then
	BROWSER='open'
elif [[ $OSTYPE =~ ^linux ]];  then
	BROWSER='google-chrome-stable'
fi

# git
alias gs='git status'

# tmux
alias t0='tmux attach -t 0'

# editor
alias v='nvim'

# system
alias q='exit'
alias p='ps aux | grep'
alias i='_lsofi(){lsof -i:$1};_lsofi'

# file
alias -s gz='tar -xzvf'
alias -s tgz='tar -xzvf'
alias -s zip='unzip'
alias -s bz2='tar -xjvf'
alias -s xz='tar -xvJf'

# compile
alias -s go='go run'

# test
alias simpleServer="$BROWSER http://localhost:8000 & ; python3 -m http.server 8000"

# more
alias ls='lsd'
alias cat='bat'
alias j='z'
alias t='tmux'

#dunst
alias cl='dunstctl history-clear'

alias icat="kitten icat"
