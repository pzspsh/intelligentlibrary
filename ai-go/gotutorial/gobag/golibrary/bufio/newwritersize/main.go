/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:08:30
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 示例：Available、Buffered、WriteString、Flush
func main() {
	buf := bufio.NewWriterSize(os.Stdout, 0)
	fmt.Println(buf.Available(), buf.Buffered()) // 4096 0

	buf.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Println(buf.Available(), buf.Buffered()) // 4070 26

	// 缓存后统一输出，避免终端频繁刷新，影响速度
	buf.Flush() // ABCDEFGHIJKLMNOPQRSTUVWXYZ
}
