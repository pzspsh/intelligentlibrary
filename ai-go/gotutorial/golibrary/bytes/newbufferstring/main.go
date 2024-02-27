/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:01:06
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("ABCDEFGH")

	fmt.Println(b.String()) // ABCDEFGH

	fmt.Println(b.Len()) // 8

	fmt.Println(string(b.Next(2))) // AB
	tmp := make([]byte, 2)

	n, _ := b.Read(tmp)
	fmt.Println(string(tmp[:n])) //CD

	nextByte, _ := b.ReadByte()
	fmt.Println(string(nextByte)) // E

	line, _ := b.ReadString('G')
	fmt.Println(line)       // FG
	fmt.Println(b.String()) // H
	b = bytes.NewBufferString("abcdefgh")

	line2, _ := b.ReadBytes('b')
	fmt.Println(string(line2)) // ab
	fmt.Println(b.String())    // cdefgh

	r, n, _ := b.ReadRune()
	fmt.Println(r, n) // 99 1

	str := string(b.Bytes())
	fmt.Println(str) // defgh
}
