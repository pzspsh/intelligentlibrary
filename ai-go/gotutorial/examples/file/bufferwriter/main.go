/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:42:09
*/
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// 打开文件，只写
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 写内存buffer到硬盘
	bufferedWriter.Flush()

	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)

	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
}
