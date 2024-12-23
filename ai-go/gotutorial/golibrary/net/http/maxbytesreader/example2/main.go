/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:43:55
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

const maxRequestBodySize = 1024 // 1 KB

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	limitedReader := http.MaxBytesReader(w, r.Body, maxRequestBodySize)
	defer r.Body.Close()

	bodyContent := make([]byte, maxRequestBodySize)
	n, err := limitedReader.Read(bodyContent)

	if err != nil && err != io.EOF {
		if err == io.ErrUnexpectedEOF {
			http.Error(w, "Request body too large", http.StatusRequestEntityTooLarge)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	fmt.Fprintf(w, "Received %d bytes\n", n)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting server on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
