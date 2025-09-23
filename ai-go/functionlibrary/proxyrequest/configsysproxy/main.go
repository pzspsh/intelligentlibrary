package main

import (
	"net/http"
	"os"
)

func main() {
	// 设置HTTP代理
	err := os.Setenv("HTTP_PROXY", "http://proxyaddress:port")
	if err != nil {
		panic(err)
	}

	// 设置HTTPS代理
	err = os.Setenv("HTTPS_PROXY", "https://proxyaddress:port")
	if err != nil {
		panic(err)
	}

	// 使用代理发起请求
	resp, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
