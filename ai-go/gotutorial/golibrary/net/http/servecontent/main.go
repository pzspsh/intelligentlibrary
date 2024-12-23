/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:36:47
*/
package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", serveContentHandler)
	http.ListenAndServe(":8080", nil)
}

func serveContentHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("example.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
