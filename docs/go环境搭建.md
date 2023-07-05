#### go 官网下载
```shell
https://golang.google.cn/dl/
```

## windows环境
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
set GOPATH=D:\go
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

```

#### linux文件下载
```shell
以go1.20.3版本为例：
wget https://golang.google.cn/dl/go1.20.3.linux-amd64.tar.gz # 先下载go文件包

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz // 如果之前安装库go环境，则直接执行该命令
# 之前没安装过，直接解压go文件包，把go文件包迁移到/usr/local目录下


环境变量配置vim /etc/profile
export GOROOT=/usr/local/go  # 安装默认的路径
export GOPATH=/home/GoProjects #  GOPATH是指你开发的路径
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

$ go version
```