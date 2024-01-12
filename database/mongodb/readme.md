# mongodb 数据库

### 一、mongodb 数据库安装

#### 1、[官网下载 mongodb 数据软件包](https://www.mongodb.com/download-center/community/releases/archive)

#### 2、解压 mongodb 数据包

#### 3、到解压的 mongodb 文件夹目录下创建 data 文件夹

#### 4、打开 cmd 终端切换到解压的 mongodb 文件夹的 bin 目录下

#### 5、启动 mongodb 数据库服务

```
1、前置启动命令为：
bin>mongod --dbpath 上面第三步创建的data文件夹路径 # 例如：D:\\mongodb-win32-x86_64-windows-5.0.3\data

2、后置启动命令为：
bin>mongod --dbpath data文件夹路径 --logpath mongodb.log文件路径 --fork --logappend
```

#### 6、编辑配置文件(进行服务启动)

```
编辑配置文件
mkdir /path/mongodb/etc/ # path路径
vi /path/mongodb/etc/mongodb.conf

dbpath=/path/mongodb/data/db/
logpath=/path/mongodb/logs/mongodb.log
logappend=true
bind_ip_all=true
port=27017
fork=true

/path/mongodb/bin/mongod --config /path/mongodb/etc/mongodb.conf或
/path/mongodb/bin/mongod -f /path/mongodb/etc/mongodb.conf
```

#### 7、打开 cmd 终端切换到解压的 mongodb 文件夹的 bin 目录下执行如下命令

```
1、执行进入数据库命令
bin>mongo

2、就可操作你要执行的命令
```

#### 8、设置密码

```
执行如下命令，在 admin 库中创建一个超级用户（对所有数据库都有读写权限）：
MongoDB 中内置角色说明：
read：提供读取所有非系统的集合（数据库）
readWrite：提供读写所有非系统的集合（数据库）和读取所有角色的所有权限
dbAdmin：提供执行管理任务的功能，例如与架构相关的任务，索引编制，收集统计信息。此角色不授予用户和角色管理权限。
dbOwner：提供对数据库执行任何管理操作的功能。此角色组合了readWrite，dbAdmin 和 userAdmin 角色授予的权限。
userAdmin ：提供在当前数据库上创建和修改角色和用户的功能。由于 userAdmin 角色允许用户向任何用户（包括他们自己）授予任何权限，因此该角色还间接提供对数据库的超级用户访问权限，或者，如果作用于管理数据库，则提供对群集的访问权限。
clusterAdmin ：提供最佳的集群管理访问。此角色组合了 clusterManager，clusterMonitor 和 hostManager 角色授予的权限。此外，该角色还提供了 dropDatabase 操作。
readAnyDatabase ：仅在 admin 数据库中使用，提供所有数据库的读权限。
readWriteAnyDatabase ：仅在 admin 数据库中使用，提供所有数据库的读写权限
userAdminAnyDatabase ：仅在 admin 数据库中使用，提供与 userAdmin 相同的用户管理操作访问权限，允许用户向任何用户（包括他们自己）授予任何权限，因此该角色还间接提供超级用户访问权限。
dbAdminAnyDatabase ：仅在 admin 数据库中使用，提供与 dbAdmin 相同的数据库管理操作访问权限，该角色还在整个群集上提供 listDatabases 操作。
root：仅在 admin 数据库中使用，提供超级权限

方式一：
use admin
db.createUser(
  {
    user: "root",
    pwd: "123",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)

方式二：
use admin
db.createUser({
  user: 'admin',    // 用户名（自定义）
  pwd: 'Abc123++',  // 密码（自定义）
  roles:[{
    role: 'root',   // 使用超级用户角色
    db: 'admin'     // 指定数据库
  }]
})

1、设置完成，可以通过指令 show users 查看是否设置成功。

2、开启权限验证
找到MongoDB安装目录下的bin目录中的mongod.cfg文件，开启权限验证功能：
security:
  authorization: enabled

3、重启MongoDB服务

4、使用超级管理员账号登录数据库的方式：
方式一：
bin>mongo
>use admin
>db.auth('admin', 'Abc123++')

方式二：
bin>mongo admin -u admin -p Abc123++

5、为数据库设置独立的登录账号
除了设置超级管理员账号以外，还可以为每个数据库单独设置账号。
例如以下指令：
use myMongoDB  // 跳转到需要添加用户的数据库
db.createUser({
  user: 'tao',          // 用户名
  pwd: 'Abc123++',      // 密码
  roles:[{
    role: 'readWrite',  // 读写权限角色
    db: 'myMongoDB'     // 数据库名
  }]
})
```

#### 9、常用命令

```
show users  // 查看当前库下的用户
db.dropUser('testadmin')  // 删除用户
db.updateUser('admin', {pwd: '654321'})  // 修改用户密码
db.auth('admin', '654321')  // 密码认证
```
