/*
@File   : demo.go
@Author : pan
@Time   : 2023-11-27 17:25:54
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p)) //123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b) //asdf
}
