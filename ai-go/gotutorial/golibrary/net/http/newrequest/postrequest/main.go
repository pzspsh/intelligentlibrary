/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:26:21
*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	req, _ := http.NewRequest("POST", "http://www.google.com/search?q=foo&q=bar&both=x&prio=1&orphan=nope&empty=not",
		strings.NewReader("z=post&both=y&prio=2&=nokey&orphan;empty=&"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//没有解析前req.Form和req.PostForm中的值为空
	fmt.Println(req)
	fmt.Println(req.Form)
	fmt.Println(req.PostForm)
	fmt.Println()

	//解析后对应的值才写入req.Form和req.PostForm
	req.ParseForm()
	fmt.Println(req)
	fmt.Println(req.Form)
	fmt.Println(req.PostForm)
}
