## 1.linux 操作命令：
```shell
ls   # 查看可显示文件
cat  # 查看文件内容
pwd  # 查看路径
top  # 查看cpu等使用状况
ls -alh # 查看文件及权限等
mv 源文件/文件夹  新文件/文件夹  # 迁移文件
ps -ef | grep python  # 查看python运行进程id号
touch 1.txt  # 创建1.txt文件
mkdir test   # 创建test文件夹
kill -9 23423 # 关闭进程id号为23423的进程
nohup python3 main.py > /dev/null 2>&1 &  # 后台运行python3 main.py的程序
wget targeturl     # 下载targeturl文件
df -hl   # 查看磁盘剩余空间
df -h    # 查看每个根路径的分区大小
du -sh [目录名]    # 返回该目录的大小，其中[目录名]可缺省，默认输出当前文件夹的总大小
du -sm [文件夹]    # 返回该文件夹总数
du -h [目录名]    # 查看指定文件夹下的所有文件大小（包含子文件夹），其中[目录名]可缺省，默认当前文件夹
du -h –max-depth=1 [目录名]    # 查看当前文件夹内部每个子文件夹的大小。max-depth可自行设置，这个参数定义了你希望看到多深的文件夹结构级别的输出
clear  # 清屏
uname -a # 显示byname 、版本、 硬件架构等信息
uname -p # 查看当前处理器类型，i686为32，x86_64为64
cat /etc/issue # 显示操作系统及版本
cat /etc/redhat-release # 显示redhat版本
cat /proc/version # 内核版本
cat /etc/issue # 查看操作系统版本和发行版本
cat /etc/redhat-release # 查看redhat的版本
cat /etc/lsb-release # 查看Ubuntu的详细版本
cat /etc/os-release # 查看操作系统基本版本
cat /proc/cpuinfo # 查看cpu信息，型号、cpu个数、频率等
cat /etc/passwd # 查看系统中的用户
cat /etc/shadow # 查看系统中的用户的口令
lscpu # 查看cpu范围
grep ^model name /proc/cpuinfo # 查看cpu型号
dmidecode -q # 查看Linux 内核版本
getconf LONG_BIT # 查看32/64位
hostname # 显示系统本地名称
dmesg | grep Memory # 显示内存条信息
ifconfig # 查询主机已经连接的IP地址
ip addr # 查询本机网卡的IP地址
fdisk -l # 查看硬盘正在使用的分区情况
lsblk # 查看本地存储设备的信息
last # 查看操作系统开机日期
ps aux # 可以查看当前系统中服务程序
netstat -anp # 查看服务的运行状态
netstat -a # 查看本机当前网络全部状态
netstat -an # 查看本机当前网络连接状态
netstat -rn # 查看本机路由信息
netstat -ap # 查看服务器端口的连接情况
netstat -ntlp # 查看某个服务占用的端口
ethtool eth0 # 查看网卡的具体设置a
chkconfig -list # 查看系统自启动服务
chkconfig --list | grep on # 查看已经开启的自启动服务
service httpd start # 启动服务
service httpd restart # 重启服务
service httpd stop # 关闭服务
shutdown -r now # 立刻重启
shutdown -r +10 # 10分钟后重启
cat /var/log/messages # 查看各种系统日志
ls /etc/init.d # 查看系统中所有服务脚本
tail -f log_file_name # 查看自定义的日志信息

sudo shutdown now # 立即关机
```
#### 2.远程命令
```shell
ssh 用户名@ip地址    # 远程连接
scp 本地文件 远程服务器用户名@远程服务器ip地址:指定拷贝到远程服务器的路径  # 远程拷贝文件
scp 远程服务器用户名@远程服务器ip地址:远程服务器文件 指定拷贝到本地电脑的路径  # 远程拷贝文件

scp -r 本地目录 远程服务器用户名@远程服务器ip地址:指定拷贝到远程服务器的路径 # 远程拷贝目标,-r 表示递归拷贝整个目录
scp -r 远程服务器用户名@远程服务器ip地址:远程服务器目录 指定拷贝到本地电脑的路径 # 远程拷贝目标,-r 表示递归拷贝整个目录
scp -r /home/<files> <remote-username>@<remote-ip>:<remote-folder> # ssh上传文件
```
#### 3.vim使用
```shell
编辑模式: i Esc
末行模式: Esc :
保存方式：
    :w  # 保存
    :wq # 保存退出
    :x # 保存退出
    :q! # 强制退出

vim 的常用命令:
    yy  # 复制光标所在行
    p  # 粘贴
    dd  # 删除/剪切当前行
    V  # 按行选中
    u  # 撤销
    ctrl+r  # 反撤销 
    >>  #  往右缩进
    <<  # 往左缩进
    :/搜索的内容  # 搜索指定内容 
    :%s/要替换的内容/替换后的内容/g  # 全局替换
    :开始行数，结束行数s/要替换的内容/替换后的内容  # 局部替换
    .  # 重复上一次命令操作
    G  # 回到最后一行
    gg  # 回到第一行
    数字+G  # 回到指定行
    shift+6  # 回到当前行的行首
    shift+4  # 回到当前的行末
    ctrl+f  # 下一屏
    ctrl+b  # 上一屏
```
#### 4.top命令
```shell
-b  # 以批处理模式操作；
-c  # 显示完整的命令；
-d  # 屏幕刷新间隔时间；
-I  # 忽略失效过程；
-s  # 保密模式；
-S  # 累积模式；
-i<时间>  # 设置间隔时间；
-u<用户名>  # 指定用户名；
-p<进程号>  # 指定进程；svn
-n<次数>  # 循环显示的次数

top -d 10
```
#### 5.kill命令
```shell
kill  [选项] 进程号 # 杀死进程
killall 进程名称 #（通过进程名称杀死进程，也支持通配符，这在系统因负载过大而变得很慢时很有用）
-9 : 表示强迫进程立即停止
```
####  6.ps -ef 命令
ps -ef 是以全格式显示当前所有的进程
-e 显示所有进程。
-f 全格式
```shell
ps -ef|grep xxxx
显示：
    UID：用户 ID
    PID：进程 ID
    PPID：父进程 ID
C：CPU 用于计算执行优先级的因子。数值越大，表明进程是 CPU 密集型运算，执行优先级会降低；数值越小，表明进程是 I/O 密集型运算，执行优先级会提高
STIME：进程启动的时间
TTY：完整的终端名称
TIME：CPU 时间
CMD：启动进程所用的命令和参数 
```

#### 7.chmod命令
chmod- 修改权限
通过chmod指令，可以修改文件或者目录的权限。
```shell
第一种修改方式：+ 、-、= 变更权限
u:所有者
g:所有组
o:其他人
a:所有人(u、g、o 的总和)
1. chmod u=rwx,g=rx,o=x 文件/目录名
2. chmod o+w  文件/目录名
3. chmod a-x  文件/目录名

1.给abc文件的所有者读写执行的权限，给所在组读执行权限，给其它组读执行权限。
chmod u=rwx,g=rx,o=rx abc
2.给abc文件的所有者除去执行的权限，增加组写的权限
chmod u-x,g+w abc
3.给abc文件的所有用户添加读的权限
chmod a+r abc


第二种修改方式：通过数字变更权限
r=4 w=2 x=1
rwx=4+2+1=7
chmod u=rwx,g=rx,o=x 文件/目录名 相当于 chmod 751 文件/目录名

1.将/home/abc.txt 文件的权限修改成 rwxr-xr-x, 使用给数字的方式实现：
　 chmod 755 /home/abc.txt
```
#### 8.打包和解压缩文件的命令
```shell
bunzip2 file1.bz2 # 解压一个叫做 'file1.bz2'的文件
bzip2 file1 # 压缩一个叫做 'file1' 的文件
gunzip file1.gz # 解压一个叫做 'file1.gz'的文件
gzip file1 # 压缩一个叫做 'file1'的文件
gzip -9 file1 # 最大程度压缩
rar a file1.rar test_file # 创建一个叫做 'file1.rar' 的包
rar a file1.rar file1 file2 dir1 # 打包 'file1', 'file2' 以及目录 'dir1'
rar x file1.rar # 解rar包
unrar x file1.rar # 解rar包
tar -cvf archive.tar file1 # 创建一个非压缩的tar包
tar -cvf archive.tar file1 file2 dir1 # 创建一个包含了 'file1', 'file2' 'dir1'的包
tar -tf archive.tar # 显示一个包中的内容
tar -xvf archive.tar # 释放一个包
tar -xvf archive.tar -C /tmp # 将压缩包释放到 /tmp目录下 （-c是指定目录）
tar -cvfj archive.tar.bz2 dir1 # 创建一个bzip2格式的压缩包
tar -xvfj archive.tar.bz2 # 解压一个bzip2格式的压缩包
tar -cvfz archive.tar.gz dir1 # 创建一个gzip格式的压缩包
tar -xvfz archive.tar.gz # 解压一个gzip格式的压缩包
zip file1.zip file1 # 创建一个zip格式的压缩包
zip -r file1.zip file1 file2 dir1 # 将几个文件和目录同时压缩成一个zip格式的压缩包
unzip file1.zip # 解压一个zip格式压缩包
```

#### 9.文件搜索命令
```shell
find / -name file1 # 从 '/' 开始进入根文件系统搜索文件和目录
find / -user user1 # 搜索属于用户 'user1' 的文件和目录
find  /home/user1  -name  \*.bin # 在目录 '/ home/user1' 中搜索带有'.bin' 结尾的文件
find /usr/bin -type f -atime +100 # 搜索在过去100天内未被使用过的执行文件
find /usr/bin -type f -mtime -10 # 搜索在10天内被创建或者修改过的文件
locate \*.ps  # 寻找以 '.ps' 结尾的文件 - 先运行 'updatedb' 命令
whereis file # 显示一个二进制文件、源码或man的位置
which file  # 显示一个二进制文件或可执行文件的完整路径
```

#### 10、复制文件/文件命令
```shell
cp -r folder1 folder2‌ # 复制一个文件夹到另一个目录
cp -r /path/to/source_folder /path/to/destination_folder # 复制一个文件夹到另一个目录
cp -rp /path/to/source_folder /path/to/destination_folder # 如果你想保留源文件的时间戳和权限，可以添加-p选项
cp source.txt destination.txt # ‌复制单个文件
cp file1.txt file2.txt /path/to/destination/ # 复制多个文件‌
cp -r source_directory/ /path/to/destination/ # ‌递归复制目录‌
cp -i file.txt /path/to/destination/ # 交互式复制‌
cp -u file.txt /path/to/destination/ # 仅复制更新的文件‌
cp -v file.txt /path/to/destination/ # ‌显示详细输出

‌-a‌：此选项保留源文件的属性，包括权限、所有者、时间戳等，通常用于复制目录时使用。
‌-f‌：强制复制，不提示是否覆盖目标文件。
‌-i‌：在覆盖目标文件前询问用户确认，避免意外覆盖。
‌-p‌：保留源文件的权限、时间戳等信息，确保复制后的文件与原始文件保持一致。
‌-r‌：递归复制目录及其所有子目录和文件。
‌-u‌：仅复制源文件中更新的部分，如果源文件比目标文件新，则进行复制。
‌-v‌：显示详细输出，列出已复制的文件，便于跟踪复制进度。
```

```bash
netstat -tulnp | grep :8080 # 查找并停止占用端口的应用
```