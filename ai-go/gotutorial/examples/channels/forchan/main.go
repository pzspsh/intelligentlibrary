/*
@File   : main.go
@Author : pan
@Time   : 2024-11-21 10:58:31
*/
package main

import (
	"fmt"
	"time"
)

func Inputchan1(data chan string) {
	var i int
	for {
		i++
		data <- fmt.Sprintf("inputchan1-%v", i)
		time.Sleep(1 * time.Second)
	}
}

func Inputchan2(data chan string) {
	var i int
	for {
		i++
		data <- fmt.Sprintf("inputchan2-%v", i)
		time.Sleep(1 * time.Second)
	}
}

func Outputchan(datas chan string) {
	for {
		if data, ok := <-datas; ok {
			fmt.Println(data)
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	data := make(chan string, 10)
	go Inputchan1(data)
	go Inputchan2(data)
	Outputchan(data)
}
