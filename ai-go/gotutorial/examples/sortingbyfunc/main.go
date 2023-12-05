/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:08:10
*/
package main

import (
	"fmt"
	"sort"
)

// 新类型是string slice的别名
type byLength []string

// 实现了排序所需的三个方法：Len、Swap、Less
// byLength实际就是string slice，是引用类型，因此不用传指针
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// byLength实现排序接口后，可执行被Sort调用
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
