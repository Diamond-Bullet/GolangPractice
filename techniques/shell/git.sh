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
git checkout -b [name] [commit_id || branch_name]
# delete local branch, -D for force delete.
git branch -d [name]
git branch -D [name]
# delete remote branch
git push origin --delete [name]
# link local branch to remote branch.
git branch --set-upstream-to=[remote_branch] [local_branch]
# change git repository
git remote set-url origin https://github.com/Diamond-Bullet/GolangPractice.git

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
git ls-files [path || file_name]
git ls-files -O [path || file_name]

## stash. https://www.atlassian.com/git/tutorials/saving-changes/git-stash
# store changes into stash list.
git stash
# save with message.
git stash save "message"
# show stash list.
git stash list
# show difference between workspace and stash.
git stash show -p stash@{$num}
# apply and not delete the stash.
git stash apply stash@{$num}
# push, pop
git stash push stash@{$num}
git stash pop stash@{$num}

## tag
# list tags
git tag
# add tag to specified commit. if commit_id is not provided, add to newest.
git tag -a V1.2 [commit_id] -m 'release 1.2'
# delete local tag
git tag -d V1.2
# push all local tags to remote repo.
git push origin --tags
# push specified tag to remote repo.
git push origin [tag_name]
# delete remote tag.
git push origin :refs/tags/V1.2
# fetch a tag
git fetch origin [tag_name]
# checkout from tag
git checkout tags/v1.0 [branch_name]