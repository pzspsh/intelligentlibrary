/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:09:32
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

	num, err := binary.ReadUvarint(bytes.NewBuffer(sbuf))
	fmt.Println(num, err) //0 EOF

	num, err = binary.ReadUvarint(bytes.NewBuffer(buf))
	fmt.Println(num, err) //1161981756374523920 <nil>

	num, err = binary.ReadUvarint(bytes.NewBuffer(bbuf))
	fmt.Println(num, err) //4620746270195064848 binary: varint overflows a 64-bit integer
}
