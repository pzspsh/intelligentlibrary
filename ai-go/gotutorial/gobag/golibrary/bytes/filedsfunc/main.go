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
	s := bytes.FieldsFunc([]byte(" hi 你啊,    is-not.good, my,boy"), func(r rune) bool {
		return r == ',' || r == '-' || r == '.' // 只要是,-. 都可以作为分隔符
	})
	for _, v := range s {
		fmt.Print(string(v) + "|") //  hi 你啊|    is|not|good| my|boy|
	}
	fmt.Println()
}
