/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 11:09:25
*/
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// func main() {
// 	server := &http.Server{
// 		Addr: ":8080",
// 	}
// 	go func() {
// 		if err := server.ListenAndServeTLS("conf/server.crt", "conf/server.key"); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("Listen: %s\n", err)
// 		}
// 	}()
// }

// 设置http重定向到https
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "dubinyang.xyz:8081",
		})

		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}

func main() {
	g := gin.Default()
	//加载中间件
	g.Use(TlsHandler())
	server := &http.Server{
		Addr: ":8080",
	}

	//监听http
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	//监听https
	go func() {
		//端口不可重复监听
		//此处更换8080端口为8081，直接用gin的RunTLS()函数进行监听
		//继续用server的ListenAndServeTLS()函数效果一样，建server2，Addr变为为":8081"即可
		if err := g.RunTLS(":8081", "conf/server.crt", "conf/server.key"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()
}
