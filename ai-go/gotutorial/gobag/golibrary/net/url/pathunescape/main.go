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
	escapedPath := "my%2Fcool+blog&about%2Cstuff"
	path, err := url.PathUnescape(escapedPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

}
