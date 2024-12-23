/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:01:11
*/
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 计算MD5哈希值
	data := []byte("hello world")
	md5Hash := md5.New()
	md5Hash.Write(data)
	md5Value := md5Hash.Sum(nil)
	fmt.Printf("MD5: %x\n", md5Value)

	// 计算SHA-1哈希值
	sha1Hash := sha1.New()
	sha1Hash.Write(data)
	sha1Value := sha1Hash.Sum(nil)
	fmt.Printf("SHA-1: %x\n", sha1Value)

	// 计算SHA-256哈希值
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	sha256Value := sha256Hash.Sum(nil)
	fmt.Printf("SHA-256: %x\n", sha256Value)
}
