/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:43:24
*/
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	//生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	//将原始数据转换为字节数组
	msg := "this is a message"
	msgBytes := []byte(msg)

	//使用SHA256算法计算摘要值
	hashed := sha256.Sum256(msgBytes)

	//使用私钥对摘要进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message: %s\n", msg)
	fmt.Printf("Message hash: %s\n", hex.EncodeToString(hashed[:]))
	fmt.Printf("Signature: %x\n", signature)

	//使用公钥验证签名的有效性
	pubKey := privateKey.PublicKey
	err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Printf("Signature verification failed: %s\n", err)
	} else {
		fmt.Println("Signature verification succeeded")
	}
}
