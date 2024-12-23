/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:53:28
*/
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func rsaSignature(plainText []byte, privateKeyFile string) ([]byte, error) {
	// 读取私钥文件并解析成rsa.PrivateKey
	file, err := os.Open(privateKeyFile)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, stat.Size())
	file.Read(buf)
	defer file.Close()

	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	// 计算原始内容的散列值
	h := sha512.New()
	h.Write(plainText)
	hValue := h.Sum(nil)

	// 通过rsa.SignPKCS1v15使用私钥对原始内容散列值进行签名
	digestSign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hValue)
	return digestSign, err
}

func rsaVerifySign(plainText []byte, publicKeyFile string, signed []byte) bool {
	// 读取公钥文件并解析成rsa.PublicKey
	file, err := os.Open(publicKeyFile)
	if err != nil {
		return false
	}
	stat, err := file.Stat()
	if err != nil {
		return false
	}
	buf := make([]byte, stat.Size())
	file.Read(buf)
	defer file.Close()

	block, _ := pem.Decode(buf)
	publicKeyInt, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := publicKeyInt.(*rsa.PublicKey)

	// 计算原始内容的散列值
	h := sha512.New()
	h.Write(plainText)
	hValue := h.Sum(nil)

	// 确认签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hValue, signed)

	return err == nil
}

func main() {
	content := []byte("Hello digest sign")
	sign, err := rsaSignature(content, "private.pem")
	if err != nil {
		return
	}
	fmt.Println("signature:", sign)
	fmt.Println("verify result:", rsaVerifySign(content, "public.pem", sign))
}
