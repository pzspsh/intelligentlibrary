# 实现文档
```go
/*######################################################*/
/*你想优雅地重启或停止 web 服务器吗？有一些方法可以做到这一点。

我们可以使用 ​fvbock/endless​ 来替换默认的 ​ListenAndServe​

router := gin.Default()
router.GET("/", handler)
// [...]
endless.ListenAndServe(":4242", router)
替代方案:

​manners​：可以优雅关机的 Go Http 服务器。
​graceful​：​Graceful是一个 Go 扩展包，可以优雅地关闭 http.Handler 服务器。
​grace​：Go 服务器平滑重启和零停机时间部署。
如果你使用的是 Go 1.8，可以不需要这些库！考虑使用 ​http.Server​ 内置的 ​Shutdown()​ 方法优雅地关机. 请参阅 gin 完整的 graceful-shutdown 示例。
*/
import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// quit := make(chan os.Signal, 1) // 无缓冲
	quit := make(chan os.Signal, 1) // 必须要用带缓冲的chan
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

/*######################################################*/
```



```go
/*######################################################*/
// Gin XML/JSON/YAML/ProtoBuf 渲染
import (
	"net/http"

	"github.com/gin-gonic/gin"
    // "google.golang.org/protobuf"
)
func main() {
	r := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
```

```go
/*######################################################*/
// 使用 ​SecureJSON防止 ​json劫持。如果给定的结构是数组值，则默认预置 ​"while(1),"​ 到响应体。
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
```



```go

/*######################################################*/
// 通常，​JSON使用 ​unicode替换特殊 ​HTML字符，例如 ​<​ 变为 ​\ u003c​。如果要按字面对这些字符进行编码，则可以使用 ​PureJSON​。​Go 1.6​ 及更低版本无法使用此功能。
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
```



```go

/*######################################################*/
// http.Pusher
import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并在 https://127.0.0.1:8080 上启动服务
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
```



```go

/*######################################################*/
// 使用 ​JSONP向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
```



```go

/*######################################################*/

// 使用 ​AsciiJSON生成具有转义的非 ​ASCII字符的 ​ASCII-only JSON​
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

```



```go

/*######################################################*/
// 绑定html复选框
import (
    "github.com/gin-gonic/gin"
)

type myForm struct {
    Colors []string `form:"colors[]"`
}

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("views/*")
    r.GET("/", indexHandler)
    r.POST("/", formHandler)

    r.Run(":8080")
}

func indexHandler(c *gin.Context) {
    c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.Bind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}

```



```go

/*######################################################*/

// 中间件中使用gorotine
import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test1", func(c *gin.Context) {
		// 拷贝一份副本在Goroutine中使用
		tmp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 这里使用的是值拷贝的tmp
			log.Println("test1已完成,路径是:" + tmp.Request.URL.Path)
		}()
	})

	r.GET("/test2", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("test2已完成,路径是:" + c.Request.URL.Path)
	})
	r.Run()
}

```



```go

/*######################################################*/
// 在Gin框架中设置和获取Cookie的方法如下
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	r.Run()
}

```



```go

/*######################################################*/
// 重定向
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.w3cschool.cn/")
	})
	r.Run()
}

func main() {
	r := gin.Default()
	r.GET("/test1", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	r.Run()
}
```



```go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	r.Run()
}


```



```go

/*######################################################*/
// 文件上传
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	r.Run()
}

```



```go
/*######################################################*/

// 参数绑定
type Userinfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

/*为了能够更方便的获取请求相关参数，提高开发效率，我们可以使用ShouldBind，它能够基于请求自动提取JSON，Form表单，Query等类型的值，并把值绑定到指定的结构体对象，具体使用方法如下*/
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var u Userinfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u)
	})
	r.Run()
}
```



```go
func main() {
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		tmp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)
			// 请注意您使用的是复制的上下文 "tmp"，这一点很重要
			log.Println("Done! in path " + tmp.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)
		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	r.Run()
}

```



```go
/*######################################################*/
// 中间件
// // 定义一个中间键m1统计请求处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	c.Next() //调用后续的处理函数
	// c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", m1, index)
	r.Run()
}

```



```go
func main() {
	r := gin.Default()
	user := r.Group("/user")
	user.GET("/index", func(c *gin.Context) {})
	user.POST("/login", func(c *gin.Context) {})
	pwd:=user.Group("/pwd")
	pwd.GET("/pwd",func(c *gin.Context) {})
	r.Run()
}
```



```go
/*######################################################*/
// 路由
func main() {
	r := gin.Default()
	user := r.Group("/user")
	user.GET("/index", func(c *gin.Context) {})
	user.POST("/login", func(c *gin.Context) {})
	r.Run()
}
```



```go
func main() {
	r := gin.Default()
	r.GET("/user/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	})
	r.Run()
}
```



```go
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html") //加载页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)

	})
	r.POST("/", func(c *gin.Context) {
		username := c.PostForm("username") //对应h5表单中的name字段
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})
	r.Run()
}

```



```go
/*######################################################*/
// 获取参数
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		pwd := c.Query("pwd")
		// fmt.Printf("name:%s ; pwd:%s",name,pwd)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})
	r.Run()
}
```



```go
/*######################################################*/
// HTML渲染
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
			"pwd":  "123456",
		})
	})
	r.Run()
}
```



```go
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

```go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// 中间件（拦截器），功能：预处理，登录授权、验证、分页、耗时统计...
// func myHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// 通过自定义中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
// 		ctx.Set("usersesion", "userid-1")
// 		ctx.Next()  // 放行
// 		ctx.Abort() // 阻止
// 	}
// }

func main() {
	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./Arctime.ico")) // 这里如果添加了东西然后再运行没有变化，请重启浏览器，浏览器有缓存

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*") // 一种是全局加载，一种是加载指定的文件

	// 加载资源文件
	ginServer.Static("/static", "./static")

	// 相应一个页面给前端

	ginServer.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This data is come from Go background.",
		})
	})

	// 能加载静态页面也可以加载测试文件

	// 获取请求中的参数

	// 传统方式：usl?userid=xxx&username=conqueror712
	// Rustful方式：/user/info/1/conqueror712

	// 下面是传统方式的例子
	ginServer.GET("/user/info", func(context *gin.Context) { // 这个格式是固定的
		userid := context.Query("userid")
		username := context.Query("username")
		// 拿到之后返回给前端
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// 此时执行代码之后，在浏览器中可以输入http://localhost:8081/user/info?userid=111&username=666
	// 就可以看到返回了JSON格式的数据

	// 下面是Rustful方式的例子
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		// 还是一样，返回给前端
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// 指定代码后，只需要在浏览器中http://localhost:8081/user/info/111/555
	// 就可以看到返回了JSON数据了，非常方便简洁

	// 序列化
	// 前端给后端传递JSON
	ginServer.POST("/json", func(ctx *gin.Context) {
		// request.body
		data, _ := ctx.GetRawData()
		var m map[string]interface{} // Go语言中object一般用空接口来表示，可以接收anything
		// 顺带一提，1.18以上，interface可以直接改成any
		_ = json.Unmarshal(data, &m)
		ctx.JSON(http.StatusOK, m)
	})
	// 用apipost或者postman写一段json传到localhost:8081/json里就可以了
	/*
		json示例：
		{
			"name": "Conqueror712",
			"age": 666,
			"address": "Mars"
		}
	*/
	// 看到后端的实时响应里面接收到数据就可以了

	// 处理表单请求 这些都是支持函数式编程，Go语言特性，可以把函数作为参数传进来
	ginServer.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由
	ginServer.GET("/test", func(ctx *gin.Context) {
		// 重定向 -> 301
		ctx.Redirect(301, "https://conqueror712.gitee.io/conqueror712.gitee.io/")
	})
	// http://localhost:8081/test

	// 404
	ginServer.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "404.html", nil)
	})

	// 路由组暂略

	// 服务器端口，用服务器端口来访问地址
	ginServer.Run(":8081") // 不写的话默认是8080，也可以更改
}
```
```go
package main

import (
	"fmt"
	"strconv"
	"time"

	// "gorm.io/driver/sqlite"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// 如何连接数据库 ? MySQL + Navicat
	// 需要更改的内容：用户名，密码，数据库名称
	dsn := "root:BqV?eGcc_1o+@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 解决复数问题
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	fmt.Println("db = ", db)
	fmt.Println("err = ", err)

	// 连接池
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟

	// 结构体
	type List struct {
		gorm.Model        // 主键
		Name       string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
		State      string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
		Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
		Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
		Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
	}

	// 迁移
	db.AutoMigrate(&List{})

	// 接口
	r := gin.Default()

	// 测试
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "请求成功",
	// 	})
	// })

	// 业务码约定：正确200，错误400

	// 增
	r.POST("/user/add", func(ctx *gin.Context) {
		// 定义一个变量指向结构体
		var data List
		// 绑定方法
		err := ctx.ShouldBindJSON(&data)
		// 判断绑定是否有错误
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg":  "添加失败",
				"data": gin.H{},
				"code": "400",
			})
		} else {
			// 数据库的操作
			db.Create(&data) // 创建一条数据
			ctx.JSON(200, gin.H{
				"msg":  "添加成功",
				"data": data,
				"code": "200",
			})
		}
	})

	// 删
	// 1. 找到对应的id对应的条目
	// 2. 判断id是否存在
	// 3. 从数据库中删除 or 返回id没有找到

	// Restful编码规范
	r.DELETE("/user/delete/:id", func(ctx *gin.Context) {
		var data []List
		// 接收id
		id := ctx.Param("id") // 如果有键值对形式的话用Query()
		// 判断id是否存在
		db.Where("id = ? ", id).Find(&data)
		if len(data) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": 400,
			})
		} else {
			// 操作数据库删除（删除id所对应的那一条）
			// db.Where("id = ? ", id).Delete(&data) <- 其实不需要这样写，因为查到的data里面就是要删除的数据
			db.Delete(&data)

			ctx.JSON(200, gin.H{
				"msg":  "删除成功",
				"code": 200,
			})
		}

	})

	// 改
	r.PUT("/user/update/:id", func(ctx *gin.Context) {
		// 1. 找到对应的id所对应的条目
		// 2. 判断id是否存在
		// 3. 修改对应条目 or 返回id没有找到
		var data List
		id := ctx.Param("id")
		// db.Where("id = ?", id).Find(&data) 可以这样写，也可以写成下面那样
		// 还可以再Where后面加上Count函数，可以查出来这个条件对应的条数
		db.Select("id").Where("id = ? ", id).Find(&data)
		if data.ID == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "用户id没有找到",
				"code": 400,
			})
		} else {
			// 绑定一下
			err := ctx.ShouldBindJSON(&data)
			if err != nil {
				ctx.JSON(200, gin.H{
					"msg":  "修改失败到",
					"code": 400,
				})
			} else {
				// db修改数据库内容
				db.Where("id = ?", id).Updates(&data)
				ctx.JSON(200, gin.H{
					"msg":  "修改成功",
					"code": 200,
				})
			}
		}
	})

	// 查
	// 第一种：条件查询，
	r.GET("/user/list/:name", func(ctx *gin.Context) {
		// 获取路径参数
		name := ctx.Param("name")
		var dataList []List
		// 查询数据库
		db.Where("name = ? ", name).Find(&dataList)
		// 判断是否查询到数据
		if len(dataList) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "没有查询到数据",
				"code": "400",
				"data": gin.H{},
			})
		} else {
			ctx.JSON(200, gin.H{
				"msg":  "查询成功",
				"code": "200",
				"data": dataList,
			})
		}
	})

	// 第二种：全部查询 / 分页查询
	r.GET("/user/list", func(ctx *gin.Context) {
		var dataList []List
		// 查询全部数据 or 查询分页数据
		pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
		pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

		// 判断是否需要分页
		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = -1
		}

		offsetVal := (pageNum - 1) * pageSize // 固定写法 记住就行
		if pageNum == -1 && pageSize == -1 {
			offsetVal = -1
		}

		// 返回一个总数
		var total int64

		// 查询数据库
		db.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)

		if len(dataList) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			ctx.JSON(200, gin.H{
				"msg":  "查询成功",
				"code": 200,
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
			})
		}

	})

	// 端口号
	PORT := "3001"
	r.Run(":" + PORT)
}
```