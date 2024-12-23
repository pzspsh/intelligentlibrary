/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:46:33
*/
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := []byte("secret key")
	data := []byte("message to authenticate")
	h := hmac.New(sha256.New, key)
	_, err := h.Write(data)
	if err != nil {
		fmt.Println("Error writing to HMAC:", err)
		return
	}
	result := h.Sum(nil)
	fmt.Printf("HMAC: %x\n", result)
}
