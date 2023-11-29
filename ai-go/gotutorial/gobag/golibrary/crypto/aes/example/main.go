/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:53:53
*/
package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
)

func EncryptAES(key string, plainText string) (string, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(plainText))
	cipher.Encrypt(out, []byte(plainText))
	return hex.EncodeToString(out), nil
}

func DecryptAES(key string, encryptText string) (string, error) {
	decodeText, _ := hex.DecodeString(encryptText)
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(decodeText))
	cipher.Decrypt(out, decodeText)

	return string(out[:]), nil
}

func main() {
	// 加密
	// cipher key
	key := "thisisakeymustmorethan16"
	// plaintext
	text := "This is a secret"
	encrypt, err := EncryptAES(key, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encrypt)
	// db647f5df56904ef3463834abc019c1d

	// 解密
	// decrypt
	decrypt := "db647f5df56904ef3463834abc019c1d"

	plaintext, err := DecryptAES(key, decrypt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(plaintext)
	// This is a secret
}
