/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:25:43
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成[0, 100)范围内的整数
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	// 随机生成[0, 1)范围内的float64小数
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 随机函数默认以当前时间戳为种子，你也可以显式为它指定种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 用同一种子生成的随机数序列必定相同
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
