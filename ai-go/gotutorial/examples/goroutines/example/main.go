/*
@File   : main.go
@Author : pan
@Time   : 2024-05-13 09:54:16
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	task := map[string]TaskStat{}
	fmt.Println(len(task))
}

type TaskStat struct {
	End     bool
	Timeout time.Time
}

func ResultStat(dateresults chan map[string]string) {
	SubBeExec := make(chan map[string]bool)
	for {
		if datas, ok := <-dateresults; ok {
			if number, ok := datas["number"]; ok {
				if len(SubBeExec) > 0 {
					if mapdata, ok := <-SubBeExec; ok {
						if ok := mapdata[number]; !ok {
							mapdata[number] = true
							SubBeExec <- mapdata
							go func(number string, submap chan map[string]bool) {
								fmt.Println("执行任务：", number)
								demap := <-submap
								delete(demap, number)
								submap <- demap
							}(number, SubBeExec)
						} else {
							SubBeExec <- mapdata
							result := map[string]string{}
							result[number] = "number"
							dateresults <- result
						}
					}
				} else {
					mapses := map[string]bool{}
					mapses[number] = true
					SubBeExec <- mapses
					go func(number string, submap chan map[string]bool) {
						fmt.Println("执行任务：", number)
						demap := <-submap
						delete(demap, number)
						submap <- demap
					}(number, SubBeExec)
				}
			}
		}
	}
}
