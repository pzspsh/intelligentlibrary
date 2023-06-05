# postgres数据库
### 一、Postgres数据库安装
#### 1、[官网下载postgres包](https://www.postgresql.org/docs/release/)
#### 2、解压postgres
#### 3、到解压的postgres文件夹目录下创建data文件夹
#### 4、打开cmd终端切换到解压的posgres文件夹目录下
#### 5、执行启动postgres服务
```
1、终端cmd切换到解压的postgres问bin目录下执行如下：
bin>pg_ctl.exe start -D 你上面创建data文件夹路径 # 例如：D:\\pgsql\data

2、终端cmd切换到解压的postgres问bin目录下执行如下:
bin>psql -U postgres  #以postgres用户登录psql
用户 postgres 的口令：输入你的密码
psql (14.0)
输入 "help" 来获取帮助信息.

postgres=#   #这一步就可以你要执行的操作了
```