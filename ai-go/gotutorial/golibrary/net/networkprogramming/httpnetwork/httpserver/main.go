/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:45
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// register callback func
	http.HandleFunc("/index", indexHandle)
	log.Println("HTTP server up at", "http://localhost:8080")
	// bind ip and start recv req
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method = ", r.Method) //请求方法
	fmt.Println("URL = ", r.URL)       // 浏览器发送请求文件路径
	fmt.Println("header = ", r.Header) // 请求头
	fmt.Println("body = ", r.Body)     // 请求包体
	fmt.Println(r.RemoteAddr, "连接成功")  //客户端网络地址

	w.Write([]byte("Hello from http server"))
	//fmt.Fprint(w, "Hello from http server")
}
