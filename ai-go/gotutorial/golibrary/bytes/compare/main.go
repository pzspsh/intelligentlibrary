/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:42:04
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Compare([]byte{}, []byte{}))   // 0
	fmt.Println(bytes.Compare([]byte{1}, []byte{2})) // -1
	fmt.Println(bytes.Compare([]byte{2}, []byte{1})) // 1
	fmt.Println(bytes.Compare([]byte{}, nil))        //0
}
