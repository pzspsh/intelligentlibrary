# CentOS教程
#### CentOS 使用firewalld打开关闭防火墙与端口
#### 1、firewalld的基本使用
```shell
systemctl start firewalld  # 启动 
systemctl stop firewalld  # 关闭
systemctl status firewalld  # 查看状态
systemctl disable firewalld  # 开机禁用 
systemctl enable firewalld  # 开机启用 
```


#### 2.systemctl是CentOS7的服务管理工具中主要的工具，它融合之前service和chkconfig的功能于一体。
```shell
systemctl start firewalld.service   # 启动一个服务
systemctl stop firewalld.service   # 关闭一个服务
systemctl restart firewalld.service   # 重启一个服务
systemctl status firewalld.service   # 显示一个服务的状态
systemctl enable firewalld.service   # 在开机时启用一个服务
systemctl disable firewalld.service   # 在开机时禁用一个服务
systemctl is-enabled firewalld.service   # 查看服务是否开机启动
systemctl list-unit-files|grep enabled   # 查看已启动的服务列表
systemctl --failed   # 查看启动失败的服务列表
```

#### 3.配置firewalld-cmd
```shell
firewall-cmd --version  # 查看版本
firewall-cmd --help   # 查看帮助
firewall-cmd --state   # 显示状态
firewall-cmd --zone=public --list-ports   # 查看所有打开的端口
firewall-cmd --reload   # 更新防火墙规则
firewall-cmd --get-active-zones    # 查看区域信息
firewall-cmd --get-zone-of-interface=eth0   # 查看指定接口所属区域
firewall-cmd --panic-on    # 拒绝所有包
firewall-cmd --panic-off   # 取消拒绝状态
firewall-cmd --query-panic   # 查看是否拒绝

如何开启一个端口:
    firewall-cmd --zone=public --add-port=80/tcp --permanent （–permanent永久生效，没有此参数重启后失效） # 添加
    firewall-cmd --reload   # 重新载入 
    firewall-cmd --zone= public --query-port=80/tcp  # 查看
    firewall-cmd --zone= public --remove-port=80/tcp --permanent  # 删除
```

### 1、基本命令操作
```shell
dnf -y update  # dnf同yum使用方法
nmcli  # 查看ip信息
nmcli c reload  网卡名  # 重启网卡 指定网卡名重启，否则重启所有网卡
nmcli device show  # 取关于已知设备的完整信息
nmcli connection show # 取活动连接配置集的概述
nmcli c up xxx  # 如果之前没有xxx的connection，则上一步reload后就已经自动生效了
nmcli connection show xxx # nmcli connection show xxx
nmcli connection show --active  # 显示所有活动连接
nmcli connection modify "System eth0"  ipv4.addresses 192.168.0.58 # 给xxx添加一个IP（IPADDR）("System eth0"为网卡的连接名,可改为自己的信息)
nmcli connection modify "System eth0"  ipv4.addresses 192.168.0.58/24 # 给xxx添加一个子网掩码（NETMASK）
nmcli connection modify xxx ipv4.dns 114.114.114.114  # 添加DNS
nmcli connection modify xxx -ipv4.dns 114.114.114.114  # 删除DNS
nmcli connection modify xxx ipv4.gateway 192.168.0.2 # 添加一个网关（GATEWAY）

vi /etc/default/grub  # 修改内核参数配置文件
wc -c file_name #查看文件内容大小
```
#### 配置网卡
# vim /etc/sysconfig/network-scripts/ifcfg-xxx # ifcfg-ens33
```shell
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=none
#BOOTPROTO=dhcp
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
NAME=xxx   # ens33
UUID=5c650620-886b-4ca8-8733-40fb2d92d2ce
DEVICE=xxxs # ens33
ONBOOT=yes


IPADDR=ip
PREFIX=24  # 子网掩码的位数长度，取值范围0~32，一般是(8、16、24),根据自己的子网掩码来批准
GATEWAY=网关IP  # 网关
DNS1=223.5.5.5
DNS2=223.6.6.6
```
#### 1.关机 (系统的关机、重启以及登出 ) 的命令
```shell
shutdown -h now # 关闭系统(1)
init 0 # 关闭系统(2)
telinit 0 # 关闭系统(3)
shutdown -h hours:minutes & # 按预定时间关闭系统
shutdown -c # 取消按预定时间关闭系统
shutdown -r now # 重启(1)
reboot # 重启(2)
logout # 注销
```
#### 2.查看系统信息的命令
```shell
arch # 显示机器的处理器架构(1)
uname -m # 显示机器的处理器架构(2)
uname -r # 显示正在使用的内核版本
dmidecode -q # 显示硬件系统部件 - (SMBIOS / DMI)
hdparm -i /dev/hda # 罗列一个磁盘的架构特性
hdparm -tT /dev/sda # 在磁盘上执行测试性读取操作
cat /proc/cpuinfo # 显示CPU info的信息
cat /proc/interrupts # 显示中断
cat /proc/meminfo # 校验内存使用
cat /proc/swaps # 显示哪些swap被使用
cat /proc/version # 显示内核的版本
cat /proc/net/dev # 显示网络适配器及统计
cat /proc/mounts # 显示已加载的文件系统
lspci -tv 罗列 PCI 设备
lsusb -tv 显示 USB 设备
date # 显示系统日期
cal 2007 # 显示2007年的日历表
date 041217002007.00 设置日期和时间 - 月日时分年.秒
clock -w # 将时间修改保存到 BIOS
```
#### 3.cd /home 进入 '/ home' 目录'
```shell
cd .. # 返回上一级目录
cd ../.. # 返回上两级目录
cd # 进入个人的主目录
cd ~user1 # 进入个人的主目录
cd - # 返回上次所在的目录
pwd # 显示工作路径
ls # 查看目录中的文件
ls -F # 查看目录中的文件
ls -l # 显示文件和目录的详细资料
ls -a # 显示隐藏文件 
mkdir dir1 # 创建一个叫做 'dir1' 的目录'
mkdir dir1 dir2 # 同时创建两个目录
mkdir -p /tmp/dir1/dir2 # 创建一个目录树
rm -f file1 # 删除一个叫做 'file1' 的文件'
rmdir dir1 # 删除一个叫做 'dir1' 的目录'
rm -rf dir1 # 删除一个叫做 'dir1' 的目录并同时删除其内容
rm -rf dir1 dir2 # 同时删除两个目录及它们的内容
mv dir1 new_dir # 重命名/移动 一个目录
cp file1 file2 # 复制一个文件
cp dir/* . # 复制一个目录下的所有文件到当前工作目录
cp -a /tmp/dir1 . # 复制一个目录到当前工作目录
cp -a dir1 dir2 # 复制一个目录
ln -s file1 lnk1 # 创建一个指向文件或目录的软链接
ln file1 lnk1 # 创建一个指向文件或目录的物理链接
touch  file1 # 创建一个文件
```
#### 4.查看文件内容
```shell
cat file1 # 从第一个字节开始正向查看文件的内容
tac file1 # 从最后一行开始反向查看一个文件的内容
more file1 # 查看一个长文件的内容
less file1 # 类似于 'more' 命令，但是它允许在文件中和正向操作一样的反向操作
head -2 file1 # 查看一个文件的前两行
tail -2 file1 # 查看一个文件的最后两行 5.挂载命令
mount /dev/hda2 /mnt/hda2 # 挂载一个叫做hda2的盘  （注：确定目录 '/ mnt/hda2' 已经存在）
umount /dev/hda2  # 卸载一个叫做hda2的盘 （先从挂载点 '/ mnt/hda2' 退出）
fuser -km /mnt/hda2 # 当设备繁忙时强制卸载
umount -n /mnt/hda2 # 运行卸载操作而不写入 /etc/mtab 文件（当文件为只读或当磁盘写满时非常有用）
mount /dev/fd0 /mnt/floppy # 挂载一个软盘
mount /dev/cdrom /mnt/cdrom # 挂载一个光盘
mount /dev/hdc /mnt/cdrecorder # 挂载一个cdrw或dvdrom
mount /dev/hdb /mnt/cdrecorder # 挂载一个cdrw或dvdrom
mount -o loop file.iso /mnt/cdrom # 挂载一个文件或ISO镜像文件
mount -t vfat /dev/hda5 /mnt/hda5 # 挂载一个Windows FAT32文件系统
mount /dev/sda1 /mnt/usbdisk # 挂载一个usb 捷盘或闪存设备
mount -t smbfs -o username=user,password=pass //WinClient/share /mnt/share # 挂载一个windows网络共享

```
#### 5.磁盘空间操作的命令
```shell
free -mh # 查看磁盘以及分区情况
df -h # 显示已经挂载的分区列表
df -ah # 人性化显示各存储空间大小
df -aT # 显示所有存储系统空间使用情况,同时显示存储系统的文件系统类型
df -ahlT # 查看本地文件，不显示网络磁盘
ls -lSr |more # 以尺寸大小排列文件和目录
du -sh dir1 # 估算目录 'dir1' 已经使用的磁盘空间'
du -sk * | sort -rn # 以容量大小为依据依次显示文件和目录的大小
du -h --max-depth=1 /home # 查看home文件夹的空间使用情况
du -ch # 看当前文件及文件中包含的子文件夹大小
du -h test1.txt # 查看某个文件容量大小
du -h test1.txt test2.txt # 查看多个文件容量大小
```
#### 6.用户和群组相关命令
```shell
groupadd group_name # 创建一个新用户组
groupdel group_name # 删除一个用户组
groupmod -n new_group_name old_group_name # 重命名一个用户组
useradd -c "Name Surname " -g admin -d /home/user1 -s /bin/bash user1 # 创建一个属于 "admin" 用户组的用户
useradd user1 # 创建一个新用户
userdel -r user1 # 删除一个用户 ( '-r' 同时删除除主目录)
passwd user1 # 修改一个用户的口令 (只允许root执行)
chage -E 2005-12-31 user1 # 设置用户口令的失效期限
ls -lh # 显示权限
chmod 777 directory1 # 设置目录的所有人(u)、群组(g)以及其他人(o)以读（r ）、写(w)和执行(x)的权限
chmod 700 directory1 # 删除群组(g)与其他人(o)对目录的读写执行权限
chown user1 file1 # 改变一个文件的所有人属性，为use1。
chown -R user1 directory1 # 改变一个目录的所有人属性并同时改变改目录下所有文件的属性都为use1所有
chgrp group1 file1 # 改变文件的群组为group1
chown user1:group1 file1 # 改变一个文件的所有人和群组属性，所属组为group1，用户为use1。
find / -perm -u+s # 罗列一个系统中所有使用了SUID控制的文件
chmod u+s /bin/file1 # 设置一个二进制文件的 SUID 位 - 运行该文件的用户也被赋予和所有者同样的权限
chmod u-s /bin/file1 # 禁用一个二进制文件的 SUID位
chmod g+s /home/public # 设置一个目录的SGID 位 - 类似SUID ，不过这是针对目录的
chmod g-s /home/public # 禁用一个目录的 SGID 位
chmod o+t /home/public # 设置一个文件的 STIKY 位 - 只允许合法所有人删除文件
chmod o-t /home/public # 禁用一个目录的 STIKY 位
```
#### 7.关于RPM 包的命令
```shell
rpm -ivh package.rpm # 安装一个rpm包
rpm -ivh --nodeeps package.rpm # 安装一个rpm包而忽略依赖关系警告
rpm -U package.rpm # 更新一个rpm包但不改变其配置文件
rpm -F package.rpm # 更新一个确定已经安装的rpm包
rpm -e package_name.rpm # 删除一个rpm包
rpm -qa # 显示系统中所有已经安装的rpm包
rpm -qa | grep httpd # 显示所有名称中包含 "httpd" 字样的rpm包
rpm -qi package_name # 获取一个已安装包的特殊信息
rpm -ql package_name # 显示一个已经安装的rpm包提供的文件列表
rpm -qc package_name # 显示一个已经安装的rpm包提供的配置文件列表
rpm -q package_name --whatrequires # 显示与一个rpm包存在依赖关系的列表
rpm -q package_name --whatprovides # 显示一个rpm包所占的体积
rpm -q package_name --scripts # 显示在安装/删除期间所执行的脚本l
rpm -q package_name --changelog # 显示一个rpm包的修改历史
rpm -qf /etc/httpd/conf/httpd.conf # 确认所给的文件由哪个rpm包所提供
rpm -qp package.rpm -l # 显示由一个尚未安装的rpm包提供的文件列表
rpm --import /media/cdrom/RPM-GPG-KEY # 导入公钥数字证书
rpm --checksig package.rpm # 确认一个rpm包的完整性
rpm -qa gpg-pubkey # 确认已安装的所有rpm包的完整性
rpm -V package_name # 检查文件尺寸、 许可、类型、所有者、群组、MD5检查以及最后修改时间
rpm -Va # 检查系统中所有已安装的rpm包- 小心使用
rpm -Vp package.rpm # 确认一个rpm包还未安装
rpm2cpio package.rpm | cpio --extract --make-directories *bin* # 从一个rpm包运行可执行文件
rpm -ivh /usr/src/redhat/RPMS/`arch`/package.rpm # 从一个rpm源码安装一个构建好的包
rpmbuild --rebuild package_name.src.rpm # 从一个rpm源码构建一个 rpm 包
```
#### 8.YUM 软件包升级器 
```shell
yum install package_name # 下载并安装一个rpm包
yum localinstall package_name.rpm # 将安装一个rpm包，使用你自己的软件仓库为你解决所有依赖关系
yum update package_name.rpm # 更新当前系统中所有安装的rpm包
yum update package_name # 更新一个rpm包
yum remove package_name # 删除一个rpm包
yum list # 列出当前系统中安装的所有包
yum search package_name # 在rpm仓库中搜寻软件包
yum clean packages # 清理rpm缓存删除下载的包
yum clean headers # 删除所有头文件
yum clean all # 删除所有缓存的包和头文件
```
#### 9.防火墙
```shell
service iptables status # 查看防火墙状态
service iptables start #打开防火墙
service iptables stop #关闭防火墙
iptables -I INPUT -p tcp –dport 80 -j ACCEPT #允许80端口tcp连接, 添加防火墙规则
```
### 2、镜像安装
