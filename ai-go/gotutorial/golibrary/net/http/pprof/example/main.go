/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:32:40
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func readMemStats() {
	// MemStats 描述内存信息的静态变量
	var ms runtime.MemStats
	// 读取某一时刻内存情况的快照
	runtime.ReadMemStats(&ms)
	// alloc占用内存情况、堆空闲情况、堆释放情况
	log.Printf("========> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

// append的扩容情况
func test() {
	container := make([]int, 8)
	log.Println("========> loop begin .")
	// 追加元素
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}
	fmt.Println(container)
	log.Println("==========> loop end.")
}

func main() {
	// 启动 pprof,这种是采用网络监听的方式，可以通过访问界面显示出测试的数据
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	log.Println("============> [start].")
	readMemStats()
	test()
	readMemStats()
	log.Println("============> [force.gc].")
	// 强制开启gc
	runtime.GC()
	log.Println("===========>[Done] .")
	readMemStats()
	go func() {
		for {
			readMemStats()
			time.Sleep(10 * time.Second)
		}
	}()
	time.Sleep(3600 * time.Second)
}
