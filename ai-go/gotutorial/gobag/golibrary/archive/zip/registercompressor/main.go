/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:20:42
*/
package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
)

func main() {
	// Override the default Deflate compressor with a higher compression level.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}
