/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:13:00
*/
package main

import (
	"crypto/ed25519"
	"fmt"
)

/**
 * 使用ed25519进行签名和验签
 * msg 签名内容
 */
func ed25519Sign(msg string) {

	// 生成公私钥
	var publicKey []byte
	var privateKey []byte
	var err error
	if publicKey, privateKey, err = ed25519.GenerateKey(nil); err != nil {
		err = fmt.Errorf("decode private key fail: %s", err.Error())
		fmt.Println(err)
		return
	}

	msgByte := []byte(msg)

	// 进行ed25519签名
	signature := ed25519.Sign(privateKey, msgByte)

	// 使用公钥进行验签
	verify := ed25519.Verify(publicKey, msgByte, signature)

	fmt.Println("公钥：", publicKey)
	fmt.Println("私钥", privateKey)
	fmt.Println("签名", signature)
	fmt.Println("验签结果：", verify)
}

func main() {
	ed25519Sign("潘")
}
