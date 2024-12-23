/*
@File   : utils_test.go
@Author : pan
@Time   : 2023-09-18 14:17:38
*/
package test

import (
	"fmt"
	"testing"
)

func TestRandomNumber(t *testing.T) {
	aaa := []string{}
	for i := 0; i < 200000; i++ {
		abc := RandomNumber("DW", 6)
		aaa = append(aaa, abc)
	}
	fmt.Println(aaa)
	fmt.Println(ArrayIntersection(aaa))
}

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
