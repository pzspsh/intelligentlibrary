/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:54:31
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(&r)
	for i := 0; i < 10; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	fmt.Println(result)
	fmt.Println(string(result))
}
