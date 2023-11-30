/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:02:09
*/
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	inputs := [][]byte{
		{0x01},
		{0x02},
		{0x7f},
		{0x80, 0x01},
		{0xff, 0x01},
		{0x80, 0x02},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b)
		if n != len(b) {
			fmt.Println("Uvarint did not consume all of in")
		}
		fmt.Println(x)
	}
}
