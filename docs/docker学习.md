## 如何用docker来打包镜像
### 如何用docker来打包镜像
### 第一部分
Docker官网：https://www.docker.com/

#### 一、帮助命令
```dockerfile
docker version 			# 显示docker的版本信息
docker info 			# 显示docker的系统信息，包括镜像和容器的数量
docker 命令 --help   	   # 帮助命令
```

#### 二、镜像命令
安装镜像的搜索网址：
https://hub.docker.com/

##### 2.1、下载镜像
```shell
docker images			 # 查看所有本地主机上的镜像
docker search mysql      # 搜索镜像
docker pull --help 		 # 下载镜像
docker pull mysql:8.1    # 指定MySQL的版本
```

##### 2.2、删除镜像
```
docker rmi -f id		 # 删除指定的镜像，remove image，加上镜像ID
docker rmi -f id id id 	 # 删除多个镜像
docker rmi -f $(docker image -aq)   # 批量删除全部的image(-f表示全部删除，即force,-ap查询命令)
```

#### 三、容器命令
下载CentOS镜像
```
docker pull centos
```

新建容器并启动
```
docker run --help
docker run [可选参数] image

# 参数说明
--name="Name"		容器名字  Tomcat1 Tomcat2 用来区分容器
-d 					后台方式运行
-it					使用交互方式运行，进入容器查看内容
-p					指定容器的端口，-p  8080:8080
	-p ip:主机端口：容器端口
	-p 主机端口：容器端口（常用）
	-p 容器端口
-P					随机指定端口
```

启动并进入容器
```
docker run -it centos /bin/bash
```

从容器中退回主机
```
exit			# 直接停止容器并退出
Ctrl + P + Q	# 容器不停止退出
ls
```

列出所有的运行容器
```
docker ps
-a 			# 列出当前正在运行的容器，带出历史运行过的容器
-n=? 		# 显示最近创建的容器
-q			# 只显示容器的编号
docker ps -a
```

删除容器
```
docker rm 容器id						# 删除指定的容器，不能删除正在运行的容器，如果要强制删除，加采纳数rm -f
docker rm -f $(docker ps -aq)	 	 # 删除所有的容器	
docker ps -a -q | xargs docker rm 	 # 删除所有的容器（管道命令）
```

启动和停止容器的操作
```
docker start 容器ID		# 启动容器
docker restart 容器ID		# 重启容器
docker stop 容器ID		# 停止正在运行的容器
docker kill 容器ID		# 强制停止当前容器
```

#### 四、常见其他命令
后台启动容器
```
docker run -d 镜像名
docker run -d centos

# 问题：docker ps，发现centos停止了
# 常见的坑：docker 容器使用后台运行，就必须要有一个前台进程,docker发现没有应用，就会自动停止
# Nginx容器启动后，发现自己没有提供服务，就会立刻停止，就是没有程序运行了
```


查看日志
```
docker logs --help
docker logs -f -t --tail 容器ID			# 容器如果没有日志

# 自己编写shell脚本
"while true;do echo OriginalCoder;sleep 1;done"
docker run -d centos /bin/sh -c "while true;do echo OriginalCoder;sleep 1;done"
docker ps # 查看容器ID

显示日志
-tf   				# 显示日志（-t:时间戳，-f:持续显示日志）
--tail number 		# 要显示的日志条数
docker logs -ft --tails 10 容器ID（10代表打印的条数）
```

查看容器中进程信息ps

```
docker ps 			# 查看容器ID
docker top 容器ID
```

查看镜像的元数据
```
docker inspect --help
docker ps -a
docker inspect 容器ID
```

进入当前正在运行的容器
```
通常情况下，容器都是后台运行的，需要进入容器中，修改一些配置

# 方式一
docker exec -it 容器ID /bin/bash
ls
ps -ef

# 方式二
docker attach 容器ID

区别
# docker exec 		进入容器后开启一个新的终端，可以在里面操作
# docker attach 	进入容器正在执行的终端，不会启动新的进程
```


从容器内拷贝文件到主机上

```
docker cp 容器ID: 容器内路径 目的主机路径
docker ps
docker images
docker run -it centos /bin/bash
docker ps 

docker attach 容器ID
cd /home
ls
容器内新建一个文件

# touch test.java
exit
docker ps -a

# 将这个文件拷贝出来到主机上
docker cp 容器ID:/home/test.java /home
ls
```

安装vim
```
apt-get install vim
```

安装apache2
```
apt-get install apache2 -y
```

安装装php和php插件
```
apt-get install php -y
apt-get install libapache2-mod-php -y  --fix-missing
apt-get install php7.0-mysql
```

## 第二部分
#### 一、实战：安装MySQL
思考：MySQL的数据持久化问题

```
# 获取镜像
docker search mysql
docker pull mysql:8.1
# 运行容器,需要做数据挂载
# 安装启动mysql，需要配置密码（注意点）
# 官方测试：docker run -it --network some-network --rm mysql mysql -hsome-mysql -uexample-user -p

# 运行
-d 后台运行
-p 端口映射
-v 卷挂载
-e 环境配置
--name 容器名字	
docker run -d -p 3310:3306 -v /home/mysql.conf:/etc/mysql/conf.d -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql01 mysql:8.1

#  启动成功之后，可以在Windows本地机测试
```

假设我们将容器删除
```
docker rm -f mysql01
docker ps
docker ps -a
```

回到宿主机查看，发现数据都还在！ 我们挂载到本地的数据卷没有丢失，这就实现了容器数据持久化功能。
#### 二、具名挂载和匿名挂载

```
# 匿名挂载
-v 容器内路径
docker run d -P --name nginx01 -v /etc/nginx nginx

# 查看所有数据卷的情况
docker volume ls
```

#### 三、DockerFile

DockerFile就是用来构建docker镜像的构建文件！命令脚本

```
mkdir docker-test-volume
ls
pwd
cd docker-test-volume/
pwd

# 通过下面这个脚本，可以生成镜像，镜像是一层一层的，脚本是一层层的命令，每一个命令都是一层
vim dockerfile1
	FROM centos
	VOLUME ['volume01','volume02']
	CMD echo "----end----"
	CMD /bin/bash
	ESC
	:wq

# 查看dockerfile1
cat dockerfile1

# 
docker build -f /home/docker-test--volume/docekrfile1 -t rich/centos:1.0 .
```

启动自己写的容器
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649901176056.png)
这个卷和外部一定有一个同步的目录

四、打包Docker镜像
1、构建SpringBoot项目
2、打包应用
3、编写DockerFile
4、构建镜像
5、发布运行
以后在使用Docker的时候，给别人一个Docker的镜像就可以



# 【Docker】从零开始将自己的应用打包到docker镜像

背景是这样：
有一个python写的web服务，希望打包到容器中，通过容器去启动。
参考了网上各种文档，都感到说的不清不楚，实际操作过程中，又遇到了不少的坑，这里摸索OK后记录一下。
docker的安装和部署此处不再赘述。以下从docker安装完成后开始讲。 
首先，我们写一个demo，使用了python的flask框架，文件名为  app.py  。
我们的目的是，将这个代码以服务的形式，打包到docker镜像中。

```python
from flask import Flask

app=Flask(__name__)
@app.route('/')
def hello():
    return 'hello world'

if __name__=='__main__':
    app.run(host='0.0.0.0',debug=True,port='7777')
```
本地路径如下图
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649902631.png)
可以看到，最外层目录是  mydocker ，内部是bdtools，app.py就放置在最内层。
首先，requirements.txt的内容如下图，这为了安装python依赖包：
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649902727.png)
然后我们开始编写Dockerfile

```
FROM python:3.8 　　　　
# 拉取一个基础镜像，基于python3.8
MAINTAINER BH8ANK　　　　
# 维护者信息
ADD ./bdtools/ /code/bdtools/ 　
# 将你的项目文件放到docker容器中的/code/bdtools文件夹，这里code是在根目录的，与/root /opt等在一个目录
# 这里的路径，可以自定义设置，主要是为了方便对项目进行管理
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
# 设置容器时间，有的容器时区与我们的时区不同，可能会带来麻烦
ENV LANG C.UTF-8　　　　
# 设置语言为utf-8
WORKDIR /code/bdtools　　　　　　
# 设置工作目录，也就是下面执行 ENTRYPOINT 后面命令的路径
RUN /usr/local/bin/pip3 install -r requirements.txt　　
# 根据requirement.txt下载好依赖包
EXPOSE 7777　　　　　　
# EXPOSE 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务。
# 此处填写7777，是因为我们上面的app.py提供的web服务就需要使用7777端口
ENTRYPOINT ["python3","app.py"]　　
```
Dockerfile编写完成后，我们就可以构建镜像了。
使用命令

```
docker build -t new Dockerfile
```
意思是，使用当前路径下的DockerFile进行构建，镜像名称为new
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649902814.png)
如上图，看到最后一行Successfully就表示构建成功了。图中红色部分报错是pip包版本不是最新的告警，不影响构建过程，可以忽略。
现在，我们可以查看一下镜像情况
使用命令

```
docker images
```
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649902877.png)
如上图，镜像已经构建出来了。
那么，开始启动容器。
执行命令

```
docker run -p 3333:7777 -dit d7d7df1b3dd5
```

这里需要说明一下：
1，-p参数，注意是小写，3333，表示宿主机的端口，7777表示容器内部的端口。
整条参数的意思是：将容器内的端口7777，映射到宿主机的3333端口。
如果我们需要从宿主机外部访问这个容器服务，只需要访问3333端口即可。

2，-dit ，这个参数我们只说-d，就是后台运行的意思。整行命令最后的那一串字符串，其实是上面构建出的镜像ID.
执行效果如下：
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649902938.png)
此刻，容器即已启动了。
我们可以通过命令查看容器的运行情况

```
docker ps -a
```
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649902986.png)
此处，我们一般关注的是容器ID、STATUS和PORTS，可以看到，容器的端口7777已经映射到宿主机的3333端口。
那么，我们如何进入到容器内部呢？
可以使用命令

```
docker exec -it 容器ID  /bin/bash
```

**需要注意，是容器ID，不是镜像ID**
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649903069.png)
执行效果如上图，可以看到，命令行提示符已经到了容器内部。

 

那么，我们还需要确认一下，这个python服务到底启动了没有。
我们首先可以在宿主机查看端口占用情况
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649903122.png)

此外，也可以在本地PC浏览器，去访问宿主机的3333端口即可。如下图
![Image text](https://github.com/pzspsh/intelligentlibrary/blob/main/images/1649903178.png)

综上，我们已经完成了从镜像制作到服务部署的全部流程。
此外，还有个别常用的docker相关的管理命令也一并贴上来

```
# 删除镜像
docker rmi  镜像ID

# 删除容器
docker rm 容器ID

# 杀容器，会将运行中的容器停下来
docker kill 容器ID

# 启动容器，并将进入容器中的bash命令行
docker run -it 镜像ID  /bin/bash  
```

当我们开发把项目打包好镜像后，可能需要给运维或后端人员进行部署测试，可以使用save -o 命令把镜像导出：
```
docker save -o .\保存的路径\镜像文件名 镜像
比如:
dicker save -o .\Desktop\hello.tar hello-world
```

vim /etc/docker/daemon.json 添加国内镜像
```
{
    "registry-mirrors": [
        "https://registry.hub.docker.com",
        "http://hub-mirror.c.163.com",
        "https://docker.mirrors.ustc.edu.cn",
        "https://registry.docker-cn.com"
    ]
}
```
docker 进入容器root权限
```
docker exec -it --user=root ID号/容器名称 bash
```