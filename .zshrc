export ZSH="$HOME/.oh-my-zsh"

DISABLE_LS_COLORS="true"

ZSH_THEME='ys'

plugins=(
	git
	extract 
	vi-mode
	z
	fzf

	zsh-syntax-highlighting
	zsh-autosuggestions
)

source "$HOME/alias.sh"

source $ZSH/oh-my-zsh.sh

# fzf
export FZF_DEFAULT_COMMAND=''
export FZF_DEFAULT_OPTS="--height=40% --info=inline --margin=0 --padding=0"
