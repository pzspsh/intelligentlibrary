/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:02:39
*/
package main

import "fmt"

func main() {
	// if条件句无需加括号
	if 7%2 == 0 {
		fmt.Println("7 is even")
		// 注意}和else必须在同一行，否则会编译报错
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if可支持多条语句，用分号分隔
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
