package main

/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:49:41
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 打开文件，创建buffered reader
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)

	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)

	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}
