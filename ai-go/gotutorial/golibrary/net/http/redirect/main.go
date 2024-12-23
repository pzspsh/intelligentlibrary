/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:15:57
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "http://www.qq.com"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, _ := client.Do(reqest)
	fmt.Println(response.StatusCode)
}
