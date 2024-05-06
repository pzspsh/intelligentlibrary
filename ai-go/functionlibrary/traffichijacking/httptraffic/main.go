/*
@File   : main.go
@Author : pan
@Time   : 2024-05-06 17:22:43
*/
package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 在这里可以修改请求，比如添加请求头、修改URL等

	// 创建一个新的请求，目标地址是实际的目标服务器
	newReq := &http.Request{
		Method: r.Method,
		URL:    r.URL,
		Header: r.Header,
		Body:   r.Body,
	}

	// 创建一个 HTTP 客户端，用于将请求转发给目标服务器
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 复制响应到客户端
	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	// 在这里可以修改响应，比如修改响应体、添加响应头等
}

func main() {
	// 监听本地端口，作为代理服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// 处理连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建读写协程，实现全双工通信
	go readFromClient(conn)
	writeToClient(conn)
}

func readFromClient(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的 HTTP 请求
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		// 将请求转发给处理函数
		reader := bytes.NewBuffer(buf[:n])
		bufferedReader := bufio.NewReader(reader)
		req, err := http.ReadRequest(bufferedReader)
		// req, err := http.ReadRequest(buf[:n])
		if err != nil {
			// http.Error(conn, err.Error(), http.StatusBadRequest)
			http.Error(http.ResponseWriter(nil), err.Error(), http.StatusBadRequest)
			return
		}

		// 创建一个新的响应记录器，用于构建响应
		rec := httptest.NewRecorder()

		// 处理请求
		handleRequest(rec, req)

		// 将响应写回客户端
		resp := rec.Result()
		resp.Write(conn)
	}
}

func writeToClient(conn net.Conn) {
	// 这里可以实现对客户端的写操作，但在这个例子中我们不需要
}
