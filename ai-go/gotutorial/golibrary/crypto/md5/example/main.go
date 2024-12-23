/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:46:00
*/
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "yangyangyang"
	//方法一
	data := []byte(str)
	has1 := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has1) //将[]byte转成16进制
	fmt.Println(md5str1)

	str2 := "yangyangyang11"
	data2 := []byte(str2)
	has2 := md5.Sum(data2)
	md5str2 := fmt.Sprintf("%x", has2) //将[]byte转成16进制
	fmt.Println(md5str2)
}
