/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:40:56
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello web server") //将字符串写入到w，即在客户端输出
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayhelloName) //设置访问的路由
	// err := http.ListenAndServe(":9090", nil) //设置监听的端口
	// if err != nil {
	//     log.Fatal("ListenAndServe : ", err)
	// }
	//如果使用的是Server，则等价于:
	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      mux,
	}
	server.ListenAndServe()
}
