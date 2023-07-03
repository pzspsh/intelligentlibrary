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
    r.Run(":5005")
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