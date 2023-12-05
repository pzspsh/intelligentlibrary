/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:05:03
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

// 示例服务端代码server/main.go
func main() {
	client := &http.Client{}
	apiURL := "http://127.0.0.1:8080/get"

	req, err := http.NewRequest("GET", apiURL, nil)
	//添加查询参数
	q := req.URL.Query()
	q.Add("username", "admin")
	q.Add("password", "123")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Printf("post failed, err:%v\n\n", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n\n", err)
		return
	}
	fmt.Println(string(b))
}
