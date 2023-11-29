/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 14:56:23
*/
package main

import (
	"bytes"
	"fmt"
)

/*
SplitAfterN个片s在每个sep实例之后分成子片，并返回这些子片的一个片。
如果sep为空，SplitAfterN在每个UTF-8序列之后进行分割。计数决定了要返回的子片的数量
*/
func main() {
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2)) // ["a," "b,c"]
}
