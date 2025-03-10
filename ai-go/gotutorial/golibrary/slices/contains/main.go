/*
@File   : main.go
@Author : pan
@Time   : 2025-03-06 14:51:20
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	NOT_REQ_PORT := []int{21, 23, 25, 22, 135, 445, 389, 873, 1433, 1521, 2181, 3306, 3389, 5432, 6379, 11211, 27017, 27018}
	port := 21
	if slices.Contains(NOT_REQ_PORT, port) {
		fmt.Println("true")
	}

	for _, v := range NOT_REQ_PORT {
		if port == v {
			fmt.Println("true")
		}
	}
}
