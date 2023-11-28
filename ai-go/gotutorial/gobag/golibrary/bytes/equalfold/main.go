/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:42:51
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.EqualFold([]byte{}, []byte{}))       // true
	fmt.Println(bytes.EqualFold([]byte{'A'}, []byte{'a'})) // true
	fmt.Println(bytes.EqualFold([]byte{'B'}, []byte{'a'})) // false
	fmt.Println(bytes.EqualFold([]byte{}, nil))            // true
}
