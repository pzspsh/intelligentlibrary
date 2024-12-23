/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 09:53:48
*/
package main

import (
	"fmt"
	"hash/crc64"
	"strconv"
)

func main() {
	m := make(map[uint64]int)
	for i := 0; i < 200; i++ {
		table := crc64.MakeTable(crc64.ISO)
		checksum := crc64.Checksum([]byte("32846561956_14:200002984#color 4"), table)
		n := fmt.Sprintf("%x", checksum)
		fmt.Println(n, len(n))
		fmt.Println(checksum, len(strconv.FormatUint(checksum, 10)))
		m[checksum] = 1
	}
	fmt.Println(len(m))

}
