/*
@File   : main.go
@Author : pan
@Time   : 2023-10-31 14:50:34
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func ProcessDemo(reader io.Reader) {
	wg := &sync.WaitGroup{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		wg.Add(1)
		go func(URL string) {
			fmt.Println(URL)
			wg.Done()
		}(text)
		time.Sleep(2 * time.Second)
	}
	wg.Wait()
}

func main() {
	file, err := os.Open(`../../pkgdemo/bufiodemo/test.txt`)
	if err != nil {
		fmt.Println("os open error:", err)
	}
	ProcessDemo(file)
}
