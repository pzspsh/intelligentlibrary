package main

import (
	"io"
	"strings"
)

func main() {
	comment := "Because these interfaces and primitives wrap lower-level operations with various implementations, " +
		"unless otherwise informed clients should not assume they are safe for parallel execution."
	basicReader := strings.NewReader(comment)
	basicWriter := new(strings.Builder)

	// 示例1。
	reader1 := io.LimitReader(basicReader, 98)
	_ = any(reader1).(io.Reader)

	// 示例2。
	reader2 := io.NewSectionReader(basicReader, 98, 89)
	_ = any(reader2).(io.Reader)
	_ = any(reader2).(io.ReaderAt)
	_ = any(reader2).(io.Seeker)

	// 示例3。
	reader3 := io.TeeReader(basicReader, basicWriter)
	_ = any(reader3).(io.Reader)

	// 示例4。
	reader4 := io.MultiReader(reader1)
	_ = any(reader4).(io.Reader)

	// 示例5。
	writer1 := io.MultiWriter(basicWriter)
	_ = any(writer1).(io.Writer)

	// 示例6。
	pReader, pWriter := io.Pipe()
	_ = any(pReader).(io.Reader)
	_ = any(pReader).(io.Closer)
	_ = any(pWriter).(io.Writer)
	_ = any(pWriter).(io.Closer)
}
