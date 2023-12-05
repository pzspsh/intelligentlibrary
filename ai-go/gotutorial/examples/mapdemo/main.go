/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:04:59
*/
package main

import "fmt"

func main() {
	// map是引用类型，也需要用make初始化
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 可以通过返回的第二个参数判断key是否存在于map中
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 可以在定义时直接初始化
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
