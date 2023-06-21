/*
@File   : readfilewithreader.go
@Author : pan
@Time   : 2023-06-21 15:07:29
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func ReadFileWithReader(r io.Reader) (chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out, nil
}

func main() {
	client := http.Client{}
	resp, _ := client.Get("http://www.baidu.com")
	result, err := ReadFileWithReader(resp.Body)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Println(<-result)
}
