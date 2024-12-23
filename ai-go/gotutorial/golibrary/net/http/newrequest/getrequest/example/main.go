/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:23:36
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpGet(requestUrl string) (err error) {
	client := &http.Client{}
	requestGet, _ := http.NewRequest("GET", requestUrl, nil)

	requestGet.Header.Add("query", "googlesearch")
	requestGet.Header.Add("content", "golang")

	resp, err := client.Do(requestGet)
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := io.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return
}

// 服务实例server2/main.go
func main() {
	var url = "http://10.10.19.200:8000/index"
	httpGet(url)
}
