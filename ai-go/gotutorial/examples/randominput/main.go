/*
@File   : main.go
@Author : pan
@Time   : 2024-08-05 10:32:54
*/
package main

import (
	"fmt"
)

func main() {
	InputExec(5)
	// 随机输入一个数，for循环一个总量数组，第0个和输入这个数的倍数更换值
}

func InputExec(input int) {
	count := 0
	value := ""
	for i := range 100 {
		if count%input == 0 {
			value = getnewvalue(fmt.Sprintf("%v", count))
		}
		fmt.Println(value+": ", i)
		count += 1
	}
}

func getnewvalue(url string) string {
	return url + "new value"
}
