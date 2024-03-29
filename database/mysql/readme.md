# mysql 数据库

## 一、安装数据库

### 1、mysql 数据库官网下载：[下载-https://dev.mysql.com/downloads/installer/](https://dev.mysql.com/downloads/installer/)

![img](../../images/mysql下载.png)

![img](../../images/mysql下载-1.png)

### 2、把下载的 mysql zip 包解压

### 3、解压的 zip 点击到 bin 路径下

### 4、复制 mysql 包的 bin 路径

### 5、配置系统环境变量(如下图所示)：

![img](../../images/环境变量.png)

![img](../../images/环境变量-1.png)

这是我安装的数据库，和上图下载的版本不是同一个版本，但操作一样。

![img](../../images/环境变量-2.png)

### 6、创建 data 文件(如下图)：

![img](../../images/mysqldata.png)

### 7、配置 mysql 的 mysql.ini 文件(如果没有 mysql.ini 文件，直接创建就行，并填上一下的数据对应内容):

![img](../../images/mysqlini.png)

```
[mysqld]
#设置3306端口
port=3306
#设置mysql的安装目录 ----------是你的文件路径
#例如
basedir=D:\Program Files\mysql-8.0.28-winx64 # 修改成你的文件夹路径
#设置mysql数据库的数据的存放目录 ---------是你的文件路径data文件夹自行创建
datadir=D:\Program Files\mysql-8.0.28-winx64\data # 修改成你的文件夹路径
#允许最大连接数
max_connections=200
#允许连接失败的次数。
max_connect_errors=10
#服务端使用的字符集默认为utf8mb4
character-set-server=utf8mb4
#创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
#默认使用“mysql_native_password”插件认证
#mysql_native_password
default_authentication_plugin=mysql_native_password
[mysql]
#设置mysql客户端默认字符集
default-character-set=utf8mb4
[client]
#设置mysql客户端连接服务端时默认使用的端口
port=3306
default-character-set=utf8mb4
```

### 8、初始化 mysql 数据库

```
mysqld --initialize --user=mysql --console # 回车

注意得到的数据中有你的初始密码。这个一定要记住，可先复制到文本中保存下来。
数据中密码出现的形式是：root@localhost:“密码”

然后执行：mysqld --install
然后开启mysql服务：net start mysql
```

### 9、然后就可以登录 mysql 数据库：

```
方式1：mysql -u root -p # 回车然后输入刚刚你保存文本的那个密码
方式2：mysql -u root -p你的密码
```

### 10、别忘了修改密码：

```
1、登录进数据库
2、set password='你的密码'
然后执行quit退出来，再执行第九步
```

### 11、mysql 开启远程访问权限

```sql
默认情况下，mysql只允许本地登录，即只能在安装MySQL环境所在的主机下访问。
1、打开终端进入mysql
>mysql -u 用户名 -p # 回车然后输入密码
2、查看数据库
mysql>show databases;
3、查看MySQL当前远程访问权限配置
mysql>use mysql;
mysql>select  User,authentication_string,Host from user;
4、开启远程访问权限
方式一：修改命令如下
mysql>update user set host='%' where user='用户名' #
方式二：授权法
通过GRANT命令可以授予主机远程访问权限
--赋予任何主机访问权限：
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'password' WITH GRANT OPTION;
--允许指定主机(IP地址)访问权限：
GRANT ALL PRIVILEGES ON *.* TO 'myuser'@'ip' IDENTIFIED BY 'root' WITH GRANT OPTION; # ip是指你允许该ip访问数据库，如：允许 192.168.100.3 的ip就把ip替换成 192.168.100.3

通过GRANT命令赋权后,需要通过FLUSH PRIVILEGES刷新权限表使修改生效：
flush privileges;

5、再次查看MySQL远程访问权限配置
select  User,authentication_string,Host from user;

注意:
出于安全性考虑，尤其是正式环境下
1.不推荐直接给root开启远程访问权限。
本案例仅以root用户作为例子做开启远程访问权限的配置,此为演示环境!

2.建议做权限细分和限制
正式环境中，推荐通过创建Mysql用户并给对应的用户赋权的形式来开放远程服务权限，并指定IP地址，赋权时根据用户需求，在GRANT命令中只开放slect、update等权限，做到权限粒度最小化。
```

## 二、数据库操作

### 1、创建数据库

```sql

```

### 2、创建数据库表

```sql

```

### 3、添加数据

```sql

```

### 4、删除数据

```sql

```

### 5、更新数据

```sql

```

### 6、查询数据

```sql

```
