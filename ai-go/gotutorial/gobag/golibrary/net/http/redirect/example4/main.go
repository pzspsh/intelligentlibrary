/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:05:43
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	var r *http.Response
	var err error
	history := make([]*http.Response, 0)

	client := http.Client{
		CheckRedirect: func(req *http.Request, hrs []*http.Request) error {
			if len(hrs) >= 10 {
				return errors.New("redirect to many times")
			}
			history = append(history, req.Response)
			return nil
		},
	}

	r, err = client.Get("http://github.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

/*
首先创建了 http.Response切片的变量，名称为 history。接着在 http.Client中为 CheckRedirect
赋予一个匿名函数，用于控制重定向的行为。CheckRedirect函数的第一个参数表示下次将要请求的Request，
第二个参数表示已经请求过的 Request。
当发生重定向时，当前的Request会保存上次请求的Response，故而此处可以将req.Response追加到
history变量中。
*/
