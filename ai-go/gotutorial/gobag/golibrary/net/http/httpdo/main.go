/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:44:13
*/
package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := &http.Cookie{Name: "userId", Value: strconv.Itoa(12345)}

	request.AddCookie(cookie) //request中添加cookie

	//设置request的header
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		r, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(r))
	}
}
