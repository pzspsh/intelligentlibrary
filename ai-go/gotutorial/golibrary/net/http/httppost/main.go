/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:43:39
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	requestUrl := "http://www.baidu.com"
	// request, err := http.Get(requestUrl)
	// request, err := http.Head(requestUrl)
	postvalue := url.Values{
		"username": {"xiaoming"},
		"address":  {"beijing"},
		"subject":  {"Hello"},
		"from":     {"china"},
	}
	// request, err := http.PostForm(requestUrl, postvalue)

	body := bytes.NewBufferString(postvalue.Encode())
	request, err := http.Post(requestUrl, "text/html", body) //Post方法
	if err != nil {
		fmt.Println(err)
	}

	defer request.Body.Close()
	fmt.Println(request.StatusCode)
	if request.StatusCode == 200 {
		rb, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println(rb)
		}
		fmt.Println(string(rb))
	}

}
