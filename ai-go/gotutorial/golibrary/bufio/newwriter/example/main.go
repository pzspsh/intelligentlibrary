/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 17:05:45
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	for _, c := range "abc123" {
		err := writer.WriteByte(byte(c))
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
	fmt.Println(buf.String())
}
