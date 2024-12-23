/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:02:09
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	var data = []any{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x", buf.Bytes())
}
