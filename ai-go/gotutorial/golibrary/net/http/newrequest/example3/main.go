/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:23:01
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.baidu.com/", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	//没有解析前req.Form和req.PostForm中的值为空
	fmt.Println(req.ProtoAtLeast(1, 0)) //true
	fmt.Println(req.ProtoAtLeast(1, 1)) //true
	fmt.Println(req.UserAgent())        //Mozilla/5.0
	fmt.Println(req.Referer())          //因为没有来源，为空
}
