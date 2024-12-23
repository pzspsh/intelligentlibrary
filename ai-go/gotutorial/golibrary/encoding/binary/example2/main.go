/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:06:12
*/
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var a int
	p := &a
	b := [10]int64{1}
	s := "adsa"
	bs := make([]byte, 10)

	fmt.Println(binary.Size(a))  // -1
	fmt.Println(binary.Size(p))  // -1
	fmt.Println(binary.Size(b))  // 80
	fmt.Println(binary.Size(s))  // -1
	fmt.Println(binary.Size(bs)) // 10
}
