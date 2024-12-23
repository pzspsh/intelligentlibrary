/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:28:02
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	// Set the directory to serve files from
	fileSystem := http.Dir("./static")

	// Create a FileServer handler
	fileServer := http.FileServer(fileSystem)

	// Set the endpoint for serving files
	http.Handle("/", fileServer)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
