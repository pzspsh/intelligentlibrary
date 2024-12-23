/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 12:28:05
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()
	router.Use(TlsHandler())

	router.RunTLS(":8080", "ssl.pem", "ssl.key")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
