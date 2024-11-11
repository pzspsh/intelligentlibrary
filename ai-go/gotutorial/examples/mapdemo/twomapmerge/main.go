/*
@File   : main.go
@Author : pan
@Time   : 2024-11-11 17:00:44
*/
package main

import (
	"fmt"
)

func MergeMap(map1, map2 map[string]string) map[string]string {
	if len(map1) > 0 {
		for key, value := range map2 {
			map1[key] = value
		}
		return map1
	} else {
		return map2
	}
}

func main() {
	map1 := map[string]string{"1": "2"}
	map2 := map[string]string{"3": "4"}
	res := MergeMap(map1, map2)
	fmt.Println(res)
}
