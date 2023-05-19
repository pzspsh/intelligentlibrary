/*
@File   : repeatedsubstringpattern.go
@Author : pan
@Time   : 2023-05-19 13:32:00
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("axbaxb"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aabcbaa"))
	fmt.Println(repeatedSubstringPattern("aa"))
}

/*
	axbaxb => xbaxbaxbax
	abab => bababa
	aa => aa
*/
// 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	size := len(s)
	ss := (s + s)[1 : size*2-1]
	return strings.Contains(ss, s)
}
