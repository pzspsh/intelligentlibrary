/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:04:02
*/
package main

import (
	"fmt"
)

func test() {
	//使用defer+ recover 来捕获和处理异常
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	//测试
	test()
	fmt.Println("main()下面的代码....")
}
