/*
@File   : main.go
@Author : pan
@Time   : 2023-11-13 17:21:16
*/
package main

import "fmt"

func main() {
	a := []string{"3", "5", "5", "6", "7", "9", "5", "5", "6", "7", "9"}
	b := removeDuplicateElement(a)
	fmt.Println(b)
}

func removeDuplicateElement(data []string) []string {
	result := make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
