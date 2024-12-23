/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 10:55:28
*/
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data))
}
