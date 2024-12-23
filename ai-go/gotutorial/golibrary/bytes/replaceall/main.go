/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 14:52:21
*/
package main

import (
	"bytes"
	"fmt"
)

/*
ReplaceAll返回片s的副本，其中所有不重叠的old实例都替换为new。如果old为空，
则在片的开头和每个UTF-8序列之后匹配，为k-rune片产生最多k+1个替换。
*/
func main() {
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))
}
