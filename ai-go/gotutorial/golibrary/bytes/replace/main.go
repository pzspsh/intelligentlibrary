/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
)

/*
Replace返回切片s的一个副本，其中旧的前n个不重叠的实例被new替换。如果old为空，
则在片的开头和每个UTF-8序列之后匹配，为k-rune片产生最多k+1个替换。如果n < 0，
则替换次数没有限制。
*/
func main() {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
	fmt.Println(bytes.Replace([]byte{1, 2, 1, 2, 3, 1, 2, 1, 2}, []byte{1, 2}, []byte{0, 0}, 3)) // [0 0 0 0 3 0 0 1 2]
}
