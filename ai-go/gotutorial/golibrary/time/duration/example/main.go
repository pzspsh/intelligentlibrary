/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:09:00
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Duration(3) * time.Hour
	fmt.Println(a.Hours())        // 3
	fmt.Println(a.Minutes())      // 180
	fmt.Println(a.Seconds())      //10800
	fmt.Println(a.Milliseconds()) //10800000
	fmt.Println(a.Microseconds()) //10800000000
	fmt.Println(a.Nanoseconds())  //10800000000000
	fmt.Println(a.String())       //3h0m0s
}
