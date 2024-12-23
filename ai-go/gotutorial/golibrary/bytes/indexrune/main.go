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
IndexRune将s解释为utf -8编码的代码点序列。它返回给定符文在s中第一次出现的字节索引。
如果rune不存在于s中，则返回-1。如果r为utf8。RuneError，它返回任何无效UTF-8字节序
列的第一个实例。
*/
func main() {
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))
	fmt.Println(bytes.IndexRune([]byte("你好吗,不太好啊,hi go go go go go go go go go"), '不')) // 9

}
