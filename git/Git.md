## git命令分类整理

### 全局设置

`git config --global user.name xxx`：设置全局用户名，信息记录在~/.gitconfig文件中
`git config --global user.email xxx@xxx.com`：设置全局邮箱地址，信息记录在~/.gitconfig文件中
`git init`：将当前目录配置成git仓库，信息记录在隐藏的.git文件夹中

### 常用命令

`git add XX` ：将XX文件添加到暂存区
`git commit -m "给自己看的备注信息"`：将暂存区的内容提交到当前分支
`git status`：查看仓库状态
`git log`：查看当前分支的所有版本
`git push -u (第一次需要-u以后不需要)` ：将当前分支推送到远程仓库
`git clone git@git.acwing.com:xxx/XXX.git`：将远程仓库XXX下载到当前目录下
`git branch`：查看所有分支和当前所处分支

### 查看命令

`git diff XX`：查看XX文件相对于暂存区修改了哪些内容
`git status`：查看仓库状态
`git log`：查看当前分支的所有版本
`git log --pretty=oneline`：用一行来显示
`git reflog`：查看HEAD指针的移动历史（包括被回滚的版本）
`git branch`：查看所有分支和当前所处分支
`git pull` ：将远程仓库的当前分支与本地仓库的当前分支合并

### 删除命令

`git rm --cached XX`：将文件从仓库索引目录中删掉，不希望管理这个文件
`git restore --staged xx`：==将xx从暂存区里移除==
`git checkout — XX`或`git restore XX`：==将XX文件尚未加入暂存区的修改全部撤销==

### 代码回滚

`git reset --hard HEAD^` 或`git reset --hard HEAD~` ：将代码库回滚到上一个版本
`git reset --hard HEAD^^`：往上回滚两次，以此类推
`git reset --hard HEAD~100`：往上回滚100个版本
`git reset --hard 版本号`：回滚到某一特定版本

### 远程仓库

`git remote add origin git@git.acwing.com:xxx/XXX.git`：将本地仓库关联到远程仓库
`git push -u (第一次需要-u以后不需要)` ：将当前分支推送到远程仓库
`git push origin branch_name`：将本地的某个分支推送到远程仓库
`git clone git@git.acwing.com:xxx/XXX.git`：将远程仓库XXX下载到当前目录下
`git push --set-upstream origin branch_name`：设置本地的branch\_name分支对应远程仓库的branch\_name分支
`git push -d origin branch_name`：删除远程仓库的branch_name分支
`git checkout -t origin/branch_name` ：将远程的branch_name分支拉取到本地
`git pull` ：将远程仓库的当前分支与本地仓库的当前分支合并
`git pull origin branch_name`：将远程仓库的branch_name分支与本地仓库的当前分支合并
`git branch --set-upstream-to=origin/branch_name1 branch_name2`：将远程的branch\_name1分支与本地的branch\_name2分支对应

### 分支命令

`git branch branch_name`：创建新分支
`git branch`：查看所有分支和当前所处分支
`git checkout -b branch_name`：创建并切换到branch_name这个分支
`git checkout branch_name`：切换到branch_name这个分支
`git merge branch_name`：将分支branch_name合并到当前分支上
`git branch -d branch_name`：删除本地仓库的branch_name分支
`git push --set-upstream origin branch_name`：设置本地的branch\_name分支对应远程仓库的branch\_name分支
`git push -d origin branch_name`：删除远程仓库的branch_name分支
`git checkout -t origin/branch_name` ：将远程的branch_name分支拉取到本地
`git pull` ：将远程仓库的当前分支与本地仓库的当前分支合并
`git pull origin branch_name`：将远程仓库的branch_name分支与本地仓库的当前分支合并
`git branch --set-upstream-to=origin/branch_name1 branch_name2`：将远程的branch\_name1分支与本地的branch\_name2分支对应

### stash暂存

`git stash`：将工作区和暂存区中尚未提交的修改存入栈中
`git stash apply`：将栈顶存储的修改恢复到当前分支，但不删除栈顶元素
`git stash drop`：删除栈顶存储的修改
`git stash pop`：将栈顶存储的修改恢复到当前分支，同时删除栈顶元素
`git stash list`：查看栈中所有元素

## 修改已经commit的作者信息

1. `git log` 查看目前所有commit id
2. `git rebase -i <最早commit>` 重新设置基准线，选择想要修改部分的上一个commit的commit id，如果是从头修改则为`git rebase -i --root`
3. 将想要修改的部分从`pick`改为`edit`
4. `git commit --amend --author="Author Name <email@address.com>"` 修改commit作者信息
5. `git rebase --continue` 移动到下个commit进行修改，重复4直到完成。
6. `git rebase --continue` 修改完成。

