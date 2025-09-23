package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 产生一个随机整数
	var num int64 = rand.Int63()
	fmt.Println(num)
}
