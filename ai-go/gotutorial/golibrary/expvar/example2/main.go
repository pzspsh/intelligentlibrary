/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:58:44
*/
package main

import (
	"expvar"
	"fmt"
	"net/http"
)

func main() {
	// 定义一个 expvar.Int 变量
	counter := expvar.NewInt("counter")

	// 设置一个 HTTP 处理函数，用于增加计数器的值
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		counter.Add(1)
		fmt.Fprintf(w, "Counter incremented")
	})

	// 设置一个 HTTP 处理函数，用于获取计数器的值
	http.HandleFunc("/value", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Counter value: %d", counter.Value())
	})

	// 启动 HTTP 服务器
	http.ListenAndServe(":8080", nil)
}

/*
在浏览器或使用工具如 cURL 访问 http://localhost:8080/increment 来增加计数器的值，
再访问 http://localhost:8080/value 来获取计数器的值。

通过使用 expvar 包，我们可以在运行时公开程序内部的变量，并通过 HTTP 接口访问这些变量。
这对于监控和调试程序非常有用，可以方便地获取和观察程序的内部状态。

*/
