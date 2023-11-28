/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:42:21
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Equal([]byte{}, []byte{}))            // true
	fmt.Println(bytes.Equal([]byte{'A', 'B'}, []byte{'a'})) // false
	fmt.Println(bytes.Equal([]byte{'a'}, []byte{'a'}))      // true
	fmt.Println(bytes.Equal([]byte{}, nil))                 // true
}
