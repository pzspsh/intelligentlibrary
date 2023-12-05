/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:07:34
*/
package main

import (
	"fmt"
	"sort"
)

/* Go内置的排序函数支持对不同数据类型slice的排序。 */
func main() {
	// 排序string和int要用不同的函数
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 判断是否已排序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
