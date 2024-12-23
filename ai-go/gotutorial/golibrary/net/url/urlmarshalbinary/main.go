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
	u, _ := url.Parse("https://example.org")
	b, err := u.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
