/*
@File   : main.go
@Author : pan
@Time   : 2023-06-07 14:27:58
*/
package main

// func main() {
// 	r := gin.Default()
// 	r.GET("/index", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

/*######################################################*/
// HTML渲染
// func main() {
// 	r := gin.Default()
// 	r.LoadHTMLGlob("../templates/*")
// 	r.GET("/index", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 			"name": "admin",
// 			"pwd":  "123456",
// 		})
// 	})
// 	r.Run()
// }

/*######################################################*/
// 获取参数
// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		name := c.Query("name")
// 		pwd := c.Query("pwd")
// 		// fmt.Printf("name:%s ; pwd:%s",name,pwd)
// 		c.JSON(http.StatusOK, gin.H{
// 			"name": name,
// 			"pwd":  pwd,
// 		})
// 	})
// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	r.LoadHTMLFiles("./login.html", "./index.html") //加载页面
// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "login.html", nil)

// 	})
// 	r.POST("/", func(c *gin.Context) {
// 		username := c.PostForm("username") //对应h5表单中的name字段
// 		password := c.PostForm("password")
// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 			"username": username,
// 			"password": password,
// 		})
// 	})
// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/user/:username", func(c *gin.Context) {
// 		username := c.Param("username")
// 		c.JSON(http.StatusOK, gin.H{
// 			"username": username,
// 		})
// 	})
// 	r.Run()
// }

/*######################################################*/
// 路由
// func main() {
// 	r := gin.Default()
// 	user := r.Group("/user")
// 	user.GET("/index", func(c *gin.Context) {})
// 	user.POST("/login", func(c *gin.Context) {})
// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	user := r.Group("/user")
// 	user.GET("/index", func(c *gin.Context) {})
// 	user.POST("/login", func(c *gin.Context) {})
// 	pwd:=user.Group("/pwd")
// 	pwd.GET("/pwd",func(c *gin.Context) {})
// 	r.Run()
// }

/*######################################################*/
// 中间件
// // 定义一个中间键m1统计请求处理函数耗时
// func m1(c *gin.Context) {
// 	fmt.Println("m1 in...")
// 	start := time.Now()
// 	c.Next() //调用后续的处理函数
// 	// c.Abort() //阻止调用后续的处理函数
// 	cost := time.Since(start)
// 	fmt.Printf("cost:%v\n", cost)
// }

// func index(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"msg": "ok",
// 	})
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/", m1, index)
// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/long_async", func(c *gin.Context) {
// 		// 创建在 goroutine 中使用的副本
// 		tmp := c.Copy()
// 		go func() {
// 			// 用 time.Sleep() 模拟一个长任务。
// 			time.Sleep(5 * time.Second)
// 			// 请注意您使用的是复制的上下文 "tmp"，这一点很重要
// 			log.Println("Done! in path " + tmp.Request.URL.Path)
// 		}()
// 	})
// 	r.GET("/long_sync", func(c *gin.Context) {
// 		// 用 time.Sleep() 模拟一个长任务。
// 		time.Sleep(5 * time.Second)
// 		// 因为没有使用 goroutine，不需要拷贝上下文
// 		log.Println("Done! in path " + c.Request.URL.Path)
// 	})
// 	r.Run()
// }

/*######################################################*/

// 参数绑定
// type Userinfo struct {
// 	Username string `form:"username"`
// 	Password string `form:"password"`
// }

// /*为了能够更方便的获取请求相关参数，提高开发效率，我们可以使用ShouldBind，它能够基于请求自动提取JSON，Form表单，Query等类型的值，并把值绑定到指定的结构体对象，具体使用方法如下*/
// func main() {
// 	r := gin.Default()
// 	r.GET("/user", func(c *gin.Context) {
// 		var u Userinfo
// 		err := c.ShouldBind(&u)
// 		if err != nil {
// 			c.JSON(http.StatusBadGateway, gin.H{
// 				"error": err.Error(),
// 			})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{
// 				"status": "ok",
// 			})
// 		}
// 		fmt.Printf("%#v\n", u)
// 	})
// 	r.Run()
// }

/*######################################################*/
// 文件上传
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	// 处理multipart forms提交文件时默认的内存限制是32 MiB
// 	// 可以通过下面的方式修改
// 	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
// 	r.POST("/upload", func(c *gin.Context) {
// 		// 单个文件
// 		file, err := c.FormFile("f1")
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		log.Println(file.Filename)
// 		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
// 		// 上传文件到指定的目录
// 		c.SaveUploadedFile(file, dst)
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
// 		})
// 	})
// 	r.Run()
// }

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	// 处理multipart forms提交文件时默认的内存限制是32 MiB
// 	// 可以通过下面的方式修改
// 	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
// 	r.POST("/upload", func(c *gin.Context) {
// 		// Multipart form
// 		form, _ := c.MultipartForm()
// 		files := form.File["file"]
// 		for index, file := range files {
// 			log.Println(file.Filename)
// 			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
// 			// 上传文件到指定的目录
// 			c.SaveUploadedFile(file, dst)
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": fmt.Sprintf("%d files uploaded!", len(files)),
// 		})
// 	})
// 	r.Run()
// }

/*######################################################*/
// 重定向
// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/test", func(c *gin.Context) {
// 		c.Redirect(http.StatusMovedPermanently, "https://www.w3cschool.cn/")
// 	})
// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/test1", func(c *gin.Context) {
// 		// 指定重定向的URL
// 		c.Request.URL.Path = "/test2"
// 		r.HandleContext(c)
// 	})
// 	r.GET("/test2", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"hello": "world"})
// 	})
// 	r.Run()
// }

/*######################################################*/
// 在Gin框架中设置和获取Cookie的方法如下
// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/cookie", func(c *gin.Context) {
// 		cookie, err := c.Cookie("gin_cookie")
// 		if err != nil {
// 			cookie = "NotSet"
// 			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
// 		}
// 		fmt.Printf("Cookie value: %s \n", cookie)
// 	})
// 	r.Run()
// }

/*######################################################*/

// 中间件中使用gorotine
// import (
// 	"log"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.GET("/test1", func(c *gin.Context) {
// 		// 拷贝一份副本在Goroutine中使用
// 		tmp := c.Copy()
// 		go func() {
// 			time.Sleep(5 * time.Second)
// 			// 这里使用的是值拷贝的tmp
// 			log.Println("test1已完成,路径是:" + tmp.Request.URL.Path)
// 		}()
// 	})

// 	r.GET("/test2", func(c *gin.Context) {
// 		time.Sleep(5 * time.Second)
// 		// 因为没有使用 goroutine，不需要拷贝上下文
// 		log.Println("test2已完成,路径是:" + c.Request.URL.Path)
// 	})
// 	r.Run()
// }

/*######################################################*/
// 绑定html复选框
// import (
// 	"github.com/gin-gonic/gin"
// )

// type myForm struct {
// 	Colors []string `form:"colors[]"`
// }

// func main() {
// 	r := gin.Default()

// 	r.LoadHTMLGlob("views/*")
// 	r.GET("/", indexHandler)
// 	r.POST("/", formHandler)

// 	r.Run(":8080")
// }

// func indexHandler(c *gin.Context) {
// 	c.HTML(200, "form.html", nil)
// }

// func formHandler(c *gin.Context) {
// 	var fakeForm myForm
// 	c.Bind(&fakeForm)
// 	c.JSON(200, gin.H{"color": fakeForm.Colors})
// }

/*######################################################*/

// 使用 ​AsciiJSON生成具有转义的非 ​ASCII字符的 ​ASCII-only JSON​
// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.GET("/someJSON", func(c *gin.Context) {
// 		data := map[string]interface{}{
// 			"lang": "GO语言",
// 			"tag":  "<br>",
// 		}

// 		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
// 		c.AsciiJSON(http.StatusOK, data)
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

/*######################################################*/
// 使用 ​JSONP向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/JSONP", func(c *gin.Context) {
// 		data := map[string]interface{}{
// 			"foo": "bar",
// 		}
// 		// /JSONP?callback=x
// 		// 将输出：x({\"foo\":\"bar\"})
// 		c.JSONP(http.StatusOK, data)
// 	})
// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

/*######################################################*/
// http.Pusher
// import (
// 	"html/template"
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// var html = template.Must(template.New("https").Parse(`
// <html>
// <head>
//   <title>Https Test</title>
//   <script src="/assets/app.js"></script>
// </head>
// <body>
//   <h1 style="color:red;">Welcome, Ginner!</h1>
// </body>
// </html>
// `))

// func main() {
// 	r := gin.Default()
// 	r.Static("/assets", "./assets")
// 	r.SetHTMLTemplate(html)

// 	r.GET("/", func(c *gin.Context) {
// 		if pusher := c.Writer.Pusher(); pusher != nil {
// 			// 使用 pusher.Push() 做服务器推送
// 			if err := pusher.Push("/assets/app.js", nil); err != nil {
// 				log.Printf("Failed to push: %v", err)
// 			}
// 		}
// 		c.HTML(200, "https", gin.H{
// 			"status": "success",
// 		})
// 	})

// 	// 监听并在 https://127.0.0.1:8080 上启动服务
// 	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
// }

/*######################################################*/
// 通常，​JSON使用 ​unicode替换特殊 ​HTML字符，例如 ​<​ 变为 ​\ u003c​。如果要按字面对这些字符进行编码，则可以使用 ​PureJSON​。​Go 1.6​ 及更低版本无法使用此功能。
// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	// 提供 unicode 实体
// 	r.GET("/json", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"html": "<b>Hello, world!</b>",
// 		})
// 	})

// 	// 提供字面字符
// 	r.GET("/purejson", func(c *gin.Context) {
// 		c.PureJSON(200, gin.H{
// 			"html": "<b>Hello, world!</b>",
// 		})
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

/*######################################################*/
// 使用 ​SecureJSON防止 ​json劫持。如果给定的结构是数组值，则默认预置 ​"while(1),"​ 到响应体。
// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	// 你也可以使用自己的 SecureJSON 前缀
// 	// r.SecureJsonPrefix(")]}',\n")

// 	r.GET("/someJSON", func(c *gin.Context) {
// 		names := []string{"lena", "austin", "foo"}

// 		// 将输出：while(1);["lena","austin","foo"]
// 		c.SecureJSON(http.StatusOK, names)
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

/*######################################################*/
// Gin XML/JSON/YAML/ProtoBuf 渲染
// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )
// func main() {
// 	r := gin.Default()

// 	// gin.H 是 map[string]interface{} 的一种快捷方式
// 	r.GET("/someJSON", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
// 	})

// 	r.GET("/moreJSON", func(c *gin.Context) {
// 		// 你也可以使用一个结构体
// 		var msg struct {
// 			Name    string `json:"user"`
// 			Message string
// 			Number  int
// 		}
// 		msg.Name = "Lena"
// 		msg.Message = "hey"
// 		msg.Number = 123
// 		// 注意 msg.Name 在 JSON 中变成了 "user"
// 		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
// 		c.JSON(http.StatusOK, msg)
// 	})

// 	r.GET("/someXML", func(c *gin.Context) {
// 		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
// 	})

// 	r.GET("/someYAML", func(c *gin.Context) {
// 		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
// 	})

// 	r.GET("/someProtoBuf", func(c *gin.Context) {
// 		reps := []int64{int64(1), int64(2)}
// 		label := "test"
// 		// protobuf 的具体定义写在 testdata/protoexample 文件中。
// 		data := &protoexample.Test{
// 			Label: &label,
// 			Reps:  reps,
// 		}
// 		// 请注意，数据在响应中变为二进制数据
// 		// 将输出被 protoexample.Test protobuf 序列化了的数据
// 		c.ProtoBuf(http.StatusOK, data)
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

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
如果你使用的是 Go 1.8，可以不需要这些库！考虑使用 ​http.Server​ 内置的 ​Shutdown()​ 方法优雅地关机. 请参阅 gin 完整的 graceful-shutdown 示例。*/
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
	quit := make(chan os.Signal)
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
