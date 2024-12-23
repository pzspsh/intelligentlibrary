/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:30:16
*/
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path:", u.Path)
	fmt.Println("RawPath:", u.RawPath)
	fmt.Println("EscapedPath:", u.EscapedPath())
}
