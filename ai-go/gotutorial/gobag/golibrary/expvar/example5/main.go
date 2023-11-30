/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:07:39
*/
package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //初始化一个gin实例

	router.GET("/debug/vars", GetCurrentRunningStats) //接口路由，如果url不是/debug/vars，则用metricBeat去获取会出问题
	s := &http.Server{
		Addr:           ":" + "6666",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe() //开始监听
}

var CuMemoryPtr *map[string]string
var BTCMemoryPtr *map[string]interface{}

// 开始时间
var start = time.Now()

// calculateUptime 计算运行时间
func calculateUptime() interface{} {
	return time.Since(start).String()
}

// currentGoVersion 当前 Golang 版本
func currentGoVersion() interface{} {
	return runtime.Version()
}

// getNumCPUs 获取 CPU 核心数量
func getNumCPUs() interface{} {
	return runtime.NumCPU()
}

// getGoOS 当前系统类型
func getGoOS() interface{} {
	return runtime.GOOS
}

// getNumGoroutins 当前 goroutine 数量
func getNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// getNumCgoCall CGo 调用次数
func getNumCgoCall() interface{} {
	return runtime.NumCgoCall()
}

// 业务特定的内存数据
func getCuMemoryMap() interface{} {
	if CuMemoryPtr == nil {
		return 0
	} else {
		return len(*CuMemoryPtr)
	}
}

// 业务特定的内存数据
func getBTCMemoryMap() interface{} {
	if BTCMemoryPtr == nil {
		return 0
	} else {
		return len(*BTCMemoryPtr)
	}
}

var lastPause uint32

// getLastGCPauseTime 获取上次 GC 的暂停时间
func getLastGCPauseTime() interface{} {
	var gcPause uint64
	ms := new(runtime.MemStats)

	statString := expvar.Get("memstats").String()
	if statString != "" {
		json.Unmarshal([]byte(statString), ms)

		if lastPause == 0 || lastPause != ms.NumGC {
			gcPause = ms.PauseNs[(ms.NumGC+255)%256]
			lastPause = ms.NumGC
		}
	}

	return gcPause
}

// GetCurrentRunningStats 返回当前运行信息
func GetCurrentRunningStats(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(c.Writer, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(c.Writer, "%q: %q", key, str)
		} else {
			fmt.Fprintf(c.Writer, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(c.Writer, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(c.Writer, "\n}\n")

	c.String(http.StatusOK, "")
}

func init() { //这些都是自定义变量，发布到expvar中，每次请求接口，expvar会自动去获取这些变量，并返回
	expvar.Publish("运行时间", expvar.Func(calculateUptime))
	expvar.Publish("version", expvar.Func(currentGoVersion))
	expvar.Publish("cores", expvar.Func(getNumCPUs))
	expvar.Publish("os", expvar.Func(getGoOS))
	expvar.Publish("cgo", expvar.Func(getNumCgoCall))
	expvar.Publish("goroutine", expvar.Func(getNumGoroutins))
	expvar.Publish("gcpause", expvar.Func(getLastGCPauseTime))
	expvar.Publish("CuMemory", expvar.Func(getCuMemoryMap))
	expvar.Publish("BTCMemory", expvar.Func(getBTCMemoryMap))
}
