/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:19:14
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	txts, err := net.LookupTXT("google.com")
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		fmt.Printf("%s\n", txt)
	}
}
