/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:19:55
*/
package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed text/*.txt
var content embed.FS

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
