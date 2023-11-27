/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 17:59:48
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCEFG")
	str := strings.NewReader("123455")
	br := bufio.NewReader(s)
	b, _ := br.ReadString('\n')
	fmt.Println(b)
	br.Reset(str)
	b, _ = br.ReadString('\n')
	fmt.Println(b)
}
