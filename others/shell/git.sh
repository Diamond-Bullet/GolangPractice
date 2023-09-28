# git config for current user.
vim ~/.gitconfig
# set a config value
git config --global http.proxy http://127.0.0.1:12333

# initialize a git repo in specified path(it is a subdirectory of current folder), `.` by default.
git init [directory]

# TODO 	usage of `.gitmodules`

## commit
# alter the message of last commit
git commit --amend -m "message"

# apply changes this commit made compared to its parent commit to current branch.
# it's useful when you just need some modifications but not all(use git merge instead).
git cherry-pick [commit_id]

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

## git reset. reset the `head` pointer and the branch pointer that `head` is pointing at.
# reset local repo.
git reset --soft [commit_id]
# reset local repo and 暂存区.
git reset --mixed [commit_id]
# reset local repo, 暂存区 and workspace.
git reset --hard [commit_id]

# cancel tracing a file.
git rm -f --cached [file]

## log
# show commit history for current branch.
git log
# show history of the file.
git log [file]
# just show commit_id and message.
git log --pretty=oneline
# show modified files.
git log --stat

## show
# show changes of the commit, and of the file if specified.
git show [commit_id] [file]
# just show changed files of the commit.
git show --stat [commit_id]

## diff. show difference between A and B. A, B can be commits, or branches.
# if specify a file, only show difference in the file.
git diff A B [file]
# just show files having difference.
git diff --stat

# list traced files in current folder. -O for untracked.
# https://blog.csdn.net/ystyaoshengting/article/details/104029519
git ls-files [path||file_name]
git ls-files -O [path||file_name]

