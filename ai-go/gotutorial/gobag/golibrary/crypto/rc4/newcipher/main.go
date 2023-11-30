/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:47:23
*/
package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	var key []byte = []byte("12345678") //初始化用于加密的KEY
	rc4obj, _ := rc4.NewCipher(key)     //返回 Cipher
	rc4str := []byte("yangyangyang")    //需要加密的字符串
	plaintext := make([]byte, len(rc4str))
	rc4obj.XORKeyStream(plaintext, rc4str)

	stringinf1 := fmt.Sprintf("%x\n", plaintext) //转换字符串
	fmt.Println(stringinf1)
}
