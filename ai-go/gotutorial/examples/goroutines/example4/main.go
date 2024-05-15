/*
@File   : main.go
@Author : pan
@Time   : 2024-05-15 10:43:33
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	datachan := make(chan map[string]string, 100)
	execmap1 := sync.Map{}
	execmap2 := sync.Map{}
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
			time.Sleep(1 * time.Second)
		}
	}(datachan)
	for {
		if data, ok := <-datachan; ok {
			if key, ok := data["test"]; ok {
				if _, ok := execmap1.Load(key); !ok {
					execmap1.Store(key, true)
					go func() {
						fmt.Println("AAAAAAAAAAA", key)
						time.Sleep(2 * time.Second)
						execmap1.Delete(key)
					}()
				} else {
					fmt.Println("CCCCCCCCCCCCC", data)
					datachan <- data
				}
			} else if key, ok := data["demo"]; ok {
				if _, ok := execmap2.Load(key); !ok {
					execmap2.Store(key, true)
					go func() {
						fmt.Println("BBBBBBBBBBB", key)
						time.Sleep(2 * time.Second)
						execmap2.Delete(key)
					}()
				} else {
					fmt.Println("DDDDDDDDDDDDDDD", data)
					datachan <- data
				}
			}
		}
	}
}
