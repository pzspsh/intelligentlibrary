/*
@File   : main.go
@Author : pan
@Time   : 2023-12-14 16:58:28
*/
package main

/*
前几天逛github发现了一个有趣的并发库-conc，其目标是：

更难出现goroutine泄漏
处理panic更友好
并发代码可读性高
从简介上看主要封装功能如下：

对waitGroup进行封装，避免了产生大量重复代码，并且也封装recover，安全性更高
提供panics.Catcher封装recover逻辑，统一捕获panic，打印调用栈一些信息
提供一个并发执行任务的worker池，可以控制并发度、goroutine可以进行复用，支持函数签名，同时提供了stream方法来保证结果有序
提供ForEach、map方法优雅的处理切片
接下来就区分模块来介绍一下这个库；
*/

import (
	"fmt"

	"github.com/sourcegraph/conc"
)

/*
func main(){
    var wg sync.WaitGroup
    for i:=0; i < 10; i++{
        wg.Add(1)
        go func() {
            defer wg.Done()
            defer func() {
                // recover panic
                err := recover()
                if err != nil {
                    fmt.Println(err)
                }
            }
            // do something
            handle()
        }
    }
    wg.Wait()
}
*/

/*
上述代码我们需要些一堆重复代码，并且需要单独在每一个func中处理recover逻辑，所以conc库对其进行了封装，
*/

// 未完成
func main() {
	wg := conc.NewWaitGroup()
	for i := 0; i < 10; i++ {
		wg.Go(doSomething)
	}
	wg.Wait()
}

func doSomething() {
	fmt.Println("test")
}
