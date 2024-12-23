# docker 部署项目

## 一、搭建好Go编译环境

省略。

## 二、docker 部署项目示例如下：

```shell
- GoProjects  # GoProjects是Go的GOPATH开发路径
	- bin  # bin文件夹
	- pkg  # pkg文件夹
	- src  # src项目文件夹
		- dockerProject  # 新建部署的docker项目dockerProject
			- main.go   # 建main.go测试docker部署项目 vim main.go
```

### 1 main.go代码如下：

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
var db = make(map[string]string)
func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
    return r
}
func main() {
    r := setupRouter()
    r.Run("0.0.0.0:8000")
}
```

### 2 切换到dockerProject文件夹下执行如下命令：

#### 1.执行如下两个命令使go代码main.go可编译：

```go
go mod init dockerproject
go mod tidy
```

#### 2.编译main.go代码如下：

```go
go build -m go_run main.go
```

##### 3.用.sh脚本文件运行，如：go_run.sh文件内容如下

```shell
#!/bin/sh
exec go_run路径/go_run
```



# [Dockerfile部署]

#### 1.创建Dockerfile文件

```shell
touch Dockerfile   # 执行vim Dockerfile进行编辑保存也行
```

#### 2.编译的Dockerfile文件内容编辑如下：

```dockerfile
# 表示拉取Go环境镜像
FROM golang
# 表示在容器执行创建dockerProject文件夹
RUN mkdir /dockerProject
# 表示设置工作目录路径
WORKDIR /dockerProject
# 表示复制之前编译的go_run到容器的文件夹dockerProject的go_app
COPY go_run /dockerProject/go_app
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
# ENV GO111MODULE=on
# 下载并安装第三方依赖到容器中
# RUN go get github.com/codegangsta/gin  
# 设置权限
# RUN chmod -R 777 /dockerProject/go_app  
# 设置编码
# ENV LANG C.UTF-8 
# 表示设置暴露的端口
EXPOSE 8000
# 表示执行go_app, 运行golang程序的命令
ENTRYPOINT ["/dockerProject/go_app"]
```

#### 3.不编译的Dockerfile文件内容编辑如下：

```dockerfile
# 基础镜像，基于golang最新镜像构建
FROM golang
# 作者
MAINTAINER pan
# 配置GOROOT环境变量（感觉安装好就可以用不用配）：
# ENV GOROOT /usr/local/go
# ENV PATH=$PATH:$GOROOT/bin
# 配置GOPATH环境变量：
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
# 创建dockerProject
RUN mkdir -p /home/GoProjects/src/dockerProject
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
# 需要把go.mod、go.sum也复制到目标目录，所以直接用点继续全部复制
# COPY main.go  $GOPATH/dockerProject/main.go
COPY . /home/GoProjects/src/dockerProject/
# 环境变量的MODULE
ENV GO111MODULE=on
# 用于代理下载go项目依赖的包
ENV GOPROXY=https://goproxy.cn,direct
# 需暴露的端口
# RUN go mod init
# RUN go mod tidy
# RUN go build -o main main.go
EXPOSE 8000
# 可外挂的目录
VOLUME ["home/GoProjects/src/dockerProject/config","home/GoProjects/src/dockerProject/log"]                                                                                     
# docker run命令触发的真实命令(相当于不编译打包，源代码直接运行)
ENTRYPOINT ["go","run","main.go"]
# 说明：如果遇到需要依赖的GCC，则需要执行如下命令操作： 
# 如果是CentOS镜像则用yum,是Ubuntu则用apt-get
# RUN yum install -y gcc
# 如：github.com/tomatome/grdp包就需要gcc才能编译，无gcc无法编译
```

#### 4.创建继续如下：

```shell
docker build -f Dockerfile  -t test-go-docker .
```

test-go-docker：表示镜像名

说明：如果执行Dockerfile有点问题，可进行调整

我遇到的问题有：空格不对、注释不对、配置docker容器的go编译环境不对等

##### 结果如下：

```shell
Sending build context to Docker daemon  11.29MB
Step 1/8 : FROM golang
 ---> 77246b1c2182
Step 2/8 : RUN mkdir /dockerProject
 ---> Running in 893702a9d2ac
Removing intermediate container 893702a9d2ac
 ---> 556af1dd2334
Step 3/8 : WORKDIR /dockerProject
 ---> Running in f098a9accc7c
Removing intermediate container f098a9accc7c
 ---> 372de302b2c8
Step 4/8 : COPY go_run /dockerProject/go_app
 ---> c903cef0197b
Step 5/8 : RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
 ---> Running in 3732dfe79e0e
Removing intermediate container 3732dfe79e0e
 ---> b44a0c51d65c
Step 6/8 : RUN echo 'Asia/Shanghai' >/etc/timezone
 ---> Running in d8d853870138
Removing intermediate container d8d853870138
 ---> 097bba84eba1
Step 7/8 : EXPOSE 8000
 ---> Running in b99f5dd2d928
Removing intermediate container b99f5dd2d928
 ---> 3bb949721054
Step 8/8 : ENTRYPOINT ["/dockerProject/go_app"]
 ---> Running in e63c09286063
Removing intermediate container e63c09286063
 ---> 954bd157c07b
Successfully built 954bd157c07b
Successfully tagged test-go-docker:latest
```

这个结果没显示安装Go环境包是我之前安装了好几次，Go环境已经安装好了，之后调通Dockerfile文件执行成功的结果

#### 4.启动镜像命令如下：

```shell
docker run -d -p 8000:8000 test-go-docker:latest
```

##### 5.用.sh文件脚本一键启动如：go_dockerfile.sh

```shell
#!/bin/bash
# 构建程序
config="export GO111MODULE=on && CGO_ENABLED=1 GOOS=linux GOARCH=amd64; go mod tidy; go build  -o 'go_run' main.go;"
bash -c "$config"
# 构建docker image
docker build -t go-test/develop .
# 构建docker container
# 删除旧容器1
docker stop go-test-docker;docker rm go-test-docker
# 新建容器1 外部8080端口映射到容器内8000端口
docker run -itd -v :/www/wwwroot/go-run/:/data/log/ -p 8080:8000 --name go-test-docker go-test/develop:latest
```

**说明**：**go_dockerfile.sh**未完成，大概就是这个格式。俺不需要所以就没进行搞



## 优化的Dockerfile--镜像占磁盘小

#### 优化1：使用alpine版本的基础镜像,Dockerfile的内容如下：

```shell
FROM golang:alpine
# 作者
MAINTAINER pan
# 配置docker容器的go环境变量
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
# 全局工作目录
RUN mkdir -p /home/GoProjects/src/dockerProject
WORKDIR /home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/GoProjects/src/dockerProject/
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译
RUN GOOS=linux GOARCH=amd64 go build main.go
# 需暴露的端口
EXPOSE 8000
# 可外挂的目录
VOLUME ["home/GoProjects/src/dockerProject/config","home/GoProjects/src/dockerProject/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]
```

#### 优化2：使用多级构建的镜像 ,Dockerfile的内容如下：

```dockerfile
# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER pan
# 配置docker容器的go环境变量
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p /home/GoProjects/src/dockerProject
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/GoProjects/src/dockerProject/
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
# 全局工作目录
WORKDIR /go/kingProject
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/main .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/config ./config
# 需暴露的端口
EXPOSE 8000
# 可外挂的目录
VOLUME ["/home/GoProjects/src/dockerProject/config","/home/GoProjects/src/dockerProject/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]
```

#### 优化3：使用多级构建+scratch基础镜像,Dockerfile的内容如下：

```dockerfile
# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER pan
# 配置docker容器的go环境变量
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p /home/GoProjects/src/dockerProject
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/GoProjects/src/dockerProject/
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go


# 使用scratch这个空镜像为基础镜像--运行阶段
FROM scratch AS runner
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/main .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/config ./config
# 需暴露的端口
EXPOSE 8000
# 可外挂的目录
VOLUME ["/home/GoProjects/src/dockerProject/config","/home/GoProjects/src/dockerProject/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]
```

#### 优化4：go编译命令去掉冗余输出,Dockerfile的内容如下：

```dockerfile
# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER pan
# 配置docker容器的go环境变量
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p /home/GoProjects/src/dockerProject
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/GoProjects/src/dockerProject/
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go


# 使用scratch这个空镜像为基础镜像--运行阶段
FROM scratch AS runner
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/main .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/config ./config
# 需暴露的端口
EXPOSE 8000
# 可外挂的目录
VOLUME ["/home/GoProjects/src/dockerProject/config","/home/GoProjects/src/dockerProject/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]
```

#### 优化5：go编译命令去掉冗余输出，解决时区问题 ,Dockerfile的内容如下：

```dockerfile
# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER pan
# 配置docker容器的go环境变量
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p /home/GoProjects/src/dockerProject
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/GoProjects/src/dockerProject/
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata

# 使用scratch这个空镜像为基础镜像--运行阶段
FROM scratch AS runner
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/main .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/config ./config
# 复制编译阶段里的时区文件到目标目录
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
# 需暴露的端口
EXPOSE 8888
# 可外挂的目录
VOLUME ["/home/GoProjects/src/dockerProject/config","/home/GoProjects/src/dockerProject/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]
```

#### 优化6：

```dockerfile
# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
# 作者
MAINTAINER pan
# 配置docker容器的go环境变量
RUN mkdir -p /home/GoProjects
RUN mkdir -p /home/GoProjects/src
RUN mkdir -p /home/GoProjects/bin
RUN mkdir -p /home/GoProjects/pkg
ENV GOPATH /home/GoProjects
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p /home/GoProjects/src/dockerProject
# 全局工作目录
WORKDIR //home/GoProjects/src/dockerProject
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /home/GoProjects/src/dockerProject/
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go


# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
# 全局工作目录
WORKDIR /home/GoProjects/src/dockerProject
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/main .
# 复制编译阶段里的config文件夹到目标目录
COPY --from=builder /home/GoProjects/src/dockerProject/config ./config
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata
# 需暴露的端口
EXPOSE 8888
# 可外挂的目录
VOLUME ["/home/GoProjects/src/dockerProject/config","/home/GoProjects/src/dockerProject/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]
```

**说明**：scratch镜像有它的缺陷，是一个真的空镜像，不支持很多命令，比如cp，sh等，如果要进入容器内部查东西，都进不去，不适合真实情况。



# [docker compose部署]

上面的步骤比较繁琐，除了构建，还需要运行各种参数，可以使用 Docker 容器管理工具 docker compose 解决此问题。

#### 目录结构如下

```go
-- Dockerfile  // 构建脚本
-- docker-compose.yml // 配置运行容器需要的命令和参数
-- main.go  // go 入口文件
```

#### 1.创建docker-compose.yml文件，命令如下

```shell
touch docker-compose.yml  # 执行命令：vim docker-compose.yml 编辑即可
```

#### 2.编辑docker-compose.yml文件的内容如下：

```shell
version: '3.8'
services:
  go-test-docker: # 容器
    restart: always # Docker 重启时，容器也重启
    build: # 构建 Docker 镜像
      context: ./ # Dockerfile 文件的目录
      dockerfile: Dockerfile # Dockerfile 文件的名称
    image: go-test/develop:latest # 镜像名称和版本号
    container_name: go-test-docker # 容器名称
    ports: # 宿主机:容器之间映射端口
      - "8080:8000"
```

#### Dockerfile 文件

```shell
FROM golang:alpine as builder
# 需要go环境
MAINTAINER pan
WORKDIR /workdir
# 源
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go
FROM alpine:latest
# 设置时区
RUN apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >  /etc/timezone
WORKDIR /server
# 复制到工作区
COPY --from=builder /work/ ./
# COPY --from=builder /work/config ./config
# 对外端口
EXPOSE 8080
# 执行
CMD ["./main"]
```

#### 执行命令

```shell
docker-compose up
```

#### docker-compose 命令

```shell
docker-compose build  # 构建（重新构建）
docker-compose up  # 启动容器
docker-compose up -d  # 后台启动并运行
docker-compose ps  # 查看容器列表
docker-compose logs  # 查看日志（-f 一直监听）
docker-compose stop  # 关闭容器
docker-compose start  # 启动容器
docker-compose restart  # 重启容器
docker-compose rm  # 删除容器
docker-compose exec xxx sh  # 进入容器
docker-compose 命令说明  # xxx是容器名称
```

### 访问

ip+端口 进行访问

### nginx 反向代理

```shell
# 反向代理配置
server
{
    listen 80;
    server_name xxx.choudalao.com;
    # 其他配置 ....
    location / {
         proxy_pass http://127.0.0.1:20239;# http://xxx.com;# 也可以是域名
    }
}
```

### 反射后访问

    反射后,可以使用域名访问

## 问题

1、Docker 启动alpine镜像中可执行程序文件遇到 not found

```
原因
由于alpine镜像使用的是musl libc而不是gnu libc，/lib64/ 是不存在的。但他们是兼容的，可以创建个软连接过去试试!
这种情况是因为动态链接库位置错误导致的，alpine镜像使用的是musl libc而不是gun libc。因而动态链接库的位置不一致 。

而一般二进制文件在linux系统下编译，动态链接库是到/lib64目录下的，在alpine镜像内无/lib64目录 。

解决方法:
Dockerfile 文件

FROM alpine:latest
# 这个是重点
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# 其他代码

```

