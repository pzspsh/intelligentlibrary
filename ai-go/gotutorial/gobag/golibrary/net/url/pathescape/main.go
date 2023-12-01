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
	path := url.PathEscape("my/cool+blog&about,stuff")
	fmt.Println(path)

}
