/*
@File   : main.go
@Author : pan
@Time   : 2024-03-15 16:09:10
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	vvsmap := make(map[string]interface{}) // 使用make初始化
	key := "key"
	value := 1
	go func() {
		for {
			add(key, value, vvsmap)
			value++
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		if len(vvsmap) > 0 {
			fmt.Println(vvsmap)
		} else {
			fmt.Println("vvsmap时空的map")
		}
		time.Sleep(2 * time.Second)
	}
}

func add(key string, value int, vvsmap map[string]interface{}) {
	vvsmap[fmt.Sprintf("%s%d", key, value)] = value
}
