/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 19:36:19
*/
package main

import (
	"fmt"
	"hash/crc32"
	"io"
)

func main() {
	//hash := crc32.NewIEEE()

	check_str := "Discard medicine more than two years old."
	//hash.write(check_str)
	//hash.Sum()
	ieee := crc32.NewIEEE()
	io.WriteString(ieee, check_str)
	s := ieee.Sum32()
	fmt.Println("IEEE(%s) = 0x%x", check_str, s)

}
