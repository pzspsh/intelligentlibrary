/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:55:06
*/
package main

import (
	"fmt"
	"unicode"
)

// Functions starting with "Is" can be used to inspect which table of range a
// rune belongs to. Note that runes may fit into more than one range.
func main() {
	// constant with mixed type runes
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {
		fmt.Printf("For %q:\n", c)
		if unicode.IsControl(c) { // 判断一个字符是否是控制字符，主要是策略C的字符和一些其他的字符如代理字符
			fmt.Println("\tis control rune")
		}
		if unicode.IsDigit(c) { // 判断一个r字符是否是十进制数字字符
			fmt.Println("\tis digit rune")
		}
		if unicode.IsGraphic(c) { // 判断一个字符是否是unicode图形。包括字母、标记、数字、符号、标点、空白，参见L、M、N、P、S、Zs
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) { // 判断一个字符是否是字母
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) { // 判断字符是否是小写字母
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) { // 判断一个字符是否是标记字符
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) { // 判断一个字符是否是数字字符
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) { // 判断一个字符是否是go的可打印字符 // 本函数基本和IsGraphic一致，只是ASCII空白字符U+0020会返回假
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) { // 判断一个字符是否是unicode标点字符
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) { // 判断一个字符是否是空白字符 // 在Latin-1字符空间中，空白字符为：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).其它的空白字符请参见策略Z和属性Pattern_White_Space
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) { // 判断一个字符是否是unicode符号字符
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) { // 判断字符是否是标题字母
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) { // 判断字符是否是大写字
			fmt.Println("\tis upper case rune")
		}
	}
	// Output:
	// For '\b':
	// is control rune
	// is not printable rune
	// For '5':
	// is digit rune
	// is graphic rune
	// is number rune
	// is printable rune
	// For 'Ὂ':
	// is graphic rune
	// is letter rune
	// is printable rune
	// is upper case rune
	// For 'g':
	// is graphic rune
	// is letter rune
	// is lower case rune
	// is printable rune
	// For '̀':
	// is graphic rune
	// is mark rune
	// is printable rune
	// For '9':
	// is digit rune
	// is graphic rune
	// is number rune
	// is printable rune
	// For '!':
	// is graphic rune
	// is printable rune
	// is punct rune
	// For ' ':
	// is graphic rune
	// is printable rune
	// is space rune
	// For '℃':
	// is graphic rune
	// is printable rune
	// is symbol rune
	// For 'ᾭ':
	// is graphic rune
	// is letter rune
	// is printable rune
	// is title case rune
	// For 'G':
	// is graphic rune
	// is letter rune
	// is printable rune
	// is upper case rune

	// 迭代在unicode标准字符映射中互相对应的unicode码值
	// 在与r对应的码值中（包括r自身），会返回最小的那个大于r的字符（如果有）；否则返回映射中最小的字符
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'
	// Output:
	// U+0061 'a'
	// U+0041 'A'
	// U+006B 'k'
	// U+212A 'K'
	// U+004B 'K'
	// U+0031 '1'

	const lcG = 'g'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG)) //转大写
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG)) //转小写
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG)) //转标题
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG)) //转大写
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG)) //转小写
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG)) //转标题
	// Output:
	// U+0047 'G'
	// U+0067 'g'
	// U+0047 'G'
	// U+0047 'G'
	// U+0067 'g'
	// U+0047 'G'

	const ucG1 = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG1))
	// Output:
	// U+0067 'g'

	const ucG2 = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG2))
	// Output:
	// U+0047 'G'

	const ucG3 = 'g'
	fmt.Printf("%#U\n", unicode.ToUpper(ucG3))
	// Output:
	// U+0047 'G'

	t := unicode.TurkishCase
	const lci = 'i'
	fmt.Printf("%#U\n", t.ToLower(lci))
	fmt.Printf("%#U\n", t.ToTitle(lci))
	fmt.Printf("%#U\n", t.ToUpper(lci))
	const uci = 'İ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToTitle(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))
	// Output:
	// U+0069 'i'
	// U+0130 'İ'
	// U+0130 'İ'
	// U+0069 'i'
	// U+0130 'İ'
	// U+0130 'İ'
}
