package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := time.Duration(rand.Intn(10))
	fmt.Printf("停止：%v秒", num)
	time.Sleep(num * time.Second)
	// time.Sleep(58 * time.Millisecond)
}
