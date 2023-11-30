/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:10:55
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)
}
