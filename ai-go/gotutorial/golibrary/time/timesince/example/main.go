/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:14:42
*/
package main

import (
	"fmt"
	"time"
)

// Calling main
func main() {
	// Defining time value
	// of Since method
	timevalue := time.Now()

	// Calling Since method
	// with its parameter
	Duration := time.Since(timevalue)

	// Prints time elapse
	fmt.Println("time elapse:", time.Since(timevalue))

	// Prints time elapse in nanoseconds
	fmt.Println("time elapse in nanoseconds:", Duration.Nanoseconds())
}
