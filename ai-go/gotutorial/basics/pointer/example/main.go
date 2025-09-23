package main

import (
	"fmt"
	"time"
)

func main() {
	vvsmap := make(map[string]any) // 使用make初始化
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

func add(key string, value int, vvsmap map[string]any) {
	vvsmap[fmt.Sprintf("%s%d", key, value)] = value
}
