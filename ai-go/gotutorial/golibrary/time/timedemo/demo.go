/*
@File   : demo.go
@Author : pan
@Time   : 2023-09-18 14:30:02
*/
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	timeLayout := "2006-01-02T15:04:05"
	timestamp, _ := time.ParseInLocation(timeLayout, "2023-08-09T11:20:54", time.Local)
	// timestamp, _ := time.ParseInLocation(time.DateTime, "2023-08-09T11:20:54", time.Local)
	fmt.Println(timestamp)
	timenow := time.Now().Unix()
	fmt.Println(reflect.TypeOf(timenow))
	fmt.Println(timenow)
}
