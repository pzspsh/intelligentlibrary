/*
@File   : output.go
@Author : pan
@Time   : 2024-07-30 16:20:31
*/

package output

import (
	"fmt"
	"time"

	"gotutorial/examples/asyncassign/common"
)

func Output() {
	for {
		if resultEvent, ok := <-common.ChanData; ok {
			fmt.Println("output: ", resultEvent)
		}
		time.Sleep(2 * time.Second)
	}
}
