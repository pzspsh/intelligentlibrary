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
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
	}
}
