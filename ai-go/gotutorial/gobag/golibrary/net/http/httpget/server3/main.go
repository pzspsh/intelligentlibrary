/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:49:31
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HttpStart(port int) {
	http.HandleFunc("/hello", helloFunc)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("监听失败：", err.Error())
	}
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("打印Header参数列表：")
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	fmt.Println("打印Form参数列表：")
	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	//验证用户名密码，如果成功则header里返回session，失败则返回StatusUnauthorized状态码
	w.WriteHeader(http.StatusOK)
	if (r.Form.Get("user") == "admin") && (r.Form.Get("pass") == "888") {
		w.Write([]byte("hello,验证成功！"))
	} else {
		w.Write([]byte("hello,验证失败了！"))
	}
}

func main() {
	HttpStart(8080)
}
