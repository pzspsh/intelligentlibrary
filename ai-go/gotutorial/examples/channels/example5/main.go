/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 15:43:16
*/
package main

import (
	"fmt"
	"time"
)

var NewData = make(chan map[string]string, 1000)

func Monitor() {
	for {
		fmt.Println("BBBBBBBBBBBBBB")
		if data, ok := <-NewData; ok {
			fmt.Println(data)
		}
	}
}

func pushData() {
	for {
		NewData <- map[string]string{"id": "1", "table": "tool"}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go pushData()
	Monitor()
}
