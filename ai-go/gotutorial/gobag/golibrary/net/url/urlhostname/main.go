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
	u, err := url.Parse("https://example.org:8000/path")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
	u, err = url.Parse("https://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
}
