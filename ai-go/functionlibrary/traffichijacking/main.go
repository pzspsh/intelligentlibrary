/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 11:02:12
*/
package main

/* import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Connect to Burp Suite proxy
	proxyConn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Error connecting to Burp Suite proxy: %v", err)
	}
	defer proxyConn.Close()

	// Forward request to Burp Suite
	err = r.WriteProxy(proxyConn)
	if err != nil {
		log.Fatalf("Error forwarding request to Burp Suite: %v", err)
	}

	// Read response from Burp Suite
	resp, err := http.ReadResponse(bufio.NewReader(proxyConn), r)
	if err != nil {
		log.Fatalf("Error reading response from Burp Suite: %v", err)
	}

	// Forward response to client
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
*/

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

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(proxyHandler))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// 打印请求的URL和方法
	fmt.Println("=========================================================================")
	fmt.Printf("URL: %s\n", r.URL.Path)
	// fmt.Printf("Method: %s\n", r.Method)

	// 打印请求头部信息
	// fmt.Println("Headers:")
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
