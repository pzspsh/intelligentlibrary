/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:53:28
*/
package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

func eccSignature(plainText []byte, privateKeyFile string) ([]byte, []byte, error) {
	// 读取私钥文件并解析成ecdsa.PrivateKey
	file, err := os.Open(privateKeyFile)
	if err != nil {
		return nil, nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	buf := make([]byte, stat.Size())
	file.Read(buf)
	defer file.Close()

	block, _ := pem.Decode(buf)
	privateKey, _ := x509.ParseECPrivateKey(block.Bytes)

	// 计算原始内容的散列值
	h := sha512.New()
	h.Write(plainText)
	hValue := h.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hValue)
	if err != nil {
		fmt.Println(err)
	}
	rText, _ := r.MarshalText()
	sText, _ := s.MarshalText()

	return rText, sText, nil
}

func eccVerifySign(plainText []byte, publicKeyFile string, rText []byte, sText []byte) bool {
	// 读取公钥文件并解析成ecdsa.PublicKey
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
	publicKey := publicKeyInt.(*ecdsa.PublicKey)

	// 计算原始内容的散列值
	h := sha512.New()
	h.Write(plainText)
	hValue := h.Sum(nil)

	var r, s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)

	return ecdsa.Verify(publicKey, hValue, &r, &s)
}

func main() {
	content := []byte("Hello digest sign")
	r, s, err := eccSignature(content, "ecc_private.pem")
	if err != nil {
		return
	}
	fmt.Println("r:", string(r))
	fmt.Println("s:", string(s))
	fmt.Println("verify result:", eccVerifySign(content, "ecc_public.pem", r, s))
}
