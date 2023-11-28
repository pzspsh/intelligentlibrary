/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:06:45
*/
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b := new(bytes.Buffer)

	b.WriteByte('a')
	fmt.Println(b.String()) // a

	b.Write([]byte{98, 99})
	fmt.Println(b.String()) // abc

	b.WriteString(" hello")
	fmt.Println(b.String()) // abc hello

	b.Truncate(3)
	fmt.Println(b.String()) // abc

	n, _ := b.WriteTo(os.Stdout) // abc
	fmt.Println(n)               // 3

	b.Reset()
	fmt.Println(b.Len(), b.String()) // 0
}
