/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 15:54:41
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	envs := os.Environ()
	for _, env := range envs {
		cache := strings.Split(env, "=")
		fmt.Printf("key=%v value=%v\n", cache[0], cache[1])
	}
}
