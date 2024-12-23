/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:55:08
*/
package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main() {
	// This is an example of writing a compressing reader.
	// This can be useful for an HTTP client body, as shown.

	const testdata = "the data to be compressed"

	// This HTTP handler is just for testing purposes.
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		zr, err := gzip.NewReader(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Just output the data for the example.
		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// The remainder is the example code.

	// The data we want to compress, as an io.Reader
	dataReader := strings.NewReader(testdata)

	// bodyReader is the body of the HTTP request, as an io.Reader.
	// httpWriter is the body of the HTTP request, as an io.Writer.
	bodyReader, httpWriter := io.Pipe()

	// Make sure that bodyReader is always closed, so that the
	// goroutine below will always exit.
	defer bodyReader.Close()

	// gzipWriter compresses data to httpWriter.
	gzipWriter := gzip.NewWriter(httpWriter)

	// errch collects any errors from the writing goroutine.
	errch := make(chan error, 1)

	go func() {
		defer close(errch)
		sentErr := false
		sendErr := func(err error) {
			if !sentErr {
				errch <- err
				sentErr = true
			}
		}

		// Copy our data to gzipWriter, which compresses it to
		// gzipWriter, which feeds it to bodyReader.
		if _, err := io.Copy(gzipWriter, dataReader); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := gzipWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := httpWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
	}()

	// Send an HTTP request to the test server.
	req, err := http.NewRequest("PUT", ts.URL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	// Note that passing req to http.Client.Do promises that it
	// will close the body, in this case bodyReader.
	resp, err := ts.Client().Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Check whether there was an error compressing the data.
	if err := <-errch; err != nil {
		log.Fatal(err)
	}

	// For this example we don't care about the response.
	resp.Body.Close()

}
