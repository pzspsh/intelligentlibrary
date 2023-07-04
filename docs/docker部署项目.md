# Docker部署项目
### Docker 部署
代码主要分为两块，一块是Go的代码块，一块是构造的脚本
```go
-- dockerfile  // docker脚本
-- go_run.sh   // 启动go脚本
-- go-build.sh // 构造脚本
-- main.go     // go 入口文件
```

##### 文件说明：
```go
main.go       //程序启动的一个Go服务，并且监听6000端口
Dockerfile    // 构造镜像
go-build.sh   // 执行构造脚本，所有的步骤都加载在里面
go_run.sh     // 执行打包好的Go文件
```

#### 执行流程
运行./go-build.sh文件,执行里面的操作,先是打包go服务,然后构建docker镜像,最后启动容器.

### 各文件内容
#### Dockerfile
```shell
FROM alpine:latest
WORKDIR /server
EXPOSE 5005
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./go_run /server/go_run
RUN  chmod u+x /server/go_run
COPY ./go_run.sh /server/go_run.sh
RUN chmod u+x /server/go_run.sh
CMD ["/server/go_run.sh"]
```

#### go_run.sh
```shell
#!/bin/sh
exec /server/go_run
```

#### go-build.sh
```shell
#!/bin/bash
# 构建程序
test="export GO111MODULE=on && CGO_ENABLED=1 GOOS=linux GOARCH=amd64; go mod tidy; go build  -o 'go_run' main.go;"
bash -c "$test"
# 构建docker image
docker build -t go-test/develop .
# 构建docker container
# 删除旧容器1
docker stop go-test-20239;docker rm go-test-20239
# 新建容器1 外部20009端口映射到容器内5005端口
docker run -itd -v :/www/wwwroot/go-run/:/data/log/ -p 20239:6000 --name go-test-20239 go-test/develop:latest
```

#### main.go
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
    r.Run(":6000")
}
```

### 构建
./go-build.sh




## docker compose 部署
上面的步骤比较繁琐，除了构建，还需要运行各种参数，可以使用 Docker 容器管理工具 docker compose 解决此问题。

#### 目录结构如下
```go
-- Dockerfile  // 构建脚本
-- docker-compose.yml // 配置运行容器需要的命令和参数
-- main.go  // go 入口文件
```

#### docker-compose.yml
```shell
version: '3.8'
services:
  go-test-20239: # 容器
    restart: always # Docker 重启时，容器也重启
    build: # 构建 Docker 镜像
      context: ./ # Dockerfile 文件的目录
      dockerfile: Dockerfile # Dockerfile 文件的名称
    image: go-test/develop:latest # 镜像名称和版本号
    container_name: go-test-20239 # 容器名称
    ports: # 宿主机:容器之间映射端口
      - "20239:6000"
```

#### Dockerfile 文件
```shell
FROM golang:alpine as builder
# 需要go环境
MAINTAINER vijay
WORKDIR /work
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
EXPOSE 6000
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
         proxy_pass http://127.0.0.1:20009;# http://xxx.com;# 也可以是域名
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