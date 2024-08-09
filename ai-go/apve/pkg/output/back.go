/*
@File   : back.go
@Author : pan
@Time   : 2024-08-09 16:27:09
*/
package output

import "fmt"

var ResultChan = make(chan Result, 10000)

type Result struct {
}

func ResultBack() {
	for {
		if data, ok := <-ResultChan; ok {
			fmt.Println(data)
		}
	}
}
