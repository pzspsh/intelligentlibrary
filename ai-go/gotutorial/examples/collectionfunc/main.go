/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:24:47
*/
package main

import (
	"fmt"
	"strings"
)

// 返回目标字符串 `t` 出现的的第一个索引位置，
// 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 如果目标字符串 `t` 在这个切片中则返回 `true`。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果这些切片中的字符串有一个满足条件 `f` 则返回 `true`。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果切片中的所有字符串都满足条件 `f` 则返回 `true`。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个包含所有切片中满足条件 `f` 的字符串的新切片。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个对原始切片中所有字符串执行函数 `f` 后的新切片。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

/*
这是一些 strings 切片的组合函数示例。你可以使用这些例子来构建自己的函数。注意有时候，
直接使用内联组合操作代码会更清晰，而不是创建并调用一个帮助函数。
*/
func main() {
	// 这里试试这些组合函数。
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 上面的例子都是用的匿名函数，但是你也可以使用类型正确的命名函数
	fmt.Println(Map(strs, strings.ToUpper))
}
