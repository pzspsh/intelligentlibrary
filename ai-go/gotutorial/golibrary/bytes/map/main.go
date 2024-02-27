/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
)

/*
Map返回字节片s的副本，其中所有字符都根据映射函数进行了修改。如果映射返回负值，
则从字节片中删除该字符，不进行替换。s中的字符和输出被解释为utf -8编码的码点。
*/
func main() {
	fmt.Println(string(bytes.Map(func(r rune) rune {
		return r + 1 // 将每一个字符都+1
	}, []byte("abc")))) // bcd
	fmt.Println("###############################################")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s\n", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
}
