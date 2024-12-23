/*
@File   : main.go
@Author : pan
@Time   : 2023-06-15 09:59:30
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"syscall"
)

const BufferSize = 100

func Readfiledata() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Readfile() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesread, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
}

func Readfile1() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bytes read: ", bytesread)
	fmt.Println("bytestream to string: ", string(buffer))
}

type chunk struct {
	bufsize int
	offset  int64
}

func Readfile2() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := int(fileinfo.Size())
	// Number of go routines we need to spawn.
	concurrency := filesize / BufferSize
	// buffer sizes that each of the go routine below should use. ReadAt
	// returns an error if the buffer size is larger than the bytes returned
	// from the file.
	chunksizes := make([]chunk, concurrency)

	// All buffer sizes are the same in the normal case. Offsets depend on the
	// index. Second go routine should start at 100, for example, given our
	// buffer size of 100.
	for i := 0; i < concurrency; i++ {
		chunksizes[i].bufsize = BufferSize
		chunksizes[i].offset = int64(BufferSize * i)
	}

	// check for any left over bytes. Add the residual number of bytes as the
	// the last chunk size.
	if remainder := filesize % BufferSize; remainder != 0 {
		c := chunk{bufsize: remainder, offset: int64(concurrency * BufferSize)}
		concurrency++
		chunksizes = append(chunksizes, c)
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunksizes []chunk, i int) {
			defer wg.Done()

			chunk := chunksizes[i]
			buffer := make([]byte, chunk.bufsize)
			bytesread, err := file.ReadAt(buffer, chunk.offset)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("bytes read, string(bytestream): ", bytesread)
			fmt.Println("bytestream to string: ", string(buffer))
		}(chunksizes, i)
	}
	wg.Wait()
}

func Readfile3() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Returns a boolean based on whether there's a next instance of `\n`
	// character in the IO stream. This step also advances the internal pointer
	// to the next position (after '\n') if it did find that token.
	for {
		read := scanner.Scan()
		if !read {
			break

		}
		fmt.Println("read byte array: ", scanner.Bytes())
		fmt.Println("read string: ", scanner.Text())
	}
}

func Readfile4() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("read lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
}

func Readfile5() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}
}

func Readfile6() {
	fd, err := syscall.Open("example.py", syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Failed on open: ", err)
	}
	defer syscall.Close(fd)

	var wg sync.WaitGroup
	wg.Add(2)
	dataChan := make(chan []byte)
	go func() {
		wg.Done()
		for {
			data := make([]byte, 100)
			n, _ := syscall.Read(fd, data)
			if n == 0 {
				break
			}
			dataChan <- data
		}
		close(dataChan)
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case data, ok := <-dataChan:
				if !ok {
					return
				}
				fmt.Println(string(data))
			default:
				fmt.Println()
			}
		}
	}()
	wg.Wait()
}

func Readfile7() {
	file, err := os.Open("filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// initial size of our wordlist
	bufferSize := 50
	words := make([]string, bufferSize)
	pos := 0

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			// This error is a non-EOF error. End the iteration if we encounter
			// an error
			fmt.Println(err)
			break
		}

		words[pos] = scanner.Text()
		pos++

		if pos >= len(words) {
			// expand the buffer by 100 again
			newbuf := make([]string, bufferSize)
			words = append(words, newbuf...)
		}
	}

	fmt.Println("word list:")
	// we are iterating only until the value of "pos" because our buffer size
	// might be more than the number of words because we increase the length by
	// a constant value. Or the scanner loop might've terminated due to an
	// error prematurely. In this case the "pos" contains the index of the last
	// successful update.
	for _, word := range words[:pos] {
		fmt.Println(word)
	}
}

func Readfile8() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建一个新的bufio.Scanner来读取文件内容
	scanner := bufio.NewScanner(file)

	// 循环读取文件的每一行
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// 检查读取过程中是否有错误发生
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Readfile9() {
	longstring := "This is a very long string. Not."
	var words []string
	scanner := bufio.NewScanner(strings.NewReader(longstring))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}
}

func Readfile10() {
	csvstring := "name, age, occupation"

	// An anonymous function declaration to avoid repeating main()
	ScanCSV := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		commaidx := bytes.IndexByte(data, ',')
		if commaidx > 0 {
			// we need to return the next position
			buffer := data[:commaidx]
			return commaidx + 1, bytes.TrimSpace(buffer), nil
		}

		// if we are at the end of the string, just return the entire buffer
		if atEOF {
			// but only do that when there is some data. If not, this might mean
			// that we've reached the end of our input CSV string
			if len(data) > 0 {
				return len(data), bytes.TrimSpace(data), nil
			}
		}
		// when 0, nil, nil is returned, this is a signal to the interface to read
		// more data in from the input reader. In this case, this input is our
		// string reader and this pretty much will never occur.
		return 0, nil, nil
	}

	scanner := bufio.NewScanner(strings.NewReader(csvstring))
	scanner.Split(ScanCSV)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadFile11() {
	content, err := os.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
}

func ReadDir() {
	filelist, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, fileinfo := range filelist {
		if fileinfo.IsDir() {
			bytes, err := os.ReadFile(fileinfo.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Bytes read: ", len(bytes))
			fmt.Println("String read: ", string(bytes))
		}
	}
}

func main() {

}
