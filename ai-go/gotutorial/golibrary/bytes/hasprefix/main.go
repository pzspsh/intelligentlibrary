package main

import (
	"bytes"
	"fmt"
)

/*
HasPrefix测试字节片s是否以prefix开头。
*/
func main() {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go"))) // true
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))  // false
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))   // true
	fmt.Println(bytes.HasPrefix([]byte{1, 2, 3}, []byte{1}))     // true
	fmt.Println(bytes.HasPrefix([]byte{1, 2, 3}, []byte{2}))     // false

}
