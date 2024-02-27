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
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}
