/*
@File   : readfilewithreaderandbuffersize.go
@Author : pan
@Time   : 2023-06-21 15:16:44
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// ReadFileWithReader with specific buffer size and stream on a channel
func ReadFileWithReaderAndBufferSize(r io.Reader, maxCapacity int) (chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(r)
		buf := make([]byte, maxCapacity)
		scanner.Buffer(buf, maxCapacity)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out, nil
}

func main() {
	// client := http.Client{}
	// resp, _ := client.Get("http://www.baidu.com")
	// result, err := ReadFileWithReaderAndBufferSize(resp.Body, 100000)
	// if err != nil {
	// 	fmt.Printf("err:%v", err)
	// }
	// fmt.Println(<-result)
	reader := io.NopCloser(bytes.NewReader([]byte("hello world")))
	result, err := ReadFileWithReaderAndBufferSize(reader, 12)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Println(<-result)
}
