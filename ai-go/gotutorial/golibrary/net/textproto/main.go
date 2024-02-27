/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:25:27
*/
package main

import (
	"fmt"
	"net/textproto"
)

func main() {
	var input string
	fmt.Print("Enter a MIME header field name: ")
	fmt.Scan(&input)
	canonical := textproto.CanonicalMIMEHeaderKey(input)
	fmt.Printf("Canonical format: %s\n", canonical)
}
