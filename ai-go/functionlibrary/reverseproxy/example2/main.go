/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 17:06:26
*/
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	// 创建一个反向代理的目标URL
	targetURL, err := url.Parse("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个新的反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	// 更新请求头，设置代理服务器的地址
	req.URL.Host = targetURL.Host
	req.URL.Scheme = targetURL.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = targetURL.Host
	// 执行反向代理请求
	proxy.ServeHTTP(res, req)
}

func main() {
	// 注册代理处理函数
	http.HandleFunc("/", handleRequestAndRedirect)
	// 启动代理服务器
	if err := http.ListenAndServe(":8080", nil); err != nil { // 访问8080端口直接跳转到http://www.baidu.com
		log.Fatal(err)
	}
}
