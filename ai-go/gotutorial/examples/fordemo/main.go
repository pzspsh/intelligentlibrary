/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:02:00
*/
package main

import "fmt"

func main() {
	i := 1
	// 类似于其他语言中的while
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 常规for循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 类似于其他语言中的while(true)
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
