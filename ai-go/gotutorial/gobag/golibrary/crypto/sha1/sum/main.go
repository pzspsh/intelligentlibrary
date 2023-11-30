/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:00:10
*/
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha1.Sum(data))
}
