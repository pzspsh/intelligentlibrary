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

/*
将片分割成所有用sep分隔的子片，并返回这些分隔符之间的子片的一个片。
如果sep为空，Split在每个UTF-8序列之后进行分割。它相当于计数为-1的SplitN。
*/
func main() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))

	s := bytes.Split([]byte(" hi 你啊,    is not good, is my boy"), []byte("is"))
	for _, v := range s {
		fmt.Print(string(v) + "|") //  |  hi 你啊,    | not good, | my boy|
	}
	fmt.Println()

	fmt.Println(bytes.Join([][]byte{{1, 1}, {2, 2}, {3, 3}}, []byte{9})) // [1 1 9 2 2 9 3 3]
}
