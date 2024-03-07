/*
@File   : main.go
@Author : pan
@Time   : 2024-03-07 14:18:33
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	now := time.Now().In(loc)
	fmt.Println(now)
	currentTime := time.Now()
	fmt.Println(currentTime)
}
