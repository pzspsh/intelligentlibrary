/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 18:02:50
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	n, err := br.Read(p)
	fmt.Printf("%-20s %-2v %v\n", p[:n], n, err)

	n, err = br.Read(p)
	fmt.Printf("%-20s %-2v %v\n", p[:n], n, err)

	n, err = br.Read(p)
	fmt.Printf("%-20s %-2v %v\n", p[:n], n, err)

	n, err = br.Read(p)
	fmt.Printf("%-20s %-2v %v\n", p[:n], n, err)

	n, err = br.Read(p)
	fmt.Printf("%-20s %-2v %v\n", p[:n], n, err)

	n, err = br.Read(p)
	fmt.Printf("%-20s %-2v %v\n", p[:n], n, err)
}
