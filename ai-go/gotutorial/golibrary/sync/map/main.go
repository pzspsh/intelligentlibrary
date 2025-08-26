/*
@File   : main.go
@Author : pan
@Time   : 2025-03-11 15:05:29
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup
	var hasValue bool
	updatech := make(chan int, 1)
	m.Store("key1", 100) // 写入键值对
	// m.Store("key2", "hello world") // 支持任意类型 (any)
	m.Store("key3", 300)
	m.Range(func(key, value any) bool { // 遍历所有键值对
		wg.Add(1)
		updatech <- 1
		go func(update chan int) {
			defer wg.Done()
			defer func() { <-update }()
			fmt.Println(key.(string), value.(int))
			time.Sleep(time.Second * 5)
		}(updatech)
		hasValue = true
		return true
	})
	wg.Wait()
	fmt.Println(hasValue)
	// m.Delete("key1") // 删除键值对 (Delete)
	// value, ok := m.Load("key1") // 读取键值对 (Load)
	// if ok {
	// 	fmt.Println(value.(int))
	// } else {
	// 	fmt.Println("not found")
	// }

	// actual, load := m.LoadOrStore("key4", 400) //存在时更新，不存在时写入 (LoadOrStore)
	// if load {
	// 	fmt.Println(actual.(int))
	// } else {
	// 	value, ok := m.Load("key4")
	// 	if ok {
	// 		fmt.Println(value.(int))
	// 	}
	// 	fmt.Println("not found")
	// }

	// value1, ok1 := m.LoadAndDelete("key1") // 读取并删除 (LoadAndDelete)
	// if ok1 {
	// 	fmt.Println(value1.(int))
	// } else {
	// 	fmt.Println("not found")
	// }

	// value, ok := m.Load("key1") // 读取键值对 (Load)
	// if ok {
	// 	fmt.Println(value.(int))
	// } else {
	// 	fmt.Println("not found")
	// }

	// ok := m.CompareAndDelete("key1", 100) // 比较并删除 (CompareAndDelete)
	// if ok {
	// 	fmt.Println("delete success")
	// } else {
	// 	fmt.Println("delete fail")
	// }

	// ok := m.CompareAndSwap("key1", 100, 102) // 比较并交换 (CompareAndSwap)
	// if ok {
	// 	fmt.Println("swap success")
	// 	value, ok := m.Load("key1")
	// 	if ok {
	// 		fmt.Println(value.(int))
	// 	}
	// } else {
	// 	fmt.Println("swap fail")
	// }

	// value, ok := m.Swap("key1", 102)
	// if ok {
	// 	fmt.Println(value.(int))
	// 	value, ok := m.Load("key1")
	// 	if ok {
	// 		fmt.Println(value.(int))
	// 	}
	// } else {
	// 	fmt.Println("swap fail")
	// }
}

func Get() {

}

func Set() {

}

func Demo() {
	var m sync.Map

	// 并发写入
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			m.Store(key, n*10)
		}(i)
	}

	// 等待写入完成
	wg.Wait()

	// 读取示例
	if val, ok := m.Load("key5"); ok {
		fmt.Println("key5:", val.(int)) // 输出: key5: 50
	}

	// 遍历所有键值对
	m.Range(func(key, value any) bool {
		fmt.Printf("%s: %d\n", key.(string), value.(int))
		return true
	})
}
