/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 13:05:34
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetHttpsSkip(url, token string) ([]byte, error) {
	// 创建各类对象
	var client *http.Client
	var request *http.Request
	var resp *http.Response
	var body []byte
	var err error

	// 这里请注意，使用 InsecureSkipVerify: true 来跳过证书验证
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	// 获取 request请求
	request, err = http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("GetHttpSkip Request Error:", err)
		return nil, nil
	}

	// 加入 token
	request.Header.Add("Authorization", token)
	resp, err = client.Do(request)
	if err != nil {
		log.Println("GetHttpSkip Response Error:", err)
		return nil, nil
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer client.CloseIdleConnections()
	return body, nil
}

func main() {
	resp, err := GetHttpsSkip("https://10.10.102.91:10250/metrics", "Bearer eyxxxxxxxxxxxxxxxxxxxx....xxxxx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
}
