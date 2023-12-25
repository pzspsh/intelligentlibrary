# echo使用教程
```go
https://github.com/labstack/echo 
https://echo.labstack.com/cookbook/hello-world
```
## 安装
```go
go get github.com/labstack/echo
```
## 使用
```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
```