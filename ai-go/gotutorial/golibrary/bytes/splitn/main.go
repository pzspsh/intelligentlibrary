/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 14:57:26
*/
package main

import (
	"bytes"
	"fmt"
)

/*
要围绕分隔符的第一个实例进行拆分，请参见切割。
*/
func main() {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2)) // ["a" "b,c"]
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)
}
