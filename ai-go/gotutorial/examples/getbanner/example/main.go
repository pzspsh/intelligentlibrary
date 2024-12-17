/*
@File   : main.go
@Author : pan
@Time   : 2024-12-16 17:45:49
*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// 要获取Banner的网站URL
	url := "http://www.example.com"

	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应的头部信息
	banner := resp.Header.Get("Server")

	// 如果Server头部不存在，尝试从其他头部获取信息
	if banner == "" {
		banner = resp.Header.Get("X-Powered-By")
	}

	// 如果仍然没有找到Banner信息，尝试解析HTML内容
	if banner == "" {
		// 读取响应的前1024字节以获取HTML内容
		body := make([]byte, 1024)
		_, err := resp.Body.Read(body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// 在HTML内容中查找Banner信息
		// 这里假设Banner信息包含在<title>标签中
		html := string(body)
		start := strings.Index(html, "<title>")
		end := strings.Index(html, "</title>")
		if start != -1 && end != -1 {
			banner = html[start+7 : end]
		}
	}

	// 输出Banner信息
	if banner != "" {
		fmt.Println("Banner:", banner)
	} else {
		fmt.Println("No banner information found.")
	}
}
