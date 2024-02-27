/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

// 使用对象接口的思维来维护
type TimeHandler struct {
} //定义结构体来实现接口

func (h *TimeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprint(response, now)
}

func main() {
	http.Handle("/time/", &TimeHandler{}) //传入实现结构的方法
	http.ListenAndServe("127.0.0.1:9999", nil)
}
