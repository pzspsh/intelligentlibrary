/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:03:56
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x", h.Sum(nil))
}
