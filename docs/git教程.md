# Git教程
```shell
1、克隆URL：
    git clone 目标url 别名
2、撤销commit，忽略文件
    # 回退commit到某个版本
    # git log查看commit日志
    git reset --mixed HEAD^
3、忽略路径中的转义字符
    git config --global core.protectNTFS false
4、禁用换行符转换
    git config --global core.autocrlf false
5、中文文件名，乱码问题。设为false的话，就不会对0x80以上的字符进行quot
    git config --global core.quotepath false
```
### 1、安装git 命令行扩展
```shell
    git lfs install
```
### 2、选择您希望Git LFS管理的文件（或直接编辑 .gitattributes）。您可以随时配置其他文件扩展名。这一步成功后会生成一个 .gitattributes 文件
```shell
    //追踪指定目录下的文件
    git lfs track "A/"

    //追踪单个文件
    git lfs track "A/XXXFramework/xxx"

    //追踪指定类型文件
    git lfs track "*.a"
```

### 3、添加 并 commit .gitattributes 文件
```shell
    git add .gitattributes
    git commit -m "提交.gitattributes 文件"
    git push 
```
### 4、现在可以正常添加、提交项目
```shell
    git add .
    git commit -m "提交test大文件文件"
    git push 
```
### 注意
在提交大文件前，最好先把 .gitattributes 文件push到github远程仓库，最后再正常提交项目大文件。
不先把 .gitattributes push到github，还是可能出现push fail的情况。
注意大文件路径不要错了，可以使用相对路径。。