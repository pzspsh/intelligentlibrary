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
	fmt.Println(bytes.Contains([]byte{1, 2, 3}, []byte{1}))    // true
	fmt.Println(bytes.Contains([]byte{1, 2, 3}, []byte{1, 3})) // false
}
