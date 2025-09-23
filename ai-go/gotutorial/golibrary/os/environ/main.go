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
