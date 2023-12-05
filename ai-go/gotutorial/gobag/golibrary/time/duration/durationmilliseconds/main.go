/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 00:05:14
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d milliseconds.\n", u.Milliseconds())
}
