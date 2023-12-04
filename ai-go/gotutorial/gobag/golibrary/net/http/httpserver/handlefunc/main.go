/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:39:07
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(req.Form) //这些信息是输出到服务端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key : ", k)
		fmt.Println("value : ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello web server") //将字符串写入到w，即在客户端输出
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}
