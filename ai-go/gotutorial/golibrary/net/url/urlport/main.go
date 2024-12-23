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
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Port())
	u, err = url.Parse("https://example.org:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Port())
}
