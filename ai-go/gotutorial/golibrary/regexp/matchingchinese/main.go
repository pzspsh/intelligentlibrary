package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "你好，世界！Hello, World!"
	re := regexp.MustCompile("[\u4e00-\u9fa5]")
	matches := re.FindAllString(str, -1)
	fmt.Println(matches)
}
