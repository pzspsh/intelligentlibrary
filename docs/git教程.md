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