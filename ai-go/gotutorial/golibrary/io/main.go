/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 17:24:38
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// 示例1。
	builder := new(strings.Builder)
	_ = any(builder).(io.Writer)
	_ = any(builder).(io.ByteWriter)
	_ = any(builder).(fmt.Stringer)

	// 示例2。
	reader := strings.NewReader("")
	_ = any(reader).(io.Reader)
	_ = any(reader).(io.ReaderAt)
	_ = any(reader).(io.ByteReader)
	_ = any(reader).(io.RuneReader)
	_ = any(reader).(io.Seeker)
	_ = any(reader).(io.ByteScanner)
	_ = any(reader).(io.RuneScanner)
	_ = any(reader).(io.WriterTo)

	// 示例3。
	buffer := bytes.NewBuffer([]byte{})
	_ = any(buffer).(io.Reader)
	_ = any(buffer).(io.ByteReader)
	_ = any(buffer).(io.RuneReader)
	_ = any(buffer).(io.ByteScanner)
	_ = any(buffer).(io.RuneScanner)
	_ = any(buffer).(io.WriterTo)

	_ = any(buffer).(io.Writer)
	_ = any(buffer).(io.ByteWriter)
	_ = any(buffer).(io.ReaderFrom)

	_ = any(buffer).(fmt.Stringer)

	// 示例4。
	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	dst := new(strings.Builder)
	written, err := io.CopyN(dst, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}
}
