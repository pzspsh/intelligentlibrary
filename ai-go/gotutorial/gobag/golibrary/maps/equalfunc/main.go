/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 13:46:24
*/
package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	m2 := map[int][]byte{
		1:    []byte("One"),
		10:   []byte("Ten"),
		1000: []byte("Thousand"),
	}
	eq := maps.EqualFunc(m1, m2, func(v1 string, v2 []byte) bool {
		return strings.ToLower(v1) == strings.ToLower(string(v2))
	})
	fmt.Println(eq)
}
