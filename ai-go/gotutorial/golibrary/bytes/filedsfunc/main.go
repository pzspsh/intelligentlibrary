/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:56:31
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
FieldsFunc将s解释为utf -8编码的代码点序列。它在每次运行代码点c满足f(c)时分割切片s，
并返回s的子切片的切片。如果s中的所有代码点都满足f(c)，或者len(s) == 0，则返回一个空切片。

FieldsFunc不保证调用f(c)的顺序，并假设对于给定的c, f总是返回相同的值。
*/
func main() {
	s := bytes.FieldsFunc([]byte(" hi 你啊,    is-not.good, my,boy"), func(r rune) bool {
		return r == ',' || r == '-' || r == '.' // 只要是,-. 都可以作为分隔符
	})
	for _, v := range s {
		fmt.Print(string(v) + "|") //  hi 你啊|    is|not|good| my|boy|
	}
	fmt.Println("###########################################")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
}
