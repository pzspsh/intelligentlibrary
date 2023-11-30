/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:09:05
*/
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	// 定义变量
	var str = "Hello World"

	// 实例化sha512
	my_sha512 := sha512.New()

	// 写入
	my_sha512.Write([]byte(str))

	// 计算哈希
	result := my_sha512.Sum(nil)

	// 以字节的形式输出
	fmt.Println(result)
	// 以字符串的形式输出
	fmt.Printf("sha512 值: %x", result)

	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(result)
	fmt.Println(hashCode)
}
