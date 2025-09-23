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
