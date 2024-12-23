/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:02:34
*/
package main

import (
	"expvar"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	ok          = expvar.NewInt("200") // 200 计算器
	notFound    = expvar.NewInt("404") // 404 计数器
	serverError = expvar.NewInt("500") // 500 计数器
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	if err != nil {
		log.Fatal(err)
	}

	ok.Add(1) // 增加 200 计数器
}

// 为了模拟 404, 500 错误
// 随机返回 Http Code [200, 404, 500]
func random(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	var code int

	switch n {
	case 0:
		code = http.StatusOK
		ok.Add(1) // 增加 200 计数器
	case 1:
		code = http.StatusNotFound
		notFound.Add(1) // 增加 404 计数器
	case 2:
		code = http.StatusInternalServerError
		serverError.Add(1) // 增加 500 计数器
	}

	w.WriteHeader(code)
}

func main() {
	http.HandleFunc("/", helloWorld)   // 默认返回地址
	http.HandleFunc("/random", random) // 随机返回状态码地址

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
首先使用 expvar 包初始化了 3 个计数器变量
接下来定义了两个 HTTP 路由回调方法
helloWorld() 负责处理默认路由 /
random() 负责处理随机返回状态码路由 / (主要用这个来做测试)
最后，分别设置了路由的回调方法，并且监听端口 8080
*/
