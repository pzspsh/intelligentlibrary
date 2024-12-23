/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:39:39
*/
package main

import (
	"encoding/hex"
	"os"
)

func main() {
	wd, _ := os.Getwd()                 // 获取当前工作目录
	file, _ := os.Create(wd + "/a.txt") // 创建目标文件
	w := hex.NewEncoder(file)           // 创建一个流io
	w.Write([]byte("abc"))              // 写入数据
	w.Write([]byte("123"))              // 写入数据
	// 上面会将字符 abc123 转换为16进制后，写入到a.txt文件中
}
