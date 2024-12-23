/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 09:46:31
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}
