/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:24:36
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("path/dummy.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
