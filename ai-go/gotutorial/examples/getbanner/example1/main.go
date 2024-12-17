/*
@File   : main.go
@Author : pan
@Time   : 2024-12-16 17:45:49
*/
package main

import (
	"fmt"
	"net"
	"net/http"
	"net/textproto"
	"time"
)

func main() {
	// 获取HTTP服务的Banner信息
	httpBanner := getHTTPBanner("http://www.example.com")
	fmt.Println("HTTP Banner:", httpBanner)

	// 获取FTP服务的Banner信息
	ftpBanner := getFTPBanner("ftp://ftp.example.com")
	fmt.Println("FTP Banner:", ftpBanner)
}

func getHTTPBanner(url string) string {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return "Error sending GET request: " + err.Error()
	}
	defer resp.Body.Close()

	// 读取响应的头部信息
	banner := resp.Header.Get("Server")
	if banner == "" {
		banner = "No banner information found."
	}
	return banner
}

func getFTPBanner(url string) string {
	// 解析FTP URL
	conn, err := net.DialTimeout("tcp", url, 5*time.Second)
	if err != nil {
		return "Error connecting to FTP server: " + err.Error()
	}

	// 创建FTP连接
	ftpConn := textproto.NewConn(conn)
	defer ftpConn.Close()

	// 发送FTP命令获取Banner信息
	ftpConn.Cmd("NOOP")
	ftpConn.StartResponse(0)
	defer ftpConn.EndResponse(0)

	// 读取FTP服务器的响应
	resp, err := ftpConn.ReadLine()
	if err != nil {
		return "Error reading FTP response: " + err.Error()
	}
	return resp
}
