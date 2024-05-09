/*
@File   : main.go
@Author : pan
@Time   : 2024-05-09 16:05:03
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	aa := map[string]time.Time{}
	start := "开始"
	i := 0
	for {
		i += 1
		startb := fmt.Sprintf("%s%d", start, i)
		fmt.Println("继续")
		aa[startb] = time.Now()
		for key, value := range aa {
			// if time.Since(value) > 5*time.Second {
			// 	fmt.Println(aa[key])
			// 	delete(aa, key)
			// }
			now := time.Now()
			// if now.After(value.Add(24 * time.Hour)) {
			if now.After(value.Add(5 * time.Second)) {
				fmt.Println(key, now, aa[key])
				delete(aa, key)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
