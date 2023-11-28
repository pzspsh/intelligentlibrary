/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:56:31
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Join([][]byte{{1, 1}, {2, 2}, {3, 3}}, []byte{9})) // [1 1 9 2 2 9 3 3]
}
