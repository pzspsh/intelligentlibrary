/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:07:16
*/
package main

import (
	"bufio"
	"os"
	"strconv"
)

// AvailableBuffer返回一个容量为b.Available()的空缓冲区。该缓冲区旨在被附加到并传递给立即后续的Write调用。该缓冲区仅在对b进行下一次写操作之前有效。
func main() {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
}
