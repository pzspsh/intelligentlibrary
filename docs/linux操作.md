## linux 操作命令：
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
```