# beego教程
```go
https://github.com/astaxie/beego 
https://beego.me/docs/install/beego.md
https://beego.me/docs 
https://github.com/beego/samples
```

# 安装bee工具
```go
go get github.com/beego/bee
```

# 使用bee创建项目
```go
bee new hello
```

# 运行项目
```go
cd hello
bee run
```

# 访问项目
```go
http://localhost:8080/
```

```go
├── main.go
└── controllers
    └── hello.go
```

main.go
```go
package main

import (
	"github.com/astaxie/beego"
	_ "hello/routers"
)

func main() {
	beego.Run()
}
```

hello.go
```go
package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Ctx.WriteString("hello world")
}
```

routers.go
```go
package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/hello",
			beego.NSInclude(
				&controllers.HelloController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
```