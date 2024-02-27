/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:57:55
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "Hello, world!" // 要计算哈希值的字符串

	// 创建一个MD5哈希对象
	hasher := md5.New()

	// 将字符串转换为字节数组并计算哈希值
	hasher.Write([]byte(str))

	// 获取计算得到的哈希值
	hash := hasher.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hash)

	// 打印结果
	fmt.Println("原始字符串：", str)
	fmt.Println("MD5哈希值（十六进制）：", hashString)
}
