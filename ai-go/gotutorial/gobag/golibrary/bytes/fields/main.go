/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:56:31
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := bytes.Fields([]byte(" hi 你啊,    is not good, my boy"))
	for _, v := range s {
		fmt.Print(string(v) + "|") // hi|你啊,|is|not|good,|my|boy|
	}
	fmt.Println()
}
