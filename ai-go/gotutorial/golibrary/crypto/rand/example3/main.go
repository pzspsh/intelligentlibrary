/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:00:55
*/
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 生成一对RSA公私钥
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// 要加密的明文
	plainText := []byte("Hello RSA!")
	// 使用公钥加密明文
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privKey.PublicKey, plainText, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cipher text: %x\n", cipherText)
	// 使用私钥解密密文
	decryptedPlainText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, cipherText, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted plain text: %s\n", decryptedPlainText)
	// 签名消息
	hashed := sha256.Sum256([]byte("message to be signed"))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("Signature: %x\n", signature)
	// 验证签名
	err = rsa.VerifyPKCS1v15(&privKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		panic(err)
	}
}
