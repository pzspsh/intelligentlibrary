/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:14:25
*/
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// 公共打印方法
func printBody(r *http.Response) {
	defer func() { _ = r.Body.Close() }()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

// 设置请求头
func requestByHead() {
	/*
			// 总超时
			client := http.Client{
		    Timeout:   time.Duration(10 * time.Second),
			}
	*/

	/*
		// 连接超时
				t := &http.Transport{
			    Dial: func(network, addr string) (net.Conn, error) {
			        timeout := time.Duration(2 * time.Second)
			        return net.DialTimeout(network, addr, timeout)
			    },
			}
	*/

	/*
		// 读取超时
		   t := &http.Transport{
		   		ResponseHeaderTimeout: time.Second * 8,
		   	}
	*/

	t := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			timeout := time.Duration(2 * time.Second)
			return net.DialTimeout(network, addr, timeout)
		},
		ResponseHeaderTimeout: time.Second * 8,
	}

	//创建一个请求request
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: t,
		Timeout:   time.Duration(10 * time.Second),
	}
	//在request上直接修改请求头
	request.Header.Add("user-agent", "IOS")
	r, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	printBody(r)
}

func main() {
	//如何设置请求头
	requestByHead()
}
