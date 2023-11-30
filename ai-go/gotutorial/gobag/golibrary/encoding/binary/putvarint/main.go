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

	for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		fmt.Print(x, "输出的可变长度为：", n, "，十六进制为：")
		fmt.Printf("%x\n", buf[:n])
	}
}
