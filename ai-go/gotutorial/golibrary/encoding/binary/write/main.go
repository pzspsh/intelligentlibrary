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
	"math"
)

func main() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}
