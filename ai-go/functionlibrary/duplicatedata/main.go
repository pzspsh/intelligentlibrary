/*
@File   : main.go
@Author : pan
@Time   : 2023-09-18 13:38:19
*/
package main

import (
	"fmt"
)

func ArrayIntersection(arr []string) []string {
	var intersection []string
	sameElem := make(map[string]int)
	for _, v := range arr {
		if _, ok := sameElem[v]; ok {
			intersection = append(intersection, v)
		} else {
			sameElem[v] = 1
		}
	}
	return intersection
}

func main() {
	arr := []string{"2", "2", "2", "3", "3", "4", "5", "6"}
	arrres := ArrayIntersection(arr)
	fmt.Println(arrres)
}
