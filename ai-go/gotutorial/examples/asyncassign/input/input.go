/*
@File   : input.go
@Author : pan
@Time   : 2024-07-30 16:20:16
*/

package input

import (
	"fmt"
	"gotutorial/examples/asyncassign/common"
	"time"
)

func Input() {
	a := 0
	for {
		abc := map[string]int{}
		a++
		abc["ABC"] = a
		fmt.Println("input", abc)
		common.ChanData <- abc
		time.Sleep(3 * time.Second)
	}
}
