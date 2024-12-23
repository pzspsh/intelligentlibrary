/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:13:22
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/cookies", nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:   "name",
		Value:  "poloxue",
		Domain: "httpbin.org",
		Path:   "/cookies",
	})

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
