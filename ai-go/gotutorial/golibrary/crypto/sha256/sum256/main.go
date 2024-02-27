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
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x", sum)
}
