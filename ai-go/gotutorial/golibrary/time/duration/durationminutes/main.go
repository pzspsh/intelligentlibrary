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
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
}
