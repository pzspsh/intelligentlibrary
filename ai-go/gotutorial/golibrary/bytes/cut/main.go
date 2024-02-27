/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:26:51
*/
package main

import (
	"bytes"
	"fmt"
)

/*
在sep的第一个实例周围切片，返回SEP之前和之后的文本。找到的结果报告sep
是否出现在s中。 如果sep未出现在s中，则cut返回snilfalse。

Cut 返回原始切片的切片，而不是副本。
*/
func main() {
	show := func(s, sep string) {
		before, after, found := bytes.Cut([]byte(s), []byte(sep))
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")
}
