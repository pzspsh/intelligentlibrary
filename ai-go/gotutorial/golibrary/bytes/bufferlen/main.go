/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:17:01
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%d", b.Len())
}
