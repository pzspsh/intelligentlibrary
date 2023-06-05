# mysql数据库
## 一、安装数据库
### 1、mysql数据库官网下载：[下载-https://dev.mysql.com/downloads/installer/](https://dev.mysql.com/downloads/installer/)
![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/mysql下载.png)

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/mysql下载-1.png)

### 2、把下载的mysql zip包解压
### 3、解压的zip点击到bin路径下
### 4、复制mysql包的bin路径
### 5、配置系统环境变量(如下图所示)：

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/环境变量.png)

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/环境变量-1.png)


这是我安装的数据库，和上图下载的版本不是同一个版本，但操作一样。

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/环境变量-2.png)

### 6、创建data文件(如下图)：

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/mysqldata.png)

### 7、配置mysql的mysql.ini文件(如果没有mysql.ini文件，直接创建就行，并填上一下的数据对应内容):

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/mysqlini.png)
```
[mysqld]
#设置3306端口
port=3306
#设置mysql的安装目录 ----------是你的文件路径
#例如
basedir=D:\Program Files\mysql-8.0.28-winx64 # 修改成你的文件路径
#设置mysql数据库的数据的存放目录 ---------是你的文件路径data文件夹自行创建
datadir=D:\Program Files\mysql-8.0.28-winx64\data # 修改成你的文件路径
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
### 8、初始化mysql数据库
```
mysqld --initialize --user=mysql --console # 回车

注意得到的数据中有你的初始密码。这个一定要记住，可先复制到文本中保存下来。
数据中密码出现的形式是：root@localhost:“密码”

然后执行：mysqld --install 
然后开启mysql服务：net start mysql
```
### 9、然后就可以登录mysql数据库：
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