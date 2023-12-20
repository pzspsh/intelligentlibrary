# Git教程
```shell
1、克隆URL：
    git clone 目标url 别名
2、
```

```shell
https://www.ipaddress.com/ 用该url进行查询
140.82.112.4   github.com
151.101.193.194 github.global.ssl.fastly.net

linux版
vim /etc/hosts # 添加以上内容
centos: nmcli c reload  # 

windows版
windows/system32/drivers/etc/hosts

ipconfig /flushdns
```
问题：fatal: 无法访问 'https://github.com/kubernetes/kubernetes/'：OpenSSL SSL_read: error:0A000126:SSL routines::unexpected eof while reading, errno 0
git config --global http.sslVerify "false"
取消设置
git config --global --unset http.sslVerify
git config --global --unset http.proxy
git config --global --unset https.proxy


```bash
Git下载
www.git-scm.com

remote
设置远程地址
git remote add origin holdlg@192.168.10.99:holdlg/front.git

切换远程地址
git remote set-url origin 
example: git remote set-url origin  holdlg@192.168.10.99:/home/holdlg/repo/dashed.git
tag
# 查看
git tag
git show <tag_name>

# 创建
git tag -a v201909 -m "这里是说明"
git tag v201909

# 删除
git tag -d v201909

# 以前的commit 打标签
git tag -a v201903 9fcec02

# 推送
git push orgin <tag_name>
git push origin --tags

# 同步所有tag

git fetch
branch

查看远程分支
git branch -v

查看所有分支
git branch -a

创建本地分支 
git branch local_branch_name 

切换本地分支 
git checkout  local_branch_name  

创建并切换到本地分支
git checkout -b  local_branch_name  

删除本地分支
git branch -d  local_branch_name  

推送本地分支到远程
git push origin <local_branch_name>:<remote_branch_name>

删除远程分支
git push origin :<remote_branch_name>

克隆远程指定分支
git clone -b <branch_name> --single-branch <repo_name>
git clone -b dashed_3.0 --single-branch http://192.168.10.90/server/transfer.git

迁出远程分支到本地
git checkout -b <local_branch_name> <remote_branch_name>
git checkout -b dashed_3.0 origin/dashed_3.0

合并指定分支到本地当前分支, local_branch_name2合并到local_branch_name1
git checkout <local_branch_name1>
git merge <local_branch_name2>

拉取远端origin/master分支并合并到当前分支
git pull origin master
迁移

git clone --bare git:url
proxy

设置
git config –global http.proxy http://127.0.0.1:1080 
git config –global https.proxy https://127.0.0.1:1080 
git config –global http.proxy 'socks5://127.0.0.1:1080'
git config –global https.proxy 'socks5://127.0.0.1:1080'

取消
git config –global –unset http.proxy 
git config –global –unset https.proxy
Git 常见Bug
Permissions 0644
error: Permissions 0644 for '/root/.ssh/id_ed25519' are too open.


@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0644 for '/root/.ssh/id_ed25519' are too open.
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
解决：


chmod 600 /root/.ssh/id_ed25519

git config --global http.sslVerify false
```