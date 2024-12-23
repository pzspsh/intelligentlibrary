/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:29:22
*/
package main

import (
	"log"
	"os"
)

func main() {
	// 裁剪一个文件到100个字节。
	// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过100个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的100个字节的文件。
	// 传入0则会清空文件。

	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
