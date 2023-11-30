/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:09:44
*/
package main

import (
	"crypto/rc4"
	"fmt"
	"log"
)

func main() {
	c, err := rc4.NewCipher([]byte("dsadsad"))
	if err != nil {
		log.Fatalln(err)
	}
	src := []byte("asdsad")
	dst := make([]byte, len(src))
	fmt.Println("Plaintext: ", src)
	c.XORKeyStream(dst, src)
	c.XORKeyStream(src, dst)
	fmt.Println("Ciphertext: ", dst)
	fmt.Println("Plaintext': ", src)
}
