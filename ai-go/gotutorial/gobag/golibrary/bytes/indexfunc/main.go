/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:45:29
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexFunc([]byte("hi go"), func(r rune) bool {
		return r == 'g'
	})) // 3
}
