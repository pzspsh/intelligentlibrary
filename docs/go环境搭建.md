#### go 官网下载

```shell
https://golang.google.cn/dl/
```

## windows 环境

```shell
1、先下载go文件包 https://golang.google.cn/dl/go1.20.3.windows-amd64.zip
2、解压go文件包
3、把解压的go文件包路径：go文件包路径/bin设置到window系统环境变量中
4、开发环境例如新建文件夹go:该目录下有文件夹bin文件、pkg文件、src文件。
5、把新建文件夹go文件的bin目录设置到windows的系统环境变量中
终端执行：go version

go环境设置：
set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOEXE=.exe
set GOEXPERIMENT=
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=D:\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=D:\GoProjects
set GOPRIVATE=
set GOPROXY=https://goproxy.cn,direct
set GOROOT=C:\Program Files\go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Program Files\go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.19.1
set GCCGO=gccgo
set GOAMD64=v1
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set GOWORK=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=C:\Users\admin\AppData\Local\Temp\go-build3924361524=/tmp/go-build -gno-record-gcc-switches

修改：
go env -w GO111MODULE=on // mod文件需求
GOPROXY=https://goproxy.cn,direct // go get 下载包需求

GO环境下载包代理设置：
    go env -w GOPROXY=https://goproxy.cn
    go env -w GOPROXY=https://goproxy.cn,direct
    go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct
```

#### linux 文件下载

```shell
以go1.20.3版本为例：
wget https://golang.google.cn/dl/go1.20.3.linux-amd64.tar.gz # 先下载go文件包

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz # 如果之前安装库go环境，则直接执行该命令
# 之前没安装过，直接解压go文件包，把go文件包迁移到/usr/local目录下


环境变量配置vim /etc/profile
export GOROOT=/usr/local/go  # 安装默认的路径
export GOPATH=/home/GoProjects #  GOPATH是指你开发的路径
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


$ source /etc/profile # 使环境变量生效
$ go version # 查看go版本
```

#### 存在Ctrl+右键进函数之后再Ctrl+右键进函数进不去问题
##### 原因：VS Code或gopls加载不到GOPATH
##### 解放办法1：
```bash
先下载好1.20.3.linux-amd64.tar.gz文件包
1、安装Go运行环境的
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz

2、安装好go编译环境，VS Code连接进入服务器，工作区是/home目录执行go version命令可查看已经安装好Go编译环境，如果工作区是/root时执行go versin命令会显示Go编译环境不可用，则需要设置全局编译环境把GOROOT添加到shell的配置文件/etc/profile或~/.bashrc文件，将编辑的内容添加到文件末尾例如下：
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

编辑好保存退出，执行source /etc/profile或source ~/.bashrc使配置生效

3、执行go env命令会看到GOPATH有默认路径不需要修改，之前我修改了好久几次GOPATH是生效的，开发的项目也可以正常使用，ctrl+右键进入系统包函数是正常的，但是ctrl+右键进入下载的依赖包函数是正常的，再在依赖包文件ctrl+右键进入函数就进不去了，所有GOPATH的默认路径可以不用修改

4、开发项目随便在你喜欢的文件夹下新键个开发目录都可正常开发运行，开发项目目录下正常生成go.mod文件，在项目目录文件下ctrl+右键都可正常进入函数

5、这是在linux上出现的问题，在windows上可以正常随便配置GOPATH不出现下载的依赖包ctrl+右键进入函数问题
```
##### 命令
```bash
# 将 GOROOT 添加到你的 Shell 配置文件中
    # 如果你使用 bash，编辑 ~/.bashrc 或 ~/.bash_profile。
    # 如果你使用 zsh，编辑 ~/.zshrc。

# Go 环境变量
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Go 模块支持,永久设置在 Shell 配置文件中添加：
export GO111MODULE=on

# Go 代理, 永久设置在 Shell 配置文件中添加：
export GOPROXY=https://goproxy.cn,direct

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

# 在GOPATH 目录下创建标准的 Go 工作区结构：
mkdir -p $GOPATH/{src,bin,pkg}

# 运行以下命令检查 GOPATH 是否设置成功：
go env GOPATH

source ~/.bashrc  # 或者 source ~/.zshrc
```

```bash
GO111MODULE='on'
GOARCH='amd64'
GOBIN=''
GOCACHE='/root/.cache/go-build'
GOENV='/root/.config/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFLAGS=''
GOHOSTARCH='amd64'
GOHOSTOS='linux'
GOINSECURE=''
GOMODCACHE='/root/go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='linux'
GOPATH='/root/go'
GOPRIVATE=''
GOPROXY='https://goproxy.cn,direct'
GOROOT='/usr/local/go'
GOSUMDB='sum.golang.org'
GOTMPDIR=''
GOTOOLCHAIN='auto'
GOTOOLDIR='/usr/local/go/pkg/tool/linux_amd64'
GOVCS=''
GOVERSION='go1.23.4'
GODEBUG=''
GOTELEMETRY='local'
GOTELEMETRYDIR='/root/.config/go/telemetry'
GCCGO='gccgo'
GOAMD64='v1'
AR='ar'
CC='gcc'
CXX='g++'
CGO_ENABLED='1'
GOMOD='/dev/null'
GOWORK=''
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
PKG_CONFIG='pkg-config'
GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1976301070=/tmp/go-build -gno-record-gcc-switches'
```