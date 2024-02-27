/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:45:29
*/
package main

import (
	"bytes"
	"fmt"
)

/*
IndexAny将s解释为utf -8编码的Unicode码点序列。它以字符的形式返回任何Unicode码点在s中
第一次出现的字节索引。如果chars为空或没有共同的代码点，则返回-1。
*/
func main() {
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aeiouy")) // 2
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))   // -1
	fmt.Println(bytes.IndexAny([]byte("hi go"), "go"))       // 3
}
