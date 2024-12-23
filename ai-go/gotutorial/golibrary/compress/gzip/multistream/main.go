/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:56:43
*/
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Date(2007, time.March, 2, 4, 5, 6, 1, time.UTC), "Hello Gophers - 2"},
	}

	for _, file := range files {
		zw.Name = file.name
		zw.Comment = file.comment
		zw.ModTime = file.modTime

		if _, err := zw.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		zw.Reset(&buf)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zr.Multistream(false)
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		fmt.Print("\n\n")

		err = zr.Reset(&buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

}
