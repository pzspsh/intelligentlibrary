/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for _, value := range rand.Perm(3) {
		fmt.Println(value)
	}

}
