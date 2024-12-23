/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:28:44
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

// 请求在认证时忽略证书校验
func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
