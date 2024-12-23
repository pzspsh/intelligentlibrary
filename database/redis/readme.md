# readi 数据库

### 一、redis 数据安装

#### 1、[官网下载 zip 包](https://redis.io/download/#redis-downloads)

#### 2、解压 redis zip 包

#### 3、打开 cmd 终端切换到解压的 redis 文件夹目录下

#### 4、执行启动 redis 服务

```
1、执行启动：
Redis>redis-server
# 指定端口启动 Redis>redis-server --port 9999

2、再打开另一个cmd终端执行如下：
Redis>redis-cli

3、执行你要执行的命令
```

#### 5、配置 redis.conf 配置文件

#### 6、配置 redis 进入命名

```
1、进入配置文件把requirepass的注释去掉
2、通过命令查看：
127.0.0.1:6379>config get requirepass

3、通过命令设置密码：
127.0.0.1:6379>config set requirepass 你设置的密码

4、 设置密码后退出，再次登录:
方式一：
Redis>redis-cli -p 6379 -a 你的密码

方式二：
Read>redis-cli
127.0.0.1:6379>auth 你的密码
```
