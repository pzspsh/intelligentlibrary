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
ln -s [原文件或目录] [软链接名] 
history -  # 显示所有的历史命令 
history 10 -  # 显示最近使用过的10个指令

```
#### 2.用户常用指令
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
#### 3.用户组常用命令
```shell
groupadd xxxxx  # 添加用户组
usermod -g usergroup username  # 修改用户组
groupdel xxxxx  # 删除用户组
```
#### 4.文件/文件夹管理
```shell
ls   # 列出当前目录文件（不包括隐含文件）
ls -a   # 列出当前目录文件（包括隐含文件）
ls -l   # 列出当前目录下文件的详细信息

cd ..   # 回当前目录的上一级目录
cd -   # 回上一次所在的目录
cd ~   # 或 cd 回当前用户的宿主目录
mkdir 目录名   # 创建一个目录
rmdir 空目录名   # 删除一个空目录
rm 文件名 文件名   # 删除一个文件或多个文件
rm -rf 非空目录名   # 删除一个非空目录下的一切

mv 路经/文件   # /经/文件移动相对路经下的文件到绝对路经下
mv 文件名 新名称   # 在当前目录下改名
find 路经 -name “字符串”   # 查找路经所在范围内满足字符串匹配的文件和目录处
```

#### 5.系统管理
```shell
fdisk -l      # 查看系统分区信息
fdisk /dev/sdb      # 为一块新的SCSI硬盘进行分区
chown root /home      # 把/home的属主改成root用户
chgrp root /home      # 把/home的属组改成root组
Useradd              # 创建一个新的用户
Groupadd 组名     # 创建一个新的组
Passwd 用户名      # 为用户创建密码
Passwd -d用户名      # 删除用户密码也能登陆
Passwd -S用户名      # 查询账号密码
Usermod -l 新用户名 老用户名    # 为用户改名
uname -a     # 查看内核版本

cat /etc/issue      # 查看ubuntu版本
lsusb      # 查看usb设备
sudo ethtool eth0      # 查看网卡状态
cat /proc/cpuinfo      # 查看cpu信息
lshw     # 查看当前硬件信息
sudo fdisk -l      # 查看磁盘信息
df -h      #查看硬盘剩余空间
free -m      # 查看当前的内存使用情况
ps -A      # 查看当前有哪些进程
kill 进程号(就是ps -A中的第一列的数字)或者 killall 进程名( 杀死一个进程)
kill -9 进程号      # 强制杀死一个进程
```
#### 6.make编译
```shell
make # 编译
make install # 安装编译好的源码包
```
#### 7.修改网络配置
nano /etc/netplan/xxx-netcfg.yaml
```shell
# This is the network config written by 'subiquity'
network:
  ethernets:
    ens33: # ens33 
      dhcp4: true
      dhcp6: false
      addresses: [ip/子网掩码位数]
      gateway4: 网关IP
      nameservers:
              addresses: [223.5.5.5, 223.6.6.6]
  version: 2
```
sudo netplan apply
 
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
