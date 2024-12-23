/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:56:57
*/
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var HTTPTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 连接超时时间
		KeepAlive: 60 * time.Second, // 保持长连接的时间
	}).DialContext, // 设置连接的参数
	MaxIdleConns:          500,              // 最大空闲连接
	IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
	ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
	MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
}

func main() {
	times := 50
	uri := "http://local.test.com/t.php"
	// uri := "http://www.baidu.com"

	// 短连接的情况
	start := time.Now()
	client := http.Client{} // 初始化http的client
	for i := 0; i < times; i++ {
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			panic("Http Req Failed " + err.Error())
		}
		resp, err := client.Do(req) // 发起请求
		if err != nil {
			panic("Http Request Failed " + err.Error())
		}
		defer resp.Body.Close()
		io.ReadAll(resp.Body)
	}
	fmt.Println("Orig GoNet Short Link", time.Since(start))

	// 长连接的情况

	start2 := time.Now()
	client2 := http.Client{Transport: HTTPTransport} // 初始化一个带有transport的http的client
	for i := 0; i < times; i++ {
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			panic("Http Req Failed " + err.Error())
		}
		resp, err := client2.Do(req)
		if err != nil {
			panic("Http Request Failed " + err.Error())
		}
		defer resp.Body.Close()
		io.ReadAll(resp.Body) // 如果不及时从请求中获取结果，此连接会占用，其他请求服务复用连接
	}
	fmt.Println("Orig GoNet Long Link", time.Since(start2))
}
