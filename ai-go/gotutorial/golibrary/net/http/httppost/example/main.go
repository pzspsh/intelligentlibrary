/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:07:57
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 示例服务端代码server/main.go
func main() {
	url := "http://127.0.0.1:8080/post"
	// 表单格式数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=王二小&amp;age=18"
	// json格式数据，一般使用结构体发送或解析json格式数据
	contentType := "application/json"
	data := `{"name":"王二小","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("post resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
