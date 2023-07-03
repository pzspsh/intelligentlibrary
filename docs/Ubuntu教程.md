# Ubuntu教程
### 1、命令操作
```shell
clear # 清屏
rm -rf 删除对象  # 删除 
df -h  # 查看整个系统磁盘使用情况   
du -sh *  # 产看某个目录磁盘使用情况 
apt-get install  # 安装对象  # 安装
sudo apt-get update  # 更新源
sudo apt-get install package  # 安装包
sudo apt-get install package --reinstall  # 重新安装包
sudo apt-get -f install  # 修复安装
sudo apt-cache search package  # 搜索软件包
sudo apt-cache show package  # 获取包信息，如说明，大小，版本等
sudo apt-get remove package  # 删除包
sudo apt-get remove package --purge  # 删除包，包括配置文件等
sudo apt-get build-dep package  # 安装相关的编译环境
sudo apt-get upgrade  # 更新已安装的包
sudo apt-get dist-upgrade  # 升级系统
sudo apt-cache depends package  # 了解使用该包依赖那些包
sudo apt-cache rdepends package  # 查看该包被哪些包依赖
sudo apt-get source package  # 下载包源代码
lsof -i:8888   # 查看端口占用 
netstat -tulpn  # 查看tcp和udp端口 
find /home/un/test -mtime +7 -name "*.*" -exec rm -rf {} \;  # 批量删除七天前的文件 
```
#### 1.用户常用指令
```shell
useradd  username  # 创建用户默认在 /home目录下
useradd -d  /xxxx username  # 创建用户并指定目录
passwd  username  # 指定和修改密码
userdel username  # 删除用户
userdel -r username  # 删除用户及目录
id username  # 查询用户信息
whoami  # 查看当前用户
su - username  # 切换用户
```
#### 2.用户组常用命令
```shell
groupadd xxxxx  # 添加用户组
usermod -g usergroup username  # 修改用户组
groupdel xxxxx  # 删除用户组
```
#### 快捷键
```shell
Ctrl+Alt+T   # 启动终端   
Ctrl+Shift+T    # 新建标签页 
Ctrl+Shift+W    # 关闭标签页 
Ctrl+Shift+Q    # 关闭终端窗口
Alt+N    # 切换到第N个标签页（N=0...9）
Ctrl+Shift++    # 放大窗口（包括窗口内的字体）
Ctrl+-    # 缩写窗口（包括窗口内的字体） 
Ctrl+0    # 普通大小（阿拉伯数字 0）    
Ctrl+Shift+C   # 复制 
Ctrl+Shift+V    # 粘贴 
Ctrl+A   # 光标移动到行的开头 
Ctrl+E   # 光标移动到行的结尾 
Ctrl+Left   # 光标移动到上一个单词的词首 
Ctrl+Right   # 光标移动到下一个单词的词尾 
Ctrl+U   # 剪切从行的开头到光标前一个位置的所有字符 
Ctrl+K   # 剪切从光标位置到行末的所有字符 
Ctrl+Y   # 粘贴ctrl+u或ctrl+k剪切的内容 
Ctrl+W   # 删除光标位置的前一个单词 
Ctrl+&   # 删除光标位置的前一个单词 
Ctrl+S   # 暂停屏幕输出 
Ctrl+Q   # 继续屏幕输出 
Alt+F2   # 命令运行对话框 
```

### 2、镜像安装
