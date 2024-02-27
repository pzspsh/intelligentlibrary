/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	// rand.Seed(time.Now().UnixNano())  // 该功能已经弃用
	fmt.Println("A number from 1-100", rand.Intn(81))
}
