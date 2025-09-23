package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(strings.Contains("seafood", "foo"))
	// fmt.Println(strings.Contains("seafood", "bar"))
	// fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("http、https、unknown、tls", "unknown"))
	fmt.Println(strings.Contains(strings.ToLower("RocketMQ <= 5.1.0 - Remote Code Execution"), strings.ToLower("RocketMq Remote")))
}
