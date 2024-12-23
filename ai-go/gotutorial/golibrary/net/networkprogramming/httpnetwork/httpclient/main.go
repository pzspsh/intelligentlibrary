/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:27
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/index")
	if err != nil {
		log.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()

	// parse resp data
	// 获取服务器端读到的数据---header
	fmt.Println("Status = ", resp.Status)         // 状态
	fmt.Println("StatusCode = ", resp.StatusCode) // 状态码
	fmt.Println("Header = ", resp.Header)         // 响应头部
	fmt.Println("Body = ", resp.Body)             // 响应包体
	// resp body
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("response body:", string(content))
}
