/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:59:39
*/
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("panic 异常后执行...")
	panic("panic 错误...") // 抛出一个panic异常
	// fmt.Println("end...")
}
