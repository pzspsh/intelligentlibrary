/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:15:14
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	jar, _ := cookiejar.New(nil)
	fmt.Println("Start Request Server")
	client := http.Client{
		Jar: jar,
	}
	url := "http://127.0.0.1:8889/test"
	req, _ := http.NewRequest("GET", url, nil)
	//第一次发请求
	client.Do(req)
	fmt.Printf("第一次 %s \n", req.Cookies())

	//第二次发请求
	client.Do(req)
	fmt.Printf("第二次 %s \n", req.Cookies())

	//第三次发请求
	client.Do(req)
	fmt.Printf("第三次 %s \n", req.Cookies())

	//第四次发请求
	client.Do(req)
	fmt.Printf("第四次 %s \n", req.Cookies())

	//第五次发请求
	client.Do(req)
	fmt.Printf("第五次 %s \n", req.Cookies())

}
