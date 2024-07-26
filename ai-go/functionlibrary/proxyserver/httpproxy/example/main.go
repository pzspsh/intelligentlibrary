/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 16:57:20
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	if conn == nil {
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024*4)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}

	request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(buf)))
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println(resp.Body)
	io.Copy(conn, resp.Body)
}

/*
Golang代理的应用场景

使用代理是解决网络编程中常见问题的有效方法。以下是一些常见的应用场景：

隐藏客户端IP地址
有些网站会通过IP地址识别并限制特定地区或特定用户的访问权限，此时使用代理可以隐藏真实的客户端IP地址，绕过地域限制。

实现负载均衡
对于高并发的系统，为了实现负载均衡，可以使用代理服务器来分发请求。代理服务器可以根据负载情况，将请求转发给可用的服务器，以达到最优的负载均衡效果。

加速数据传输
一些代理服务器通过优化数据传输的方式，可以加速数据传输速度。例如goproxy_cn库针对中国网络环境进行了优化，可提供更快的数据传输速度。
*/
