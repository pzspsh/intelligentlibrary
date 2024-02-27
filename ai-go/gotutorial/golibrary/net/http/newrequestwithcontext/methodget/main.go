/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 11:52:06
*/
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost:8080/user?name=tom", nil)
	if err != nil {
		return
	}
	request.Header.Set("headerParam", "header")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(string(bytes)) // {"code":0,"data":{"list":[{"name":"小明","age":20},{"name":"小红","age":18}]},"message":"success"}
}
