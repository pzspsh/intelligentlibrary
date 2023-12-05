/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:58:01
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 解码p中最后一个utf-8编码序列，返回该码值和编码序列的长度 utf8.DecodeLastRune(b)
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[:len(b)-size]
	}
	// Output:
	// 界 3
	// 世 3
	// 1
	// , 1
	// o 1
	// l 1
	// l 1
	// e 1
	// H 1

	// 类似DecodeLastRune但输入参数是字符串 utf8.DecodeLastRuneInString(string(b))
	str := "Hello, 世界"
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c %v\n", r, size)
		str = str[:len(str)-size]
	}
	// Output:
	// 界 3
	// 世 3
	// 1
	// , 1
	// o 1
	// l 1
	// l 1
	// e 1
	// H 1

	// 解码p开始位置的第一个utf-8编码的码值，返回该码值和编码的字节数
	// 如果编码不合法，会返回(RuneError, 1)。该返回值在正确的utf-8编码情况下是不可能返回的 utf8.DecodeRune(b)
	b1 := []byte("Hello, 世界")
	for len(b1) > 0 {
		r, size := utf8.DecodeRune(b1)
		fmt.Printf("%c %v\n", r, size)
		b1 = b1[size:]
	}
	// Output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	// , 1
	// 1
	// 世 3
	// 界 3

	// 类似DecodeRune但输入参数是字符串 utf8.DecodeRuneInString(string(b))
	str1 := "Hello, 世界"
	for len(str1) > 0 {
		r, size := utf8.DecodeRuneInString(str1)
		fmt.Printf("%c %v\n", r, size)
		str1 = str1[size:]
	}
	// Output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	// , 1
	// 1
	// 世 3
	// 界 3

	// 将r的utf-8编码序列写入p（p必须有足够的长度），并返回写入的字节数 utf8.EncodeRune(b, 'H')
	r := '世'
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, r)
	fmt.Println(buf)
	fmt.Println(n)
	// Output:
	// [228 184 150]
	// 3

	// 判断切片p是否以一个码值的完整utf-8编码开始
	// 不合法的编码因为会被转换为宽度1的错误码值而被视为完整的
	// 如中文字符占3位byte，一位byte判断为false，完整的3位为true utf8.FullRune(b)
	buf1 := []byte{228, 184, 150} // 世
	fmt.Println(utf8.FullRune(buf1))
	fmt.Println(utf8.FullRune(buf1[:2]))
	// Output:
	// true
	// false

	// 类似FullRune但输入参数是字符串 utf8.FullRuneInString(string(b))
	str2 := "世"
	fmt.Println(utf8.FullRuneInString(str2))
	fmt.Println(utf8.FullRuneInString(str2[:2]))
	// Output:
	// true
	// false

	// 返回p中的utf-8编码的码值的个数。错误或者不完整的编码会被视为宽度1字节的单个码值 utf8.RuneCount(b)
	buf2 := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf2))
	fmt.Println("runes =", utf8.RuneCount(buf2))
	// Output:
	// bytes = 13
	// runes = 9

	// 类似RuneCount但输入参数是一个字符串 utf8.RuneCountInString(string(b))
	str3 := "Hello, 世界"
	fmt.Println("bytes =", len(str3))
	fmt.Println("runes =", utf8.RuneCountInString(str3))
	// Output:
	// bytes = 13
	// runes = 9

	// 返回r编码后的字节数。如果r不是一个合法的可编码为utf-8序列的值，会返回-1 utf8.RuneLen('世')
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
	// Output:
	// 1
	// 3

	// 判断字节b是否可以作为某个rune编码后的第一个字节。第二个即之后的字节总是将左端两个字位设为10 utf8.RuneStart('世')
	buf3 := []byte("a界")
	fmt.Println(utf8.RuneStart(buf3[0]))
	fmt.Println(utf8.RuneStart(buf3[1]))
	fmt.Println(utf8.RuneStart(buf3[2]))
	// Output:
	// true
	// true
	// false

	// 判断切片p是否包含完整且合法的utf-8编码序列 utf8.Valid(b)
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
	// Output:
	// true
	// false

	// 判断r是否可以编码为合法的utf-8序列 utf8.ValidRune('H')
	valid1 := 'a'
	invalid1 := rune(0xfffffff)
	fmt.Println(utf8.ValidRune(valid1))
	fmt.Println(utf8.ValidRune(invalid1))
	// Output:
	// true
	// false

	// 判断s是否包含完整且合法的utf-8编码序列
	utf8.ValidString(string(b))
	valid2 := "Hello, 世界"
	invalid2 := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Println(utf8.ValidString(valid2))
	fmt.Println(utf8.ValidString(invalid2))
	// Output:
	// true
	// false
}
