/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:05:27
*/
package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 遍历slice返回的变量，第一个是下标，第二个是对应下标的元素
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 遍历map返回的变量，第一个是key，第二个是value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 遍历string返回的变量，第一个是对应rune的byte的起始下标，第二个是对应的rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	for i, c := range "我们" {
		fmt.Println(i, c)
	}
}
