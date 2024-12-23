/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:19:18
*/
package main

import (
	"fmt"
	"mime"
)

func main() {
	mediatype, params, err := mime.ParseMediaType("text/html; charset=utf-8")
	if err != nil {
		panic(err)
	}

	fmt.Println("type:", mediatype)
	fmt.Println("charset:", params["charset"])
}
