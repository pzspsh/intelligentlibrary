/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 09:58:39
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}
