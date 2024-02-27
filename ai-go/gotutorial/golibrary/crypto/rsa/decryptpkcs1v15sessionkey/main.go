/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:12:20
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"os"
)

// 未完成
func main() {
	// The hybrid scheme should use at least a 16-byte symmetric key. Here
	// we read the random key that will be used if the RSA decryption isn't
	// well-formed.
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic("RNG failure")
	}

	rsaPrivateKey := &rsa.PrivateKey{} // 赋值
	rsaCiphertext, _ := hex.DecodeString("aabbccddeeff")

	if err := rsa.DecryptPKCS1v15SessionKey(nil, rsaPrivateKey, rsaCiphertext, key); err != nil {
		// Any errors that result will be “public” – meaning that they
		// can be determined without any secret information. (For
		// instance, if the length of key is impossible given the RSA
		// public key.)
		fmt.Fprintf(os.Stderr, "Error from RSA decryption: %s\n", err)
		return
	}

	// Given the resulting key, a symmetric scheme can be used to decrypt a
	// larger ciphertext.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes.NewCipher failed: " + err.Error())
	}

	// Since the key is random, using a fixed nonce is acceptable as the
	// (key, nonce) pair will still be unique, as required.
	var zeroNonce [12]byte
	aead, err := cipher.NewGCM(block)
	if err != nil {
		panic("cipher.NewGCM failed: " + err.Error())
	}
	ciphertext, _ := hex.DecodeString("00112233445566")
	plaintext, err := aead.Open(nil, zeroNonce[:], ciphertext, nil)
	if err != nil {
		// The RSA ciphertext was badly formed; the decryption will
		// fail here because the AES-GCM key will be incorrect.
		fmt.Fprintf(os.Stderr, "Error decrypting: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))
}
