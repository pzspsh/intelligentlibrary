/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:22:26
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func httpGet(requestUrl string) (err error) {
	Url, err := url.Parse(requestUrl)
	if err != nil {
		fmt.Printf("requestUrl parse failed, err:[%s]", err.Error())
		return
	}

	params := url.Values{}
	params.Set("query", "googlesearch")
	params.Set("content", "golang")
	Url.RawQuery = params.Encode()

	requestUrl = Url.String()
	fmt.Printf("requestUrl:[%s]\n", requestUrl)

	resp, err := http.Get(requestUrl)
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

// 把一些参数做成变量，然后放到 url 中，可以参考下面的方式
// 服务示例server2/main.go
func main() {
	var url = "http://10.10.19.200:8000/index"
	httpGet(url)
}
