/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:12:20
*/
package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

// 未完成
func main() {
	message := []byte("message to be signed")

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha256.Sum256(message)
	rsaPrivateKey := &rsa.PrivateKey{} // 赋值
	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	fmt.Printf("Signature: %x\n", signature)
}
