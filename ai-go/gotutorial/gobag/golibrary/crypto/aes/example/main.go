/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:41:40
*/
package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// 填充
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	fmt.Println(string(pad))
	return append(src, pad...)
}

// 去掉填充
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptDES(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 解密
func decryptDES(src []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func main() {
	d := []byte("hello world")
	key := []byte("12345678")
	fmt.Println("加密前:", string(d))
	x1, err := encryptDES(d, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("加密后:", string(x1))
	x2, err := decryptDES(x1, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("解密后:", string(x2))
}
