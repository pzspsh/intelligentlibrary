/*
@File   : main.go
@Author : pan
@Time   : 2024-05-15 10:36:27
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := 0
	go func() {
		for {
			t++
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		fmt.Println(t)
		time.Sleep(2 * time.Second)
	}
}
