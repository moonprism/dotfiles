## init
if   [[ $OSTYPE =~ ^darwin ]]; then
	BROWSER='open'
elif [[ $OSTYPE =~ ^linux ]];  then
	BROWSER='google-chrome'
fi

# git
alias gs='git status'

# tmux
alias t0='tmux attach -t 0'

# editor
alias vi='nvim'
alias vim="nvim -c ':call LoadSession()'"

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
alias simpleServer="$BROWSER http://localhost:8000; python -m SimpleHTTPServer"
alias phpServer="$BROWSER http://localhost:8088; php -S 0.0.0.0:8088"

# more
#alias ls='exa'
alias cat='bat'
alias j='z'
