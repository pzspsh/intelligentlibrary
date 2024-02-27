/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

// Http的三步骤
// 1、处理器函数
// 2、绑定URL关系
// 3、启动Server服务
func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		ctx, err := os.ReadFile("HTTP/index.html")
		if err != nil {
			fmt.Fprintf(response, "%s", "网页读取出现错误")
			fmt.Println(err)
			fmt.Println(string(ctx))
		} else {
			fmt.Fprintf(response, "%s", string(ctx))
		}
	})
	http.ListenAndServe("127.0.0.1:9999", nil)
}
