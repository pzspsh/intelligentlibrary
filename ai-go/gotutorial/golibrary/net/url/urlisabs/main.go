/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:30:16
*/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := url.URL{Host: "example.com", Path: "foo"}
	fmt.Println(u.IsAbs())
	u.Scheme = "http"
	fmt.Println(u.IsAbs())
}
