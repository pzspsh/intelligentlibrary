/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:46:32
*/
package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(b)
	fmt.Println(bytes.Equal(b, make([]byte, c)))
}
