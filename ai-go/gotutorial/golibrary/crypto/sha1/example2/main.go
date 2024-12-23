/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:02:36
*/
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func main() {
	fmt.Println(SHA1("123456"))
}
