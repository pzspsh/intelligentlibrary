/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:51:19
*/
package main

import (
	"crypto/des"
	"fmt"
)

func main() {
	ede2Key := []byte("example key 1234")
	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)
	desCipher, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}
	var inputData = []byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34}
	out := make([]byte, len(inputData))
	desCipher.Encrypt(out, inputData)
	fmt.Printf("Encrypted data : %#v\n", out) //Encrypted data : []byte{0x39, 0x9e, 0xbe, 0xa9, 0xc3, 0xfa, 0x77, 0x5e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

	plain := make([]byte, len(inputData))
	desCipher.Decrypt(plain, out)
	fmt.Printf("Decrypted data : %#v\n", plain) //Decrypted data : []byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
}
