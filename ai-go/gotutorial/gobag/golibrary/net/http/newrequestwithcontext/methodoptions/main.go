/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:34:59
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodOptions, "http://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
