/*
@File   : main.go
@Author : pan
@Time   : 2023-10-27 14:17:35
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	taskch := make(chan []byte, 100)
	stop := make(chan bool)
	go func() {
		// var periodtime int64 = 86400
		var periodtime int64 = 4
		var starttime int64 = time.Now().Unix()
		var endtime int64
		for {
			if endtime-starttime >= periodtime {
				stop <- true
				starttime = time.Now().Unix()
			} else {
				endtime = time.Now().Unix()
				time.Sleep(2 * time.Second)
			}
		}
	}()
	for {
		select {
		case <-stop:
			fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAA")
			return
		case <-taskch:
			fmt.Println("exec run script scan")
		default:
			fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBB")
			time.Sleep(1 * time.Second)
		}
	}
}
