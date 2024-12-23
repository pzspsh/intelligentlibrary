/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:12:26
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now().Add(-2 * time.Hour)
	timeDiff := time.Since(startTime)
	fmt.Printf("time diff: %s\n", timeDiff)
}
