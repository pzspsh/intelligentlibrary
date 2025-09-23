package main

import (
	"mime/quotedprintable"
	"os"
)

func main() {
	w := quotedprintable.NewWriter(os.Stdout)
	w.Write([]byte("These symbols will be escaped: = \t"))
	w.Close()

}
