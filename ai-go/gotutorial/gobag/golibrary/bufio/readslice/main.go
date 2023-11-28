/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 18:06:10
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

// 示例：ReadSlice
func main() {
	// 尾部有换行标记
	buf := bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	// "ABCDEFG\n"   <nil>
	// ""   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	// "ABCDEFG"   EOF
}
