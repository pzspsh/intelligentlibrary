/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:46:13
*/
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	str := "JBSWY3DPFQQHO33SNRSCC==="
	dst := make([]byte, base32.StdEncoding.DecodedLen(len(str)))
	n, err := base32.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	dst = dst[:n]
	fmt.Printf("%q\n", dst)
}
