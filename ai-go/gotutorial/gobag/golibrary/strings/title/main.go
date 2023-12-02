/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:08:40
*/
package main

import (
	"fmt"
	"strings"
)

/*
字符串。自Go 1.18以来，Title已被弃用，自Go 1.0以来，有一个替代方案:Title用于单
词边界的规则不能正确处理Unicode标点符号。请使用golang.org/x/text/cases。
*/
func main() {
	// Compare this example to the ToTitle example.
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("loud noises"))
	fmt.Println(strings.Title("хлеб"))
}
