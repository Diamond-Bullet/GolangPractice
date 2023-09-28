# git config for current user.
vim ~/.gitconfig

# set a config value
git config --global http.proxy http://127.0.0.1:12333

## commit
# alter the message of last commit
git commit --amend -m "message"


## branch
# list all branches
git branch -a

# create a branch
git branch [name]

# create a branch from specified commit or branch. from current branch by default.
git checkout -b [name] [commit_id||branch_name]

# delete local branch, -D for force delete.
git branch -d [name]
git branch -D [name]

# delete remote branch
git push origin --delete [name]

# link local branch to remote branch.
git branch --set-upstream-to=[remote_branch] [local_branch]

# change git repository
git remote set-url origin https://github.com/szq-123/codingPractice.git



