/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:09:59
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var sbuf []byte
	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, err := binary.ReadVarint(bytes.NewBuffer(sbuf))
	fmt.Println(num, err) //0 EOF

	num, err = binary.ReadVarint(bytes.NewBuffer(buf))
	fmt.Println(num, err) //580990878187261960 <nil>

	num, err = binary.ReadVarint(bytes.NewBuffer(bbuf))
	fmt.Println(num, err) //2310373135097532424 binary: varint overflows a 64-bit integer
}
