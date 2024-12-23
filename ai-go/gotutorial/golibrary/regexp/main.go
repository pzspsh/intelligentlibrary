/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:22:16
*/
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 直接判断正则表达式和文字是否匹配
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// 解析正则表达式。注意对于特殊符号，无需像Java一样用反斜杠转义
	r, _ := regexp.Compile("p([a-z]+)ch")
	// 再用解析好的正则，判断文字是否匹配
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))
	// 强制解析正则，失败则会panic
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
