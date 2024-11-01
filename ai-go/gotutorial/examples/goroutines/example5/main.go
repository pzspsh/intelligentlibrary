/*
@File   : main.go
@Author : pan
@Time   : 2024-05-15 13:31:58
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	datachan := make(chan map[string]string, 100)
	execmap1 := map[string]int{}
	execmap2 := map[string]int{}
	go func(datachan chan map[string]string) {
		i := 0
		for {
			i++
			dmap := map[string]string{}
			dmap["test"] = fmt.Sprintf("test:%v", i)
			dmap2 := map[string]string{}
			dmap2["demo"] = fmt.Sprintf("demo:%v", i)
			datachan <- dmap
			datachan <- dmap2
			if i > 9 {
				i = 0
			}
			time.Sleep(100 * time.Microsecond)
		}
	}(datachan)
	go func() {
		for {
			if data, ok := <-datachan; ok {
				if key, ok := data["test"]; ok {
					if value, ok := execmap1[key]; !ok {
						execmap1[key] = 1
					} else {
						execmap1[key] = value + 1
					}
				} else if key, ok := data["demo"]; ok {
					if value, ok := execmap2[key]; !ok {
						execmap2[key] = 1
					} else {
						execmap2[key] = value + 1
					}
				}
			}
		}
	}()
	for {
		if len(execmap1) > 0 {
			for number, value := range execmap1 {
				go func(number string, value int) {
					fmt.Println("AAAAAAAA", number, value)
					delete(execmap1, number)
				}(number, value)
			}
		}
		if len(execmap2) > 0 {
			for number, value := range execmap2 {
				go func(number string, value int) {
					fmt.Println("BBBBBBBBB", number, value)
					delete(execmap2, number)
				}(number, value)
			}
		}
		time.Sleep(2 * time.Second)
	}
}
