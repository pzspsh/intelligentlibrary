# 5 种文件变更时自动重载 Go 程序的方法

本文假设已安装 Go 编译器，并且已将 GOPATH/bin 路径添加到 PATH 环境变量中。

在开始之前，我们先创建一个简单的 web 服务器，可以返回响应内容”Hello，World”。

```go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World"))
    })

    http.ListenAndServe(":5000", nil)
}
```

#### Method 1: 使用 Air

[Air](https://github.com/cosmtrek/air) 是一个命令行程序，可以为 Go 程序提供实时代码加载。

通过运行以下命令来安装 Air。

```bash
go get -u github.com/cosmtrek/air
```

下一步，在使用项目的根目录中创建 Air 配置文件.air.conf。

```bash
# .air.conf
# toml配置文件来源于 [Air](https://github.com/cosmtrek/air)

# 工作区间
# .(当前目录)或绝对路径, 注意这些目录都在根目录下面.
root = "." 
tmp_dir = "tmp"

[build]
# 只是普通的shell命令。 可以使用`make`。
cmd = "go build -o ./tmp/main ."
# `cmd`配置命令输出的二进制文件的位置。
bin = "tmp/main"
# 自定义二进制输出。
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# 监听的文件扩展后缀列表。
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录。
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 如果指定，则监听包含这些文件。
include_dir = []
# 忽略文件列表.
exclude_file = []
# 排除监听的文件的正则模式
exclude_regex = []
# 是否排除未变化的文件
exclude_unchanged = true
# 是否跟随文件链接
follow_symlink = true
# 如果文件修改太频繁，则不必在每次修改时都立刻触发构建，设置触发等待时间。
delay = 1000 # ms
# 发生编译错误时，是否停止旧的二进制程序。
stop_on_error = true
# 该日志文件放置在tmp_dir中。
log = "air_errors.log"
# 是否在项进程发送 kill 之前发送 interrupt 信号
send_interrupt = false
# 发送 interrupt 后到发送 kill 之间的的延迟时间
kill_delay = 500 # ms

[log]
# 日志是否显示时间
time = false

[color]
# 自定义每类输出的颜色。 如果找不到这个颜色，使用原本的日志输出演示。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时是否删除临时目录
clean_on_exit = true
```

配置是简单明了的，请根据你的项目情况去调整。

最后，不要使用常用的 go run 命令来运行 Go 程序，而应使用 air 命令来启动程序。

#### Method 2: docker 运行 Ari

这种方法需要使用 docker，如果你没有安装，可以按照

我们仍将使用 Air 库，因此仍然需要 Air 配置文件。 如果你还没有，请创建一个配置文件。

Docker 镜像 cosmtrek/air 附带安装了 Air 命令，并且 GOPATH 环境变量设置为 /go。

我们只需要将我们的项目目录挂载到 Docker 容器的 GOPATH 中，并使用 - p 暴露需要使用的端口即可。 我们可以通过运行 docker run 命令来实现这一点：

```bash
docker run -it --rm -w <WORKING_DIR> -v <PROJECT_FOLDER>:<MOUNT_POINT> -p <HOST_PORT>:<CONTAINER_PORT> <IMAGE_NAME>
```

就我而言，我需要运行以下命令：

```bash
docker run -it --rm -w /go/src/github.com/praveen001/live-reloading -v /go/src/github.com/praveen001/live-reloading:/go/src/github.com/praveen001/live-reloading -p 5000:5000 cosmtrek/air
```

解释：

使用 - v 参数将项目目录 /home/praveen/go/src/github.com/praveen001/live-reloading 挂载到容器里面的 GOPATH 中的目录 /go/src/github.com/praveen001/live-reloading。

```bash
-v /home/praveen/go/src/github.com/praveen001/live-reloading:/go/src/github.com/praveen001/live-reloading 
```

使用 -w 参数指定挂载目录成工作目录。

```
-w /go/src/github.com/praveen001/live-reloading
```

Web 服务器正在监听端口 5000，因此需要使用 -p 标志将容器端口 5000 暴露到主机端口 5000。

```shell
-p 5000:5000
```

最后，指定 docker 镜像名称 cosmtrek / air。

#### Method 3: 使用 Gin

[Gin](https://github.com/codegangsta/gin) 是另一个用于实时重新加载 Go 应用程序的命令行程序。

通过运行以下命令来安装 Gin。

```bash
go get github.com/codegangsta/gin
```

而不是使用通常的 go run main.go 命令运行应用程序，而是使用 gin 命令。

就我而言，--appPort 参数告诉 Gin 监听端口 5000，--port 参数告诉 Gin 代理监听端口 3000 端口

```bash
gin --appPort 5000 --port 3000
```

现在使用地址 http://localhost:3000 访问 Gin 程序.

如果要排除监听那个目录可以使用 --excludeDir 参数，例如：

```bash
gin --appPort 5000 --port 3000 --excludeDir ./frontend
```

如果你项使用 Gin 实现加载没有启动端口监听的程序，你们必须使用 --immediate 参数。但是 Gin 仍然会去 5000 端口。

你可以在这找到所有受支持的参数 [Gin](https://github.com/codegangsta/gin) 的 Github.

#### Method 5: 使用 Fresh

[Fresh](https://github.com/gravityblast/fresh) 是另一个 GO 实现的用于实时重新加载 Go 的程序

安装 Fresh

```bash
go get github.com/pilu/fresh
```

而不是使用常用的 `go run main.go` 命令来运行应用程序，而是使用 `fresh` 命令。

```
fresh
```

要配置 Fresh，需要在项目的根目录中创建一个配置文件 `runner.conf`。

这是一个示例配置文件。

```bash
root:              .
tmp_path:          ./tmp
build_name:        runner-build
build_log:         runner-build-errors.log
valid_ext:         .go, .tpl, .tmpl, .html
no_rebuild_ext:    .tpl, .tmpl, .html
ignored:           assets, tmp
build_delay:       600
colors:            1
log_color_main:    cyan
log_color_build:   yellow
log_color_runner:  green
log_color_watcher: magenta
log_color_app:
```

