/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"net/http"
	"strconv"
	"strings"
)

func SetCook(cookie string) map[string]string {
	cookieMap := make(map[string]string)
	vals := strings.Split(cookie, ";") //切分cookie
	for _, val := range vals {
		cook := strings.Split(val, "=") //将cookie分为map类型
		cookieMap[strings.TrimSpace(cook[0])] = strings.TrimSpace(cook[1])
	}
	return cookieMap
}

// 上传文件 =>multipart/form-data
func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		cookieMap := SetCook(req.Header.Get("cookie")) //处理cookie函数
		counter := 0
		if v, err := strconv.Atoi(cookieMap["counter"]); err == nil {
			counter = v
		}
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter + 1),
			HttpOnly: true,
		} //当访问这个网页的时候，服务器返回一个cookie回去
		http.SetCookie(resp, counterCookie) //将cookie的值使用响应体的方式返回
	})
	http.ListenAndServe("127.0.0.1:9999", nil)
}
