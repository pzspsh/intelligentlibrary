/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:00:19
*/
package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	aliceKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	bobKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	alicePubkey := aliceKey.PublicKey()

	shared, _ := bobKey.ECDH(alicePubkey)
	bobShared := sha256.Sum256(shared)
	fmt.Printf("秘钥哈希(Bob)  %x\n", bobShared)
	// 秘钥哈希(Bob)  a74e7949e71ead5f3bd4de031e2ad45c3f5b80b48ccf50e50eb86f4bdb025c3a

	bobPubkey := bobKey.PublicKey()
	shared, _ = aliceKey.ECDH(bobPubkey)
	aliceShared := sha256.Sum256(shared)
	fmt.Printf("秘钥哈希(Alice)  %x\n", aliceShared)
	// 秘钥哈希(Alice)  a74e7949e71ead5f3bd4de031e2ad45c3f5b80b48ccf50e50eb86f4bdb025c3a
}
