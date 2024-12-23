/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:07:34
*/
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// 创建代理服务器的URL
	// proxyUrl, _ := url.Parse("http://proxy.server.address:port")
	proxyUrl, _ := url.Parse("http://127.0.0.1:1080")

	// 创建ReverseProxy
	reverseProxy := httputil.NewSingleHostReverseProxy(proxyUrl)

	// 创建HTTP服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置代理信息
		r.URL.Host = "www.google.com"
		r.URL.Scheme = "https"
		// 将请求转发给代理服务器
		reverseProxy.ServeHTTP(w, r)
	})

	// 启动HTTP服务器
	http.ListenAndServe(":8080", nil)
}
