/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:15:14
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var i int

func receiveReq(w http.ResponseWriter, r *http.Request) {
	//i记录请求的次数
	i++
	fmt.Printf("第 %d 次请求的Cookie： ", i)
	fmt.Println(r.Cookies())
	//设置cookie
	tNow := time.Now()
	cookie := &http.Cookie{
		Name:    "XMEN",
		Value:   "STORM" + strconv.Itoa(i),
		Expires: tNow.AddDate(1, 0, i),
	}
	http.SetCookie(w, cookie)
	//返回信息
	w.Write(([]byte("your cookie has been received")))

}

func main() {
	fmt.Println("Server Start Now!")
	http.HandleFunc("/test", receiveReq)
	err := http.ListenAndServe("127.0.0.1:8889", nil)
	if err != nil {
		fmt.Println("ListenAndServe ERROR: ", err)
	}

}
