/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:05:44
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var (
		piVar   float64
		boolVar bool
	)
	piByte := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	boolByte := []byte{0x00}
	piBuffer := bytes.NewReader(piByte)
	boolBuffer := bytes.NewReader(boolByte)
	binary.Read(piBuffer, binary.LittleEndian, &piVar)
	binary.Read(boolBuffer, binary.LittleEndian, &boolByte)
	fmt.Println("pi", piVar)     // pi 3.141592653589793
	fmt.Println("bool", boolVar) // bool false
}
