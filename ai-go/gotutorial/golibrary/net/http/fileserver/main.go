/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:49:51
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}
