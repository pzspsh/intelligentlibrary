/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:43:38
*/
package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

func main() {
	// parameters 是私钥的参数
	var param dsa.Parameters
	// L1024N160是一个枚举，根据L1024N160来决定私钥的长度（L N）
	dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)
	// 定义私钥的变量
	var privateKey dsa.PrivateKey
	// 设置私钥的参数
	privateKey.Parameters = param
	// 生成密钥对
	dsa.GenerateKey(&privateKey, rand.Reader)
	// 公钥是存在在私钥中的，从私钥中读取公钥
	publicKey := privateKey.PublicKey
	message := []byte("hello,dsa签名")

	fmt.Println(publicKey)
	fmt.Println(privateKey)

	// 进入签名操作
	r, s, _ := dsa.Sign(rand.Reader, &privateKey, message)

	// 进入验证
	flag := dsa.Verify(&publicKey, message, r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
	flag = dsa.Verify(&publicKey, []byte("hello"), r, s)
	if flag {
		fmt.Println("数据未被修改")
	} else {
		fmt.Println("数据被修改")
	}
}
