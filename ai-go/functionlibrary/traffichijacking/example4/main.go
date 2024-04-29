/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 15:46:25
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.ListenAndServe(":8080", http.HandlerFunc(proxyHandler))
	http.HandleFunc("/", proxyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// 打印请求的URL和方法
	fmt.Println("=========================================================================")
	fmt.Printf("URL: %s %s\n", r.URL.Path, r.Method)
	// fmt.Printf("Method: %s\n", r.Method)

	// fmt.Println("Headers:")
	for name, headers := range r.Header { // 打印请求头部信息
		for _, h := range headers {
			fmt.Printf("%s: %s\n", name, h)
		}
	}
	fmt.Println(r.Body)
	buf := make([]byte, 1024) // 打印请求体
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err != nil {
			break
		}
	}
}

/* import (
    "fmt"
    "log"
    "net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
    // 打印请求的URL和方法
    fmt.Printf("URL: %s\n", r.URL.Path)
    fmt.Printf("Method: %s\n", r.Method)

    // 打印请求头部信息
    fmt.Println("Headers:")
    for name, headers := range r.Header {
        for _, h := range headers {
            fmt.Printf("%s: %s\n", name, h)
        }
    }

    // 打印请求体
    buf := make([]byte, 1024)
    for {
        n, err := r.Body.Read(buf)
        if n > 0 {
            fmt.Print(string(buf[:n]))
        }
        if err != nil {
            break
        }
    }
}

func main() {
    // 注册路由处理函数
    http.HandleFunc("/", requestHandler)

    // 启动HTTP服务器并监听特定端口
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
} */
