# expand commands search paths
export PATH="$PATH:/c/Users/royalsong/AppData/Local/Programs/Python/Python313/Scripts/:/c/Users/royalsong/AppData/Local/Programs/Python/Python313/:/c/Users/royalsong/AppData/Local/Programs/Python/Python313/Launcher/"

# customize command-line prompt
source /mingw64/share/git/completion/git-prompt.sh
export PS1="\[\e[32m\]\w\[\e[31m\]$(__git_ps1 " (%s)")\[\e[0m\]\[\e[35m\]\$ \[\e[0m\]"

# Increase history size
HISTSIZE=5000
HISTFILESIZE=20000

# Avoid duplicate entries
HISTCONTROL=ignoredups:erasedups

# Append to history, don't overwrite it
shopt -s histappend

# Include timestamps in history
export HISTTIMEFORMAT="%d/%m/%y %T "

# Share history across all sessions
export PROMPT_COMMAND="history -a; history -c; history -r; $PROMPT_COMMAND"
