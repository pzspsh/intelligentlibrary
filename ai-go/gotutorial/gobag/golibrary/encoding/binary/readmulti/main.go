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
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	r := bytes.NewReader(b)

	var data struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(data.PI)
	fmt.Println(data.Uate)
	fmt.Printf("% x\n", data.Mine)
	fmt.Println(data.Too)
}
