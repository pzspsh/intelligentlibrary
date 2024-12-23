/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:24:20
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}
