/*
@File   : constant.go
@Author : pan
@Time   : 2023-06-01 09:59:06
*/
package main

import "fmt"

const x, y int = 1, 2     // 多常量初始化
const s = "Hello, World!" // 类型推断

const ( //常量组
	a, b      = 10, 100
	c    bool = false
)

func main() {
	fmt.Println("x和y的初始化值：", x, y)
	fmt.Println("s的初始化值：", s)
	fmt.Println("a、b、c的常量组值：", a, b, c)
}
