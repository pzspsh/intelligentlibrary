package main

import (
	"bytes"
	"fmt"
)

/*
LastIndexAny解释为utf -8编码的Unicode码点序列。它以字符的形式返回任何Unicode
码点在s中最后出现的字节索引。如果chars为空或没有共同的代码点，则返回-1。
*/
func main() {
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "MüQp")) // 5
	fmt.Println(bytes.LastIndexAny([]byte("go 地鼠"), "地大"))       // 3
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "z,!.")) // -1
}
