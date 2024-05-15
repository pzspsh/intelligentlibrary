/*
@File   : main.go
@Author : pan
@Time   : 2024-05-14 17:36:29
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskStat struct {
	End     bool
	Timeout time.Time
}

// SharedData 定义了共享数据的结构
type SharedData struct {
	mu    sync.Mutex
	Value map[string]TaskStat
}

// Set 更新共享数据的值
func (sd *SharedData) Set(value map[string]TaskStat) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	sd.Value = value
}

// Get 获取共享数据的值
func (sd *SharedData) Get() map[string]TaskStat {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	return sd.Value
}

func main() {
	// 创建一个SharedData实例
	sharedData := &SharedData{
		Value: map[string]TaskStat{},
	}

	// 启动第一个goroutine，它将修改共享数据
	go func() {
		i := 0
		for {
			i++
			time.Sleep(1 * time.Second) // 模拟工作负载
			newValue := sharedData.Get()
			newValue[fmt.Sprintf("task:%v", i)] = TaskStat{End: false, Timeout: time.Now().Add(10 * time.Second)}
			fmt.Println("Goroutine 1: Setting value to", newValue)
			sharedData.Set(newValue)
		}
	}()

	// 启动第二个goroutine，它将读取共享数据
	go func() {
		for {
			value := sharedData.Get()
			fmt.Println("Goroutine 2: Current value is", value)
			for key := range value { // 检查是否应该结束
				delete(value, key)
			}
			sharedData.Set(value)
			time.Sleep(1 * time.Second) // 模拟不同的工作负载
		}
	}()

	// 等待，防止主goroutine退出
	select {}
}
