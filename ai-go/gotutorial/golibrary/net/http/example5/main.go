/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method, req.URL, req.Proto)
		fmt.Println(req.Header) //获取请求头的信息
		fmt.Println(req.Header.Get("User-Agent"))
	})
	http.ListenAndServe("127.0.0.1:9999", nil)
}
