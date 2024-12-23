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

/*
Equal报告a和b的长度和字节数是否相同。nil参数相当于空切片。
*/
func main() {
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))    // true
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))   // false
	fmt.Println(bytes.Equal([]byte{}, []byte{}))            // true
	fmt.Println(bytes.Equal([]byte{'A', 'B'}, []byte{'a'})) // false
	fmt.Println(bytes.Equal([]byte{'a'}, []byte{'a'}))      // true
	fmt.Println(bytes.Equal([]byte{}, nil))                 // true
}
